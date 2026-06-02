package slicex

// MergeUint merges multiple slices into one.
// MergeUint 合并多个切片为一个。
func MergeUint(slices ...[]uint) []uint {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertUint inserts elements at the specified position.
// InsertUint 按位置插入元素。
func InsertUint(slice []uint, pos int, target ...uint) []uint {
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
	rs := make([]uint, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadUint inserts elements at the head.
// InsertHeadUint 在头部插入元素。
func InsertHeadUint(slice []uint, target ...uint) []uint {
	return InsertUint(slice, 0, target...)
}

// InsertTailUint inserts elements at the tail.
// InsertTailUint 在尾部插入元素。
func InsertTailUint(slice []uint, target ...uint) []uint {
	return InsertUint(slice, len(slice), target...)
}

// RemoveValueUint removes the first element equal to target.
// RemoveValueUint 按值删除第一个匹配元素。
func RemoveValueUint(slice []uint, target uint) ([]uint, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueUint removes all elements equal to target.
// RemoveAllValueUint 按值删除所有匹配元素。
func RemoveAllValueUint(slice []uint, target uint) ([]uint, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtUint removes the element at the specified index.
// RemoveAtUint 按索引删除元素。
func RemoveAtUint(slice []uint, pos int) ([]uint, uint, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadUint removes the first element.
// RemoveHeadUint 删除头部元素。
func RemoveHeadUint(slice []uint) ([]uint, uint, bool) {
	return RemoveAtUint(slice, 0)
}

// RemoveTailUint removes the last element.
// RemoveTailUint 删除尾部元素。
func RemoveTailUint(slice []uint) ([]uint, uint, bool) {
	return RemoveAtUint(slice, len(slice)-1)
}

// RemoveFromUint removes length elements starting at startPos.
// RemoveFromUint 从 startPos 起删除 length 个元素。
func RemoveFromUint(slice []uint, startPos int, length int) (result []uint, removed []uint, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint(slice, startPos, endPos)
}

// RemoveRangeUint removes elements in [startPos, endPos).
// RemoveRangeUint 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeUint(slice []uint, startPos int, endPos int) (result []uint, removed []uint, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsUint reports whether slice contains target.
// ContainsUint 判断切片是否包含目标值。
func ContainsUint(slice []uint, target uint) bool {
	_, ok := IndexUint(slice, target)
	return ok
}

// IndexUint returns the index of the first matching element.
// IndexUint 从头部查找第一个匹配元素的索引。
func IndexUint(slice []uint, target uint) (int, bool) {
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

// LastIndexUint returns the index of the last matching element.
// LastIndexUint 从尾部查找最后一个匹配元素的索引。
func LastIndexUint(slice []uint, target uint) (int, bool) {
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

// ReverseUint reverses elements in place.
// ReverseUint 原地倒序元素。
func ReverseUint(slice []uint) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualUint reports whether two slices are equal element-wise.
// EqualUint 逐元素比较两个切片是否相等。
func EqualUint(a, b []uint) bool {
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

// SumUint returns the sum of all elements.
// SumUint 求和所有元素。
func SumUint(slice []uint) uint {
	rs := uint(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
