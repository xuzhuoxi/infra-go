package cmdx

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ArgSet stores key=value arguments parsed from -key=value or --key=value syntax.
// ArgSet 存储从 -key=value 或 --key=value 语法解析出的键值参数。
type ArgSet struct {
	kv map[string]string
}

// Len returns the number of parsed key=value arguments.
// Len 返回已解析的 key=value 参数数量。
func (o *ArgSet) Len() int {
	if nil == o.kv {
		return 0
	}
	return len(o.kv)
}

// Get returns the string value for the given key, or an empty string if not found.
// Get 返回给定键的字符串值；未找到时返回空字符串。
func (o *ArgSet) Get(key string) string {
	v, ok := o.getValue(key)
	if !ok {
		return ""
	}
	return v
}

// ParseArgs parses a slice of -key=value or --key=value argument strings.
// ParseArgs 解析一组 -key=value 或 --key=value 参数字符串。
func (o *ArgSet) ParseArgs(args []string) error {
	o.kv = make(map[string]string)
	if len(args) == 0 {
		return nil
	}
	for _, s := range args {
		name, value, ok, err := o.parseOne(s)
		if err != nil {
			return err
		}
		if !ok {
			continue
		}
		o.kv[strings.ToLower(name)] = value
	}
	return nil
}

// CheckKey reports whether the given key exists in parsed arguments.
// CheckKey 检查给定键是否存在于已解析参数中。
func (o *ArgSet) CheckKey(key string) bool {
	_, ok := o.getValue(key)
	return ok
}

// GetBool returns a bool value for the given key; "0" and "false" are treated as false.
// GetBool 返回给定键的 bool 值；"0" 与 "false" 视为 false。
func (o *ArgSet) GetBool(key string) (val bool, ok bool) {
	v, ok1 := o.getValue(key)
	if !ok1 {
		return false, false
	}
	v = strings.ToLower(v)
	return v != "0" && v != "false", true
}

// GetString returns the string value for the given key.
// GetString 返回给定键的 string 值。
func (o *ArgSet) GetString(key string) (val string, ok bool) {
	return o.getValue(key)
}

// GetDuration returns a time.Duration value parsed from the given key's string value.
// GetDuration 从给定键的字符串值解析并返回 time.Duration。
func (o *ArgSet) GetDuration(key string) (val time.Duration, ok bool) {
	v, ok1 := o.getValue(key)
	if !ok1 {
		return
	}
	d, err := time.ParseDuration(v)
	if nil != err {
		return
	}
	return d, true
}

// GetInt returns an int value parsed from the given key's string value.
// GetInt 从给定键的字符串值解析并返回 int。
func (o *ArgSet) GetInt(key string) (val int, ok bool) {
	i, ok1 := o.GetInt64(key)
	if !ok1 {
		return
	}
	return int(i), true
}

// GetUint returns a uint value parsed from the given key's string value.
// GetUint 从给定键的字符串值解析并返回 uint。
func (o *ArgSet) GetUint(key string) (val uint, ok bool) {
	i, ok1 := o.GetUint64(key)
	if !ok1 {
		return
	}
	return uint(i), true
}

// GetFloat64 returns a float64 value parsed from the given key's string value.
// GetFloat64 从给定键的字符串值解析并返回 float64。
func (o *ArgSet) GetFloat64(key string) (val float64, ok bool) {
	v, ok1 := o.getValue(key)
	if !ok1 {
		return
	}
	i, err := strconv.ParseFloat(v, 64)
	if nil != err {
		return
	}
	return i, true
}

// GetInt64 returns an int64 value parsed from the given key's string value.
// GetInt64 从给定键的字符串值解析并返回 int64。
func (o *ArgSet) GetInt64(key string) (val int64, ok bool) {
	v, ok1 := o.getValue(key)
	if !ok1 {
		return
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if nil != err {
		return
	}
	return i, true
}

// GetUint64 returns a uint64 value parsed from the given key's string value.
// GetUint64 从给定键的字符串值解析并返回 uint64。
func (o *ArgSet) GetUint64(key string) (val uint64, ok bool) {
	v, ok1 := o.getValue(key)
	if !ok1 {
		return
	}
	i, err := strconv.ParseUint(v, 10, 64)
	if nil != err {
		return
	}
	return i, true
}

func (o *ArgSet) getValue(key string) (value string, ok bool) {
	value, ok = o.kv[strings.ToLower(key)]
	return
}

func (o *ArgSet) parseOne(s string) (name, value string, ok bool, err error) {
	if len(s) < 2 || s[0] != '-' {
		ok = false
		return
	}
	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 {
			ok = false
			return
		}
	}
	name = s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		ok, err = false, fmt.Errorf("bad flag syntax: %s", s)
		return
	}
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			value = name[i+1:]
			name = name[0:i]
			ok = true
		}
	}
	return
}
