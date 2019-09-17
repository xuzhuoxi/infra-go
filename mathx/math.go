package mathx

import (
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

//十进制转换不确定进制
//value:十进制数值
//system:不确定进制数组
//return;由十进制数据组成的数组
func SystemTo(value int, system []int) []int {
	rs := make([]int, len(system))
	temp := value
	for i, sys := range system {
		rs[i] = temp % sys
		temp = temp / sys
	}
	return rs
}

//十进制数 转为 26进制字符表示
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

//26进制字符 转为 十进制数
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
			return 0
		}
		n += int(char-64) * j
		j *= 26
	}
	return n
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
