// Package jsonx
// Created by xuzhuoxi
// on 2019-02-24.
// @author xuzhuoxi
//
package jsonx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

func NewJsonBuffEncoder(handler bytex.IDataBlockHandler) encodingx.IBuffEncoder {
	return encodingx.NewBuffEncoder(handler, NewJsonCodingHandlerAsync())
}

func NewJsonBuffDecoder(handler bytex.IDataBlockHandler) encodingx.IBuffDecoder {
	return encodingx.NewBuffDecoder(handler, NewJsonCodingHandlerAsync())
}

func NewJsonBuffCodecs(handler bytex.IDataBlockHandler) encodingx.IBuffCodecs {
	return encodingx.NewBuffCodecs(handler, NewJsonCodingHandlerAsync())
}
