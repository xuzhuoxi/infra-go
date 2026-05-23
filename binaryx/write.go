package binaryx

import (
	"bytes"
	"encoding/binary"
	"io"
)

// WriteString writes a length-prefixed string to w.
// WriteString 向 Writer 中写入一个 string 数据。
// Format: [length + payload]; length is uint16 by default.
// 格式：[长度 + 数据]；长度格式默认为 uint16。
func WriteString(w io.Writer, order binary.ByteOrder, str string) error {
	buff := bytes.NewBuffer(nil)
	if err := WriteLen(buff, order, len(str)); nil != err {
		return err
	}
	if _, err := buff.Write([]byte(str)); nil != err {
		return err
	}
	if _, err := w.Write(buff.Bytes()); nil != err {
		return err
	}
	return nil
}

// Write writes data to w.
// Write 向 Writer 中写入一个数据。
// Currently supported types for data: bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128, string and their pointer types; int and uint follow bit32 (int32/int64 or uint32/uint64).
// 目前 data 支持的类型为 bool、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、float32、float64、complex64、complex128、string 及相应指针类型；int、uint 类型会根据 bit32 值选择 int32/int64 或 uint32/uint64。
func Write(w io.Writer, order binary.ByteOrder, data interface{}) error {
	tempData := data
	if dataPtr, ok := data.(*interface{}); ok {
		tempData = *dataPtr
	}
	switch d := tempData.(type) {
	case string:
		return WriteString(w, order, d)
	case int:
		if bit32 {
			return binary.Write(w, order, int32(d))
		} else {
			return binary.Write(w, order, int64(d))
		}
	case uint:
		if bit32 {
			return binary.Write(w, order, uint32(d))
		} else {
			return binary.Write(w, order, uint64(d))
		}
	default:
		return binary.Write(w, order, tempData)
	}
}
