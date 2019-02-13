//
//Created by xuzhuoxi
//on 2019-02-13.
//@author xuzhuoxi
//
package encodingx

import "unsafe"

//-------------------------------------

type StringCodeHandler struct {
}

func (h *StringCodeHandler) HandleEncode(data interface{}) []byte {
	return StringToByteArray(data.(string))
}

func (h *StringCodeHandler) HandleDecode(bytes []byte) interface{} {
	return ByteArrayToString(bytes)
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
