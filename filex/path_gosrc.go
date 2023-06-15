// Package filex
// Create on 2022/7/8
// @author xuzhuoxi
package filex

import "strings"

// IsAbs reports whether the path is absolute.
func isUnixAbs(path string) bool {
	return strings.HasPrefix(path, "/")
}

// reservedNames lists reserved Windows names. Search for PRN in
// https://docs.microsoft.com/en-us/windows/desktop/fileio/naming-a-file
// for details.
var reservedNames = []string{
	"CON", "PRN", "AUX", "NUL",
	"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
	"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
}

func isWinSlash(c uint8) bool {
	return c == WindowsSeparator || c == '/'
}

// isReservedName returns true, if path is Windows reserved name.
// See reservedNames for the full list.
func isReservedName(path string) bool {
	if len(path) == 0 {
		return false
	}
	for _, reserved := range reservedNames {
		if strings.EqualFold(path, reserved) {
			return true
		}
	}
	return false
}

// IsAbs reports whether the path is absolute.
func isWinAbs(path string) (b bool) {
	if isReservedName(path) {
		return true
	}
	l := volumeNameLen(path)
	if l == 0 {
		return false
	}
	path = path[l:]
	if path == "" {
		return false
	}
	return isWinSlash(path[0])
}

// volumeNameLen returns length of the leading volume name on Windows.
// It returns 0 elsewhere.
func volumeNameLen(path string) int {
	if len(path) < 2 {
		return 0
	}
	// with drive letter
	c := path[0]
	if path[1] == ':' && ('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') {
		return 2
	}
	// is it UNC? https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx
	if l := len(path); l >= 5 && isWinSlash(path[0]) && isWinSlash(path[1]) &&
		!isWinSlash(path[2]) && path[2] != '.' {
		// first, leading `\\` and next shouldn't be `\`. its server name.
		for n := 3; n < l-1; n++ {
			// second, next '\' shouldn't be repeated.
			if isWinSlash(path[n]) {
				n++
				// third, following something characters. its share name.
				if !isWinSlash(path[n]) {
					if path[n] == '.' {
						break
					}
					for ; n < l; n++ {
						if isWinSlash(path[n]) {
							break
						}
					}
					return n
				}
				break
			}
		}
	}
	return 0
}
