package slicex

// MergeComplex128
// 合并
func MergeComplex128(slices ...[]complex128) []complex128 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]complex128, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertComplex128
// 按位置插入
func InsertComplex128(slice []complex128, pos int, target ...complex128) []complex128 {
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
	rs := make([]complex128, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadComplex128  头插入
func InsertHeadComplex128(slice []complex128, target ...complex128) []complex128 {
	return InsertComplex128(slice, 0, target...)
}

// InsertTailComplex128
// 尾插入
func InsertTailComplex128(slice []complex128, target ...complex128) []complex128 {
	return InsertComplex128(slice, len(slice), target...)
}

// RemoveValueComplex128
// 按值删除
func RemoveValueComplex128(slice []complex128, target complex128) ([]complex128, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexComplex128(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtComplex128(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueComplex128
// 按值删除
func RemoveAllValueComplex128(slice []complex128, target complex128) ([]complex128, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]complex128, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtComplex128
// 按点删除
func RemoveAtComplex128(slice []complex128, pos int) ([]complex128, complex128, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]complex128, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadComplex128
// 删除头
func RemoveHeadComplex128(slice []complex128) ([]complex128, complex128, bool) {
	return RemoveAtComplex128(slice, 0)
}

// RemoveTailComplex128
// 删除尾
func RemoveTailComplex128(slice []complex128) ([]complex128, complex128, bool) {
	return RemoveAtComplex128(slice, len(slice)-1)
}

// RemoveFromComplex128
// 删除区间
func RemoveFromComplex128(slice []complex128, startPos int, length int) (result []complex128, removed []complex128, ok bool) {
	endPos := startPos + length
	return RemoveRangeComplex128(slice, startPos, endPos)
}

// RemoveRangeComplex128
// 删除区间
func RemoveRangeComplex128(slice []complex128, startPos int, endPos int) (result []complex128, removed []complex128, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]complex128, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsComplex128
// 包含
func ContainsComplex128(slice []complex128, target complex128) bool {
	_, ok := IndexComplex128(slice, target)
	return ok
}

// IndexComplex128
// 从头部查找
func IndexComplex128(slice []complex128, target complex128) (int, bool) {
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

// LastIndexComplex128
// 从尾部查找
func LastIndexComplex128(slice []complex128, target complex128) (int, bool) {
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

// ReverseComplex128
// 倒序
func ReverseComplex128(slice []complex128) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualComplex128
// 比较
func EqualComplex128(a, b []complex128) bool {
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
