// Package yamlx
// Created by xuzhuoxi
// on 2019-03-25.
// @author xuzhuoxi
//
package yamlx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

var (
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

	DefaultPoolYamlCodingHandler = encodingx.NewPoolCodingHandler()
	DefaultPoolYamlBuffEncoder   = encodingx.NewPoolBuffEncoder()
	DefaultPoolYamlBuffDecoder   = encodingx.NewPoolBuffDecoder()
	DefaultPoolYamlBuffCodecs    = encodingx.NewPoolBuffCodecs()
)

func init() {
	DefaultPoolYamlCodingHandler.Register(NewDefaultYamlCodingHandler)
	DefaultPoolYamlBuffEncoder.Register(NewDefaultYamlBuffEncoder)
	DefaultPoolYamlBuffDecoder.Register(NewDefaultYamlBuffDecoder)
	DefaultPoolYamlBuffCodecs.Register(NewDefaultYamlBuffCodecs)
}

func NewDefaultYamlCodingHandler() encodingx.ICodingHandler {
	return NewYamlCodingHandlerAsync()
}

func NewDefaultYamlBuffEncoder() encodingx.IBuffEncoder {
	return NewYamlBuffEncoder(DefaultDataBlockHandler)
}

func NewDefaultYamlBuffDecoder() encodingx.IBuffDecoder {
	return NewYamlBuffDecoder(DefaultDataBlockHandler)
}

func NewDefaultYamlBuffCodecs() encodingx.IBuffCodecs {
	return NewYamlBuffCodecs(DefaultDataBlockHandler)
}
