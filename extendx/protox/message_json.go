// Package protox
// Create on 2023/7/22
// @author xuzhuoxi
package protox

import (
	"encoding/json"
)

var (
	marshalHandler func(v interface{}) ([]byte, error)
)

func jsonMarshal(obj interface{}) ([]byte, error) {
	if nil == marshalHandler {
		return json.Marshal(obj)
	} else {
		return marshalHandler(obj)
	}
}

func SetJsonMarshalHandler(handler func(v interface{}) ([]byte, error)) {
	marshalHandler = handler
}

func toJson(o interface{}) (json string, err error) {
	switch d := o.(type) {
	case string:
		return d, nil
	default:
		bs, err1 := jsonMarshal(d)
		if nil != err1 {
			return "", err1
		}
		return string(bs), nil
	}
}
