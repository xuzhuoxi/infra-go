package mathx

import "fmt"

// BitString
// 二进制输出位信息
func BitString(value int) string {
	return fmt.Sprintf("%b", value)
}

// BitValid
// <summary>
/// 是否有效，1为效，0为无效
/// </summary>
/// <returns><c>true</c>, if valid was ised, <c>false</c> otherwise.</returns>
/// <param name="value">Value.</param>
/// <param name="bitIndex">从低位开始，第1位为0</param>
func BitValid(value int, bitIndex uint) bool {
	return 0 != ((1 << bitIndex) & value)
}

// BitFit
// 检查位值重复
// isValid=true时，检查有效的重复位
// isValid=false时，检查无效的重复位
func BitFit(value int, checkingValue int, isValid bool) bool {
	if isValid {
		return value&checkingValue != 0
	} else {
		return ^(value & checkingValue) != 0
	}
}

// BitValidAnd
// 全部有效
func BitValidAnd(value int, bitIndex ...uint) bool {
	for _, index := range bitIndex {
		if !BitValid(value, index) {
			return false
		}
	}
	return true
}

// BitValidOr
// 单个有效
func BitValidOr(value int, bitIndex ...uint) bool {
	for _, index := range bitIndex {
		if BitValid(value, index) {
			return true
		}
	}
	return false
}

// BitSet
// 设置位
func BitSet(value int, bitIndex uint, isValid bool) int {
	v := 1 << bitIndex
	if isValid {
		return value | v
	} else {
		return value & ^v
	}
}
