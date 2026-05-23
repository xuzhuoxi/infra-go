package encodingx

import (
	"errors"
	"unsafe"
)

var (
	// ErrNotString is returned when data is not a string.
	// ErrNotString 在数据不是 string 类型时返回。
	ErrNotString = errors.New("data is not string type! ")
)

// NewUtf8StringCodingHandle returns a UTF-8 string ICodingHandler.
// NewUtf8StringCodingHandle 返回 UTF-8 字符串编解码处理器。
func NewUtf8StringCodingHandle() ICodingHandler {
	return &utf8StringHandler{}
}

//-------------------------------

type utf8StringHandler struct {
}

func (*utf8StringHandler) HandleEncode(data interface{}) (bs []byte, err error) {
	if str, ok := data.(string); ok {
		return []byte(str), nil
	}
	return nil, ErrNotString
}

func (*utf8StringHandler) HandleDecode(bytes []byte, data interface{}) error {
	if _, ok := data.(*string); ok {
		data = string(bytes)
		return nil
	}
	return ErrNotString
}

//-------------------------------

// StringToByteArray converts a string to a byte slice without copying (unsafe).
// StringToByteArray 将字符串转为字节切片（unsafe，无拷贝）。
func StringToByteArray(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h1 := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h1))
}

// ByteArrayToString converts a byte slice to a string without copying (unsafe).
// ByteArrayToString 将字节切片转为字符串（unsafe，无拷贝）。
func ByteArrayToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
