//
//Created by xuzhuoxi
//on 2019-03-20.
//@author xuzhuoxi
//
package binaryx

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestPtr(t *testing.T) {
	var data interface{} = false
	var data2 *interface{} = &data
	fmt.Println("data:", data, *data2)
}
func TestBuff(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	var data interface{} = float32(23)
	binary.Write(buff, binary.BigEndian, &data)
	fmt.Println(buff.Bytes())
}

func TestSize(t *testing.T) {
	fmt.Println("int:", binary.Size(1111))
	fmt.Println("int16:", binary.Size(int16(1111)))
	fmt.Println("uint:", binary.Size(uint(1111)))
	fmt.Println("uint16:", binary.Size(uint16(1111)))
}

func TestType(t *testing.T) {
	catchType := func(e interface{}) {
		switch e := e.(type) {
		case bool:
			fmt.Println("bool")
		case *bool:
			fmt.Println("*bool")
		default:
			fmt.Println("default", e)
		}
	}
	var data interface{} = true
	var pdata *interface{} = &data
	catchType(data)
	catchType(pdata)
	fmt.Println("无敌分界线——————————————")
	var data2 = true
	catchType(data2)
	catchType(&data2)
	fmt.Println("无敌分界线——————————————")

	fmt.Println(unsafe.Sizeof(data), unsafe.Sizeof(&data), unsafe.Sizeof(data2), unsafe.Sizeof(&data2))
	fmt.Println(unsafe.Sizeof(struct{}{}), unsafe.Sizeof(make(map[string]struct{})), unsafe.Sizeof(make(map[string]string)))

	//这个interface{}真麻烦，具体类型转为interface{}时好像被内嵌了
}

func TestReflect(t *testing.T) {
	var data interface{} = true
	var data2 = true
	var data3 = &data
	var data4 = &data2
	fmt.Println(reflect.TypeOf(data3).Kind() == reflect.Ptr)
	fmt.Println(reflect.TypeOf(data4).Kind() == reflect.Ptr)
}
