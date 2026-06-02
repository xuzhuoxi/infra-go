package slicex

// MergeInt8 merges multiple slices into one.
// MergeInt8 合并多个切片为一个。
func MergeInt8(slices ...[]int8) []int8 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int8, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertInt8 inserts elements at the specified position.
// InsertInt8 按位置插入元素。
func InsertInt8(slice []int8, pos int, target ...int8) []int8 {
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
	rs := make([]int8, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadInt8 inserts elements at the head.
// InsertHeadInt8 在头部插入元素。
func InsertHeadInt8(slice []int8, target ...int8) []int8 {
	return InsertInt8(slice, 0, target...)
}

// InsertTailInt8 inserts elements at the tail.
// InsertTailInt8 在尾部插入元素。
func InsertTailInt8(slice []int8, target ...int8) []int8 {
	return InsertInt8(slice, len(slice), target...)
}

// RemoveValueInt8 removes the first element equal to target.
// RemoveValueInt8 按值删除第一个匹配元素。
func RemoveValueInt8(slice []int8, target int8) ([]int8, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt8(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt8(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueInt8 removes all elements equal to target.
// RemoveAllValueInt8 按值删除所有匹配元素。
func RemoveAllValueInt8(slice []int8, target int8) ([]int8, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int8, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtInt8 removes the element at the specified index.
// RemoveAtInt8 按索引删除元素。
func RemoveAtInt8(slice []int8, pos int) ([]int8, int8, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int8, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadInt8 removes the first element.
// RemoveHeadInt8 删除头部元素。
func RemoveHeadInt8(slice []int8) ([]int8, int8, bool) {
	return RemoveAtInt8(slice, 0)
}

// RemoveTailInt8 removes the last element.
// RemoveTailInt8 删除尾部元素。
func RemoveTailInt8(slice []int8) ([]int8, int8, bool) {
	return RemoveAtInt8(slice, len(slice)-1)
}

// RemoveFromInt8 removes length elements starting at startPos.
// RemoveFromInt8 从 startPos 起删除 length 个元素。
func RemoveFromInt8(slice []int8, startPos int, length int) (result []int8, removed []int8, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt8(slice, startPos, endPos)
}

// RemoveRangeInt8 removes elements in [startPos, endPos).
// RemoveRangeInt8 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeInt8(slice []int8, startPos int, endPos int) (result []int8, removed []int8, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int8, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsInt8 reports whether slice contains target.
// ContainsInt8 判断切片是否包含目标值。
func ContainsInt8(slice []int8, target int8) bool {
	_, ok := IndexInt8(slice, target)
	return ok
}

// IndexInt8 returns the index of the first matching element.
// IndexInt8 从头部查找第一个匹配元素的索引。
func IndexInt8(slice []int8, target int8) (int, bool) {
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

// LastIndexInt8 returns the index of the last matching element.
// LastIndexInt8 从尾部查找最后一个匹配元素的索引。
func LastIndexInt8(slice []int8, target int8) (int, bool) {
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

// ReverseInt8 reverses elements in place.
// ReverseInt8 原地倒序元素。
func ReverseInt8(slice []int8) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualInt8 reports whether two slices are equal element-wise.
// EqualInt8 逐元素比较两个切片是否相等。
func EqualInt8(a, b []int8) bool {
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

// SumInt8 returns the sum of all elements.
// SumInt8 求和所有元素。
func SumInt8(slice []int8) int8 {
	rs := int8(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
