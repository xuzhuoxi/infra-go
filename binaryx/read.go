// Package binaryx
// Create on 2023/7/18
// @author xuzhuoxi
package binaryx

import (
	"encoding/binary"
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

// ReadInt8
//从一个Reader中读取一个Int8数据
// r: Reader
// order: 大小端设定
func ReadInt8(r io.Reader, order binary.ByteOrder) (int8, error) {
	var rs int8
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

// ReadInt32
// 从一个Reader中读取一个Int32数据
// r: Reader
// order: 大小端设定
func ReadInt32(r io.Reader, order binary.ByteOrder) (int32, error) {
	var rs int32
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

// ReadUInt8
// 从一个Reader中读取一个Uint8数据
// r: Reader
// order: 大小端设定
func ReadUInt8(r io.Reader, order binary.ByteOrder) (uint8, error) {
	var rs uint8
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUInt16
// 从一个Reader中读取一个Uint16数据
// r: Reader
// order: 大小端设定
func ReadUInt16(r io.Reader, order binary.ByteOrder) (uint16, error) {
	var rs uint16
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUInt32
//从一个Reader中读取一个Uint32数据
// r: Reader
// order: 大小端设定
func ReadUInt32(r io.Reader, order binary.ByteOrder) (uint32, error) {
	var rs uint32
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUInt64
// 从一个Reader中读取一个Uint64数据
// r: Reader
// order: 大小端设定
func ReadUInt64(r io.Reader, order binary.ByteOrder) (uint64, error) {
	var rs uint64
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadUInt
// 从一个Reader中读取一个Uint数据, 32位则读取unt32,64位则读取unt64
// r: Reader
// order: 大小端设定
func ReadUInt(r io.Reader, order binary.ByteOrder) (uint, error) {
	if bit32 {
		val, err := ReadUInt32(r, order)
		return uint(val), err
	} else {
		val, err := ReadUInt64(r, order)
		return uint(val), err
	}
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

// ReadFloat64
// 从一个Reader中读取一个Float64数据
// r: Reader
// order: 大小端设定
func ReadFloat64(r io.Reader, order binary.ByteOrder) (float64, error) {
	var rs float64
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

// ReadComplex128
// 从一个Reader中读取一个Complex128数据
// r: Reader
// order: 大小端设定
func ReadComplex128(r io.Reader, order binary.ByteOrder) (complex128, error) {
	var rs complex128
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
			*dataPtr, err = ReadUInt(r, order)
		case uint8:
			*dataPtr, err = ReadUInt8(r, order)
		case uint16:
			*dataPtr, err = ReadUInt16(r, order)
		case uint32:
			*dataPtr, err = ReadUInt32(r, order)
		case uint64:
			*dataPtr, err = ReadUInt64(r, order)
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
