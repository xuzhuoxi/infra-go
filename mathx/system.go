// Package mathx
// Create on 2023/7/23
// @author xuzhuoxi
package mathx

import "strings"

// SystemTo
// 十进制转换不确定进制
// value:十进制数值
// system:不确定进制数组
// return;由十进制数据组成的数组
func SystemTo(value int, system []int) []int {
	rs := make([]int, len(system))
	temp := value
	for i, sys := range system {
		rs[i] = temp % sys
		temp = temp / sys
	}
	return rs
}

// System10To26
// 十进制数 转为 26进制字符表示
func System10To26(n int) string {
	if 0 == n {
		return ""
	}
	s := ""
	for n > 0 {
		m := n % 26
		if m == 0 {
			m = 26
		}
		s = string(rune(int32(m)+64)) + s
		n = (n - m) / 26
	}
	return s
}

// System26To10
// 26进制字符 转为 十进制数
func System26To10(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	str := []rune(strings.ToUpper(s))
	l := len(str)
	j := 1
	for index := l - 1; index >= 0; index-- {
		char := str[index]
		if char < 'A' || char > 'Z' {
			return -1
		}
		n += int(char-64) * j
		j *= 26
	}
	return n
}

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// System10To36
// 十进制数 转为 36进制字符表示
func System10To36(n int) string {
	// 如果n为0，直接返回"0"
	if n == 0 {
		return "0"
	}
	// 定义一个字符串切片，用来存放结果
	var res []string
	// 定义一个变量，用来存放余数
	var r int
	// 循环除以36，直到商为0
	for n > 0 {
		// 计算余数
		r = n % 36
		// 将余数对应的字符添加到切片的头部
		res = append([]string{string(base36[r])}, res...)
		// 更新商
		n = n / 36
	}
	// 将切片拼接为字符串并返回
	return strings.Join(res, "")
}

// System36To10
// 36进制字符 转为 十进制数
func System36To10(s string) int {
	// 如果s为空，直接返回0
	if s == "" {
		return 0
	}
	// 将s转换为大写
	s = strings.ToUpper(s)
	// 定义一个变量，用来存放结果
	var res int
	// 定义一个变量，用来存放幂指数
	var p int
	// 从右往左遍历s中的每个字符
	for i := len(s) - 1; i >= 0; i-- {
		// 获取当前字符在base36中的索引，即对应的十进制数
		n := strings.Index(base36, string(s[i]))
		// 如果索引为-1，说明s中有非法字符，直接返回0
		if n == -1 {
			return 0
		}
		// 将当前字符对应的十进制数乘以36的p次方，并累加到结果中
		res += n * PowInt(36, p)
		// 更新幂指数
		p++
	}
	// 返回结果
	return res
}
