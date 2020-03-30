package slicex

//合并
func MergeFloat64(slices ...[]float64) []float64 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]float64, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertFloat64(slice []float64, pos int, target ...float64) []float64 {
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
	rs := make([]float64, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadFloat64(slice []float64, target ...float64) []float64 {
	return InsertFloat64(slice, 0, target...)
}

//尾插入
func InsertTailFloat64(slice []float64, target ...float64) []float64 {
	return InsertFloat64(slice, len(slice), target...)
}

//按值删除
func RemoveValueFloat64(slice []float64, target float64) ([]float64, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexFloat64(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtFloat64(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueFloat64(slice []float64, target float64) ([]float64, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]float64, sl)
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
func RemoveAtFloat64(slice []float64, pos int) ([]float64, float64, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]float64, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除头
func RemoveHeadFloat64(slice []float64) ([]float64, float64, bool) {
	return RemoveAtFloat64(slice, 0)
}

//删除尾
func RemoveTailFloat64(slice []float64) ([]float64, float64, bool) {
	return RemoveAtFloat64(slice, len(slice)-1)
}

//删除区间
func RemoveFromFloat64(slice []float64, startPos int, length int) (result []float64, removed []float64, ok bool) {
	endPos := startPos + length
	return RemoveRangeFloat64(slice, startPos, endPos)
}

//删除区间
func RemoveRangeFloat64(slice []float64, startPos int, endPos int) (result []float64, removed []float64, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]float64, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsFloat64(slice []float64, target float64) bool {
	_, ok := IndexFloat64(slice, target)
	return ok
}

//从头部查找
func IndexFloat64(slice []float64, target float64) (int, bool) {
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
func LastIndexFloat64(slice []float64, target float64) (int, bool) {
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
func ReverseFloat64(slice []float64) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

//比较
func EqualFloat64(a, b []float64) bool {
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
func SumFloat64(slice []float64) float64 {
	rs := float64(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
