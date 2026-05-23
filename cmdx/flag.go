package cmdx

import (
	"errors"
	"flag"
	"os"
	"reflect"
	"strings"
	"time"
)

// NewDefaultFlagSetExtend creates a FlagSetExtend using os.Args[0] as the name and ContinueOnError handling.
// NewDefaultFlagSetExtend 使用 os.Args[0] 作为名称并以 ContinueOnError 模式创建 FlagSetExtend。
func NewDefaultFlagSetExtend() *FlagSetExtend {
	name := os.Args[0]
	errorHandling := flag.ContinueOnError
	return NewFlagSetExtend(name, errorHandling)
}

// NewFlagSetExtend creates a FlagSetExtend with the given name and error handling mode.
// NewFlagSetExtend 使用给定名称与错误处理模式创建 FlagSetExtend。
func NewFlagSetExtend(name string, errorHandling flag.ErrorHandling) *FlagSetExtend {
	return &FlagSetExtend{
		FlagSet: *flag.NewFlagSet(name, errorHandling), errorHandling: errorHandling}
}

// FlagSetExtend embeds flag.FlagSet and provides unified access to flags and key=value arguments.
// FlagSetExtend 嵌入 flag.FlagSet，并统一提供对 flag 与 key=value 参数的访问。
type FlagSetExtend struct {
	flag.FlagSet
	argSet ArgSet
	//keyList       []string
	errorHandling flag.ErrorHandling
}

// Parse parses predefined flags and key=value arguments.
// Parse 解析已预定义的标志与 key=value 参数。
// Only predefined parameters can be used.
// 只有预定义了参数才能使用。
func (fs *FlagSetExtend) Parse(arguments []string) error {
	if fs.FlagSet.Parsed() {
		return errors.New("Parsed! ")
	}
	err := fs.FlagSet.Parse(arguments)
	if nil != err {
		return err
	}
	return fs.argSet.ParseArgs(arguments)
}

// CheckKey reports whether the given key exists in flags or parsed arguments.
// CheckKey 检查给定键是否存在于 flag 或已解析参数中。
func (fs *FlagSetExtend) CheckKey(key string) bool {
	return fs.argSet.CheckKey(key)
}

// GetReflectValue returns the reflect.Value of a registered flag by key.
// GetReflectValue 按键返回已注册 flag 的 reflect.Value。
func (fs *FlagSetExtend) GetReflectValue(key string) (val reflect.Value, ok bool) {
	key = strings.ToLower(key)
	f := fs.Lookup(key)
	if nil != f {
		return reflect.ValueOf(f.Value).Elem(), true
	}
	return reflect.Value{}, false
}

// GetBool returns a bool value for the given key from arguments or registered flags.
// GetBool 从参数或已注册 flag 中获取给定键的 bool 值。
func (fs *FlagSetExtend) GetBool(key string) (val bool, ok bool) {
	if val, ok = fs.argSet.GetBool(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Bool {
			return val.Bool(), true
		}
	}
	return false, false
}

// GetString returns a string value for the given key from arguments or registered flags.
// GetString 从参数或已注册 flag 中获取给定键的 string 值。
func (fs *FlagSetExtend) GetString(key string) (val string, ok bool) {
	if val, ok = fs.argSet.GetString(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.String {
			return val.String(), true
		}
	}
	return "", false
}

// GetDuration returns a time.Duration value for the given key from arguments or registered flags.
// GetDuration 从参数或已注册 flag 中获取给定键的 time.Duration 值。
func (fs *FlagSetExtend) GetDuration(key string) (val time.Duration, ok bool) {
	if val, ok = fs.argSet.GetDuration(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Int64 {
			return time.Duration(val.Int()), true
		}
	}
	return 0, false
}

// GetFloat64 returns a float64 value for the given key from arguments or registered flags.
// GetFloat64 从参数或已注册 flag 中获取给定键的 float64 值。
func (fs *FlagSetExtend) GetFloat64(key string) (val float64, ok bool) {
	if val, ok = fs.argSet.GetFloat64(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Float64 || val.Kind() == reflect.Float32 {
			return float64(val.Float()), true
		}
	}
	return 0, false
}

// GetInt64 returns an int64 value for the given key from arguments or registered flags.
// GetInt64 从参数或已注册 flag 中获取给定键的 int64 值。
func (fs *FlagSetExtend) GetInt64(key string) (val int64, ok bool) {
	if val, ok = fs.argSet.GetInt64(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Int || val.Kind() == reflect.Int8 || val.Kind() == reflect.Int16 || val.Kind() == reflect.Int32 || val.Kind() == reflect.Int64 {
			return int64(val.Int()), true
		}
	}
	return 0, false
}

// GetUint64 returns a uint64 value for the given key from arguments or registered flags.
// GetUint64 从参数或已注册 flag 中获取给定键的 uint64 值。
func (fs *FlagSetExtend) GetUint64(key string) (val uint64, ok bool) {
	if val, ok = fs.argSet.GetUint64(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Uint || val.Kind() == reflect.Uint8 || val.Kind() == reflect.Uint16 || val.Kind() == reflect.Uint32 || val.Kind() == reflect.Uint64 {
			return uint64(val.Uint()), true
		}
	}
	return 0, false
}

// GetInt returns an int value for the given key from arguments or registered flags.
// GetInt 从参数或已注册 flag 中获取给定键的 int 值。
func (fs *FlagSetExtend) GetInt(key string) (val int, ok bool) {
	if val, ok = fs.argSet.GetInt(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Int || val.Kind() == reflect.Int8 || val.Kind() == reflect.Int16 || val.Kind() == reflect.Int32 || val.Kind() == reflect.Int64 {
			return int(val.Int()), true
		}
	}
	return 0, false
}

// GetUint returns a uint value for the given key from arguments or registered flags.
// GetUint 从参数或已注册 flag 中获取给定键的 uint 值。
func (fs *FlagSetExtend) GetUint(key string) (val uint, ok bool) {
	if val, ok = fs.argSet.GetUint(key); ok {
		return
	}
	if val, ok := fs.GetReflectValue(key); ok {
		if val.Kind() == reflect.Uint || val.Kind() == reflect.Uint8 || val.Kind() == reflect.Uint16 || val.Kind() == reflect.Uint32 || val.Kind() == reflect.Uint64 {
			return uint(val.Uint()), true
		}
	}
	return 0, false
}
