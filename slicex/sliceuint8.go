package slicex

//合并
func MergeUint8(slices ...[]uint8) []uint8 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint8, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertUint8(slice []uint8, pos int, target ...uint8) []uint8 {
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
	rs := make([]uint8, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadUint8(slice []uint8, target ...uint8) []uint8 {
	return InsertUint8(slice, 0, target...)
}

//尾插入
func InsertUint8ailUint8(slice []uint8, target ...uint8) []uint8 {
	return InsertUint8(slice, len(slice), target...)
}

//按值删除
func RemoveValueUint8(slice []uint8, target uint8) ([]uint8, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint8(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint8(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueUint8(slice []uint8, target uint8) ([]uint8, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint8, sl)
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
func RemoveAtUint8(slice []uint8, pos int) ([]uint8, uint8, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint8, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadUint8(slice []uint8, pos int) ([]uint8, uint8, bool) {
	return RemoveAtUint8(slice, 0)
}

//删除头
func RemoveUint8ailUint8(slice []uint8, pos int) ([]uint8, uint8, bool) {
	return RemoveAtUint8(slice, len(slice)-1)
}

//删除区间
func RemoveFromUint8(slice []uint8, startPos int, length int) (result []uint8, removed []uint8, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint8(slice, startPos, endPos)
}

//删除区间
func RemoveRangeUint8(slice []uint8, startPos int, endPos int) (result []uint8, removed []uint8, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint8, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsUint8(slice []uint8, target uint8) bool {
	_, ok := IndexUint8(slice, target)
	return ok
}

//从头部查找
func IndexUint8(slice []uint8, target uint8) (int, bool) {
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
func LastIndexUint8(slice []uint8, target uint8) (int, bool) {
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
func ReverseUint8(slice []uint8) []uint8 {
	if nil == slice {
		return nil
	}
	ln := len(slice)
	if 0 == ln {
		return []uint8{}
	}
	rs := make([]uint8, ln)
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = slice[j], slice[i]
	}
	return rs
}
