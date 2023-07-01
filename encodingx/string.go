// Package encodingx
// Created by xuzhuoxi
// on 2019-02-13.
// @author xuzhuoxi
//
package encodingx

import (
	"errors"
	"unsafe"
)

func NewUtf8StringCodingHandle() ICodingHandler {
	return &utf8StringHandler{}
}

//-------------------------------

type utf8StringHandler struct {
}

func (*utf8StringHandler) HandleEncode(data interface{}) []byte {
	if str, ok := data.(string); ok {
		return []byte(str)
	}
	return nil
}

func (*utf8StringHandler) HandleDecode(bytes []byte, data interface{}) error {
	if _, ok := data.(*string); ok {
		data = string(bytes)
		return nil
	}
	return errors.New("data is not string type! ")
}

//-------------------------------

func StringToByteArray(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h1 := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h1))
}

func ByteArrayToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
