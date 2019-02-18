package slicex

//合并
func MergeComplex64(slices ...[]complex64) []complex64 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]complex64, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertComplex64(slice []complex64, pos int, target ...complex64) []complex64 {
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
	rs := make([]complex64, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadComplex64(slice []complex64, target ...complex64) []complex64 {
	return InsertComplex64(slice, 0, target...)
}

//尾插入
func InsertTailComplex64(slice []complex64, target ...complex64) []complex64 {
	return InsertComplex64(slice, len(slice), target...)
}

//按值删除
func RemoveValueComplex64(slice []complex64, target complex64) ([]complex64, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexComplex64(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtComplex64(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueComplex64(slice []complex64, target complex64) ([]complex64, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]complex64, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

//按点删除
func RemoveAtComplex64(slice []complex64, pos int) ([]complex64, complex64, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]complex64, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadComplex64(slice []complex64, pos int) ([]complex64, complex64, bool) {
	return RemoveAtComplex64(slice, 0)
}

//删除头
func RemoveComplex64ailComplex64(slice []complex64, pos int) ([]complex64, complex64, bool) {
	return RemoveAtComplex64(slice, len(slice)-1)
}

//删除区间
func RemoveFromComplex64(slice []complex64, startPos int, length int) (result []complex64, removed []complex64, ok bool) {
	endPos := startPos + length
	return RemoveRangeComplex64(slice, startPos, endPos)
}

//删除区间
func RemoveRangeComplex64(slice []complex64, startPos int, endPos int) (result []complex64, removed []complex64, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]complex64, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsComplex64(slice []complex64, target complex64) bool {
	_, ok := IndexComplex64(slice, target)
	return ok
}

//从头部查找
func IndexComplex64(slice []complex64, target complex64) (int, bool) {
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

//从尾部查找
func LastIndexComplex64(slice []complex64, target complex64) (int, bool) {
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

//倒序
func ReverseComplex64(slice []complex64) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
