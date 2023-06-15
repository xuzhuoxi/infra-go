// Package encodingx
// Created by xuzhuoxi
// on 2019-03-20.
// @author xuzhuoxi
//
package encodingx

import (
	"fmt"
	"testing"
)

func TestCoding(t *testing.T) {
	cm := NewCodingMap()
	cm.Set("key0", false)
	cm.Set("key1", true)
	cm.Set("key2", 222)
	cm.Set("key3", 222.5)
	cm.Set("key4", "aaa，哈哈")
	cm.Set("key5", []bool{false, true})
	cm.Set("key6", []uint16{222, 333})
	cm.Set("key7", []int{222, 333})
	cm.Set("key8", []string{"", "o只"})

	bs := cm.EncodeToBytes()
	fmt.Println("序列化：", bs)

	var cm2 = NewCodingMap()
	cm2.DecodeFromBytes(bs)
	fmt.Println("反序列：", cm2)
}
