// Package binaryx
// Create on 2023/7/18
// @author xuzhuoxi
package binaryx

import (
	"encoding/binary"
	"errors"
	"io"
	"reflect"
)

// WriteSliceStringWithLen
// 向Writer中写入一组string数据[长度+数据]
// 格式：长度 + [长度+数据]...
// 长度格式默认为uint16
func WriteSliceStringWithLen(w io.Writer, order binary.ByteOrder, str []string) error {
	if err := writeLen(w, order, str); nil != err {
		return err
	}
	return WriteSliceString(w, order, str)
}

// WriteSliceString
// 向Writer中写入一组string数据[长度+数据]
// 格式：长度 + [长度+数据]...
// 长度格式默认为uint16
func WriteSliceString(w io.Writer, order binary.ByteOrder, str []string) error {
	for index := 0; index < len(str); index++ {
		if err := WriteString(w, order, str[index]); nil != err {
			return err
		}
	}
	return nil
}

func WriteSliceWithLen(w io.Writer, order binary.ByteOrder, data interface{}) error {
	if err := writeLen(w, order, data); nil != err {
		return err
	}
	return WriteSlice(w, order, data)
}

// WriteSlice
// 向Writer中写入一组数据
// 目前data支持的类型为[]bool,[]int,[]int8,[]int16,[]int32,[]int64,[]uint,[]uint8,[]uint16,[]uint32,[]uint64,[]float32,[]float64,[]complex64,[]complex128,[]string及相应指针类型
// int,uint类型会根据bit32值进行选择为int32,int64,uint32,uint64中一个
func WriteSlice(w io.Writer, order binary.ByteOrder, data interface{}) error {
	tempData := data
	if dataPtr, ok := data.(*interface{}); ok {
		tempData = *dataPtr
	}
	switch d := tempData.(type) {
	case []string:
		return WriteSliceString(w, order, d)
	case []int:
		if bit32 {
			var val []int32
			for _, v := range d {
				val = append(val, int32(v))
			}
			return binary.Write(w, order, val)
		} else {
			var val []int64
			for _, v := range d {
				val = append(val, int64(v))
			}
			return binary.Write(w, order, val)
		}
	case []uint:
		if bit32 {
			var val []uint32
			for _, v := range d {
				val = append(val, uint32(v))
			}
			return binary.Write(w, order, val)
		} else {
			var val []uint64
			for _, v := range d {
				val = append(val, uint64(v))
			}
			return binary.Write(w, order, val)
		}
	default:
		return binary.Write(w, order, tempData)
	}
}

func writeLen(w io.Writer, order binary.ByteOrder, data interface{}) error {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Array {
		return errors.New("data is not array! ")
	}
	v := reflect.ValueOf(data)
	if err := WriteLen(w, order, v.Len()); nil != err {
		return err
	}
	return nil
}
