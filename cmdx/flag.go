//
//Created by xuzhuoxi
//on 2019-06-07.
//@author xuzhuoxi
//
package cmdx

import (
	"errors"
	"flag"
	"github.com/xuzhuoxi/infra-go/slicex"
	"os"
	"reflect"
	"strings"
	"time"
)

func NewDefaultFlagSetExtend() *FlagSetExtend {
	name := os.Args[0]
	errorHandling := flag.ContinueOnError
	return NewFlagSetExtend(name, errorHandling)
}

func NewFlagSetExtend(name string, errorHandling flag.ErrorHandling) *FlagSetExtend {
	return &FlagSetExtend{
		FlagSet: *flag.NewFlagSet(name, errorHandling), errorHandling: errorHandling}
}

type FlagSetExtend struct {
	flag.FlagSet
	keyList       []string
	errorHandling flag.ErrorHandling
}

// 只有预定义了参数才能使用
func (fs *FlagSetExtend) Parse(arguments []string) error {
	if fs.FlagSet.Parsed() {
		return errors.New("Parsed! ")
	}
	err := fs.FlagSet.Parse(arguments)
	if nil != err {
		return err
	}
	for _, kv := range arguments {
		s := strings.IndexByte(kv, '-')
		e := strings.LastIndexByte(kv, '=')
		if -1 == s || -1 == e {
			return errors.New("bad flag syntax: " + kv)
		}
		key := kv[s+1 : e]
		fs.keyList = append(fs.keyList, key)
	}
	return nil
}

func (fs *FlagSetExtend) CheckKey(key string) bool {
	_, ok := slicex.IndexString(fs.keyList, strings.ToLower(key))
	return ok
}

func (fs *FlagSetExtend) GetReflectValue(key string) (val reflect.Value, ok bool) {
	key = strings.ToLower(key)
	flag := fs.Lookup(key)
	if nil != flag {
		return reflect.ValueOf(flag.Value).Elem(), true
	}
	return reflect.Value{}, false
}

func (fs *FlagSetExtend) GetBool(key string) (val bool, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Bool {
			return val.Bool(), true
		}
	}
	return false, false
}

func (fs *FlagSetExtend) GetString(key string) (val string, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.String {
			return val.String(), true
		}
	}
	return "", false
}

func (fs *FlagSetExtend) GetDuration(key string) (val time.Duration, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Int64 {
			return time.Duration(val.Int()), true
		}
	}
	return 0, false
}

func (fs *FlagSetExtend) GetFloat64(key string) (val float64, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Float64 || val.Kind() == reflect.Float32 {
			return float64(val.Float()), true
		}
	}
	return 0, false
}

func (fs *FlagSetExtend) GetInt64(key string) (val int64, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Int || val.Kind() == reflect.Int8 || val.Kind() == reflect.Int16 || val.Kind() == reflect.Int32 || val.Kind() == reflect.Int64 {
			return int64(val.Int()), true
		}
	}
	return 0, false
}

func (fs *FlagSetExtend) GetUint64(key string) (val uint64, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Uint || val.Kind() == reflect.Uint8 || val.Kind() == reflect.Uint16 || val.Kind() == reflect.Uint32 || val.Kind() == reflect.Uint64 {
			return uint64(val.Uint()), true
		}
	}
	return 0, false
}

func (fs *FlagSetExtend) GetInt(key string) (val int, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Int || val.Kind() == reflect.Int8 || val.Kind() == reflect.Int16 || val.Kind() == reflect.Int32 || val.Kind() == reflect.Int64 {
			return int(val.Int()), true
		}
	}
	return 0, false
}

func (fs *FlagSetExtend) GetUint(key string) (val uint, ok bool) {
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Uint || val.Kind() == reflect.Uint8 || val.Kind() == reflect.Uint16 || val.Kind() == reflect.Uint32 || val.Kind() == reflect.Uint64 {
			return uint(val.Uint()), true
		}
	}
	return 0, false
}
