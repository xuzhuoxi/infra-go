package slicex

// MergeInt32 merges multiple slices into one.
// MergeInt32 合并多个切片为一个。
func MergeInt32(slices ...[]int32) []int32 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int32, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertInt32 inserts elements at the specified position.
// InsertInt32 按位置插入元素。
func InsertInt32(slice []int32, pos int, target ...int32) []int32 {
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
	rs := make([]int32, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadInt32 inserts elements at the head.
// InsertHeadInt32 在头部插入元素。
func InsertHeadInt32(slice []int32, target ...int32) []int32 {
	return InsertInt32(slice, 0, target...)
}

// InsertTailInt32 inserts elements at the tail.
// InsertTailInt32 在尾部插入元素。
func InsertTailInt32(slice []int32, target ...int32) []int32 {
	return InsertInt32(slice, len(slice), target...)
}

// RemoveValueInt32 removes the first element equal to target.
// RemoveValueInt32 按值删除第一个匹配元素。
func RemoveValueInt32(slice []int32, target int32) ([]int32, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt32(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt32(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueInt32 removes all elements equal to target.
// RemoveAllValueInt32 按值删除所有匹配元素。
func RemoveAllValueInt32(slice []int32, target int32) ([]int32, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int32, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtInt32 removes the element at the specified index.
// RemoveAtInt32 按索引删除元素。
func RemoveAtInt32(slice []int32, pos int) ([]int32, int32, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int32, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadInt32 removes the first element.
// RemoveHeadInt32 删除头部元素。
func RemoveHeadInt32(slice []int32) ([]int32, int32, bool) {
	return RemoveAtInt32(slice, 0)
}

// RemoveTailInt32 removes the last element.
// RemoveTailInt32 删除尾部元素。
func RemoveTailInt32(slice []int32) ([]int32, int32, bool) {
	return RemoveAtInt32(slice, len(slice)-1)
}

// RemoveFromInt32 removes length elements starting at startPos.
// RemoveFromInt32 从 startPos 起删除 length 个元素。
func RemoveFromInt32(slice []int32, startPos int, length int) (result []int32, removed []int32, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt32(slice, startPos, endPos)
}

// RemoveRangeInt32 removes elements in [startPos, endPos).
// RemoveRangeInt32 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeInt32(slice []int32, startPos int, endPos int) (result []int32, removed []int32, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int32, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsInt32 reports whether slice contains target.
// ContainsInt32 判断切片是否包含目标值。
func ContainsInt32(slice []int32, target int32) bool {
	_, ok := IndexInt32(slice, target)
	return ok
}

// IndexInt32 returns the index of the first matching element.
// IndexInt32 从头部查找第一个匹配元素的索引。
func IndexInt32(slice []int32, target int32) (int, bool) {
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

// LastIndexInt32 returns the index of the last matching element.
// LastIndexInt32 从尾部查找最后一个匹配元素的索引。
func LastIndexInt32(slice []int32, target int32) (int, bool) {
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

// ReverseInt32 reverses elements in place.
// ReverseInt32 原地倒序元素。
func ReverseInt32(slice []int32) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualInt32 reports whether two slices are equal element-wise.
// EqualInt32 逐元素比较两个切片是否相等。
func EqualInt32(a, b []int32) bool {
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

// SumInt32 returns the sum of all elements.
// SumInt32 求和所有元素。
func SumInt32(slice []int32) int32 {
	rs := int32(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
