//
//Created by xuzhuoxi
//on 2019-04-03.
//@author xuzhuoxi
//
package listx

import (
	"github.com/xuzhuoxi/infra-go/lang"
)

// 实例化一个用数组实现的List
// ArrayList
// maxSize: 列表允许的最大长度
// initCap: 列表数组的初始化cap
// maxSize >= initCap
func NewArrayList(maxSize, initCap int) *ArrayList {
	if initCap > maxSize {
		return nil
	}
	rs := &ArrayList{maxSize: maxSize, initCap: initCap}
	rs.Clear()
	return rs
}

type ArrayList struct {
	maxSize  int
	initCap  int
	elements []interface{}
}

func (l *ArrayList) Len() int {
	return len(l.elements)
}

func (l *ArrayList) Swap(i, j int) {
	l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
}

func (l *ArrayList) Clear() {
	if l.initCap > 0 {
		l.elements = make([]interface{}, 0, l.initCap)
	} else {
		l.elements = nil
	}
}

func (l *ArrayList) Add(ele ...interface{}) (suc bool) {
	if l.maxSize > 0 && len(ele)+len(l.elements) > l.maxSize {
		return false
	}
	l.elements = append(l.elements, ele...)
	return true
}

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

func (l *ArrayList) AddAll(index int, list IList) (suc bool) {
	if all := list.GetAll(); all != nil {
		return l.AddAt(index, all...)
	}
	return true
}

func (l *ArrayList) RemoveAt(index int) (ele interface{}, suc bool) {
	if index >= 0 && index < l.Len() {
		rs := l.elements[index]
		l.elements = append(l.elements[:index], l.elements[index+1:]...)
		return rs, true
	}
	return nil, false
}

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

func (l *ArrayList) RemoveLast() (ele interface{}, suc bool) {
	return l.RemoveAt(l.Len() - 1)
}

func (l *ArrayList) RemoveLastMulti(amount int) (ele []interface{}, suc bool) {
	startIndex := l.Len() - amount
	return l.RemoveMultiAt(startIndex, amount)
}

func (l *ArrayList) RemoveFirst() (ele interface{}, suc bool) {
	return l.RemoveAt(0)
}

func (l *ArrayList) RemoveFirstMulti(amount int) (ele []interface{}, suc bool) {
	return l.RemoveMultiAt(0, amount)
}

func (l *ArrayList) Get(index int) (ele interface{}, ok bool) {
	ln := l.Len()
	if ln > 0 && index >= 0 && index < ln {
		return l.elements[index], true
	}
	return nil, false
}

func (l *ArrayList) GetMulti(index int, amount int) (ele []interface{}, ok bool) {
	if amount > 0 && index >= 0 && index+amount <= l.Len() {
		return l.elements[index : index+amount], true
	}
	return nil, false
}

func (l *ArrayList) GetMultiLast(lastIndex int, amount int) (ele []interface{}, ok bool) {
	startIndex := lastIndex - amount + 1
	return l.GetMulti(startIndex, amount)
}

func (l *ArrayList) GetAll() []interface{} {
	return l.elements
}

func (l *ArrayList) First() (ele interface{}, ok bool) {
	return l.Get(0)
}

func (l *ArrayList) FirstMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMulti(0, amount)
}

func (l *ArrayList) Last() (ele interface{}, ok bool) {
	return l.Get(l.Len() - 1)
}

func (l *ArrayList) LastMulti(amount int) (ele []interface{}, ok bool) {
	return l.GetMultiLast(l.Len()-1, amount)
}

func (l *ArrayList) IndexOf(ele interface{}) (index int, ok bool) {
	ln := len(l.elements)
	for index := 0; index < ln; index++ {
		if lang.Equal(l.elements[index], ele) {
			return index, true
		}
	}
	return -1, false
}

func (l *ArrayList) LastIndexOf(ele interface{}) (index int, ok bool) {
	for index := len(l.elements) - 1; index >= 0; index-- {
		if lang.Equal(ele, l.elements[index]) {
			return index, true
		}
	}
	return -1, false
}

func (l *ArrayList) Contains(ele interface{}) (contains bool) {
	_, contains = l.IndexOf(ele)
	return
}

func (l *ArrayList) ForEach(each func(index int, ele interface{}) (stop bool)) {
	ln := len(l.elements)
	for index := 0; index < ln; index++ {
		if each(index, l.elements[index]) {
			break
		}
	}
}

func (l *ArrayList) ForEachLast(each func(index int, ele interface{}) (stop bool)) {
	for index := len(l.elements) - 1; index >= 0; index-- {
		if each(index, l.elements[index]) {
			break
		}
	}
}
