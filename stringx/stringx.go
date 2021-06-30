package stringx

import (
	"strings"
	"unicode/utf8"
)

//取字符串的字符个数
func GetCharCount(s string) int {
	if "" == s {
		return 0
	}
	return utf8.RuneCountInString(s)
	//return len([]rune(s))
}

//取单个字符的首位置
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

//取单个字符的尾位置
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

//截取子字符串
func SubStr(s string, startIndex, length int) string {
	return SubString(s, startIndex, startIndex+length)
}

//截取子字符串
func SubString(s string, startIndex, endIndex int) string {
	runes := []rune(s)
	if endIndex > len(runes) {
		endIndex = len(runes)
	}
	return string(runes[startIndex:endIndex])
}

//截取前部分
func SubPrefix(s string, index int) string {
	prefix, _ := CutString(s, index, true)
	return prefix
}

//截取后部分
func SubSuffix(s string, index int) string {
	_, suffix := CutString(s, index, true)
	return suffix
}

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
