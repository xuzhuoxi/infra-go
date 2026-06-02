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

// Print concatenates args into a string using a shared builder.
// Print 使用共享 builder 将参数拼接成字符串。
func Print(args ...interface{}) string {
	builder.Reset()
	return printX(builder, args...)
}

// PrintSafe concatenates args into a string with mutex-protected shared builder.
// PrintSafe 在互斥锁保护下使用共享 builder 将参数拼接成字符串。
func PrintSafe(args ...interface{}) string {
	builderLock.Lock()
	defer builderLock.Unlock()
	builder.Reset()
	return printX(builder, args...)
}

// GetRuneCount returns the number of Unicode code points in str.
// GetRuneCount 返回 str 中 Unicode 码点（字符）的个数。
func GetRuneCount(str string) int {
	if "" == str {
		return 0
	}
	return utf8.RuneCountInString(str)
}

// IndexOfString returns the first rune index of char in s, or -1 if not found.
// IndexOfString 返回 char 在 s 中首次出现的 rune 索引，未找到返回 -1。
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

// LastIndexOfString returns the last rune index of char in s, or -1 if not found.
// LastIndexOfString 返回 char 在 s 中最后一次出现的 rune 索引，未找到返回 -1。
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

// SubStr returns a substring starting at startIndex with the given rune length.
// SubStr 从 startIndex 起按 rune 长度截取子字符串。
func SubStr(s string, startIndex, length int) string {
	return SubString(s, startIndex, startIndex+length)
}

// SubString returns the substring from startIndex to endIndex in runes.
// SubString 按 rune 索引从 startIndex 到 endIndex 截取子字符串。
func SubString(s string, startIndex, endIndex int) string {
	runes := []rune(s)
	if endIndex > len(runes) {
		endIndex = len(runes)
	}
	return string(runes[startIndex:endIndex])
}

// SubPrefix returns the prefix of s before the rune at index.
// SubPrefix 返回 s 在 rune index 之前的前缀部分。
func SubPrefix(s string, index int) string {
	prefix, _ := CutString(s, index, true)
	return prefix
}

// SubSuffix returns the suffix of s from the rune at index onward.
// SubSuffix 返回 s 从 rune index 起的后缀部分。
func SubSuffix(s string, index int) string {
	_, suffix := CutString(s, index, true)
	return suffix
}

// CutString splits s into two parts at runeIndex.
// CutString 在 runeIndex 处把字符串一分为二。
// When keepIndex is true, the rune at index remains as the first rune of the second return value.
// 当 keepIndex=true 时，index 所在字符会留在第二个返回字符串的第一个位置。
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

// SplitToken splits str by separator with optional trimming and empty-token filtering.
// SplitToken 按分隔符拆分字符串，可选去除空白并忽略空 token。
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
