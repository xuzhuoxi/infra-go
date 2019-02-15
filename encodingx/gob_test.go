//
//Created by xuzhuoxi
//on 2019-02-16.
//@author xuzhuoxi
//
package encodingx

import (
	"fmt"
	"testing"
)

type testData struct {
	A int
	B string
}

func TestGobHandler(t *testing.T) {
	for index := 0; index < 3; index++ {
		td := testData{A: 1, B: "顶你个肺"}
		handler := NewGobCodingHandler()
		encode := handler.HandleEncode(td)
		fmt.Println(encode)
		var rs testData
		fmt.Println(rs)
		handler.HandleDecode(encode, &rs)
		fmt.Println(rs)
	}
}

func TestGobHandlerSync(t *testing.T) {
	for index := 0; index < 3; index++ {
		td := testData{A: 2, B: "顶你个肺"}
		handler := NewGobCodingHandlerSync()
		encode := handler.HandleEncode(td)
		fmt.Println(encode)
		var rs testData
		fmt.Println(rs)
		handler.HandleDecode(encode, &rs)
		fmt.Println(rs)
	}
}

func TestGobHandlerAsync(t *testing.T) {
	for index := 0; index < 3; index++ {
		td := testData{A: 3, B: "顶你个肺"}
		handler := NewGobCodingHandlerAsync()
		encode := handler.HandleEncode(td)
		fmt.Println(encode)
		var rs testData
		fmt.Println(rs)
		handler.HandleDecode(encode, &rs)
		fmt.Println(rs)
	}
}
