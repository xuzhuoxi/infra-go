//
//Created by xuzhuoxi
//on 2019-04-03.
//@author xuzhuoxi
//
package astar

import "github.com/xuzhuoxi/infra-go/lang/listx"

func newSortedQueue() *sortedQueue {
	return &sortedQueue{list: listx.NewLinkedList()}
}

func newQueue() *queue {
	return &queue{list: listx.NewArrayList(0, 1024)}
}

//----------------------------

type queue struct {
	list listx.IList
}

func (q *queue) Push(pos Position) {
	q.list.Add(pos)
}

func (q *queue) Shift() (pos Position, ok bool) {
	if e, ok := q.list.RemoveFirst(); ok {
		return e.(Position), true
	}
	return Position{}, false
}

//----------------------------

type sortedQueue struct {
	list listx.IList
}

func (q *sortedQueue) getAll() []interface{} {
	return q.list.GetAll()
}

//选择插入
func (q *sortedQueue) Push(x, y, priority int) {
	pp := NewPriorityPosition(x, y, priority)
	q.PushPriorityPosition(pp)
}

//选择插入
func (q *sortedQueue) PushPriorityPosition(pp *PriorityPosition) {
	if q.list.Len() == 0 {
		q.list.Add(pp)
		return
	}
	added := false
	q.list.ForEachLast(func(index int, ele interface{}) (stop bool) {
		if pp.Priority >= ele.(*PriorityPosition).Priority {
			q.list.AddAt(index+1, pp)
			added = true
			return true
		}
		return false
	})
	if !added {
		q.list.AddAt(0, pp)
	}
}

//取尾
func (q *sortedQueue) Pop() (pp *PriorityPosition, ok bool) {
	if last, ok := q.list.RemoveLast(); ok {
		return last.(*PriorityPosition), true
	}
	return nil, false
}

//取头
func (q *sortedQueue) Shift() (pp *PriorityPosition, ok bool) {
	if first, ok := q.list.RemoveFirst(); ok {
		return first.(*PriorityPosition), true
	}
	return nil, false
}

func (q *sortedQueue) Clear() {
	q.list.Clear()
}

func (q *sortedQueue) Len() int {
	return q.list.Len()
}
