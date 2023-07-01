// Package jsonx
// Created by xuzhuoxi
// on 2019-02-24.
// @author xuzhuoxi
//
package jsonx

import (
	"github.com/json-iterator/go"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

func NewJsonCodingHandlerAsync() encodingx.ICodingHandler {
	return jsonHandlerAsync{}
}

func NewJsonCodingHandlerSync() encodingx.ICodingHandler {
	return jsonHandlerAsync{}
}

type jsonHandlerAsync struct{}

func (c jsonHandlerAsync) HandleEncode(data interface{}) []byte {
	bs, err := jsoniter.Marshal(data)
	if nil != err {
		return nil
	}
	return bs
}

func (c jsonHandlerAsync) HandleDecode(bs []byte, data interface{}) error {
	return jsoniter.Unmarshal(bs, data)
}
