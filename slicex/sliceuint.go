package slicex

//合并
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

//按位置插入
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

//头插入
func InsertHeadUint(slice []uint, target ...uint) []uint {
	return InsertUint(slice, 0, target...)
}

//尾插入
func InsertUintailUint(slice []uint, target ...uint) []uint {
	return InsertUint(slice, len(slice), target...)
}

//按值删除
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

//按值删除
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

//按点删除
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

//删除尾
func RemoveHeadUint(slice []uint, pos int) ([]uint, uint, bool) {
	return RemoveAtUint(slice, 0)
}

//删除头
func RemoveUintailUint(slice []uint, pos int) ([]uint, uint, bool) {
	return RemoveAtUint(slice, len(slice)-1)
}

//删除区间
func RemoveFromUint(slice []uint, startPos int, length int) (result []uint, removed []uint, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint(slice, startPos, endPos)
}

//删除区间
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

//包含
func ContainsUint(slice []uint, target uint) bool {
	_, ok := IndexUint(slice, target)
	return ok
}

//从头部查找
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

//从尾部查找
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

//倒序
func ReverseUint(slice []uint) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
