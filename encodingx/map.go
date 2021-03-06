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

type IKeyValue interface {
	//序列化接口
	ICodingData

	//键值对数量
	Len() int
	//设置键值
	Set(key string, value interface{}) (IKeyValue, bool)
	//取值
	Get(key string) (interface{}, bool)
	//删除键值
	Delete(key string) (interface{}, bool)
	//检查键值是否存在
	Check(key string) bool

	//合并
	Merge(vs IKeyValue) IKeyValue
	//遍历
	ForEach(handler func(key string, value interface{}))
}

//-------------------------

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

func (v CodingMap) Len() int {
	return len(v)
}

func (v CodingMap) Set(key string, value interface{}) (IKeyValue, bool) {
	if v2, ok := v[key]; ok && lang.Equal(v2, value) {
		return nil, false
	}
	v[key] = value
	rs := NewCodingMap()
	rs[key] = value
	return rs, true
}

func (v CodingMap) Get(key string) (interface{}, bool) {
	value, ok := v[key]
	return value, ok
}

func (v CodingMap) Delete(key string) (interface{}, bool) {
	if v.Check(key) {
		rs, ok := v.Get(key)
		delete(v, key)
		return rs, ok
	}
	return nil, false
}

func (v CodingMap) Check(key string) bool {
	_, ok := v[key]
	return ok
}

func (v CodingMap) Merge(vs IKeyValue) IKeyValue {
	if nil == vs {
		return nil
	}
	var rm []string
	vs.ForEach(func(key string, value interface{}) {
		if v2, ok := v[key]; ok && lang.Equal(v2, value) {
			rm = append(rm, key)
			return
		}
		v[key] = value
	})
	if len(rm) > 0 { //有重复
		for _, key := range rm {
			vs.Delete(key)
		}
	}
	if vs.Len() == 0 {
		return nil
	}
	return vs
}

func (v CodingMap) ForEach(handler func(key string, value interface{})) {
	for key, value := range v {
		handler(key, value)
	}
}
