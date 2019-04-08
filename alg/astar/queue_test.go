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
	queue.Push2D(0, 0, 0)
	queue.Push2D(4, 0, 4)
	queue.Push2D(1, 0, 1)
	queue.Push2D(5, 0, 5)
	queue.Push2D(3, 0, 3)
	queue.Push2D(2, 0, 2)
	queue.Push2D(6, 0, 6)
	fmt.Println(queue.getAll())
	fmt.Println(queue.Len())
	fmt.Println(queue.Shift())
	fmt.Println(queue.getAll())
	fmt.Println(queue.Len())
	fmt.Println(queue.Pop())
	fmt.Println(queue.getAll())
	fmt.Println(queue.Len())
}
