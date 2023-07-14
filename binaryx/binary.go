// Package binaryx
// Created by xuzhuoxi
// on 2019-03-19.
// @author xuzhuoxi
//
package binaryx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// ReadBool
// 从一个Reader中读取一个bool数据
// r: Reader
// order: 大小端设定
func ReadBool(r io.Reader, order binary.ByteOrder) (bool, error) {
	var rs bool
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceBool
// 从一个Reader中读取一组bool数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceBool(r io.Reader, order binary.ByteOrder, ln int) ([]bool, error) {
	var rs = make([]bool, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadInt
// 从一个Reader中读取一个Int数据, 32位则读取int32,64位则读取int64
// r: Reader
// order: 大小端设定
func ReadInt(r io.Reader, order binary.ByteOrder) (int, error) {
	if bit32 {
		val, err := ReadInt32(r, order)
		return int(val), err
	} else {
		val, err := ReadInt64(r, order)
		return int(val), err
	}
}

// ReadSliceInt
// 从一个Reader中读取一组Int数据, 32位则读取int32,64位则读取int64
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt(r io.Reader, order binary.ByteOrder, ln int) ([]int, error) {
	var rs []int
	if bit32 {
		val, err := ReadSliceInt32(r, order, ln)
		for _, v := range val {
			rs = append(rs, int(v))
		}
		return rs, err
	} else {
		val, err := ReadSliceInt64(r, order, ln)
		for _, v := range val {
			rs = append(rs, int(v))
		}
		return rs, err
	}
}

// ReadInt8
//从一个Reader中读取一个Int8数据
// r: Reader
// order: 大小端设定
func ReadInt8(r io.Reader, order binary.ByteOrder) (int8, error) {
	var rs int8
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt8
// 从一个Reader中读取一组Int8数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt8(r io.Reader, order binary.ByteOrder, ln int) ([]int8, error) {
	var rs = make([]int8, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadInt16
//从一个Reader中读取一个Int16数据
// r: Reader
// order: 大小端设定
func ReadInt16(r io.Reader, order binary.ByteOrder) (int16, error) {
	var rs int16
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt16
// 从一个Reader中读取一组Int16数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt16(r io.Reader, order binary.ByteOrder, ln int) ([]int16, error) {
	var rs = make([]int16, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadInt32
// 从一个Reader中读取一个Int32数据
// r: Reader
// order: 大小端设定
func ReadInt32(r io.Reader, order binary.ByteOrder) (int32, error) {
	var rs int32
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt32
// 从一个Reader中读取一组Int32数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt32(r io.Reader, order binary.ByteOrder, ln int) ([]int32, error) {
	var rs = make([]int32, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadInt64
// 从一个Reader中读取一个Int64数据
// r: Reader
// order: 大小端设定
func ReadInt64(r io.Reader, order binary.ByteOrder) (int64, error) {
	var rs int64
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt64
// 从一个Reader中读取一组Int64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt64(r io.Reader, order binary.ByteOrder, ln int) ([]int64, error) {
	var rs = make([]int64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUint
// 从一个Reader中读取一个Uint数据, 32位则读取unt32,64位则读取unt64
// r: Reader
// order: 大小端设定
func ReadUint(r io.Reader, order binary.ByteOrder) (uint, error) {
	if bit32 {
		val, err := ReadUint32(r, order)
		return uint(val), err
	} else {
		val, err := ReadUint64(r, order)
		return uint(val), err
	}
}

// ReadSliceUint
// 从一个Reader中读取一组Uint数据, 32位则读取unt32,64位则读取unt64
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUint(r io.Reader, order binary.ByteOrder, ln int) ([]uint, error) {
	var rs []uint
	if bit32 {
		val, err := ReadSliceUint32(r, order, ln)
		for _, v := range val {
			rs = append(rs, uint(v))
		}
		return rs, err
	} else {
		val, err := ReadSliceUint64(r, order, ln)
		for _, v := range val {
			rs = append(rs, uint(v))
		}
		return rs, err
	}
}

// ReadUint8
// 从一个Reader中读取一个Uint8数据
// r: Reader
// order: 大小端设定
func ReadUint8(r io.Reader, order binary.ByteOrder) (uint8, error) {
	var rs uint8
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUint8
// 从一个Reader中读取一组Uint8数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUint8(r io.Reader, order binary.ByteOrder, ln int) ([]uint8, error) {
	var rs = make([]uint8, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUint16
// 从一个Reader中读取一个Uint16数据
// r: Reader
// order: 大小端设定
func ReadUint16(r io.Reader, order binary.ByteOrder) (uint16, error) {
	var rs uint16
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUint16
//从一个Reader中读取一组Uint16数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUint16(r io.Reader, order binary.ByteOrder, ln int) ([]uint16, error) {
	var rs = make([]uint16, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUint32
//从一个Reader中读取一个Uint32数据
// r: Reader
// order: 大小端设定
func ReadUint32(r io.Reader, order binary.ByteOrder) (uint32, error) {
	var rs uint32
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUint32
// 从一个Reader中读取一组Uint32数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUint32(r io.Reader, order binary.ByteOrder, ln int) ([]uint32, error) {
	var rs = make([]uint32, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUint64
// 从一个Reader中读取一个Uint64数据
// r: Reader
// order: 大小端设定
func ReadUint64(r io.Reader, order binary.ByteOrder) (uint64, error) {
	var rs uint64
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUint64
// 从一个Reader中读取一组Uint64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUint64(r io.Reader, order binary.ByteOrder, ln int) ([]uint64, error) {
	var rs = make([]uint64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadFloat32
// 从一个Reader中读取一个Float32数据
// r: Reader
// order: 大小端设定
func ReadFloat32(r io.Reader, order binary.ByteOrder) (float32, error) {
	var rs float32
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceFloat32
// 从一个Reader中读取一组Float32数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceFloat32(r io.Reader, order binary.ByteOrder, ln int) ([]float32, error) {
	var rs = make([]float32, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadFloat64
// 从一个Reader中读取一个Float64数据
// r: Reader
// order: 大小端设定
func ReadFloat64(r io.Reader, order binary.ByteOrder) (float64, error) {
	var rs float64
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceFloat64
// 从一个Reader中读取一组Float64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceFloat64(r io.Reader, order binary.ByteOrder, ln int) ([]float64, error) {
	var rs = make([]float64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadComplex64
// 从一个Reader中读取一个Complex64数据
// r: Reader
// order: 大小端设定
func ReadComplex64(r io.Reader, order binary.ByteOrder) (complex64, error) {
	var rs complex64
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceComplex64
// 从一个Reader中读取一组Complex64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceComplex64(r io.Reader, order binary.ByteOrder, ln int) ([]complex64, error) {
	var rs = make([]complex64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadComplex128
// 从一个Reader中读取一个Complex128数据
// r: Reader
// order: 大小端设定
func ReadComplex128(r io.Reader, order binary.ByteOrder) (complex128, error) {
	var rs complex128
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceComplex128
// 从一个Reader中读取一组Complex128数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceComplex128(r io.Reader, order binary.ByteOrder, ln int) ([]complex128, error) {
	var rs = make([]complex128, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadString
// 从一个Reader中读取一个string数据
// r: Reader
// order: 大小端设定
func ReadString(r io.Reader, order binary.ByteOrder) (string, error) {
	var ln int
	var err error
	if ln, err = ReadLen(r, order); nil != err {
		return "", err
	}
	var bs = make([]byte, ln)
	if err := binary.Read(r, order, &bs); nil != err {
		return "", err
	}
	return string(bs), nil
}

// ReadSliceString
//从一个Reader中读取一组string数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceString(r io.Reader, order binary.ByteOrder, ln int) ([]string, error) {
	var rs []string
	for index := ln - 1; index >= 0; index-- {
		str, err := ReadString(r, order)
		if nil != err {
			return nil, err
		}
		rs = append(rs, str)
	}
	return rs, nil
}

// Read
// 从Reader中读取数据，
// 目前data支持的类型为*bool,*int,*int8,*int16,*int32,*int64,*uint,*uint8,*uint16,*uint32,*uint64,*float32,*float64,*complex64,*complex128,*string
func Read(r io.Reader, order binary.ByteOrder, data interface{}) (err error) {
	if dataPtr, ok := data.(*interface{}); ok {
		isCatch := true
		switch (*dataPtr).(type) {
		case bool:
			*dataPtr, err = ReadBool(r, order)
		case int:
			*dataPtr, err = ReadInt(r, order)
		case int8:
			*dataPtr, err = ReadInt8(r, order)
		case int16:
			*dataPtr, err = ReadInt16(r, order)
		case int32:
			*dataPtr, err = ReadInt32(r, order)
		case int64:
			*dataPtr, err = ReadInt64(r, order)
		case uint:
			*dataPtr, err = ReadUint(r, order)
		case uint8:
			*dataPtr, err = ReadUint8(r, order)
		case uint16:
			*dataPtr, err = ReadUint16(r, order)
		case uint32:
			*dataPtr, err = ReadUint32(r, order)
		case uint64:
			*dataPtr, err = ReadUint64(r, order)
		case float32:
			*dataPtr, err = ReadFloat32(r, order)
		case float64:
			*dataPtr, err = ReadFloat64(r, order)
		case complex64:
			*dataPtr, err = ReadComplex64(r, order)
		case complex128:
			*dataPtr, err = ReadComplex128(r, order)
		case string:
			*dataPtr, err = ReadString(r, order)
		default:
			isCatch = false
		}
		if isCatch {
			return
		}
	}
	return binary.Read(r, order, data)
}

// ReadSlice
// 从Reader中读取数据，
// 目前data支持的类型为*[]bool,*[]int,*[]int8,*[]int16,*[]int32,*[]int64,*[]uint,*[]uint8,*[]uint16,*[]uint32,*[]uint64,*[]float32,*[]float64,*[]complex64,*[]complex128,*[]string
func ReadSlice(r io.Reader, order binary.ByteOrder, data interface{}, ln int) (err error) {
	if dataPtr, ok := data.(*interface{}); ok {
		isCatch := true
		switch (*dataPtr).(type) {
		case []bool:
			*dataPtr, err = ReadSliceBool(r, order, ln)
		case []int:
			*dataPtr, err = ReadSliceInt(r, order, ln)
		case []int8:
			*dataPtr, err = ReadSliceInt8(r, order, ln)
		case []int16:
			*dataPtr, err = ReadSliceInt16(r, order, ln)
		case []int32:
			*dataPtr, err = ReadSliceInt32(r, order, ln)
		case []int64:
			*dataPtr, err = ReadSliceInt64(r, order, ln)
		case []uint:
			*dataPtr, err = ReadSliceUint(r, order, ln)
		case []uint8:
			*dataPtr, err = ReadSliceUint8(r, order, ln)
		case []uint16:
			*dataPtr, err = ReadSliceUint16(r, order, ln)
		case []uint32:
			*dataPtr, err = ReadSliceUint32(r, order, ln)
		case []uint64:
			*dataPtr, err = ReadSliceUint64(r, order, ln)
		case []float32:
			*dataPtr, err = ReadSliceFloat32(r, order, ln)
		case []float64:
			*dataPtr, err = ReadSliceFloat64(r, order, ln)
		case []complex64:
			*dataPtr, err = ReadSliceComplex64(r, order, ln)
		case []complex128:
			*dataPtr, err = ReadSliceComplex128(r, order, ln)
		case []string:
			*dataPtr, err = ReadSliceString(r, order, ln)
		default:
			isCatch = false
		}
		if isCatch {
			return
		}
	}
	return binary.Read(r, order, data)
}

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

// WriteSliceString
// 向Writer中写入一组string数据[长度+数据]
// 格式：长度 + [长度+数据]...
// 长度格式默认为uint16
func WriteSliceString(w io.Writer, order binary.ByteOrder, str []string) error {
	buff := bytes.NewBuffer(nil)
	for index := 0; index < len(str); index++ {
		if err := WriteString(buff, order, str[index]); nil != err {
			return err
		}
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

//-------------------------------------

// ReadLen
// 从Reader中读取一个长度数据
// 按uint16格式读取，强制转换为int值返回
// r: Reader
// order: 大小端设定
func ReadLen(r io.Reader, order binary.ByteOrder) (int, error) {
	var ln uint16
	if err := ReadLenTo(r, order, &ln); nil != err {
		return 0, err
	}
	return int(ln), nil
}

// ReadLenTo
// 从Reader中读取一个长度数据
// r: Reader
// order: 大小端设定
// ln: 长度值，只支持*int8, *int16, *int32, *int64, *uint8, *uint16, *uint32, *uint64
func ReadLenTo(r io.Reader, order binary.ByteOrder, ln interface{}) error {
	switch ln.(type) {
	case *int8, *int16, *int32, *int64:
		return binary.Read(r, order, ln)
	case *uint8, *uint16, *uint32, *uint64:
		return binary.Read(r, order, ln)
	}
	return errors.New("ln type error! ")
}

// WriteLen
//把长度写入到Writer中
// w: Writer
// order: 大小端设定
// ln: 长度值，按uint16写入
func WriteLen(w io.Writer, order binary.ByteOrder, ln int) error {
	return binary.Write(w, order, uint16(ln))
}

// WriteLenTo
// 把长度写入到Writer中
// w: Writer
// order: 大小端设定
// ln: 长度值，只支持int8, int16, int32, int64, uint8, uint16, uint32, uint64及其指针类型
func WriteLenTo(w io.Writer, order binary.ByteOrder, ln interface{}) error {
	switch ln.(type) {
	case int8, int16, int32, int64, *int8, *int16, *int32, *int64:
		return binary.Write(w, order, ln)
	case uint8, uint16, uint32, uint64, *uint8, *uint16, *uint32, *uint64:
		return binary.Write(w, order, ln)
	}
	return errors.New("ln type error! ")
}
