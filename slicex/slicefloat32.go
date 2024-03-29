package slicex

// MergeFloat32
// 合并
func MergeFloat32(slices ...[]float32) []float32 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]float32, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertFloat32
// 按位置插入
func InsertFloat32(slice []float32, pos int, target ...float32) []float32 {
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
	rs := make([]float32, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadFloat32
// 头插入
func InsertHeadFloat32(slice []float32, target ...float32) []float32 {
	return InsertFloat32(slice, 0, target...)
}

// InsertTailFloat32
// 尾插入
func InsertTailFloat32(slice []float32, target ...float32) []float32 {
	return InsertFloat32(slice, len(slice), target...)
}

// RemoveValueFloat32
// 按值删除
func RemoveValueFloat32(slice []float32, target float32) ([]float32, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexFloat32(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtFloat32(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueFloat32
// 按值删除
func RemoveAllValueFloat32(slice []float32, target float32) ([]float32, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]float32, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtFloat32
// 按点删除
func RemoveAtFloat32(slice []float32, pos int) ([]float32, float32, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]float32, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadFloat32
// 删除头
func RemoveHeadFloat32(slice []float32) ([]float32, float32, bool) {
	return RemoveAtFloat32(slice, 0)
}

// RemoveTailFloat32
// 删除尾
func RemoveTailFloat32(slice []float32) ([]float32, float32, bool) {
	return RemoveAtFloat32(slice, len(slice)-1)
}

// RemoveFromFloat32
// 删除区间
func RemoveFromFloat32(slice []float32, startPos int, length int) (result []float32, removed []float32, ok bool) {
	endPos := startPos + length
	return RemoveRangeFloat32(slice, startPos, endPos)
}

// RemoveRangeFloat32
// 删除区间
func RemoveRangeFloat32(slice []float32, startPos int, endPos int) (result []float32, removed []float32, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]float32, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsFloat32
// 包含
func ContainsFloat32(slice []float32, target float32) bool {
	_, ok := IndexFloat32(slice, target)
	return ok
}

// IndexFloat32
// 从头部查找
func IndexFloat32(slice []float32, target float32) (int, bool) {
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

// LastIndexFloat32
// 从尾部查找
func LastIndexFloat32(slice []float32, target float32) (int, bool) {
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

// ReverseFloat32
// 倒序
func ReverseFloat32(slice []float32) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualFloat32
// 比较
func EqualFloat32(a, b []float32) bool {
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

// SumFloat32
// 求和
func SumFloat32(slice []float32) float32 {
	rs := float32(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
