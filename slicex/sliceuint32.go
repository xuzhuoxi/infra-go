package slicex

// MergeUint32 merges multiple slices into one.
// MergeUint32 合并多个切片为一个。
func MergeUint32(slices ...[]uint32) []uint32 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint32, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertUint32 inserts elements at the specified position.
// InsertUint32 按位置插入元素。
func InsertUint32(slice []uint32, pos int, target ...uint32) []uint32 {
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
	rs := make([]uint32, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadUint32 inserts elements at the head.
// InsertHeadUint32 在头部插入元素。
func InsertHeadUint32(slice []uint32, target ...uint32) []uint32 {
	return InsertUint32(slice, 0, target...)
}

// InsertTailUint32 inserts elements at the tail.
// InsertTailUint32 在尾部插入元素。
func InsertTailUint32(slice []uint32, target ...uint32) []uint32 {
	return InsertUint32(slice, len(slice), target...)
}

// RemoveValueUint32 removes the first element equal to target.
// RemoveValueUint32 按值删除第一个匹配元素。
func RemoveValueUint32(slice []uint32, target uint32) ([]uint32, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint32(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint32(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueUint32 removes all elements equal to target.
// RemoveAllValueUint32 按值删除所有匹配元素。
func RemoveAllValueUint32(slice []uint32, target uint32) ([]uint32, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint32, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtUint32 removes the element at the specified index.
// RemoveAtUint32 按索引删除元素。
func RemoveAtUint32(slice []uint32, pos int) ([]uint32, uint32, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint32, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadUint32 removes the first element.
// RemoveHeadUint32 删除头部元素。
func RemoveHeadUint32(slice []uint32) ([]uint32, uint32, bool) {
	return RemoveAtUint32(slice, 0)
}

// RemoveTailUint32 removes the last element.
// RemoveTailUint32 删除尾部元素。
func RemoveTailUint32(slice []uint32) ([]uint32, uint32, bool) {
	return RemoveAtUint32(slice, len(slice)-1)
}

// RemoveFromUint32 removes length elements starting at startPos.
// RemoveFromUint32 从 startPos 起删除 length 个元素。
func RemoveFromUint32(slice []uint32, startPos int, length int) (result []uint32, removed []uint32, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint32(slice, startPos, endPos)
}

// RemoveRangeUint32 removes elements in [startPos, endPos).
// RemoveRangeUint32 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeUint32(slice []uint32, startPos int, endPos int) (result []uint32, removed []uint32, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint32, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsUint32 reports whether slice contains target.
// ContainsUint32 判断切片是否包含目标值。
func ContainsUint32(slice []uint32, target uint32) bool {
	_, ok := IndexUint32(slice, target)
	return ok
}

// IndexUint32 returns the index of the first matching element.
// IndexUint32 从头部查找第一个匹配元素的索引。
func IndexUint32(slice []uint32, target uint32) (int, bool) {
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

// LastIndexUint32 returns the index of the last matching element.
// LastIndexUint32 从尾部查找最后一个匹配元素的索引。
func LastIndexUint32(slice []uint32, target uint32) (int, bool) {
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

// ReverseUint32 reverses elements in place.
// ReverseUint32 原地倒序元素。
func ReverseUint32(slice []uint32) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualUint32 reports whether two slices are equal element-wise.
// EqualUint32 逐元素比较两个切片是否相等。
func EqualUint32(a, b []uint32) bool {
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

// SumUint32 returns the sum of all elements.
// SumUint32 求和所有元素。
func SumUint32(slice []uint32) uint32 {
	rs := uint32(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
