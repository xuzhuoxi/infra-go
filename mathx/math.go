package mathx

import (
	"strconv"
	"strings"
)

type SizeUnit uint64

const (
	Byte SizeUnit = 1
	KB            = 1024 * Byte
	MB            = 1024 * KB
	GB            = 1024 * MB
	TB            = 1024 * GB
	PB            = 1024 * TB
	EB            = 1024 * PB
)

func ParseSize(sizeStr string) SizeUnit {
	sizeStr = strings.ToUpper(strings.TrimSpace(sizeStr))
	if len(sizeStr) == 0 {
		return 0
	}
	var size SizeUnit
	var err error
	if strings.HasSuffix(sizeStr, "KB") {
		size, err = parseSize(sizeStr, KB)
	} else if strings.HasSuffix(sizeStr, "MB") {
		size, err = parseSize(sizeStr, MB)
	} else if strings.HasSuffix(sizeStr, "GB") {
		size, err = parseSize(sizeStr, GB)
	} else if strings.HasSuffix(sizeStr, "TB") {
		size, err = parseSize(sizeStr, TB)
	} else if strings.HasSuffix(sizeStr, "PB") {
		size, err = parseSize(sizeStr, PB)
	} else if strings.HasSuffix(sizeStr, "EB") {
		size, err = parseSize(sizeStr, EB)
	} else {
		value, err := strconv.ParseFloat(strings.TrimSpace(sizeStr), 64)
		if err != nil {
			return 0
		}
		return SizeUnit(value)
	}
	if err != nil {
		return 0
	}
	return size
}

func parseSize(sizeStr string, unit SizeUnit) (size SizeUnit, err error) {
	sizeStr = sizeStr[:len(sizeStr)-2]
	sizeStr = strings.TrimSpace(sizeStr)
	value, err1 := strconv.ParseFloat(sizeStr, 64)
	if nil != err1 {
		return 0, err1
	}
	return SizeUnit(value * float64(unit)), nil
}

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

// PowInt
// 定义一个函数，计算x的y次方
func PowInt(x, y int) int {
	// 如果y为0，直接返回1
	if y == 0 {
		return 1
	}
	// 定义一个变量，用来存放结果，初始值为x
	var res = x
	// 循环y-1次，每次将res乘以x
	for i := 1; i < y; i++ {
		res *= x
	}
	// 返回结果
	return res
}

func MinInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min3Int(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}
func Max3Int(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func MinUint(a, b uint) uint {
	if a > b {
		return b
	} else {
		return a
	}
}
func MaxUint(a, b uint) uint {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min3Uint(a, b, c uint) uint {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}
func Max3Uint(a, b, c uint) uint {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func MinFloat64(a, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min3Float64(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}
func Max3Float64(a, b, c float64) float64 {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}
