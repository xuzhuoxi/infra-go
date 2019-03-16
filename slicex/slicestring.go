package slicex

//合并
func MergeString(slices ...[]string) []string {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]string, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

//按位置插入
func InsertString(slice []string, pos int, target ...string) []string {
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
	rs := make([]string, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadString(slice []string, target ...string) []string {
	return InsertString(slice, 0, target...)
}

//尾插入
func InsertTailString(slice []string, target ...string) []string {
	return InsertString(slice, len(slice), target...)
}

//按值删除
func RemoveValueString(slice []string, target string) ([]string, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexString(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtString(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueString(slice []string, target string) ([]string, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]string, sl)
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
func RemoveAtString(slice []string, pos int) ([]string, string, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, "", false
	}
	obj := slice[pos]
	rs := make([]string, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadString(slice []string, pos int) ([]string, string, bool) {
	return RemoveAtString(slice, 0)
}

//删除头
func RemoveStringailString(slice []string, pos int) ([]string, string, bool) {
	return RemoveAtString(slice, len(slice)-1)
}

//删除区间
func RemoveFromString(slice []string, startPos int, length int) (result []string, removed []string, ok bool) {
	endPos := startPos + length
	return RemoveRangeString(slice, startPos, endPos)
}

//删除区间
func RemoveRangeString(slice []string, startPos int, endPos int) (result []string, removed []string, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]string, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsString(slice []string, target string) bool {
	_, ok := IndexString(slice, target)
	return ok
}

//从头部查找
func IndexString(slice []string, target string) (int, bool) {
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
func LastIndexString(slice []string, target string) (int, bool) {
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
func ReverseString(slice []string) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

//克隆
func CopyString(slice []string) []string {
	if nil == slice {
		return nil
	}
	ln := len(slice)
	if 0 == ln {
		return []string{}
	}
	rs := make([]string, ln)
	copy(rs, slice)
	return rs
}

//比较
func EqualString(a, b []string) bool {
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
