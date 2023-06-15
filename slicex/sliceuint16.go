package slicex

// MergeUint16
// 合并
func MergeUint16(slices ...[]uint16) []uint16 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint16, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertUint16
// 按位置插入
func InsertUint16(slice []uint16, pos int, target ...uint16) []uint16 {
	ln := len(slice)
	if pos < 0 {
		pos = 0
		goto Start
	}
	if pos > ln {
		pos = ln
	}
Start:
	tl := len(target)
	rs := make([]uint16, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadUint16
// 头插入
func InsertHeadUint16(slice []uint16, target ...uint16) []uint16 {
	return InsertUint16(slice, 0, target...)
}

// InsertTailUint16
// 尾插入
func InsertTailUint16(slice []uint16, target ...uint16) []uint16 {
	return InsertUint16(slice, len(slice), target...)
}

// RemoveValueUint16
// 按值删除
func RemoveValueUint16(slice []uint16, target uint16) ([]uint16, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint16(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint16(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueUint16
// 按值删除
func RemoveAllValueUint16(slice []uint16, target uint16) ([]uint16, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint16, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtUint16
// 按点删除
func RemoveAtUint16(slice []uint16, pos int) ([]uint16, uint16, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint16, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadUint16
// 删除头
func RemoveHeadUint16(slice []uint16) ([]uint16, uint16, bool) {
	return RemoveAtUint16(slice, 0)
}

// RemoveTailUint16
// 删除尾
func RemoveTailUint16(slice []uint16) ([]uint16, uint16, bool) {
	return RemoveAtUint16(slice, len(slice)-1)
}

// RemoveFromUint16
// 删除区间
func RemoveFromUint16(slice []uint16, startPos int, length int) (result []uint16, removed []uint16, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint16(slice, startPos, endPos)
}

// RemoveRangeUint16
// 删除区间
func RemoveRangeUint16(slice []uint16, startPos int, endPos int) (result []uint16, removed []uint16, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint16, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsUint16
// 包含
func ContainsUint16(slice []uint16, target uint16) bool {
	_, ok := IndexUint16(slice, target)
	return ok
}

// IndexUint16
// 从头部查找
func IndexUint16(slice []uint16, target uint16) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	for index, value := range slice {
		if value == target {
			return index, true
		}
	}
	return -1, false
}

// LastIndexUint16
// 从尾部查找
func LastIndexUint16(slice []uint16, target uint16) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	for index := len(slice) - 1; index >= 0; index-- {
		if slice[index] == target {
			return index, true
		}
	}
	return -1, false
}

// ReverseUint16
// 倒序
func ReverseUint16(slice []uint16) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualUint16
// 比较
func EqualUint16(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}
	for index, val := range a {
		if val != b[index] {
			return false
		}
	}
	return true
}

// SumUint16
// 求和
func SumUint16(slice []uint16) uint16 {
	rs := uint16(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
