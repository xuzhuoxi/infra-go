package mathx

import "fmt"

// BitString returns the binary string representation of value.
// BitString 以二进制字符串形式输出位信息。
func BitString(value int) string {
	return fmt.Sprintf("%b", value)
}

// BitValid reports whether the bit at bitIndex is set (1 = valid, 0 = invalid).
// BitValid 判断指定位是否有效，1 为有效，0 为无效。
// bitIndex counts from the least significant bit; the first bit is index 0.
// bitIndex 从低位开始计数，第 1 位为 0。
func BitValid(value int, bitIndex uint) bool {
	return 0 != ((1 << bitIndex) & value)
}

// BitFit checks for overlapping bits between value and checkingValue.
// BitFit 检查位值是否重复。
// When isValid is true, checks for overlapping valid (set) bits; when false, checks invalid bits.
// isValid 为 true 时检查有效位的重复；为 false 时检查无效位的重复。
func BitFit(value int, checkingValue int, isValid bool) bool {
	if isValid {
		return value&checkingValue != 0
	} else {
		return ^(value & checkingValue) != 0
	}
}

// BitValidAnd reports whether all listed bits are set in value.
// BitValidAnd 判断所列位是否全部有效。
func BitValidAnd(value int, bitIndex ...uint) bool {
	for _, index := range bitIndex {
		if !BitValid(value, index) {
			return false
		}
	}
	return true
}

// BitValidOr reports whether any listed bit is set in value.
// BitValidOr 判断所列位是否至少有一位有效。
func BitValidOr(value int, bitIndex ...uint) bool {
	for _, index := range bitIndex {
		if BitValid(value, index) {
			return true
		}
	}
	return false
}

// BitSet sets or clears the bit at bitIndex and returns the updated value.
// BitSet 设置或清除指定位并返回更新后的值。
func BitSet(value int, bitIndex uint, isValid bool) int {
	v := 1 << bitIndex
	if isValid {
		return value | v
	} else {
		return value & ^v
	}
}
