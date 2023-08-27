// Package cmdx
// Create on 2023/8/27
// @author xuzhuoxi
package cmdx

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ArgSet struct {
	kv map[string]string
}

func (o *ArgSet) Len() int {
	if nil == o.kv {
		return 0
	}
	return len(o.kv)
}

func (o *ArgSet) Get(key string) string {
	v, ok := o.getValue(key)
	if !ok {
		return ""
	}
	return v
}

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

func (o *ArgSet) CheckKey(key string) bool {
	_, ok := o.getValue(key)
	return ok
}

func (o *ArgSet) GetBool(key string) (val bool, ok bool) {
	v, ok1 := o.getValue(key)
	if !ok1 {
		return false, false
	}
	v = strings.ToLower(v)
	return v != "0" && v != "false", true
}

func (o *ArgSet) GetString(key string) (val string, ok bool) {
	return o.getValue(key)
}

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

func (o *ArgSet) GetInt(key string) (val int, ok bool) {
	i, ok1 := o.GetInt64(key)
	if !ok1 {
		return
	}
	return int(i), true
}

func (o *ArgSet) GetUint(key string) (val uint, ok bool) {
	i, ok1 := o.GetUint64(key)
	if !ok1 {
		return
	}
	return uint(i), true
}

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
