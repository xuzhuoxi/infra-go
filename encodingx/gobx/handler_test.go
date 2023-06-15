// Package gobx
// Created by xuzhuoxi
// on 2019-02-16.
// @author xuzhuoxi
//
package gobx

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/xuzhuoxi/infra-go/slicex"
	"testing"
	"time"
)

type testData1 struct {
	A int
	B string
}

type testData2 struct {
	A string
	B int
}

type testData3 struct {
	A string
}

var (
	data1 = testData1{A: 1, B: "顶你个肺"}
	data2 = testData2{B: 1, A: "顶你个肺"}
	data3 = testData3{A: "顶你个肺哈哈"}
)

var (
	//包含TypeDescriptor
	bst1 = []byte{35, 255, 129, 3, 1, 1, 9, 116, 101, 115, 116, 68, 97, 116, 97, 49, 1, 255, 130, 0, 1, 2, 1, 1, 65, 1, 4, 0, 1, 1, 66, 1, 12, 0, 0, 0, 19, 255, 130, 1, 2, 1, 12, 233, 161, 182, 228, 189, 160, 228, 184, 170, 232, 130, 186, 0}
	bst2 = []byte{35, 255, 131, 3, 1, 1, 9, 116, 101, 115, 116, 68, 97, 116, 97, 50, 1, 255, 132, 0, 1, 2, 1, 1, 65, 1, 12, 0, 1, 1, 66, 1, 4, 0, 0, 0, 19, 255, 132, 1, 12, 233, 161, 182, 228, 189, 160, 228, 184, 170, 232, 130, 186, 1, 2, 0}
	bst3 = []byte{29, 255, 133, 3, 1, 1, 9, 116, 101, 115, 116, 68, 97, 116, 97, 51, 1, 255, 134, 0, 1, 1, 1, 1, 65, 1, 12, 0, 0, 0, 23, 255, 134, 1, 18, 233, 161, 182, 228, 189, 160, 228, 184, 170, 232, 130, 186, 229, 147, 136, 229, 147, 136, 0}
)

var (
	//不包含TypeDescriptor
	bs1 = []byte{19, 255, 130, 1, 2, 1, 12, 233, 161, 182, 228, 189, 160, 228, 184, 170, 232, 130, 186, 0}
	bs2 = []byte{19, 255, 132, 1, 12, 233, 161, 182, 228, 189, 160, 228, 184, 170, 232, 130, 186, 1, 2, 0}
	bs3 = []byte{23, 255, 134, 1, 18, 233, 161, 182, 228, 189, 160, 228, 184, 170, 232, 130, 186, 229, 147, 136, 229, 147, 136, 0}
)

func TestSysData(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	decoder := gob.NewEncoder(buff)
	decoder.Encode("顶你胩肺")
	fmt.Println(buff.Bytes())
	buff.Reset()
	decoder.Encode("顶你胩肺")
	fmt.Println(buff.Bytes())
	//系统基本数据类型不会增加TypeDescriptor
}

func TestDataChangeTypeDescriptor(t *testing.T) {
	buff := bytes.NewBuffer(bst1)
	buff.Write(bst2)
	decoder := gob.NewDecoder(buff)
	var d1 testData1
	var d2 testData2
	decoder.Decode(&d1)
	decoder.Decode(&d2)
	fmt.Println(d1, d2)
	buff.Write(bs1)
	var d11 testData1
	decoder.Decode(&d11)
	fmt.Println(d11)
}

func TestDataTypeDescriptor(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buff)
	decoder := gob.NewDecoder(buff)
	encoder.Encode(data1)
	buff.Reset()
	buff.Write(bs1)
	var d1 testData1
	decoder.Decode(&d1)
	fmt.Println(d1)
	//TypeDescriptor是分别记录在Encoder和Decoder中的,相互不影响
}

func TestDataNoTypeDescriptor(t *testing.T) {
	buffD := bytes.NewBuffer(bst1)
	buff := bytes.NewBuffer(bs1)
	decoderD := gob.NewDecoder(buffD)
	decoder := gob.NewDecoder(buff)
	var d1, d2 testData1
	decoderD.Decode(&d1)
	decoder.Decode(&d2)
	fmt.Println(d1, d2)
}

func TestDataFirstTypeDescriptor(t *testing.T) {
	buff := bytes.NewBuffer(bst1)
	buff.Write(bs1)
	decoder := gob.NewDecoder(buff)
	var d1, d2 testData1
	decoder.Decode(&d1)
	decoder.Decode(&d2)
	fmt.Println(d1, d2)
}

func TestGobSameBuff(t *testing.T) {
	var buff1 bytes.Buffer
	encoder := gob.NewEncoder(&buff1)
	decoder := gob.NewDecoder(&buff1)

	gob.Register(testData1{})

	encoder.Encode(data1)
	fmt.Println(buff1.Bytes())
	buff1.Reset()

	encoder.Encode(data1)
	fmt.Println(buff1.Bytes())
	buff1.Reset()

	encoder.Encode(data2)
	encodedBytes := buff1.Bytes()
	fmt.Println(buff1.Bytes())
	fmt.Println(encodedBytes)

	buff1.Write(encodedBytes)
	var rs1 testData1
	var rs2 testData2
	decoder.Decode(&rs1)
	decoder.Decode(&rs2)
	fmt.Println(data1, rs1, rs2)
}

func TestGobDiffBuff(t *testing.T) {
	var buff1, buff2 bytes.Buffer
	encoder := gob.NewEncoder(&buff1)
	decoder := gob.NewDecoder(&buff2)

	encoder.Encode(data1)
	encoder.Encode(data2)
	encodedBytes := buff1.Bytes()

	buff2.Write(encodedBytes)
	var rs1 testData1
	var rs2 testData2
	decoder.Decode(&rs1)
	decoder.Decode(&rs2)
	fmt.Println(data1, rs1, rs2)
}

func TestGobHandler(t *testing.T) {
	handler := NewGobCodingHandler()
	for index := 0; index < 3; index++ {
		bs1 := handler.HandleEncode(data1)
		var rs1 testData1
		handler.HandleDecode(bs1, &rs1)

		bs2 := handler.HandleEncode(data2)
		var rs2 testData2
		handler.HandleDecode(bs2, &rs2)

		fmt.Println(bs1, bs1)
		fmt.Println(rs1, rs2)
		fmt.Println("++++++++++++++")
	}
	//}()
	//go func() {
	for index := 0; index < 3; index++ {
		bs1 := slicex.CopyUint8(handler.HandleEncode(data1))
		bs2 := slicex.CopyUint8(handler.HandleEncode(data2))
		var rs1 testData1
		var rs2 testData2
		handler.HandleDecode(bs1, &rs1)
		handler.HandleDecode(bs2, &rs2)
		fmt.Println(bs1, bs2)
		fmt.Println(rs1, rs2)
		fmt.Println("-------------")
	}
	//}()
	time.Sleep(time.Second)
}
func TestGobHandlerSync(t *testing.T) {
	handler := NewGobCodingHandlerSync()
	var f1 = func() {
		for index := 0; index < 3; index++ {
			bs1 := handler.HandleEncode(data1)
			var rs1 testData1
			handler.HandleDecode(bs1, &rs1)

			bs2 := handler.HandleEncode(data2)
			var rs2 testData2
			handler.HandleDecode(bs2, &rs2)

			fmt.Println("++", bs1, bs2)
			fmt.Println("++", rs1, rs2)
			fmt.Println("++++++++++++++")
		}
	}
	var f2 = func() {
		for index := 0; index < 3; index++ {
			bs1 := slicex.CopyUint8(handler.HandleEncode(data1))
			bs2 := slicex.CopyUint8(handler.HandleEncode(data2))
			var rs1 testData1
			var rs2 testData2
			handler.HandleDecode(bs1, &rs1)
			handler.HandleDecode(bs2, &rs2)
			fmt.Println("--", bs1, bs2)
			fmt.Println("--", rs1, rs2)
			fmt.Println("-------------")
		}
	}
	go f1()
	go f2()
	time.Sleep(time.Second)
}

func TestGobHandlerAsync(t *testing.T) {
	handler := NewGobCodingHandlerAsync()
	var f1 = func() {
		for index := 0; index < 3; index++ {
			bs1 := handler.HandleEncode(data1)
			var rs1 testData1
			handler.HandleDecode(bs1, &rs1)

			bs2 := handler.HandleEncode(data2)
			var rs2 testData2
			handler.HandleDecode(bs2, &rs2)

			fmt.Println("++", bs1, bs2)
			fmt.Println("++", rs1, rs2)
			fmt.Println("++++++++++++++")
		}
	}
	var f2 = func() {
		for index := 0; index < 3; index++ {
			bs1 := handler.HandleEncode(data1)
			bs2 := handler.HandleEncode(data2)
			var rs1 testData1
			var rs2 testData2
			handler.HandleDecode(bs1, &rs1)
			handler.HandleDecode(bs2, &rs2)
			fmt.Println("--", bs1, bs2)
			fmt.Println("--", rs1, rs2)
			fmt.Println("-------------")
		}
	}
	go f1()
	go f2()
	time.Sleep(time.Second)
}

func TestInfoWithTypeDescriptor(t *testing.T) {
	data := []interface{}{data1, data2, data3}
	for index := 0; index < len(data); index++ {
		buff := bytes.NewBuffer(nil)
		encoder := gob.NewEncoder(buff)
		encoder.Encode(data[index])
		fmt.Println(buff.Bytes())
		buff.Reset()
		encoder.Encode(data[index])
		fmt.Println(buff.Bytes())
	}
}
