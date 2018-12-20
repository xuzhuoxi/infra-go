package stringsx

//取字符串的字符个数
func GetCharCount(s string) int {
	if "" == s {
		return 0
	}
	return len([]rune(s))
}

//取单个字符的首位置
func IndexOfChar(s string, char string) int {
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
func LastIndexOfChar(s string, char string) int {
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

//把字符串一分为二
func CutString(s string, index int, keepIndex bool) (string, string) {
	runes := []rune(s)
	if index < 0 {
		return "", s
	}
	if index >= len(runes) {
		return s, ""
	}
	if keepIndex {
		return string(runes[:index]), string(runes[index:])
	} else {
		return string(runes[:index]), string(runes[index+1:])
	}
}
