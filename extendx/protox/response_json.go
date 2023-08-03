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

func (resp *SockResponse) AppendJson(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		jsonStr, err1 := resp.toJson(data[index])
		if nil != err1 {
			return err1
		}
		err2 := resp.AppendString(jsonStr)
		if nil != err2 {
			return err2
		}
	}
	return nil
}

func (resp *SockResponse) SendJsonResponse(data ...interface{}) error {
	resp.PrepareResponse()
	err := resp.AppendJson(data...)
	if nil != err {
		return err
	}
	return resp.SendResponse()
}

func (resp *SockResponse) toJson(o interface{}) (json string, err error) {
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
