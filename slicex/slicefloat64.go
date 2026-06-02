package slicex

// MergeFloat64 merges multiple slices into one.
// MergeFloat64 合并多个切片为一个。
func MergeFloat64(slices ...[]float64) []float64 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]float64, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertFloat64 inserts elements at the specified position.
// InsertFloat64 按位置插入元素。
func InsertFloat64(slice []float64, pos int, target ...float64) []float64 {
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
	rs := make([]float64, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadFloat64 inserts elements at the head.
// InsertHeadFloat64 在头部插入元素。
func InsertHeadFloat64(slice []float64, target ...float64) []float64 {
	return InsertFloat64(slice, 0, target...)
}

// InsertTailFloat64 inserts elements at the tail.
// InsertTailFloat64 在尾部插入元素。
func InsertTailFloat64(slice []float64, target ...float64) []float64 {
	return InsertFloat64(slice, len(slice), target...)
}

// RemoveValueFloat64 removes the first element equal to target.
// RemoveValueFloat64 按值删除第一个匹配元素。
func RemoveValueFloat64(slice []float64, target float64) ([]float64, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexFloat64(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtFloat64(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueFloat64 removes all elements equal to target.
// RemoveAllValueFloat64 按值删除所有匹配元素。
func RemoveAllValueFloat64(slice []float64, target float64) ([]float64, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]float64, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtFloat64 removes the element at the specified index.
// RemoveAtFloat64 按索引删除元素。
func RemoveAtFloat64(slice []float64, pos int) ([]float64, float64, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]float64, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadFloat64 removes the first element.
// RemoveHeadFloat64 删除头部元素。
func RemoveHeadFloat64(slice []float64) ([]float64, float64, bool) {
	return RemoveAtFloat64(slice, 0)
}

// RemoveTailFloat64 removes the last element.
// RemoveTailFloat64 删除尾部元素。
func RemoveTailFloat64(slice []float64) ([]float64, float64, bool) {
	return RemoveAtFloat64(slice, len(slice)-1)
}

// RemoveFromFloat64 removes length elements starting at startPos.
// RemoveFromFloat64 从 startPos 起删除 length 个元素。
func RemoveFromFloat64(slice []float64, startPos int, length int) (result []float64, removed []float64, ok bool) {
	endPos := startPos + length
	return RemoveRangeFloat64(slice, startPos, endPos)
}

// RemoveRangeFloat64 removes elements in [startPos, endPos).
// RemoveRangeFloat64 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeFloat64(slice []float64, startPos int, endPos int) (result []float64, removed []float64, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]float64, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsFloat64 reports whether slice contains target.
// ContainsFloat64 判断切片是否包含目标值。
func ContainsFloat64(slice []float64, target float64) bool {
	_, ok := IndexFloat64(slice, target)
	return ok
}

// IndexFloat64 returns the index of the first matching element.
// IndexFloat64 从头部查找第一个匹配元素的索引。
func IndexFloat64(slice []float64, target float64) (int, bool) {
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

// LastIndexFloat64 returns the index of the last matching element.
// LastIndexFloat64 从尾部查找最后一个匹配元素的索引。
func LastIndexFloat64(slice []float64, target float64) (int, bool) {
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

// ReverseFloat64 reverses elements in place.
// ReverseFloat64 原地倒序元素。
func ReverseFloat64(slice []float64) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualFloat64 reports whether two slices are equal element-wise.
// EqualFloat64 逐元素比较两个切片是否相等。
func EqualFloat64(a, b []float64) bool {
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

// SumFloat64 returns the sum of all elements.
// SumFloat64 求和所有元素。
func SumFloat64(slice []float64) float64 {
	rs := float64(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
