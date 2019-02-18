package slicex

//合并
func MergeInt32(slices ...[]int32) []int32 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int32, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertInt32(slice []int32, pos int, target ...int32) []int32 {
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
	rs := make([]int32, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadInt32(slice []int32, target ...int32) []int32 {
	return InsertInt32(slice, 0, target...)
}

//尾插入
func InsertTailInt32(slice []int32, target ...int32) []int32 {
	return InsertInt32(slice, len(slice), target...)
}

//按值删除
func RemoveValueInt32(slice []int32, target int32) ([]int32, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt32(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt32(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueInt32(slice []int32, target int32) ([]int32, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int32, sl)
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
func RemoveAtInt32(slice []int32, pos int) ([]int32, int32, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int32, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadInt32(slice []int32, pos int) ([]int32, int32, bool) {
	return RemoveAtInt32(slice, 0)
}

//删除头
func RemoveInt32ailInt32(slice []int32, pos int) ([]int32, int32, bool) {
	return RemoveAtInt32(slice, len(slice)-1)
}

//删除区间
func RemoveFromInt32(slice []int32, startPos int, length int) (result []int32, removed []int32, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt32(slice, startPos, endPos)
}

//删除区间
func RemoveRangeInt32(slice []int32, startPos int, endPos int) (result []int32, removed []int32, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int32, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsInt32(slice []int32, target int32) bool {
	_, ok := IndexInt32(slice, target)
	return ok
}

//从头部查找
func IndexInt32(slice []int32, target int32) (int, bool) {
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
func LastIndexInt32(slice []int32, target int32) (int, bool) {
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
func ReverseInt32(slice []int32) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
