package slicex

//合并
func MergeUint32(slices ...[]uint32) []uint32 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint32, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertUint32(slice []uint32, pos int, target ...uint32) []uint32 {
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
	rs := make([]uint32, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadUint32(slice []uint32, target ...uint32) []uint32 {
	return InsertUint32(slice, 0, target...)
}

//尾插入
func InsertUint32ailUint32(slice []uint32, target ...uint32) []uint32 {
	return InsertUint32(slice, len(slice), target...)
}

//按值删除
func RemoveValueUint32(slice []uint32, target uint32) ([]uint32, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint32(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint32(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueUint32(slice []uint32, target uint32) ([]uint32, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint32, sl)
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
func RemoveAtUint32(slice []uint32, pos int) ([]uint32, uint32, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint32, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadUint32(slice []uint32, pos int) ([]uint32, uint32, bool) {
	return RemoveAtUint32(slice, 0)
}

//删除头
func RemoveUint32ailUint32(slice []uint32, pos int) ([]uint32, uint32, bool) {
	return RemoveAtUint32(slice, len(slice)-1)
}

//删除区间
func RemoveFromUint32(slice []uint32, startPos int, length int) (result []uint32, removed []uint32, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint32(slice, startPos, endPos)
}

//删除区间
func RemoveRangeUint32(slice []uint32, startPos int, endPos int) (result []uint32, removed []uint32, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint32, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsUint32(slice []uint32, target uint32) bool {
	_, ok := IndexUint32(slice, target)
	return ok
}

//从头部查找
func IndexUint32(slice []uint32, target uint32) (int, bool) {
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
func LastIndexUint32(slice []uint32, target uint32) (int, bool) {
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
func ReverseUint32(slice []uint32) []uint32 {
	if nil == slice {
		return nil
	}
	ln := len(slice)
	if 0 == ln {
		return []uint32{}
	}
	rs := make([]uint32, ln)
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = slice[j], slice[i]
	}
	return rs
}
