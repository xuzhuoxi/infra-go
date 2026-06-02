package slicex

// MergeInt16 merges multiple slices into one.
// MergeInt16 合并多个切片为一个。
func MergeInt16(slices ...[]int16) []int16 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int16, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertInt16 inserts elements at the specified position.
// InsertInt16 按位置插入元素。
func InsertInt16(slice []int16, pos int, target ...int16) []int16 {
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
	rs := make([]int16, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadInt16 inserts elements at the head.
// InsertHeadInt16 在头部插入元素。
func InsertHeadInt16(slice []int16, target ...int16) []int16 {
	return InsertInt16(slice, 0, target...)
}

// InsertTailInt16 inserts elements at the tail.
// InsertTailInt16 在尾部插入元素。
func InsertTailInt16(slice []int16, target ...int16) []int16 {
	return InsertInt16(slice, len(slice), target...)
}

// RemoveValueInt16 removes the first element equal to target.
// RemoveValueInt16 按值删除第一个匹配元素。
func RemoveValueInt16(slice []int16, target int16) ([]int16, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt16(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt16(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueInt16 removes all elements equal to target.
// RemoveAllValueInt16 按值删除所有匹配元素。
func RemoveAllValueInt16(slice []int16, target int16) ([]int16, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int16, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtInt16 removes the element at the specified index.
// RemoveAtInt16 按索引删除元素。
func RemoveAtInt16(slice []int16, pos int) ([]int16, int16, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int16, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadInt16 removes the first element.
// RemoveHeadInt16 删除头部元素。
func RemoveHeadInt16(slice []int16) ([]int16, int16, bool) {
	return RemoveAtInt16(slice, 0)
}

// RemoveTailInt1 removes the last element.
// RemoveTailInt1 删除尾部元素。
func RemoveTailInt1(slice []int16) ([]int16, int16, bool) {
	return RemoveAtInt16(slice, len(slice)-1)
}

// RemoveFromInt16 removes length elements starting at startPos.
// RemoveFromInt16 从 startPos 起删除 length 个元素。
func RemoveFromInt16(slice []int16, startPos int, length int) (result []int16, removed []int16, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt16(slice, startPos, endPos)
}

// RemoveRangeInt16 removes elements in [startPos, endPos).
// RemoveRangeInt16 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeInt16(slice []int16, startPos int, endPos int) (result []int16, removed []int16, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int16, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsInt16 reports whether slice contains target.
// ContainsInt16 判断切片是否包含目标值。
func ContainsInt16(slice []int16, target int16) bool {
	_, ok := IndexInt16(slice, target)
	return ok
}

// IndexInt16 returns the index of the first matching element.
// IndexInt16 从头部查找第一个匹配元素的索引。
func IndexInt16(slice []int16, target int16) (int, bool) {
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

// LastIndexInt16 returns the index of the last matching element.
// LastIndexInt16 从尾部查找最后一个匹配元素的索引。
func LastIndexInt16(slice []int16, target int16) (int, bool) {
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

// ReverseInt16 reverses elements in place.
// ReverseInt16 原地倒序元素。
func ReverseInt16(slice []int16) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualInt16 reports whether two slices are equal element-wise.
// EqualInt16 逐元素比较两个切片是否相等。
func EqualInt16(a, b []int16) bool {
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

// SumInt16 returns the sum of all elements.
// SumInt16 求和所有元素。
func SumInt16(slice []int16) int16 {
	rs := int16(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
