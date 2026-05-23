package encodingx

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/lang"
)

// IKeyValue is a string-keyed map with binary encode/decode and map operations.
// IKeyValue 是支持二进制编解码与常用映射操作的字符串键值容器。
type IKeyValue interface {
	ICodingData

	// Len returns the number of key-value pairs.
	// Len 返回键值对数量。
	Len() int
	// Clear removes all key-value pairs.
	// Clear 清理全部键值对。
	Clear()
	// Set stores value under key and returns the previous value if replaced.
	// Set 设置键值，若发生替换则返回旧值。
	Set(key string, value interface{}) (old interface{}, ok bool)
	// Get returns the value for key.
	// Get 根据键取值。
	Get(key string) (interface{}, bool)
	// Delete removes the key and returns its former value.
	// Delete 删除键值。
	Delete(key string) (interface{}, bool)
	// Check reports whether key exists.
	// Check 检查键是否存在。
	Check(key string) bool

	// Merge merges entries from vs; nil values delete keys; returns updated keys.
	// Merge 合并 vs 中的条目；值为 nil 时删除键；返回发生更新的键列表。
	Merge(vs IKeyValue, ignoreSame bool) (updated []string)
	// MergeArray merges parallel keys and vals; nil values delete keys.
	// MergeArray 按 keys/vals 并行合并；值为 nil 时删除键。
	MergeArray(keys []string, vals []interface{}, ignoreSame bool) (updated []string)
	// ForEach invokes handler for each key-value pair.
	// ForEach 遍历所有键值对。
	ForEach(handler func(key string, value interface{}))
	// Clone returns a shallow copy of the map.
	// Clone 克隆当前映射。
	Clone() IKeyValue
	// CloneEmpty returns a new empty map of the same concrete type.
	// CloneEmpty 克隆一个同类型的空映射。
	CloneEmpty() IKeyValue
	// CopyFrom copies selected keys from from into the receiver.
	// CopyFrom 从 from 复制指定键到当前映射。
	CopyFrom(from IKeyValue, keys []string) (n int)
}

//-------------------------

// NewCodingMap returns an empty CodingMap.
// NewCodingMap 创建空的 CodingMap。
func NewCodingMap() CodingMap {
	return make(map[string]interface{})
}

// CodingMap is a string-keyed map with binaryx wire encoding.
// CodingMap 是基于 binaryx 线格式的 string 键值映射。
type CodingMap map[string]interface{}

// String returns a sorted debug representation of the map.
// String 返回按键排序的调试字符串。
func (v CodingMap) String() string {
	if len(v) == 0 {
		return "{}"
	}
	var keys []string
	for key := range v {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	builder := &strings.Builder{}
	builder.WriteString("{")
	ln := len(keys)
	for index, key := range keys {
		builder.WriteString(key + ":" + fmt.Sprint(v[key]))
		if index < ln-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("}")
	return builder.String()
}

// EncodeToBytes serializes the map to bytes.
// EncodeToBytes 将映射序列化为字节。
//
// Format:
//
//	Key + Value per entry
//	Key: string (uint16 length + []byte)
//	Value: Kind [+ Len] + payload
//	Other: []byte... or string (uint16+[]byte)...
//
// 格式：
//
//	每项为 Key + Value
//	Key：string（uint16 长度 + []byte）
//	Value：Kind [+ Len] + 载荷
//	Other：[]byte... 或 string（uint16+[]byte）...
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

// DecodeFromBytes deserializes bytes into the map.
// DecodeFromBytes 从字节反序列化并更新映射。
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

// Clear removes all entries.
// Clear 清理全部键值对。
func (v CodingMap) Clear() {
	for key := range v {
		delete(v, key)
	}
}

// Len returns the number of entries.
// Len 返回键值对数量。
func (v CodingMap) Len() int {
	return len(v)
}

// Set stores value under key.
// Set 设置键值。
func (v CodingMap) Set(key string, value interface{}) (old interface{}, ok bool) {
	old, ok = v[key]
	if !ok && nil == value {
		return nil, false
	}
	if ok && lang.Equal(old, value) {
		return nil, false
	}
	v[key] = value
	return old, true
}

// Get returns the value for key.
// Get 根据键取值。
func (v CodingMap) Get(key string) (interface{}, bool) {
	value, ok := v[key]
	return value, ok
}

// Delete removes key and returns its former value.
// Delete 删除键值。
func (v CodingMap) Delete(key string) (interface{}, bool) {
	if v.Check(key) {
		rs, ok := v.Get(key)
		delete(v, key)
		return rs, ok
	}
	return nil, false
}

// Check reports whether key exists.
// Check 检查键是否存在。
func (v CodingMap) Check(key string) bool {
	_, ok := v[key]
	return ok
}

// Merge merges entries from vs.
// Merge 合并 vs 中的条目。
func (v CodingMap) Merge(vs IKeyValue, ignoreSame bool) (updated []string) {
	if nil == vs {
		return
	}
	vs.ForEach(func(key string, value interface{}) {
		if value == nil {
			updated = append(updated, key)
			v.Delete(key)
			return
		}
		_, ok := v.Set(key, value)
		if ok || ignoreSame {
			updated = append(updated, key)
		}
	})
	return
}

// MergeArray merges parallel keys and vals.
// MergeArray 按 keys/vals 并行合并。
func (v CodingMap) MergeArray(keys []string, vals []interface{}, ignoreSame bool) (updated []string) {
	for index := range keys {
		if vals[index] == nil {
			updated = append(updated, keys[index])
			v.Delete(keys[index])
			return
		}
		_, ok := v.Set(keys[index], vals[index])
		if ok || ignoreSame {
			updated = append(updated, keys[index])
		}
	}
	return
}

// ForEach invokes handler for each entry.
// ForEach 遍历所有键值对。
func (v CodingMap) ForEach(handler func(key string, value interface{})) {
	for key, value := range v {
		handler(key, value)
	}
}

// Clone returns a shallow copy.
// Clone 克隆当前映射。
func (v CodingMap) Clone() IKeyValue {
	rs := make(CodingMap)
	for key, val := range v {
		rs[key] = val
	}
	return rs
}

// CloneEmpty returns a new empty CodingMap.
// CloneEmpty 克隆一个空的 CodingMap。
func (v CodingMap) CloneEmpty() IKeyValue {
	return make(CodingMap)
}

// CopyFrom copies selected keys from from.
// CopyFrom 从 from 复制指定键。
func (v CodingMap) CopyFrom(from IKeyValue, keys []string) (n int) {
	for index := range keys {
		if val, ok := from.Get(keys[index]); ok {
			v.Set(keys[index], val)
			n++
		}
	}
	return
}
