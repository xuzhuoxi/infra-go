package slicex

// MergeInt64 merges multiple slices into one.
// MergeInt64 合并多个切片为一个。
func MergeInt64(slices ...[]int64) []int64 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int64, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertInt64 inserts elements at the specified position.
// InsertInt64 按位置插入元素。
func InsertInt64(slice []int64, pos int, target ...int64) []int64 {
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
	rs := make([]int64, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadInt64 inserts elements at the head.
// InsertHeadInt64 在头部插入元素。
func InsertHeadInt64(slice []int64, target ...int64) []int64 {
	return InsertInt64(slice, 0, target...)
}

// InsertTailInt64 inserts elements at the tail.
// InsertTailInt64 在尾部插入元素。
func InsertTailInt64(slice []int64, target ...int64) []int64 {
	return InsertInt64(slice, len(slice), target...)
}

// RemoveValueInt64 removes the first element equal to target.
// RemoveValueInt64 按值删除第一个匹配元素。
func RemoveValueInt64(slice []int64, target int64) ([]int64, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt64(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt64(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueInt64 removes all elements equal to target.
// RemoveAllValueInt64 按值删除所有匹配元素。
func RemoveAllValueInt64(slice []int64, target int64) ([]int64, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int64, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtInt64 removes the element at the specified index.
// RemoveAtInt64 按索引删除元素。
func RemoveAtInt64(slice []int64, pos int) ([]int64, int64, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int64, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadInt64 removes the first element.
// RemoveHeadInt64 删除头部元素。
func RemoveHeadInt64(slice []int64) ([]int64, int64, bool) {
	return RemoveAtInt64(slice, 0)
}

// RemoveTailInt64 removes the last element.
// RemoveTailInt64 删除尾部元素。
func RemoveTailInt64(slice []int64) ([]int64, int64, bool) {
	return RemoveAtInt64(slice, len(slice)-1)
}

// RemoveFromInt64 removes length elements starting at startPos.
// RemoveFromInt64 从 startPos 起删除 length 个元素。
func RemoveFromInt64(slice []int64, startPos int, length int) (result []int64, removed []int64, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt64(slice, startPos, endPos)
}

// RemoveRangeInt64 removes elements in [startPos, endPos).
// RemoveRangeInt64 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeInt64(slice []int64, startPos int, endPos int) (result []int64, removed []int64, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int64, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsInt64 reports whether slice contains target.
// ContainsInt64 判断切片是否包含目标值。
func ContainsInt64(slice []int64, target int64) bool {
	_, ok := IndexInt64(slice, target)
	return ok
}

// IndexInt64 returns the index of the first matching element.
// IndexInt64 从头部查找第一个匹配元素的索引。
func IndexInt64(slice []int64, target int64) (int, bool) {
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

// LastIndexInt64 returns the index of the last matching element.
// LastIndexInt64 从尾部查找最后一个匹配元素的索引。
func LastIndexInt64(slice []int64, target int64) (int, bool) {
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

// ReverseInt64 reverses elements in place.
// ReverseInt64 原地倒序元素。
func ReverseInt64(slice []int64) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualInt64 reports whether two slices are equal element-wise.
// EqualInt64 逐元素比较两个切片是否相等。
func EqualInt64(a, b []int64) bool {
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

// SumInt64 returns the sum of all elements.
// SumInt64 求和所有元素。
func SumInt64(slice []int64) int64 {
	rs := int64(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
