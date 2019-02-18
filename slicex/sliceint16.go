package slicex

//合并
func MergeInt16(slices ...[]int16) []int16 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int16, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertInt16(slice []int16, pos int, target ...int16) []int16 {
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
	rs := make([]int16, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadInt16(slice []int16, target ...int16) []int16 {
	return InsertInt16(slice, 0, target...)
}

//尾插入
func InsertTailInt16(slice []int16, target ...int16) []int16 {
	return InsertInt16(slice, len(slice), target...)
}

//按值删除
func RemoveValueInt16(slice []int16, target int16) ([]int16, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt16(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt16(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueInt16(slice []int16, target int16) ([]int16, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int16, sl)
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
func RemoveAtInt16(slice []int16, pos int) ([]int16, int16, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int16, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadInt16(slice []int16, pos int) ([]int16, int16, bool) {
	return RemoveAtInt16(slice, 0)
}

//删除头
func RemoveInt16ailInt16(slice []int16, pos int) ([]int16, int16, bool) {
	return RemoveAtInt16(slice, len(slice)-1)
}

//删除区间
func RemoveFromInt16(slice []int16, startPos int, length int) (result []int16, removed []int16, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt16(slice, startPos, endPos)
}

//删除区间
func RemoveRangeInt16(slice []int16, startPos int, endPos int) (result []int16, removed []int16, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int16, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsInt16(slice []int16, target int16) bool {
	_, ok := IndexInt16(slice, target)
	return ok
}

//从头部查找
func IndexInt16(slice []int16, target int16) (int, bool) {
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
func LastIndexInt16(slice []int16, target int16) (int, bool) {
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
func ReverseInt16(slice []int16) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
