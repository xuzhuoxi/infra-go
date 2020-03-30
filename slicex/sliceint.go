package slicex

//合并
func MergeInt(slices ...[]int) []int {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]int, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertInt(slice []int, pos int, target ...int) []int {
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
	rs := make([]int, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadInt(slice []int, target ...int) []int {
	return InsertInt(slice, 0, target...)
}

//尾插入
func InsertTailInt(slice []int, target ...int) []int {
	return InsertInt(slice, len(slice), target...)
}

//按值删除
func RemoveValueInt(slice []int, target int) ([]int, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexInt(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtInt(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueInt(slice []int, target int) ([]int, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]int, sl)
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
func RemoveAtInt(slice []int, pos int) ([]int, int, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]int, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除头
func RemoveHeadInt(slice []int) ([]int, int, bool) {
	return RemoveAtInt(slice, 0)
}

//删除尾
func RemoveTailInt(slice []int) ([]int, int, bool) {
	return RemoveAtInt(slice, len(slice)-1)
}

//删除区间
func RemoveFromInt(slice []int, startPos int, length int) (result []int, removed []int, ok bool) {
	endPos := startPos + length
	return RemoveRangeInt(slice, startPos, endPos)
}

//删除区间
func RemoveRangeInt(slice []int, startPos int, endPos int) (result []int, removed []int, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]int, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsInt(slice []int, target int) bool {
	_, ok := IndexInt(slice, target)
	return ok
}

//从头部查找
func IndexInt(slice []int, target int) (int, bool) {
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
func LastIndexInt(slice []int, target int) (int, bool) {
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
func ReverseInt(slice []int) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

//比较
func EqualInt(a, b []int) bool {
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

// 求和
func SumInt(slice []int) int {
	rs := 0
	for _, val := range slice {
		rs += val
	}
	return rs
}
