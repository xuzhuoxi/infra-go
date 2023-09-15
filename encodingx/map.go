// Package encodingx
// Created by xuzhuoxi
// on 2019-03-20.
// @author xuzhuoxi
//
package encodingx

import (
	"bytes"
	"fmt"
	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/lang"
	"strings"
)

type IKeyValue interface {
	// ICodingData
	// 序列化接口
	ICodingData

	// Len 键值对数量
	Len() int
	// Set 设置键值
	Set(key string, value interface{}) (old interface{}, ok bool)
	// Get 取值
	Get(key string) (interface{}, bool)
	// Delete 删除键值
	Delete(key string) (interface{}, bool)
	// Check 检查键值是否存在
	Check(key string) bool

	// Merge 合并
	Merge(vs IKeyValue) (update IKeyValue, del []string)
	// MergeArray 合并
	MergeArray(keys []string, vals []interface{}) (update IKeyValue, del []string)
	// ForEach 遍历
	ForEach(handler func(key string, value interface{}))
	// Clone 克隆
	Clone() IKeyValue
	// CloneEmpty 克隆空
	CloneEmpty() IKeyValue
}

//-------------------------

func NewCodingMap() CodingMap {
	return make(map[string]interface{})
}

type CodingMap map[string]interface{}

func (v CodingMap) String() string {
	if len(v) == 0 {
		return "{}"
	}
	builder := &strings.Builder{}
	builder.WriteString("{")
	index := 0
	ln := len(v)
	for key, val := range v {
		builder.WriteString(key + ":" + fmt.Sprint(val))
		index++
		if index < ln {
			builder.WriteString(",")
		}
	}
	builder.WriteString("}")
	return builder.String()
}

// EncodeToBytes
// 序列化
// 格式:
// 	Key + Value
//	Key : string(uint16+[]byte)
//	Value: Kind [+ Len] + Other
//	Other: []byte... 或 string(uint16+[]byte)...
func (v CodingMap) EncodeToBytes() (bs []byte, err error) {
	buff := bytes.NewBuffer(nil)
	err = binaryx.WriteLen(buff, DefaultOrder, len(v))
	if nil != err {
		return
	}
	for key, val := range v {
		if !binaryx.CheckValue(val) { //非法值
			continue
		}
		err = binaryx.WriteString(buff, DefaultOrder, key) //Key
		if nil != err {
			return
		}
		kind, ln := binaryx.GetValueKind(val)
		err = binaryx.Write(buff, DefaultOrder, kind) //Kind
		if nil != err {
			return
		}

		if binaryx.IsSliceKind(kind) {
			err = binaryx.WriteLen(buff, DefaultOrder, ln)
			if nil != err {
				return
			}
			err = binaryx.WriteSlice(buff, DefaultOrder, val)
			if nil != err {
				return
			}
		} else {
			err = binaryx.Write(buff, DefaultOrder, val)
			if nil != err {
				return
			}
		}
	}
	return buff.Bytes(), nil
}

func (v CodingMap) DecodeFromBytes(bs []byte) error {
	buff := bytes.NewBuffer(bs)
	ln, err := binaryx.ReadLen(buff, DefaultOrder)
	if nil != err {
		return err
	}
	for ln >= 0 && buff.Len() > 0 {
		key, _ := binaryx.ReadString(buff, DefaultOrder)
		var kind binaryx.ValueKind
		_ = binaryx.Read(buff, DefaultOrder, &kind) //Kind
		//fmt.Println("Decode Kind:", kind)
		var val interface{}
		if binaryx.IsSliceKind(kind) {
			ln, _ := binaryx.ReadLen(buff, DefaultOrder) //Len
			switch kind {
			case binaryx.KindSliceString: //Value=[]string
				val, err = binaryx.ReadSliceStringBy(buff, DefaultOrder, ln)
			default: //Value
				val = binaryx.GetKindValue(kind, ln)
				err = binaryx.ReadSliceBy(buff, DefaultOrder, &val, ln)
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
		if nil != err {
			return err
		}
		v.Set(key, val)
		ln--
	}
	return nil
}

func (v CodingMap) Len() int {
	return len(v)
}

func (v CodingMap) Set(key string, value interface{}) (old interface{}, ok bool) {
	if old, ok = v[key]; ok && lang.Equal(old, value) {
		return nil, false
	}
	v[key] = value
	return old, true
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

func (v CodingMap) Merge(vs IKeyValue) (update IKeyValue, del []string) {
	if nil == vs {
		return nil, nil
	}
	var rm []string
	vs.ForEach(func(key string, value interface{}) {
		if value == nil {
			del = append(del, key)
			v.Delete(key)
			return
		}
		_, _ = v.Set(key, value)
	})
	if len(rm) > 0 { //有重复
		for _, key := range rm {
			vs.Delete(key)
		}
	}
	return vs, del
}

func (v CodingMap) MergeArray(keys []string, vals []interface{}) (update IKeyValue, del []string) {
	update = NewCodingMap()
	for index := range keys {
		if vals[index] == nil {
			del = append(del, keys[index])
			v.Delete(keys[index])
			return
		}
		_, _ = v.Set(keys[index], vals[index])
		_, _ = update.Set(keys[index], vals[index])
	}
	return
}

func (v CodingMap) ForEach(handler func(key string, value interface{})) {
	for key, value := range v {
		handler(key, value)
	}
}

func (v CodingMap) Clone() IKeyValue {
	rs := make(CodingMap)
	for key, val := range v {
		rs[key] = val
	}
	return rs
}

func (v CodingMap) CloneEmpty() IKeyValue {
	return make(CodingMap)
}
