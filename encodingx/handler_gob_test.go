//
//Created by xuzhuoxi
//on 2019-02-16.
//@author xuzhuoxi
//
package encodingx

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

func TestGobSameBuff(t *testing.T) {
	var buff1 bytes.Buffer
	encoder := gob.NewEncoder(&buff1)
	decoder := gob.NewDecoder(&buff1)

	encoder.Encode(data1)
	encoder.Encode(data2)
	encodedBytes := buff1.Bytes()

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
