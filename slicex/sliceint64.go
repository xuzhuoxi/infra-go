package slicex

// MergeInt64
// 合并
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

// InsertInt64
// 按位置插入
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

// InsertHeadInt64
// 头插入
func InsertHeadInt64(slice []int64, target ...int64) []int64 {
	return InsertInt64(slice, 0, target...)
}

// InsertTailInt64
// 尾插入
func InsertTailInt64(slice []int64, target ...int64) []int64 {
	return InsertInt64(slice, len(slice), target...)
}

// RemoveValueInt64
// 按值删除
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

// RemoveAllValueInt64
// 按值删除
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

// RemoveAtInt64
// 按点删除
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

// RemoveHeadInt64
// 删除头
func RemoveHeadInt64(slice []int64) ([]int64, int64, bool) {
	return RemoveAtInt64(slice, 0)
}

// RemoveTailInt64
// 删除尾
func RemoveTailInt64(slice []int64) ([]int64, int64, bool) {
	return RemoveAtInt64(slice, len(slice)-1)
}

// RemoveFromInt64
// 删除区间
func RemoveFromInt64(slice []int64, startPos int, length int) (result []int64, removed []int64, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt64(slice, startPos, endPos)
}

// RemoveRangeInt64
// 删除区间
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

// ContainsInt64
// 包含
func ContainsInt64(slice []int64, target int64) bool {
	_, ok := IndexInt64(slice, target)
	return ok
}

// IndexInt64
// 从头部查找
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

// LastIndexInt64
// 从尾部查找
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

// ReverseInt64
// 倒序
func ReverseInt64(slice []int64) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualInt64
// 比较
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

// SumInt64
// 求和
func SumInt64(slice []int64) int64 {
	rs := int64(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
