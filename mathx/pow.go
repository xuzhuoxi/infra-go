// Package mathx
// Create on 2023/7/23
// @author xuzhuoxi
package mathx

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
