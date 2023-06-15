// Package gobx
// Created by xuzhuoxi
// on 2019-02-22.
// @author xuzhuoxi
//
package gobx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/slicex"
	"testing"
)

type A struct {
	Name  string
	Value int
}

func TestCodecsGob(t *testing.T) {
	a := A{Name: "哈哈", Value: 234}
	encoder := NewDefaultGobBuffEncoder()
	decoder := NewDefaultGobBuffDecoder()

	encoder.EncodeDataToBuff(a)
	bs1 := encoder.ReadBytes()
	fmt.Println("E:", bs1)

	decoder.WriteBytes(bs1)
	var a1 A
	decoder.DecodeDataFromBuff(&a1)
	fmt.Println("D:", a1)

	fmt.Println("------------")

	encoder.EncodeDataToBuff(a)
	bs2 := encoder.ReadBytes()
	fmt.Println("E:", bs2)

	decoder.WriteBytes(bs2)
	var a2 A
	decoder.DecodeDataFromBuff(&a2)
	fmt.Println("D:", a2)
}

func TestCodecsGob2(t *testing.T) {
	a := A{Name: "哈哈", Value: 234}
	encoder := NewDefaultGobBuffEncoder()
	decoder := NewDefaultGobBuffDecoder()

	encoder.EncodeDataToBuff(a)
	bs1 := slicex.CopyUint8(encoder.ReadBytes())
	fmt.Println("E1:", bs1)

	encoder.EncodeDataToBuff(a)
	bs2 := slicex.CopyUint8(encoder.ReadBytes())
	fmt.Println("E2:", bs2)

	fmt.Println("------------")

	decoder.WriteBytes(bs1)
	decoder.WriteBytes(bs2)

	var a1 A
	decoder.DecodeDataFromBuff(&a1)
	fmt.Println("D:", a1)

	var a2 A
	decoder.DecodeDataFromBuff(&a2)
	fmt.Println("D:", a2)
}
