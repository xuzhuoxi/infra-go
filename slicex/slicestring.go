package slicex

import "strings"

// MergeString
// 合并
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

// InsertString
// 按位置插入
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

// InsertHeadString
// 头插入
func InsertHeadString(slice []string, target ...string) []string {
	return InsertString(slice, 0, target...)
}

// InsertTailString
// 尾插入
func InsertTailString(slice []string, target ...string) []string {
	return InsertString(slice, len(slice), target...)
}

// RemoveValueString
// 按值删除
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

// RemoveAllValueString
// 按值删除
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

// RemoveAtString
// 按点删除
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

// RemoveHeadString
// 删除头
func RemoveHeadString(slice []string) ([]string, string, bool) {
	return RemoveAtString(slice, 0)
}

// RemoveTailString
// 删除尾
func RemoveTailString(slice []string) ([]string, string, bool) {
	return RemoveAtString(slice, len(slice)-1)
}

// RemoveFromString
// 删除区间
func RemoveFromString(slice []string, startPos int, length int) (result []string, removed []string, ok bool) {
	endPos := startPos + length
	return RemoveRangeString(slice, startPos, endPos)
}

// RemoveRangeString
// 删除区间
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

// ContainsString
// 包含
func ContainsString(slice []string, target string) bool {
	_, ok := IndexString(slice, target)
	return ok
}

// ContainsCaseString
// 包含
// ignoreCase 是否忽略大小写
func ContainsCaseString(slice []string, target string, ignoreCase bool) bool {
	if !ignoreCase {
		return ContainsString(slice, target)
	}
	_, ok := IndexCaseString(slice, target, true)
	return ok
}

// IndexString
// 从头部查找
func IndexString(slice []string, target string) (int, bool) {
	return IndexCaseString(slice, target, false)
}

// IndexCaseString
// 从头部查找
// ignoreCase 是否忽略大小写
func IndexCaseString(slice []string, target string, ignoreCase bool) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	if !ignoreCase {
		for index, value := range slice {
			if value == target {
				return index, true
			}
		}
	} else {
		for index, value := range slice {
			if strings.EqualFold(value, target) {
				return index, true
			}
		}
	}
	return -1, false
}

// LastIndexString
// 从尾部查找
func LastIndexString(slice []string, target string) (int, bool) {
	return LastIndexCaseString(slice, target, false)
}

// LastIndexCaseString
// 从尾部查找
// ignoreCase 是否忽略大小写
func LastIndexCaseString(slice []string, target string, ignoreCase bool) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	if !ignoreCase {
		for index := len(slice) - 1; index >= 0; index-- {
			if slice[index] == target {
				return index, true
			}
		}
	} else {
		for index := len(slice) - 1; index >= 0; index-- {
			if strings.EqualFold(slice[index], target) {
				return index, true
			}
		}
	}
	return -1, false
}

// ReverseString
// 倒序
func ReverseString(slice []string) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// CopyString
// 克隆
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

// EqualString
// 比较
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

// ClearDuplicateString
// 清除重复值
func ClearDuplicateString(slice []string) {
	if nil == slice || len(slice) < 2 {
		return
	}
	var headIndex = 0
	for tailIndex := len(slice) - 1; tailIndex >= 0; tailIndex-- {
		for headIndex = 0; headIndex < tailIndex; headIndex++ {
			if slice[tailIndex] == slice[headIndex] {
				slice = append(slice[:tailIndex], slice[tailIndex+1:]...)
				break
			}
		}
	}
}
