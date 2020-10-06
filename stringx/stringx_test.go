package stringx

import (
	"fmt"
	"strings"
	"testing"
)

func TestCutString(t *testing.T) {
	str := "aaa顶你个肺bbb"
	c := GetCharCount(str)
	fmt.Println("长度：", c)
	fmt.Println(CutString(str, 0, true))
	fmt.Println(CutString(str, 0, false))
	fmt.Println(CutString(str, 1, true))
	fmt.Println(CutString(str, 1, false))
	fmt.Println()
	fmt.Println(CutString(str, c, true))
	fmt.Println(CutString(str, c, false))
	fmt.Println(CutString(str, c-1, true))
	fmt.Println(CutString(str, c-1, false))
	fmt.Println()
	fmt.Println(CutString(str, 6, true))
	fmt.Println(CutString(str, 6, false))
	fmt.Println(CutString(str, 9, true))
	fmt.Println(CutString(str, 9, false))
}

func TestSub(t *testing.T) {
	str := "aaa顶你个肺bbb"
	fmt.Println(SubPrefix(str, -1))
	fmt.Println(SubPrefix(str, 0))
	fmt.Println(SubPrefix(str, 1))
	fmt.Println(SubPrefix(str, 5))
	fmt.Println(SubPrefix(str, 8))
	fmt.Println(SubPrefix(str, 9))
	fmt.Println(SubPrefix(str, 10))
}

func TestIndex(t *testing.T) {
	str := "ada顶你个肺bdb"
	fmt.Println(IndexOfString(str, "d"))
	fmt.Println(LastIndexOfString(str, "d"))
}

func TestOther(t *testing.T) {
	str := "ada顶你个肺bdb"
	fmt.Println(strings.LastIndex(str, "d"))
	fmt.Println(strings.Index(str, "d"))
	fmt.Println(strings.LastIndex(str, "你"))
	fmt.Println(strings.Index(str, "肺"))
}
