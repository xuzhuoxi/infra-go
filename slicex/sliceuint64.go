package slicex

// MergeUint64 merges multiple slices into one.
// MergeUint64 合并多个切片为一个。
func MergeUint64(slices ...[]uint64) []uint64 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint64, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertUint64 inserts elements at the specified position.
// InsertUint64 按位置插入元素。
func InsertUint64(slice []uint64, pos int, target ...uint64) []uint64 {
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
	rs := make([]uint64, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadUint64 inserts elements at the head.
// InsertHeadUint64 在头部插入元素。
func InsertHeadUint64(slice []uint64, target ...uint64) []uint64 {
	return InsertUint64(slice, 0, target...)
}

// InsertTailUint64 inserts elements at the tail.
// InsertTailUint64 在尾部插入元素。
func InsertTailUint64(slice []uint64, target ...uint64) []uint64 {
	return InsertUint64(slice, len(slice), target...)
}

// RemoveValueUint64 removes the first element equal to target.
// RemoveValueUint64 按值删除第一个匹配元素。
func RemoveValueUint64(slice []uint64, target uint64) ([]uint64, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint64(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint64(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueUint64 removes all elements equal to target.
// RemoveAllValueUint64 按值删除所有匹配元素。
func RemoveAllValueUint64(slice []uint64, target uint64) ([]uint64, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint64, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtUint64 removes the element at the specified index.
// RemoveAtUint64 按索引删除元素。
func RemoveAtUint64(slice []uint64, pos int) ([]uint64, uint64, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint64, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadUint64 removes the first element.
// RemoveHeadUint64 删除头部元素。
func RemoveHeadUint64(slice []uint64) ([]uint64, uint64, bool) {
	return RemoveAtUint64(slice, 0)
}

// RemoveTailUint64 removes the last element.
// RemoveTailUint64 删除尾部元素。
func RemoveTailUint64(slice []uint64) ([]uint64, uint64, bool) {
	return RemoveAtUint64(slice, len(slice)-1)
}

// RemoveFromUint64 removes length elements starting at startPos.
// RemoveFromUint64 从 startPos 起删除 length 个元素。
func RemoveFromUint64(slice []uint64, startPos int, length int) (result []uint64, removed []uint64, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint64(slice, startPos, endPos)
}

// RemoveRangeUint64 removes elements in [startPos, endPos).
// RemoveRangeUint64 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeUint64(slice []uint64, startPos int, endPos int) (result []uint64, removed []uint64, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint64, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsUint64 reports whether slice contains target.
// ContainsUint64 判断切片是否包含目标值。
func ContainsUint64(slice []uint64, target uint64) bool {
	_, ok := IndexUint64(slice, target)
	return ok
}

// IndexUint64 returns the index of the first matching element.
// IndexUint64 从头部查找第一个匹配元素的索引。
func IndexUint64(slice []uint64, target uint64) (int, bool) {
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

// LastIndexUint64 returns the index of the last matching element.
// LastIndexUint64 从尾部查找最后一个匹配元素的索引。
func LastIndexUint64(slice []uint64, target uint64) (int, bool) {
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

// ReverseUint64 reverses elements in place.
// ReverseUint64 原地倒序元素。
func ReverseUint64(slice []uint64) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualUint64 reports whether two slices are equal element-wise.
// EqualUint64 逐元素比较两个切片是否相等。
func EqualUint64(a, b []uint64) bool {
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

// UintToUint64 converts []uint to []uint64.
// UintToUint64 将 []uint 转换为 []uint64。
func UintToUint64(source []uint) []uint64 {
	if nil == source {
		return nil
	}
	if len(source) == 0 {
		return []uint64{}
	}
	var rs []uint64
	for _, val := range source {
		rs = append(rs, uint64(val))
	}
	return rs
}

// Uint64ToUint converts []uint64 to []uint.
// Uint64ToUint 将 []uint64 转换为 []uint。
func Uint64ToUint(source []uint64) []uint {
	if nil == source {
		return nil
	}
	if len(source) == 0 {
		return []uint{}
	}
	var rs []uint
	for _, val := range source {
		rs = append(rs, uint(val))
	}
	return rs
}

// SumUint64 returns the sum of all elements.
// SumUint64 求和所有元素。
func SumUint64(slice []uint64) uint64 {
	rs := uint64(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
