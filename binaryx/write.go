// Package binaryx
// Create on 2023/7/18
// @author xuzhuoxi
package binaryx

import (
	"bytes"
	"encoding/binary"
	"io"
)

// WriteString
// 向Writer中写入一个string数据
// 格式：[长度+数据]
// 长度格式默认为uint16
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

// Write
// 向Writer中写入一个数据
// 目前data支持的类型为bool,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64,complex64,complex128,string及相应指针类型
// int,uint类型会根据bit32值进行选择为int32,int64,uint32,uint64中一个
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
