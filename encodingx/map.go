//
//Created by xuzhuoxi
//on 2019-03-20.
//@author xuzhuoxi
//
package encodingx

import (
	"bytes"
	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/lang"
)

func NewCodingMap() CodingMap {
	return make(map[string]interface{})
}

type CodingMap map[string]interface{}

//序列化
//格式:
// 	Key + Value
//	Key : string(uint16+[]byte)
//	Value: Kind [+ Len] + Other
//	Other: []byte... 或 string(uint16+[]byte)...
func (v CodingMap) EncodeToBytes() []byte {
	buff := bytes.NewBuffer(nil)
	for key, val := range v {
		if !binaryx.CheckValue(val) { //非法值
			continue
		}
		binaryx.WriteString(buff, DefaultOrder, key) //Key
		kind, ln := binaryx.GetValueKind(val)
		//fmt.Println("Encode Kind:", kind, ln)
		binaryx.Write(buff, DefaultOrder, kind) //Kind
		if binaryx.IsSliceKind(kind) {
			binaryx.WriteLen(buff, DefaultOrder, ln) //Len
			switch kind {
			case binaryx.KindSliceString: //Value=[]string
				binaryx.WriteSliceString(buff, DefaultOrder, val.([]string))
			default: //Value
				binaryx.WriteSlice(buff, DefaultOrder, val)
			}
		} else {
			switch kind {
			case binaryx.KindString: //Value=string
				binaryx.WriteString(buff, DefaultOrder, val.(string))
			default: //Value
				binaryx.Write(buff, DefaultOrder, val)
			}
		}
		//fmt.Println("EncodeToBytes:", buff.Bytes())
	}
	return buff.Bytes()
}

func (v CodingMap) DecodeFromBytes(bs []byte) bool {
	buff := bytes.NewBuffer(bs)
	var err error
	for buff.Len() > 0 {
		key, _ := binaryx.ReadString(buff, DefaultOrder)
		var kind binaryx.ValueKind
		binaryx.Read(buff, DefaultOrder, &kind) //Kind
		//fmt.Println("Decode Kind:", kind)
		var val interface{}
		if binaryx.IsSliceKind(kind) {
			ln, _ := binaryx.ReadLen(buff, DefaultOrder) //Len
			switch kind {
			case binaryx.KindSliceString: //Value=[]string
				val, err = binaryx.ReadSliceString(buff, DefaultOrder, ln)
			default: //Value
				val = binaryx.GetKindValue(kind, ln)
				err = binaryx.ReadSlice(buff, DefaultOrder, &val, ln)
			}
		} else {
			switch kind {
			case binaryx.KindString: //Value=string
				val, err = binaryx.ReadString(buff, DefaultOrder)
			default: //Value
				val = binaryx.GetKindValue(kind, 0)
				err = binaryx.Read(buff, DefaultOrder, &val)
			}
		}
		if nil == err {
			v.Set(key, val)
		}
		//fmt.Println("DecodeFromBytes", buff.Len())
	}
	return true
}

func (v CodingMap) Merge(vs CodingMap) CodingMap {
	var rm []string
	for key, val := range vs {
		if v2, ok := v[key]; ok && lang.Equal(v2, val) {
			rm = append(rm, key)
			continue
		}
		v[key] = val
	}
	if len(rm) > 0 { //有重复
		for _, key := range rm {
			delete(vs, key)
		}
	}
	if len(vs) == 0 {
		return nil
	}
	return vs
}
func (v CodingMap) Set(key string, value interface{}) CodingMap {
	if v2, ok := v[key]; ok && lang.Equal(v2, value) {
		return nil
	}
	v[key] = value
	rs := NewCodingMap()
	rs[key] = value
	return rs
}
