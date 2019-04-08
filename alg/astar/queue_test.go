//
//Created by xuzhuoxi
//on 2019-04-05.
//@author xuzhuoxi
//
package astar

import (
	"fmt"
	"testing"
)

func TestSortedQueue_Push(t *testing.T) {
	queue := newSortedQueue()
	queue.Push(0, 0, 0)
	queue.Push(4, 0, 4)
	queue.Push(1, 0, 1)
	queue.Push(5, 0, 5)
	queue.Push(3, 0, 3)
	queue.Push(2, 0, 2)
	queue.Push(6, 0, 6)
	fmt.Println(queue.getAll())
	fmt.Println(queue.Len())
	fmt.Println(queue.Shift())
	fmt.Println(queue.getAll())
	fmt.Println(queue.Len())
	fmt.Println(queue.Pop())
	fmt.Println(queue.getAll())
	fmt.Println(queue.Len())
}
