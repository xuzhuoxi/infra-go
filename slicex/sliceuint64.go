package slicex

//合并
func MergeUint64(slices ...[]uint64) []uint64 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint64, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertUint64(slice []uint64, pos int, target ...uint64) []uint64 {
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
	rs := make([]uint64, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadUint64(slice []uint64, target ...uint64) []uint64 {
	return InsertUint64(slice, 0, target...)
}

//尾插入
func InsertTailUint64(slice []uint64, target ...uint64) []uint64 {
	return InsertUint64(slice, len(slice), target...)
}

//按值删除
func RemoveValueUint64(slice []uint64, target uint64) ([]uint64, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint64(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint64(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueUint64(slice []uint64, target uint64) ([]uint64, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint64, sl)
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
func RemoveAtUint64(slice []uint64, pos int) ([]uint64, uint64, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint64, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadUint64(slice []uint64, pos int) ([]uint64, uint64, bool) {
	return RemoveAtUint64(slice, 0)
}

//删除头
func RemoveUint64ailUint64(slice []uint64, pos int) ([]uint64, uint64, bool) {
	return RemoveAtUint64(slice, len(slice)-1)
}

//删除区间
func RemoveFromUint64(slice []uint64, startPos int, length int) (result []uint64, removed []uint64, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint64(slice, startPos, endPos)
}

//删除区间
func RemoveRangeUint64(slice []uint64, startPos int, endPos int) (result []uint64, removed []uint64, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint64, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsUint64(slice []uint64, target uint64) bool {
	_, ok := IndexUint64(slice, target)
	return ok
}

//从头部查找
func IndexUint64(slice []uint64, target uint64) (int, bool) {
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
func LastIndexUint64(slice []uint64, target uint64) (int, bool) {
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
func ReverseUint64(slice []uint64) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

//比较
func EqualUint64(a, b []uint64) bool {
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

func UintToUint64(source []uint) []uint64 {
	if nil == source {
		return nil
	}
	if len(source) == 0 {
		return []uint64{}
	}
	var rs []uint64
	for _, val := range source {
		rs = append(rs, uint64(val))
	}
	return rs
}

func Uint64ToUint(source []uint64) []uint {
	if nil == source {
		return nil
	}
	if len(source) == 0 {
		return []uint{}
	}
	var rs []uint
	for _, val := range source {
		rs = append(rs, uint(val))
	}
	return rs
}
