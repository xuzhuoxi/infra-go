// Package binaryx
// Create on 2023/7/18
// @author xuzhuoxi
package binaryx

import (
	"encoding/binary"
	"io"
)

// ReadSliceBool
// 从一个Reader中读取一组bool数据
// r: Reader
// order: 大小端设定
func ReadSliceBool(r io.Reader, order binary.ByteOrder) ([]bool, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceBoolBy(r, order, byLen)
}

// ReadSliceBoolBy
// 从一个Reader中读取一组bool数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceBoolBy(r io.Reader, order binary.ByteOrder, byLen int) ([]bool, error) {
	var rs = make([]bool, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt8
// 从一个Reader中读取一组Int8数据
// r: Reader
// order: 大小端设定
func ReadSliceInt8(r io.Reader, order binary.ByteOrder) ([]int8, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt8By(r, order, byLen)
}

// ReadSliceInt8By
// 从一个Reader中读取一组Int8数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt8By(r io.Reader, order binary.ByteOrder, byLen int) ([]int8, error) {
	var rs = make([]int8, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt16
// 从一个Reader中读取一组Int16数据
// r: Reader
// order: 大小端设定
func ReadSliceInt16(r io.Reader, order binary.ByteOrder) ([]int16, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt16By(r, order, byLen)
}

// ReadSliceInt16By
// 从一个Reader中读取一组Int16数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt16By(r io.Reader, order binary.ByteOrder, byLen int) ([]int16, error) {
	var rs = make([]int16, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt32
// 从一个Reader中读取一组Int32数据
// r: Reader
// order: 大小端设定
func ReadSliceInt32(r io.Reader, order binary.ByteOrder) ([]int32, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt32By(r, order, byLen)
}

// ReadSliceInt32By
// 从一个Reader中读取一组Int32数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt32By(r io.Reader, order binary.ByteOrder, byLen int) ([]int32, error) {
	var rs = make([]int32, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt64
// 从一个Reader中读取一组Int64数据
// r: Reader
// order: 大小端设定
func ReadSliceInt64(r io.Reader, order binary.ByteOrder) ([]int64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt64By(r, order, byLen)
}

// ReadSliceInt64By
// 从一个Reader中读取一组Int64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt64By(r io.Reader, order binary.ByteOrder, byLen int) ([]int64, error) {
	var rs = make([]int64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt
// 从一个Reader中读取一组Int数据, 32位则读取int32,64位则读取int64
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceInt(r io.Reader, order binary.ByteOrder) ([]int, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceIntBy(r, order, byLen)
}

// ReadSliceIntBy
// 从一个Reader中读取一组Int数据, 32位则读取int32,64位则读取int64
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceIntBy(r io.Reader, order binary.ByteOrder, byLen int) ([]int, error) {
	var rs []int
	if bit32 {
		val, err := ReadSliceInt32By(r, order, byLen)
		for _, v := range val {
			rs = append(rs, int(v))
		}
		return rs, err
	} else {
		val, err := ReadSliceInt64By(r, order, byLen)
		for _, v := range val {
			rs = append(rs, int(v))
		}
		return rs, err
	}
}

// ReadSliceUInt8
// 从一个Reader中读取一组Uint8数据
// r: Reader
// order: 大小端设定
func ReadSliceUInt8(r io.Reader, order binary.ByteOrder) ([]uint8, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt8By(r, order, byLen)
}

// ReadSliceUInt8By
// 从一个Reader中读取一组Uint8数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUInt8By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint8, error) {
	var rs = make([]uint8, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt16
//从一个Reader中读取一组Uint16数据
// r: Reader
// order: 大小端设定
func ReadSliceUInt16(r io.Reader, order binary.ByteOrder) ([]uint16, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt16By(r, order, byLen)
}

// ReadSliceUInt16By
//从一个Reader中读取一组Uint16数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUInt16By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint16, error) {
	var rs = make([]uint16, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt32
// 从一个Reader中读取一组Uint32数据
// r: Reader
// order: 大小端设定
func ReadSliceUInt32(r io.Reader, order binary.ByteOrder) ([]uint32, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt32By(r, order, byLen)
}

// ReadSliceUInt32By
// 从一个Reader中读取一组Uint32数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUInt32By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint32, error) {
	var rs = make([]uint32, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt64
// 从一个Reader中读取一组Uint64数据
// r: Reader
// order: 大小端设定
func ReadSliceUInt64(r io.Reader, order binary.ByteOrder) ([]uint64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt64By(r, order, byLen)
}

// ReadSliceUInt64By
// 从一个Reader中读取一组Uint64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUInt64By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint64, error) {
	var rs = make([]uint64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt
// 从一个Reader中读取一组Uint数据, 32位则读取unt32,64位则读取unt64
// r: Reader
// order: 大小端设定
func ReadSliceUInt(r io.Reader, order binary.ByteOrder) ([]uint, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUIntBy(r, order, byLen)
}

// ReadSliceUIntBy
// 从一个Reader中读取一组Uint数据, 32位则读取unt32,64位则读取unt64
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceUIntBy(r io.Reader, order binary.ByteOrder, byLen int) ([]uint, error) {
	var rs []uint
	if bit32 {
		val, err := ReadSliceUInt32By(r, order, byLen)
		for _, v := range val {
			rs = append(rs, uint(v))
		}
		return rs, err
	} else {
		val, err := ReadSliceUInt64By(r, order, byLen)
		for _, v := range val {
			rs = append(rs, uint(v))
		}
		return rs, err
	}
}

// ReadSliceFloat32
// 从一个Reader中读取一组Float32数据
// r: Reader
// order: 大小端设定
func ReadSliceFloat32(r io.Reader, order binary.ByteOrder) ([]float32, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceFloat32By(r, order, byLen)
}

// ReadSliceFloat32By
// 从一个Reader中读取一组Float32数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceFloat32By(r io.Reader, order binary.ByteOrder, byLen int) ([]float32, error) {
	var rs = make([]float32, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceFloat64
// 从一个Reader中读取一组Float64数据
// r: Reader
// order: 大小端设定
func ReadSliceFloat64(r io.Reader, order binary.ByteOrder) ([]float64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceFloat64By(r, order, byLen)
}

// ReadSliceFloat64By
// 从一个Reader中读取一组Float64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceFloat64By(r io.Reader, order binary.ByteOrder, byLen int) ([]float64, error) {
	var rs = make([]float64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceComplex64
// 从一个Reader中读取一组Complex64数据
// r: Reader
// order: 大小端设定
func ReadSliceComplex64(r io.Reader, order binary.ByteOrder) ([]complex64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceComplex64By(r, order, byLen)
}

// ReadSliceComplex64By
// 从一个Reader中读取一组Complex64数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceComplex64By(r io.Reader, order binary.ByteOrder, byLen int) ([]complex64, error) {
	var rs = make([]complex64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceComplex128
// 从一个Reader中读取一组Complex128数据
// r: Reader
// order: 大小端设定
func ReadSliceComplex128(r io.Reader, order binary.ByteOrder) ([]complex128, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceComplex128By(r, order, byLen)
}

// ReadSliceComplex128By
// 从一个Reader中读取一组Complex128数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceComplex128By(r io.Reader, order binary.ByteOrder, byLen int) ([]complex128, error) {
	var rs = make([]complex128, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceString
//从一个Reader中读取一组string数据
// r: Reader
// order: 大小端设定
func ReadSliceString(r io.Reader, order binary.ByteOrder) ([]string, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceStringBy(r, order, byLen)
}

// ReadSliceStringBy
//从一个Reader中读取一组string数据
// r: Reader
// order: 大小端设定
// ln: 数据长度
func ReadSliceStringBy(r io.Reader, order binary.ByteOrder, byLen int) ([]string, error) {
	var rs []string
	for index := byLen - 1; index >= 0; index-- {
		str, err := ReadString(r, order)
		if nil != err {
			return nil, err
		}
		rs = append(rs, str)
	}
	return rs, nil
}

// ReadSlice
// 从Reader中读取数据，先读取长度
// 目前data支持的类型为*[]bool,*[]int,*[]int8,*[]int16,*[]int32,*[]int64,*[]uint,*[]uint8,*[]uint16,*[]uint32,*[]uint64,*[]float32,*[]float64,*[]complex64,*[]complex128,*[]string
func ReadSlice(r io.Reader, order binary.ByteOrder, data interface{}) (err error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return errL
	}
	return ReadSliceBy(r, order, data, byLen)
}

// ReadSliceBy
// 从Reader中读取数据
// 目前data支持的类型为*[]bool,*[]int,*[]int8,*[]int16,*[]int32,*[]int64,*[]uint,*[]uint8,*[]uint16,*[]uint32,*[]uint64,*[]float32,*[]float64,*[]complex64,*[]complex128,*[]string
func ReadSliceBy(r io.Reader, order binary.ByteOrder, data interface{}, byLen int) (err error) {
	if dataPtr, ok := data.(*interface{}); ok {
		isCatch := true
		switch (*dataPtr).(type) {
		case []bool:
			*dataPtr, err = ReadSliceBoolBy(r, order, byLen)
		case []int:
			*dataPtr, err = ReadSliceIntBy(r, order, byLen)
		case []int8:
			*dataPtr, err = ReadSliceInt8By(r, order, byLen)
		case []int16:
			*dataPtr, err = ReadSliceInt16By(r, order, byLen)
		case []int32:
			*dataPtr, err = ReadSliceInt32By(r, order, byLen)
		case []int64:
			*dataPtr, err = ReadSliceInt64By(r, order, byLen)
		case []uint:
			*dataPtr, err = ReadSliceUIntBy(r, order, byLen)
		case []uint8:
			*dataPtr, err = ReadSliceUInt8By(r, order, byLen)
		case []uint16:
			*dataPtr, err = ReadSliceUInt16By(r, order, byLen)
		case []uint32:
			*dataPtr, err = ReadSliceUInt32By(r, order, byLen)
		case []uint64:
			*dataPtr, err = ReadSliceUInt64By(r, order, byLen)
		case []float32:
			*dataPtr, err = ReadSliceFloat32By(r, order, byLen)
		case []float64:
			*dataPtr, err = ReadSliceFloat64By(r, order, byLen)
		case []complex64:
			*dataPtr, err = ReadSliceComplex64By(r, order, byLen)
		case []complex128:
			*dataPtr, err = ReadSliceComplex128By(r, order, byLen)
		case []string:
			*dataPtr, err = ReadSliceStringBy(r, order, byLen)
		default:
			isCatch = false
		}
		if isCatch {
			return
		}
	}
	return binary.Read(r, order, data)
}
