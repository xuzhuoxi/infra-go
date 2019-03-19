//
//Created by xuzhuoxi
//on 2019-03-18.
//@author xuzhuoxi
//
package binaryx

import (
	"strconv"
)

type ValueKind uint8

const (
	//简单类型
	KindNone ValueKind = iota
	KindBool
	KindInt8
	KindInt16
	KindInt32
	KindInt64
	KindUint8
	KindUint16
	KindUint32
	KindUint64
	KindFloat32
	KindFloat64
	KindComplex64
	KindComplex128
	KindInt
	KindUint
	KindString
)
const (
	//数组类型
	KindSliceNone ValueKind = iota + 128
	KindSliceBool
	KindSliceInt8
	KindSliceInt16
	KindSliceInt32
	KindSliceInt64
	KindSliceUint8
	KindSliceUint16
	KindSliceUint32
	KindSliceUint64
	KindSliceFloat32
	KindSliceFloat64
	KindSliceComplex64
	KindSliceComplex128
	KindSliceInt
	KindSliceUint
	KindSliceString
)

func IsForbidKind(kind ValueKind) bool {
	return kind == KindNone || kind == KindSliceNone
}

func IsSimpleKind(kind ValueKind) bool {
	return kind > KindNone && kind < KindSliceNone
}

func IsSliceKind(kind ValueKind) bool {
	return kind > KindSliceNone
}

//取定义的默认值
func GetKindValue(kind ValueKind, ln int) interface{} {
	if value := getKindValue(kind, ln); nil != value {
		return value
	}
	panic("No ValueKind Define:" + strconv.Itoa(int(kind)))
}

//检查值是为合法
func CheckValue(value interface{}) bool {
	kind, _ := getValueKind(value)
	return kind != KindNone
}

//取值对应的定义
func GetValueKind(value interface{}) (ValueKind, int) {
	return getValueKind(value)
}

func getKindValue(kind ValueKind, ln int) interface{} {
	switch kind {
	case KindBool:
		return false
	case KindInt8:
		return int8(0)
	case KindInt16:
		return int16(0)
	case KindInt32:
		return int32(0)
	case KindInt64:
		return int64(0)
	case KindUint8:
		return uint8(0)
	case KindUint16:
		return uint16(0)
	case KindUint32:
		return uint32(0)
	case KindUint64:
		return uint64(0)
	case KindFloat32:
		return float32(0)
	case KindFloat64:
		return float64(0)
	case KindComplex64:
		return complex64(0)
	case KindComplex128:
		return complex128(0)
	case KindString:
		return ""

	case KindSliceBool:
		return make([]bool, ln)
	case KindSliceInt8:
		return make([]int8, ln)
	case KindSliceInt16:
		return make([]int16, ln)
	case KindSliceInt32:
		return make([]int32, ln)
	case KindSliceInt64:
		return make([]int64, ln)
	case KindSliceUint8:
		return make([]uint8, ln)
	case KindSliceUint16:
		return make([]uint16, ln)
	case KindSliceUint32:
		return make([]uint32, ln)
	case KindSliceUint64:
		return make([]uint64, ln)
	case KindSliceFloat32:
		return make([]float32, ln)
	case KindSliceFloat64:
		return make([]float64, ln)
	case KindSliceComplex64:
		return make([]complex64, ln)
	case KindSliceComplex128:
		return make([]complex128, ln)
	case KindSliceString:
		return make([]string, ln)

	case KindInt:
		if bit32 {
			return int32(0)
		} else {
			return int64(0)
		}
	case KindUint:
		if bit32 {
			return uint32(0)
		} else {
			return uint64(0)
		}
	case KindSliceInt:
		if bit32 {
			return make([]int32, ln)
		} else {
			return make([]int64, ln)
		}
	case KindSliceUint:
		if bit32 {
			return make([]uint32, ln)
		} else {
			return make([]uint64, ln)
		}
	}
	return nil
}

func getValueKind(value interface{}) (ValueKind, int) {
	switch t := value.(type) {
	case *bool, bool:
		return KindBool, 1
	case []bool:
		return KindSliceBool, len(t)
	case *int, int:
		if bit32 {
			return KindInt, 4
		} else {
			return KindInt, 8
		}
	case []int:
		return KindSliceInt, len(t)
	case *int8, int8:
		return KindInt8, 1
	case []int8:
		return KindSliceInt8, len(t)
	case *int16, int16:
		return KindInt16, 2
	case []int16:
		return KindSliceInt16, len(t)
	case *int32, int32:
		return KindInt32, 4
	case []int32:
		return KindSliceInt32, len(t)
	case *int64, int64:
		return KindInt64, 8
	case []int64:
		return KindSliceInt64, len(t)
	case *uint, uint:
		if bit32 {
			return KindInt, 4
		} else {
			return KindInt, 8
		}
	case []uint:
		return KindSliceUint, len(t)
	case *uint8, uint8:
		return KindUint8, 1
	case []uint8:
		return KindSliceUint8, len(t)
	case *uint16, uint16:
		return KindUint16, 2
	case []uint16:
		return KindSliceUint16, len(t)
	case *uint32, uint32:
		return KindUint32, 4
	case []uint32:
		return KindSliceUint32, len(t)
	case *uint64, uint64:
		return KindUint64, 8
	case []uint64:
		return KindSliceUint64, len(t)
	case *float32, float32:
		return KindFloat32, 4
	case []float32:
		return KindSliceFloat32, len(t)
	case *float64, float64:
		return KindFloat64, 8
	case []float64:
		return KindSliceFloat64, len(t)
	case *complex64, complex64:
		return KindComplex64, 8
	case []complex64:
		return KindSliceComplex64, len(t)
	case *complex128, complex128:
		return KindComplex128, 16
	case []complex128:
		return KindSliceComplex128, len(t)
	case *string:
		return KindString, len(*t)
	case string:
		return KindString, len(t)
	case []string:
		return KindSliceString, len(t)
	default:
		return KindNone, 0
	}
}
