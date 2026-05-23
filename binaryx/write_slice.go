package binaryx

import (
	"encoding/binary"
	"errors"
	"io"
	"reflect"
)

// WriteSliceStringWithLen writes the slice length followed by length-prefixed strings to w.
// WriteSliceStringWithLen 向 Writer 中写入一组 string 数据 [长度 + 数据]。
// Format: count + [length + payload]...; length fields use uint16 by default.
// 格式：长度 + [长度 + 数据]...；长度格式默认为 uint16。
func WriteSliceStringWithLen(w io.Writer, order binary.ByteOrder, str []string) error {
	if err := writeLen(w, order, str); nil != err {
		return err
	}
	return WriteSliceString(w, order, str)
}

// WriteSliceString writes length-prefixed strings to w without a leading slice length.
// WriteSliceString 向 Writer 中写入一组 string 数据 [长度 + 数据]。
// Format: [length + payload] for each element; length fields use uint16 by default.
// 格式：[长度 + 数据]...；长度格式默认为 uint16。
func WriteSliceString(w io.Writer, order binary.ByteOrder, str []string) error {
	for index := 0; index < len(str); index++ {
		if err := WriteString(w, order, str[index]); nil != err {
			return err
		}
	}
	return nil
}

// WriteSliceWithLen writes the slice length followed by slice payload to w.
// WriteSliceWithLen 向 Writer 中写入一组数据，先写入长度再写入元素。
func WriteSliceWithLen(w io.Writer, order binary.ByteOrder, data interface{}) error {
	if err := writeLen(w, order, data); nil != err {
		return err
	}
	return WriteSlice(w, order, data)
}

// WriteSlice writes slice data to w without a leading slice length.
// WriteSlice 向 Writer 中写入一组数据。
// Currently supported types for data: []bool, []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float32, []float64, []complex64, []complex128, []string and their pointer types; []int and []uint follow bit32 (int32/int64 or uint32/uint64).
// 目前 data 支持的类型为 []bool、[]int、[]int8、[]int16、[]int32、[]int64、[]uint、[]uint8、[]uint16、[]uint32、[]uint64、[]float32、[]float64、[]complex64、[]complex128、[]string 及相应指针类型；[]int、[]uint 类型会根据 bit32 值选择 int32/int64 或 uint32/uint64。
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
