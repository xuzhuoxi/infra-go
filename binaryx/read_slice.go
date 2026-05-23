package binaryx

import (
	"encoding/binary"
	"io"
)

// ReadSliceBool reads a length-prefixed []bool from r.
// ReadSliceBool 从一个 Reader 中读取一组 bool 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceBool(r io.Reader, order binary.ByteOrder) ([]bool, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceBoolBy(r, order, byLen)
}

// ReadSliceBoolBy reads byLen bool values from r.
// ReadSliceBoolBy 从一个 Reader 中读取一组 bool 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceBoolBy(r io.Reader, order binary.ByteOrder, byLen int) ([]bool, error) {
	var rs = make([]bool, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt8 reads a length-prefixed []int8 from r.
// ReadSliceInt8 从一个 Reader 中读取一组 Int8 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceInt8(r io.Reader, order binary.ByteOrder) ([]int8, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt8By(r, order, byLen)
}

// ReadSliceInt8By reads byLen int8 values from r.
// ReadSliceInt8By 从一个 Reader 中读取一组 Int8 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceInt8By(r io.Reader, order binary.ByteOrder, byLen int) ([]int8, error) {
	var rs = make([]int8, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt16 reads a length-prefixed []int16 from r.
// ReadSliceInt16 从一个 Reader 中读取一组 Int16 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceInt16(r io.Reader, order binary.ByteOrder) ([]int16, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt16By(r, order, byLen)
}

// ReadSliceInt16By reads byLen int16 values from r.
// ReadSliceInt16By 从一个 Reader 中读取一组 Int16 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceInt16By(r io.Reader, order binary.ByteOrder, byLen int) ([]int16, error) {
	var rs = make([]int16, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt32 reads a length-prefixed []int32 from r.
// ReadSliceInt32 从一个 Reader 中读取一组 Int32 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceInt32(r io.Reader, order binary.ByteOrder) ([]int32, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt32By(r, order, byLen)
}

// ReadSliceInt32By reads byLen int32 values from r.
// ReadSliceInt32By 从一个 Reader 中读取一组 Int32 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceInt32By(r io.Reader, order binary.ByteOrder, byLen int) ([]int32, error) {
	var rs = make([]int32, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt64 reads a length-prefixed []int64 from r.
// ReadSliceInt64 从一个 Reader 中读取一组 Int64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceInt64(r io.Reader, order binary.ByteOrder) ([]int64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceInt64By(r, order, byLen)
}

// ReadSliceInt64By reads byLen int64 values from r.
// ReadSliceInt64By 从一个 Reader 中读取一组 Int64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceInt64By(r io.Reader, order binary.ByteOrder, byLen int) ([]int64, error) {
	var rs = make([]int64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceInt reads a length-prefixed []int from r; uses int32 on 32-bit platforms and int64 on 64-bit platforms.
// ReadSliceInt 从一个 Reader 中读取一组 Int 数据，32 位则读取 int32，64 位则读取 int64。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceInt(r io.Reader, order binary.ByteOrder) ([]int, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceIntBy(r, order, byLen)
}

// ReadSliceIntBy reads byLen int values from r; uses int32 on 32-bit platforms and int64 on 64-bit platforms.
// ReadSliceIntBy 从一个 Reader 中读取一组 Int 数据，32 位则读取 int32，64 位则读取 int64。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
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

// ReadSliceUInt8 reads a length-prefixed []uint8 from r.
// ReadSliceUInt8 从一个 Reader 中读取一组 Uint8 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceUInt8(r io.Reader, order binary.ByteOrder) ([]uint8, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt8By(r, order, byLen)
}

// ReadSliceUInt8By reads byLen uint8 values from r.
// ReadSliceUInt8By 从一个 Reader 中读取一组 Uint8 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceUInt8By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint8, error) {
	var rs = make([]uint8, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt16 reads a length-prefixed []uint16 from r.
// ReadSliceUInt16 从一个 Reader 中读取一组 Uint16 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceUInt16(r io.Reader, order binary.ByteOrder) ([]uint16, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt16By(r, order, byLen)
}

// ReadSliceUInt16By reads byLen uint16 values from r.
// ReadSliceUInt16By 从一个 Reader 中读取一组 Uint16 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceUInt16By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint16, error) {
	var rs = make([]uint16, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt32 reads a length-prefixed []uint32 from r.
// ReadSliceUInt32 从一个 Reader 中读取一组 Uint32 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceUInt32(r io.Reader, order binary.ByteOrder) ([]uint32, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt32By(r, order, byLen)
}

// ReadSliceUInt32By reads byLen uint32 values from r.
// ReadSliceUInt32By 从一个 Reader 中读取一组 Uint32 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceUInt32By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint32, error) {
	var rs = make([]uint32, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt64 reads a length-prefixed []uint64 from r.
// ReadSliceUInt64 从一个 Reader 中读取一组 Uint64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceUInt64(r io.Reader, order binary.ByteOrder) ([]uint64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUInt64By(r, order, byLen)
}

// ReadSliceUInt64By reads byLen uint64 values from r.
// ReadSliceUInt64By 从一个 Reader 中读取一组 Uint64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceUInt64By(r io.Reader, order binary.ByteOrder, byLen int) ([]uint64, error) {
	var rs = make([]uint64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceUInt reads a length-prefixed []uint from r; uses uint32 on 32-bit platforms and uint64 on 64-bit platforms.
// ReadSliceUInt 从一个 Reader 中读取一组 Uint 数据，32 位则读取 uint32，64 位则读取 uint64。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceUInt(r io.Reader, order binary.ByteOrder) ([]uint, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceUIntBy(r, order, byLen)
}

// ReadSliceUIntBy reads byLen uint values from r; uses uint32 on 32-bit platforms and uint64 on 64-bit platforms.
// ReadSliceUIntBy 从一个 Reader 中读取一组 Uint 数据，32 位则读取 uint32，64 位则读取 uint64。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
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

// ReadSliceFloat32 reads a length-prefixed []float32 from r.
// ReadSliceFloat32 从一个 Reader 中读取一组 Float32 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceFloat32(r io.Reader, order binary.ByteOrder) ([]float32, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceFloat32By(r, order, byLen)
}

// ReadSliceFloat32By reads byLen float32 values from r.
// ReadSliceFloat32By 从一个 Reader 中读取一组 Float32 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceFloat32By(r io.Reader, order binary.ByteOrder, byLen int) ([]float32, error) {
	var rs = make([]float32, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceFloat64 reads a length-prefixed []float64 from r.
// ReadSliceFloat64 从一个 Reader 中读取一组 Float64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceFloat64(r io.Reader, order binary.ByteOrder) ([]float64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceFloat64By(r, order, byLen)
}

// ReadSliceFloat64By reads byLen float64 values from r.
// ReadSliceFloat64By 从一个 Reader 中读取一组 Float64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceFloat64By(r io.Reader, order binary.ByteOrder, byLen int) ([]float64, error) {
	var rs = make([]float64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceComplex64 reads a length-prefixed []complex64 from r.
// ReadSliceComplex64 从一个 Reader 中读取一组 Complex64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceComplex64(r io.Reader, order binary.ByteOrder) ([]complex64, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceComplex64By(r, order, byLen)
}

// ReadSliceComplex64By reads byLen complex64 values from r.
// ReadSliceComplex64By 从一个 Reader 中读取一组 Complex64 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceComplex64By(r io.Reader, order binary.ByteOrder, byLen int) ([]complex64, error) {
	var rs = make([]complex64, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceComplex128 reads a length-prefixed []complex128 from r.
// ReadSliceComplex128 从一个 Reader 中读取一组 Complex128 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceComplex128(r io.Reader, order binary.ByteOrder) ([]complex128, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceComplex128By(r, order, byLen)
}

// ReadSliceComplex128By reads byLen complex128 values from r.
// ReadSliceComplex128By 从一个 Reader 中读取一组 Complex128 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
func ReadSliceComplex128By(r io.Reader, order binary.ByteOrder, byLen int) ([]complex128, error) {
	var rs = make([]complex128, byLen)
	err := binary.Read(r, order, &rs)
	return rs, err
}

// ReadSliceString reads a length-prefixed []string from r.
// ReadSliceString 从一个 Reader 中读取一组 string 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadSliceString(r io.Reader, order binary.ByteOrder) ([]string, error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return nil, errL
	}
	return ReadSliceStringBy(r, order, byLen)
}

// ReadSliceStringBy reads byLen length-prefixed strings from r.
// ReadSliceStringBy 从一个 Reader 中读取一组 string 数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
// byLen: element count.
// byLen: 数据长度。
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

// ReadSlice reads a length-prefixed slice from r into data.
// ReadSlice 从 Reader 中读取数据，先读取长度。
// Currently supported types for data: *[]bool, *[]int, *[]int8, *[]int16, *[]int32, *[]int64, *[]uint, *[]uint8, *[]uint16, *[]uint32, *[]uint64, *[]float32, *[]float64, *[]complex64, *[]complex128, *[]string.
// 目前 data 支持的类型为 *[]bool、*[]int、*[]int8、*[]int16、*[]int32、*[]int64、*[]uint、*[]uint8、*[]uint16、*[]uint32、*[]uint64、*[]float32、*[]float64、*[]complex64、*[]complex128、*[]string。
func ReadSlice(r io.Reader, order binary.ByteOrder, data interface{}) (err error) {
	byLen, errL := ReadLen(r, order)
	if nil != errL {
		return errL
	}
	return ReadSliceBy(r, order, data, byLen)
}

// ReadSliceBy reads byLen elements from r into data.
// ReadSliceBy 从 Reader 中读取数据。
// Currently supported types for data: *[]bool, *[]int, *[]int8, *[]int16, *[]int32, *[]int64, *[]uint, *[]uint8, *[]uint16, *[]uint32, *[]uint64, *[]float32, *[]float64, *[]complex64, *[]complex128, *[]string.
// 目前 data 支持的类型为 *[]bool、*[]int、*[]int8、*[]int16、*[]int32、*[]int64、*[]uint、*[]uint8、*[]uint16、*[]uint32、*[]uint64、*[]float32、*[]float64、*[]complex64、*[]complex128、*[]string。
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
