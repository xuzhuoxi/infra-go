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

func (c jsonHandlerAsync) HandleEncode(data interface{}) (bs []byte, err error) {
	return jsoniter.Marshal(data)
}

func (c jsonHandlerAsync) HandleDecode(bs []byte, data interface{}) error {
	return jsoniter.Unmarshal(bs, data)
}
