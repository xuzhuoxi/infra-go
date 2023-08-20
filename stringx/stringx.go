package stringx

import (
	"fmt"
	"strings"
	"sync"
	"unicode/utf8"
)

var (
	builder     = &strings.Builder{}
	builderLock sync.RWMutex
)

func init() {
	builder.Grow(1024)
}

func printX(sb *strings.Builder, args ...interface{}) string {
	if len(args) == 0 {
		return ""
	}
	for index := range args {
		if args[index] == nil {
			sb.WriteString("[nil]")
		}
		if str, ok := args[index].(string); ok {
			sb.WriteString(str)
			continue
		}
		if sg, ok := args[index].(fmt.Stringer); ok {
			sb.WriteString(sg.String())
			continue
		}
		fmt.Fprintf(sb, "%v", args[index])
		fmt.Print()
	}
	return sb.String()
}

// Print 字符串拼接
func Print(args ...interface{}) string {
	builder.Reset()
	return printX(builder, args...)
}

// PrintSafe 字符串拼接
func PrintSafe(args ...interface{}) string {
	builderLock.Lock()
	defer builderLock.Unlock()
	builder.Reset()
	return printX(builder, args...)
}

// GetRuneCount
// 取字符串的字符个数
func GetRuneCount(str string) int {
	if "" == str {
		return 0
	}
	return utf8.RuneCountInString(str)
}

// IndexOfString
// 取单个字符的首位置
func IndexOfString(s string, char string) int {
	if "" == char || "" == s {
		return -1
	}
	sRune := []rune(s)
	charRune := []rune(char)
	if len(sRune) == 0 || len(charRune) != 1 {
		return -1
	}
	for index, rune := range sRune {
		if rune == charRune[0] {
			return index
		}
	}
	return -1
}

// LastIndexOfString
// 取单个字符的尾位置
func LastIndexOfString(s string, char string) int {
	if "" == char || "" == s {
		return -1
	}
	sRune := []rune(s)
	charRune := []rune(char)
	if len(sRune) == 0 || len(charRune) != 1 {
		return -1
	}
	for index := len(sRune) - 1; index >= 0; index-- {
		if sRune[index] == charRune[0] {
			return index
		}
	}
	return -1
}

// SubStr
// 截取子字符串
func SubStr(s string, startIndex, length int) string {
	return SubString(s, startIndex, startIndex+length)
}

// SubString
// 截取子字符串
func SubString(s string, startIndex, endIndex int) string {
	runes := []rune(s)
	if endIndex > len(runes) {
		endIndex = len(runes)
	}
	return string(runes[startIndex:endIndex])
}

// SubPrefix
// 截取前部分
func SubPrefix(s string, index int) string {
	prefix, _ := CutString(s, index, true)
	return prefix
}

// SubSuffix
// 截取后部分
func SubSuffix(s string, index int) string {
	_, suffix := CutString(s, index, true)
	return suffix
}

// CutString
// 把字符串一分为二
// 当keepIndex=true是，index所在字符会留在第二个返回字符的第一个
func CutString(s string, runeIndex int, keepIndex bool) (string, string) {
	runes := []rune(s)
	if runeIndex < 0 {
		return "", s
	}
	if runeIndex >= len(runes) {
		return s, ""
	}
	if keepIndex {
		return string(runes[:runeIndex]), string(runes[runeIndex:])
	} else {
		return string(runes[:runeIndex]), string(runes[runeIndex+1:])
	}
}

func SplitToken(str string, separator string, trimTokens bool, ignoreEmptyTokens bool) []string {
	if len(str) == 0 {
		return nil
	}
	arr := strings.Split(str, separator)
	if len(arr) == 0 {
		return nil
	}
	var rs []string = nil
	for index := range arr {
		if trimTokens {
			arr[index] = strings.TrimSpace(arr[index])
		}
		if !ignoreEmptyTokens || len(arr[index]) > 0 {
			rs = append(rs, arr[index])
		}
	}
	return rs
}

func HasSuffixAt(str string, prefix string, tOffset int) bool {
	if tOffset >= len(str) || tOffset < 0 {
		return false
	}
	checkStr := str[tOffset:]
	return strings.HasSuffix(checkStr, prefix)
}
