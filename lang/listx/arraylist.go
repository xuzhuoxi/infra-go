package listx

import (
	"github.com/xuzhuoxi/infra-go/lang"
)

// NewArrayList creates a slice-backed list with optional max size and initial capacity.
// NewArrayList 创建一个基于切片实现的列表，可指定最大长度与初始容量。
// maxSize: maximum allowed length; 0 means unlimited. initCap: initial slice capacity.
// maxSize: 列表允许的最大长度，0 表示不限制；initCap: 底层切片初始容量。
// Returns nil if maxSize or initCap is negative. If maxSize > 0 and maxSize < initCap, maxSize is raised to initCap.
// 若 maxSize 或 initCap 为负数则返回 nil。若 maxSize > 0 且 maxSize < initCap，则将 maxSize 调整为 initCap。
func NewArrayList(maxSize, initCap int) *ArrayList {
	if maxSize < 0 || initCap < 0 {
		return nil
	}
	if maxSize > 0 && initCap > maxSize {
		maxSize = initCap
	}
	rs := &ArrayList{maxSize: maxSize, initCap: initCap}
	rs.Clear()
	return rs
}

// ArrayList is an IList backed by a Go slice.
// ArrayList 是基于 Go 切片的 IList 实现。
type ArrayList struct {
	maxSize  int
	initCap  int
	elements []interface{}
}

// Len returns the number of elements in the list.
// Len 返回列表中的元素个数。
func (l *ArrayList) Len() int {
	return len(l.elements)
}

// Swap swaps the elements at indexes i and j.
// Swap 交换索引 i 与 j 处的元素。
func (l *ArrayList) Swap(i, j int) {
	l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
}

// Clear removes all elements from the list.
// Clear 移除列表中的全部元素。
func (l *ArrayList) Clear() {
	if l.initCap > 0 {
		l.elements = make([]interface{}, 0, l.initCap)
	} else {
		l.elements = nil
	}
}

// Add appends elements to the end of the list.
// Add 在列表末尾追加元素。
func (l *ArrayList) Add(ele ...interface{}) (suc bool) {
	if l.maxSize > 0 && len(ele)+len(l.elements) > l.maxSize {
		return false
	}
	l.elements = append(l.elements, ele...)
	return true
}

// AddAt inserts elements at the specified index.
// AddAt 在指定索引处插入元素。
func (l *ArrayList) AddAt(index int, ele ...interface{}) (suc bool) {
	addLn := len(ele)
	if l.maxSize > 0 && addLn+len(l.elements) > l.maxSize {
		return false
	}
	l.elements = append(l.elements, ele...)
	copy(l.elements[index+addLn:], l.elements[index:])
	copy(l.elements[index:index+addLn], ele)
	return true
}

// AddAll inserts all elements from list starting at index.
// AddAll 从 index 起插入 list 中的全部元素。
func (l *ArrayList) AddAll(index int, list IList) (suc bool) {
	if all := list.GetAll(); all != nil {
		return l.AddAt(index, all...)
	}
	return true
}

// RemoveAt removes the element at the specified index.
// RemoveAt 移除指定索引处的元素。
func (l *ArrayList) RemoveAt(index int) (ele interface{}, suc bool) {
	if index >= 0 && index < l.Len() {
		rs := l.elements[index]
		l.elements = append(l.elements[:index], l.elements[index+1:]...)
		return rs, true
	}
	return nil, false
}

// RemoveMultiAt removes amount elements starting at index.
// RemoveMultiAt 从 index 起移除 amount 个元素。
func (l *ArrayList) RemoveMultiAt(index int, amount int) (ele []interface{}, suc bool) {
	ln := l.Len()
	endIndex := index + amount
	if amount > 0 && index >= 0 && endIndex <= ln {
		removes := l.elements[index : index+amount]
		l.elements = append(l.elements, removes...)
		ele = l.elements[ln : ln+amount]
		l.elements = append(l.elements[:index], l.elements[index+amount:ln]...)
		suc = true
		return
	}
	return nil, false
}

// RemoveLast removes the last element.
// RemoveLast 移除最后一个元素。
func (l *ArrayList) RemoveLast() (ele interface{}, suc bool) {
	return l.RemoveAt(l.Len() - 1)
}

// RemoveLastMulti removes amount elements from the end.
// RemoveLastMulti 从末尾起移除 amount 个元素。
func (l *ArrayList) RemoveLastMulti(amount int) (ele []interface{}, suc bool) {
	startIndex := l.Len() - amount
	return l.RemoveMultiAt(startIndex, amount)
}

// RemoveFirst removes the first element.
// RemoveFirst 移除第一个元素。
func (l *ArrayList) RemoveFirst() (ele interface{}, suc bool) {
	return l.RemoveAt(0)
}

// RemoveFirstMulti removes amount elements from the front.
// RemoveFirstMulti 从开头起移除 amount 个元素。
func (l *ArrayList) RemoveFirstMulti(amount int) (ele []interface{}, suc bool) {
	return l.RemoveMultiAt(0, amount)
}

// Get returns the element at the specified index.
// Get 返回指定索引处的元素。
func (l *ArrayList) Get(index int) (ele interface{}, ok bool) {
	ln := l.Len()
	if ln > 0 && index >= 0 && index < ln {
		return l.elements[index], true
	}
	return nil, false
}

// GetMulti returns amount elements starting at index.
// GetMulti 从 index 起返回 amount 个元素。
func (l *ArrayList) GetMulti(index int, amount int) (ele []interface{}, ok bool) {
	if amount > 0 && index >= 0 && index+amount <= l.Len() {
		return l.elements[index : index+amount], true
	}
	return nil, false
}

// GetMultiLast returns amount elements ending at lastIndex.
// GetMultiLast 以 lastIndex 为终点返回 amount 个元素。
func (l *ArrayList) GetMultiLast(lastIndex int, amount int) (ele []interface{}, ok bool) {
	startIndex := lastIndex - amount + 1
	return l.GetMulti(startIndex, amount)
}

// GetAll returns all elements in the list.
// GetAll 返回列表中的全部元素。
func (l *ArrayList) GetAll() []interface{} {
	return l.elements
}

// First returns the first element.
// First 返回第一个元素。
func (l *ArrayList) First() (ele interface{}, ok bool) {
	return l.Get(0)
}

// FirstMulti returns amount elements from the front.
// FirstMulti 从开头返回 amount 个元素。
func (l *ArrayList) FirstMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMulti(0, amount)
}

// Last returns the last element.
// Last 返回最后一个元素。
func (l *ArrayList) Last() (ele interface{}, ok bool) {
	return l.Get(l.Len() - 1)
}

// LastMulti returns amount elements from the end.
// LastMulti 从末尾返回 amount 个元素。
func (l *ArrayList) LastMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMultiLast(l.Len()-1, amount)
}

// IndexOf returns the index of the first matching element.
// IndexOf 返回首个匹配元素的索引。
func (l *ArrayList) IndexOf(ele interface{}) (index int, ok bool) {
	ln := len(l.elements)
	for index := 0; index < ln; index++ {
		if lang.Equal(l.elements[index], ele) {
			return index, true
		}
	}
	return -1, false
}

// LastIndexOf returns the index of the last matching element.
// LastIndexOf 返回最后一个匹配元素的索引。
func (l *ArrayList) LastIndexOf(ele interface{}) (index int, ok bool) {
	for index := len(l.elements) - 1; index >= 0; index-- {
		if lang.Equal(ele, l.elements[index]) {
			return index, true
		}
	}
	return -1, false
}

// Contains reports whether the list contains ele.
// Contains 判断列表是否包含 ele。
func (l *ArrayList) Contains(ele interface{}) (contains bool) {
	_, contains = l.IndexOf(ele)
	return
}

// ForEach iterates elements from front to back until stop is true.
// ForEach 从前向后遍历元素，stop 为 true 时停止。
func (l *ArrayList) ForEach(each func(index int, ele interface{}) (stop bool)) {
	ln := len(l.elements)
	for index := 0; index < ln; index++ {
		if each(index, l.elements[index]) {
			break
		}
	}
}

// ForEachLast iterates elements from back to front until stop is true.
// ForEachLast 从后向前遍历元素，stop 为 true 时停止。
func (l *ArrayList) ForEachLast(each func(index int, ele interface{}) (stop bool)) {
	for index := len(l.elements) - 1; index >= 0; index-- {
		if each(index, l.elements[index]) {
			break
		}
	}
}
