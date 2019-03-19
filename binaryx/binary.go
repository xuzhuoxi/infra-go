//
//Created by xuzhuoxi
//on 2019-03-19.
//@author xuzhuoxi
//
package binaryx

import (
	"encoding/binary"
	"io"
)

var bit32 = true

func SetLangBit(bit64 bool) {
	bit32 = !bit64
}

func ReadBool(r io.Reader, order binary.ByteOrder) (bool, error) {
	var rs bool
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceBool(r io.Reader, order binary.ByteOrder, ln int) ([]bool, error) {
	var rs = make([]bool, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadInt(r io.Reader, order binary.ByteOrder) (int, error) {
	if bit32 {
		val, err := ReadInt32(r, order)
		return int(val), err
	} else {
		val, err := ReadInt64(r, order)
		return int(val), err
	}
}
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
func ReadInt8(r io.Reader, order binary.ByteOrder) (int8, error) {
	var rs int8
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceInt8(r io.Reader, order binary.ByteOrder, ln int) ([]int8, error) {
	var rs = make([]int8, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadInt16(r io.Reader, order binary.ByteOrder) (int16, error) {
	var rs int16
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceInt16(r io.Reader, order binary.ByteOrder, ln int) ([]int16, error) {
	var rs = make([]int16, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadInt32(r io.Reader, order binary.ByteOrder) (int32, error) {
	var rs int32
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceInt32(r io.Reader, order binary.ByteOrder, ln int) ([]int32, error) {
	var rs = make([]int32, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadInt64(r io.Reader, order binary.ByteOrder) (int64, error) {
	var rs int64
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceInt64(r io.Reader, order binary.ByteOrder, ln int) ([]int64, error) {
	var rs = make([]int64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadUint(r io.Reader, order binary.ByteOrder) (uint, error) {
	if bit32 {
		val, err := ReadUint32(r, order)
		return uint(val), err
	} else {
		val, err := ReadUint64(r, order)
		return uint(val), err
	}
}
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
func ReadUint8(r io.Reader, order binary.ByteOrder) (uint8, error) {
	var rs uint8
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceUint8(r io.Reader, order binary.ByteOrder, ln int) ([]uint8, error) {
	var rs = make([]uint8, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadUint16(r io.Reader, order binary.ByteOrder) (uint16, error) {
	var rs uint16
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceUint16(r io.Reader, order binary.ByteOrder, ln int) ([]uint16, error) {
	var rs = make([]uint16, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadUint32(r io.Reader, order binary.ByteOrder) (uint32, error) {
	var rs uint32
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceUint32(r io.Reader, order binary.ByteOrder, ln int) ([]uint32, error) {
	var rs = make([]uint32, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadUint64(r io.Reader, order binary.ByteOrder) (uint64, error) {
	var rs uint64
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceUint64(r io.Reader, order binary.ByteOrder, ln int) ([]uint64, error) {
	var rs = make([]uint64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadFloat32(r io.Reader, order binary.ByteOrder) (float32, error) {
	var rs float32
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceFloat32(r io.Reader, order binary.ByteOrder, ln int) ([]float32, error) {
	var rs = make([]float32, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadFloat64(r io.Reader, order binary.ByteOrder) (float64, error) {
	var rs float64
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceFloat64(r io.Reader, order binary.ByteOrder, ln int) ([]float64, error) {
	var rs = make([]float64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadComplex64(r io.Reader, order binary.ByteOrder) (complex64, error) {
	var rs complex64
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceComplex64(r io.Reader, order binary.ByteOrder, ln int) ([]complex64, error) {
	var rs = make([]complex64, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadComplex128(r io.Reader, order binary.ByteOrder) (complex128, error) {
	var rs complex128
	err := binary.Read(r, order, &rs)
	return rs, err
}
func ReadSliceComplex128(r io.Reader, order binary.ByteOrder, ln int) ([]complex128, error) {
	var rs = make([]complex128, ln)
	err := binary.Read(r, order, &rs)
	return rs, err
}

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
		default:
			isCatch = false
		}
		if isCatch {
			return
		}
	}
	return binary.Read(r, order, data)
}

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
		default:
			isCatch = false
		}
		if isCatch {
			return
		}
	}
	return binary.Read(r, order, data)
}

func Write(w io.Writer, order binary.ByteOrder, data interface{}) error {
	tempData := data
	if dataPtr, ok := data.(*interface{}); ok {
		tempData = *dataPtr
	}
	switch d := tempData.(type) {
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
