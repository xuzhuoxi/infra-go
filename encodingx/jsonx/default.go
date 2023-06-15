// Package jsonx
// Created by xuzhuoxi
// on 2019-03-25.
// @author xuzhuoxi
//
package jsonx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

var (
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

	DefaultPoolJsonCodingHandler = encodingx.NewPoolCodingHandler()
	DefaultPoolJsonBuffEncoder   = encodingx.NewPoolBuffEncoder()
	DefaultPoolJsonBuffDecoder   = encodingx.NewPoolBuffDecoder()
	DefaultPoolJsonBuffCodecs    = encodingx.NewPoolBuffCodecs()
)

func init() {
	DefaultPoolJsonCodingHandler.Register(NewDefaultJsonCodingHandler)
	DefaultPoolJsonBuffEncoder.Register(NewDefaultJsonBuffEncoder)
	DefaultPoolJsonBuffDecoder.Register(NewDefaultJsonBuffDecoder)
	DefaultPoolJsonBuffCodecs.Register(NewDefaultJsonBuffCodecs)
}

func NewDefaultJsonCodingHandler() encodingx.ICodingHandler {
	return NewJsonCodingHandlerAsync()
}

func NewDefaultJsonBuffEncoder() encodingx.IBuffEncoder {
	return NewJsonBuffEncoder(DefaultDataBlockHandler)
}

func NewDefaultJsonBuffDecoder() encodingx.IBuffDecoder {
	return NewJsonBuffDecoder(DefaultDataBlockHandler)
}

func NewDefaultJsonBuffCodecs() encodingx.IBuffCodecs {
	return NewJsonBuffCodecs(DefaultDataBlockHandler)
}
