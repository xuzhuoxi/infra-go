package lang

import (
	"bytes"
	"reflect"
)

func Equal(a, b interface{}) bool {
	if &a == &b {
		return true
	}
	if te, T := TypeEqual(a, b); te {
		switch T.Kind() {
		case reflect.Bool:
			return a == b
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return a == b
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return a == b
		case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
			return a == b
		case reflect.String:
			return a == b
		case reflect.Array, reflect.Slice:
			if as, ok := a.([]byte); ok {
				return bytes.Equal(as, b.([]byte))
			}
			va := reflect.ValueOf(a)
			vb := reflect.ValueOf(b)
			if va.Len() != vb.Len() {
				return false
			}
			for i := 0; i < va.Len(); i++ {
				if va.Index(i).Interface() != vb.Index(i).Interface() {
					return false
				}
			}
			return true
		}
	}
	return false
}

func TypeEqual(a, b interface{}) (bool, reflect.Type) {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(b)
	if t1 != t2 { //就算是两个nil，也是会判断类型的，因为nil自带类型
		return false, nil
	}
	return true, t1
}
