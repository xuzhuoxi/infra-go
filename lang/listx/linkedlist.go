package listx

import (
	"container/list"

	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/mathx"
)

// NewLinkedList creates a doubly-linked IList backed by container/list.
// NewLinkedList 创建一个基于 container/list 的双向链表 IList。
func NewLinkedList() *LinkedList {
	rs := &LinkedList{list: list.New()}
	return rs
}

// LinkedList is an IList backed by container/list.List.
// LinkedList 是基于 container/list.List 的 IList 实现。
type LinkedList struct {
	list *list.List
}

// Len returns the number of elements in the list.
// Len 返回列表中的元素个数。
func (l *LinkedList) Len() int {
	return l.list.Len()
}

// Swap swaps the elements at indexes i and j.
// Swap 交换索引 i 与 j 处的元素。
func (l *LinkedList) Swap(i, j int) {
	if i == j {
		return
	}
	before := mathx.MinInt(i, j)
	after := mathx.MaxInt(i, j)
	if eleBefore, okBefore := l.getElement(before); okBefore {
		if eleAfter, okAfter := l.getElement(after); okAfter {
			switch {
			case 0 == before && 1 == after: //只有两个元素
				l.list.MoveToFront(eleAfter)
			case 0 == before && after == l.Len()-1: //头和尾
				l.list.MoveToBack(eleBefore)
				l.list.MoveToFront(eleAfter)
			default:
				if 0 == before { //头，不是尾
					beforeNext := eleBefore.Next() //必须
					afterNext := eleAfter.Next()
					l.list.MoveBefore(eleAfter, beforeNext)
					l.list.MoveBefore(eleBefore, afterNext)
				} else { //不是头，可能是尾
					afterPrev := eleAfter.Prev() //必须
					beforePrev := eleBefore.Prev()
					l.list.MoveAfter(eleBefore, afterPrev)
					l.list.MoveAfter(eleAfter, beforePrev)
				}
			}
		}
	}
}

// Clear removes all elements from the list.
// Clear 移除列表中的全部元素。
func (l *LinkedList) Clear() {
	l.list.Init()
}

// Add appends elements to the end of the list.
// Add 在列表末尾追加元素。
func (l *LinkedList) Add(ele ...interface{}) (suc bool) {
	ln := len(ele)
	if ln == 0 {
		return false
	}
	for index := 0; index < ln; index++ {
		l.list.PushBack(ele[index])
	}
	return true
}

// AddAt inserts elements at the specified index.
// AddAt 在指定索引处插入元素。
func (l *LinkedList) AddAt(index int, ele ...interface{}) (suc bool) {
	if index == l.Len() {
		return l.Add(ele...)
	}
	if tempEle, ok := l.getElement(index); ok {
		pre := tempEle.Prev()
		if nil == pre {
			for index := len(ele) - 1; index >= 0; index-- {
				l.list.PushFront(ele[index])
			}
		} else {
			for index := len(ele) - 1; index >= 0; index-- {
				l.list.InsertAfter(ele[index], pre)
			}
		}
		return true
	}
	return false
}

// AddAll inserts all elements from list starting at index.
// AddAll 从 index 起插入 list 中的全部元素。
func (l *LinkedList) AddAll(index int, list IList) (suc bool) {
	if nil == list || list.Len() == 0 {
		return false
	}
	return l.AddAt(index, list.GetAll()...)
}

// RemoveAt removes the element at the specified index.
// RemoveAt 移除指定索引处的元素。
func (l *LinkedList) RemoveAt(index int) (ele interface{}, suc bool) {
	if tempEle, ok := l.getElement(index); ok {
		ele = l.list.Remove(tempEle)
		suc = true
		return
	}
	return
}

// RemoveMultiAt removes amount elements starting at index.
// RemoveMultiAt 从 index 起移除 amount 个元素。
func (l *LinkedList) RemoveMultiAt(index int, amount int) (ele []interface{}, suc bool) {
	if eleArr, ok := l.getElements(index, amount); ok {
		arr := make([]interface{}, amount)
		for index, ele := range eleArr {
			arr[index] = l.list.Remove(ele)
		}
		return arr[:], true
	}
	return
}

// RemoveLast removes the last element.
// RemoveLast 移除最后一个元素。
func (l *LinkedList) RemoveLast() (ele interface{}, suc bool) {
	if tempEle := l.list.Back(); tempEle != nil {
		ele = l.list.Remove(tempEle)
		suc = true
		return
	}
	return nil, false
}

// RemoveLastMulti removes amount elements from the end.
// RemoveLastMulti 从末尾起移除 amount 个元素。
func (l *LinkedList) RemoveLastMulti(amount int) (ele []interface{}, suc bool) {
	startIndex := l.Len() - amount
	return l.RemoveMultiAt(startIndex, amount)
}

// RemoveFirst removes the first element.
// RemoveFirst 移除第一个元素。
func (l *LinkedList) RemoveFirst() (ele interface{}, suc bool) {
	if tempEle := l.list.Front(); tempEle != nil {
		ele = l.list.Remove(tempEle)
		suc = true
		return
	}
	return nil, false
}

// RemoveFirstMulti removes amount elements from the front.
// RemoveFirstMulti 从开头起移除 amount 个元素。
func (l *LinkedList) RemoveFirstMulti(amount int) (ele []interface{}, suc bool) {
	return l.RemoveMultiAt(0, amount)
}

// Get returns the element at the specified index.
// Get 返回指定索引处的元素。
func (l *LinkedList) Get(index int) (ele interface{}, ok bool) {
	if tempEle, ok1 := l.getElement(index); ok1 {
		return tempEle.Value, true
	}
	return nil, false
}

// GetMulti returns amount elements starting at index.
// GetMulti 从 index 起返回 amount 个元素。
func (l *LinkedList) GetMulti(index int, amount int) (ele []interface{}, ok bool) {
	if amount <= 0 || index < 0 || index+amount > l.Len() {
		return nil, false
	}
	arr := make([]interface{}, amount)
	ele = arr[0:0]
	tempEle, _ := l.getElement(index)
	ele = append(ele, tempEle.Value)
	for i := 1; i < amount; i++ {
		tempEle = tempEle.Next()
		ele = append(ele, tempEle.Value)
	}
	return
}

// GetMultiLast returns amount elements ending at lastIndex.
// GetMultiLast 以 lastIndex 为终点返回 amount 个元素。
func (l *LinkedList) GetMultiLast(lastIndex int, amount int) (ele []interface{}, ok bool) {
	startIndex := lastIndex - amount + 1
	return l.GetMulti(startIndex, amount)
}

// GetAll returns all elements in the list.
// GetAll 返回列表中的全部元素。
func (l *LinkedList) GetAll() []interface{} {
	arr := make([]interface{}, l.Len())
	rs := arr[0:0]
	for e := l.list.Front(); e != nil; e = e.Next() {
		rs = append(rs, e.Value)
	}
	return rs
}

// First returns the first element.
// First 返回第一个元素。
func (l *LinkedList) First() (ele interface{}, ok bool) {
	return l.Get(0)
}

// FirstMulti returns amount elements from the front.
// FirstMulti 从开头返回 amount 个元素。
func (l *LinkedList) FirstMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMulti(0, amount)
}

// Last returns the last element.
// Last 返回最后一个元素。
func (l *LinkedList) Last() (ele interface{}, ok bool) {
	return l.Get(l.Len() - 1)
}

// LastMulti returns amount elements from the end.
// LastMulti 从末尾返回 amount 个元素。
func (l *LinkedList) LastMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMultiLast(l.Len()-1, amount)
}

// IndexOf returns the index of the first matching element.
// IndexOf 返回首个匹配元素的索引。
func (l *LinkedList) IndexOf(ele interface{}) (index int, ok bool) {
	index = -1
	each := func(eleIndex int, tempEle interface{}) (stop bool) {
		if lang.Equal(tempEle, ele) {
			index = eleIndex
			return true
		}
		return false
	}
	l.ForEach(each)
	ok = -1 != index
	return
}

// LastIndexOf returns the index of the last matching element.
// LastIndexOf 返回最后一个匹配元素的索引。
func (l *LinkedList) LastIndexOf(ele interface{}) (index int, ok bool) {
	index = -1
	each := func(eleIndex int, tempEle interface{}) (stop bool) {
		if lang.Equal(tempEle, ele) {
			index = eleIndex
			return true
		}
		return false
	}
	l.ForEachLast(each)
	ok = -1 != index
	return
}

// Contains reports whether the list contains ele.
// Contains 判断列表是否包含 ele。
func (l *LinkedList) Contains(ele interface{}) (contains bool) {
	_, contains = l.IndexOf(ele)
	return
}

// ForEach iterates elements from front to back until stop is true.
// ForEach 从前向后遍历元素，stop 为 true 时停止。
func (l *LinkedList) ForEach(each func(index int, ele interface{}) (stop bool)) {
	index := 0
	for tempEle := l.list.Front(); tempEle != nil; tempEle = tempEle.Next() {
		if each(index, tempEle.Value) {
			break
		}
		index++
	}
}

// ForEachLast iterates elements from back to front until stop is true.
// ForEachLast 从后向前遍历元素，stop 为 true 时停止。
func (l *LinkedList) ForEachLast(each func(index int, ele interface{}) (stop bool)) {
	index := l.list.Len() - 1
	for tempEle := l.list.Back(); tempEle != nil; tempEle = tempEle.Prev() {
		if each(index, tempEle.Value) {
			break
		}
		index--
	}
}

func (l *LinkedList) getElement(index int) (ele *list.Element, ok bool) {
	ln := l.Len()
	if 0 == ln || index < 0 || index >= ln {
		return nil, false
	}
	if 0 == index {
		ele = l.list.Front()
		ok = nil != ele
		return
	}
	if ln-1 == index {
		ele = l.list.Back()
		ok = nil != ele
		return
	}
	var i int
	if index <= ln>>1 {
		i = 0
		for tempEle := l.list.Front(); tempEle != nil; tempEle = tempEle.Next() {
			if i == index {
				return tempEle, true
			}
			i++
		}
	} else {
		i = ln - 1
		for tempEle := l.list.Back(); tempEle != nil; tempEle = tempEle.Prev() {
			if i == index {
				return tempEle, true
			}
			i--
		}
	}
	return nil, false
}

func (l *LinkedList) getElements(index int, amount int) (ele []*list.Element, ok bool) {
	ln := l.Len()
	if 0 == ln || amount <= 0 || index < 0 || index+amount > ln {
		return nil, false
	}
	arr := make([]*list.Element, amount)
	tempEle, _ := l.getElement(index)
	arr[0] = tempEle
	for i := 1; i < amount; i++ {
		ele := tempEle.Next() //使用copy
		arr[i] = ele
	}
	return arr[:], true
}
