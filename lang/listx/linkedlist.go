// Package listx
// Created by xuzhuoxi
// on 2019-04-03.
// @author xuzhuoxi
//
package listx

import (
	"container/list"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/mathx"
)

// NewLinkedList
// 实例化一个基于链表实现的List
// LinkedList
func NewLinkedList() *LinkedList {
	rs := &LinkedList{list: list.New()}
	return rs
}

type LinkedList struct {
	list *list.List
}

func (l *LinkedList) Len() int {
	return l.list.Len()
}

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

func (l *LinkedList) Clear() {
	l.list.Init()
}

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

func (l *LinkedList) AddAll(index int, list IList) (suc bool) {
	if nil == list || list.Len() == 0 {
		return false
	}
	return l.AddAt(index, list.GetAll()...)
}

func (l *LinkedList) RemoveAt(index int) (ele interface{}, suc bool) {
	if tempEle, ok := l.getElement(index); ok {
		ele = l.list.Remove(tempEle)
		suc = true
		return
	}
	return
}

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

func (l *LinkedList) RemoveLast() (ele interface{}, suc bool) {
	if tempEle := l.list.Back(); tempEle != nil {
		ele = l.list.Remove(tempEle)
		suc = true
		return
	}
	return nil, false
}

func (l *LinkedList) RemoveLastMulti(amount int) (ele []interface{}, suc bool) {
	startIndex := l.Len() - amount
	return l.RemoveMultiAt(startIndex, amount)
}

func (l *LinkedList) RemoveFirst() (ele interface{}, suc bool) {
	if tempEle := l.list.Front(); tempEle != nil {
		ele = l.list.Remove(tempEle)
		suc = true
		return
	}
	return nil, false
}

func (l *LinkedList) RemoveFirstMulti(amount int) (ele []interface{}, suc bool) {
	return l.RemoveMultiAt(0, amount)
}

func (l *LinkedList) Get(index int) (ele interface{}, ok bool) {
	if tempEle, ok1 := l.getElement(index); ok1 {
		return tempEle.Value, true
	}
	return nil, false
}

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

func (l *LinkedList) GetMultiLast(lastIndex int, amount int) (ele []interface{}, ok bool) {
	startIndex := lastIndex - amount + 1
	return l.GetMulti(startIndex, amount)
}

func (l *LinkedList) GetAll() []interface{} {
	arr := make([]interface{}, l.Len())
	rs := arr[0:0]
	for e := l.list.Front(); e != nil; e = e.Next() {
		rs = append(rs, e.Value)
	}
	return rs
}

func (l *LinkedList) First() (ele interface{}, ok bool) {
	return l.Get(0)
}

func (l *LinkedList) FirstMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMulti(0, amount)
}

func (l *LinkedList) Last() (ele interface{}, ok bool) {
	return l.Get(l.Len() - 1)
}

func (l *LinkedList) LastMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMultiLast(l.Len()-1, amount)
}

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

func (l *LinkedList) Contains(ele interface{}) (contains bool) {
	_, contains = l.IndexOf(ele)
	return
}

func (l *LinkedList) ForEach(each func(index int, ele interface{}) (stop bool)) {
	index := 0
	for tempEle := l.list.Front(); tempEle != nil; tempEle = tempEle.Next() {
		if each(index, tempEle.Value) {
			break
		}
		index++
	}
}

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
