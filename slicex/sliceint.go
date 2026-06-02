package slicex

// MergeInt merges multiple slices into one.
// MergeInt 合并多个切片为一个。
func MergeInt(slices ...[]int) []int {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertInt inserts elements at the specified position.
// InsertInt 按位置插入元素。
func InsertInt(slice []int, pos int, target ...int) []int {
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
	rs := make([]int, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadInt inserts elements at the head.
// InsertHeadInt 在头部插入元素。
func InsertHeadInt(slice []int, target ...int) []int {
	return InsertInt(slice, 0, target...)
}

// InsertTailInt inserts elements at the tail.
// InsertTailInt 在尾部插入元素。
func InsertTailInt(slice []int, target ...int) []int {
	return InsertInt(slice, len(slice), target...)
}

// RemoveValueInt removes the first element equal to target.
// RemoveValueInt 按值删除第一个匹配元素。
func RemoveValueInt(slice []int, target int) ([]int, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueInt removes all elements equal to target.
// RemoveAllValueInt 按值删除所有匹配元素。
func RemoveAllValueInt(slice []int, target int) ([]int, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtInt removes the element at the specified index.
// RemoveAtInt 按索引删除元素。
func RemoveAtInt(slice []int, pos int) ([]int, int, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadInt removes the first element.
// RemoveHeadInt 删除头部元素。
func RemoveHeadInt(slice []int) ([]int, int, bool) {
	return RemoveAtInt(slice, 0)
}

// RemoveTailInt removes the last element.
// RemoveTailInt 删除尾部元素。
func RemoveTailInt(slice []int) ([]int, int, bool) {
	return RemoveAtInt(slice, len(slice)-1)
}

// RemoveFromInt removes length elements starting at startPos.
// RemoveFromInt 从 startPos 起删除 length 个元素。
func RemoveFromInt(slice []int, startPos int, length int) (result []int, removed []int, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt(slice, startPos, endPos)
}

// RemoveRangeInt removes elements in [startPos, endPos).
// RemoveRangeInt 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeInt(slice []int, startPos int, endPos int) (result []int, removed []int, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsInt reports whether slice contains target.
// ContainsInt 判断切片是否包含目标值。
func ContainsInt(slice []int, target int) bool {
	_, ok := IndexInt(slice, target)
	return ok
}

// IndexInt returns the index of the first matching element.
// IndexInt 从头部查找第一个匹配元素的索引。
func IndexInt(slice []int, target int) (int, bool) {
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

// LastIndexInt returns the index of the last matching element.
// LastIndexInt 从尾部查找最后一个匹配元素的索引。
func LastIndexInt(slice []int, target int) (int, bool) {
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

// ReverseInt reverses elements in place.
// ReverseInt 原地倒序元素。
func ReverseInt(slice []int) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualInt reports whether two slices are equal element-wise.
// EqualInt 逐元素比较两个切片是否相等。
func EqualInt(a, b []int) bool {
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

// SumInt returns the sum of all elements.
// SumInt 求和所有元素。
func SumInt(slice []int) int {
	rs := 0
	for _, val := range slice {
		rs += val
	}
	return rs
}
