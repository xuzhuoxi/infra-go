// Package encodingx
// Created by xuzhuoxi
// on 2019-03-20.
// @author xuzhuoxi
//
package encodingx

import (
	"encoding/binary"
	"github.com/xuzhuoxi/infra-go/bytex"
)

var (
	DefaultOrder            = binary.BigEndian
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()
)

func NewDefaultBuffEncoder(encodeHandler IEncodeHandler) IBuffEncoder {
	return newBuffCodecs(DefaultDataBlockHandler, encodeHandler, nil)
}

func NewDefaultBuffDecoder(decodeHandler IDecodeHandler) IBuffDecoder {
	return newBuffCodecs(DefaultDataBlockHandler, nil, decodeHandler)
}

func NewDefaultBuffCodecs(codingHandler ICodingHandler) IBuffCodecs {
	return newBuffCodecs(DefaultDataBlockHandler, codingHandler, codingHandler)
}
