//
//Created by xuzhuoxi
//on 2019-03-23.
//@author xuzhuoxi
//
package lang

import (
	"fmt"
	"testing"
)

type testPoolObj struct {
	A int
}

func TestNewPool(t *testing.T) {
	pool := NewObjectPoolAsync()
	pool.Register(func() interface{} {
		return &testPoolObj{A: 1}
	}, func(instance interface{}) bool {
		if _, ok := instance.(*testPoolObj); ok {
			return true
		}
		return false
	})
	obj1 := pool.GetInstance().(*testPoolObj)
	fmt.Println(obj1)
	obj1.A = 2
	pool.Recycle(obj1)
}
