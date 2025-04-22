// Package poolx
// Create on 2025/4/23
// @author xuzhuoxi
package poolx

import (
	"fmt"
	"testing"
)

func TestByteSlicePool(t *testing.T) {
	pool := NewByteSlicePool(1024)
	s1 := pool.Get()
	s2, err2 := pool.GetL(64)
	s3, err3 := pool.GetL(2048)
	fmt.Println("1", s1)
	fmt.Println("2", s2, err2)
	fmt.Println("3", s3, err3)
	*s1 = append(*s1, 1)
	(*s2)[0] = 2
	fmt.Println("1", s1)
	fmt.Println("2", s2, err2)
	pool.Put(s1)
	pool.Put(s2)
	pool.Put(s3)
}
