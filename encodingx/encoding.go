// Package encodingx
// Created by xuzhuoxi
// on 2019-02-11.
// @author xuzhuoxi
//
package encodingx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
)

//-------------------------------

type IEncodingData interface {
	// EncodeToBytes 序列化
	EncodeToBytes() []byte
}

type IDecodingData interface {
	// DecodeFromBytes 反序列化更新
	DecodeFromBytes(bs []byte) bool
}

type ICodingData interface {
	IEncodingData
	IDecodingData
}

//------------------------------------

type IEncodeHandler interface {
	HandleEncode(data interface{}) []byte
}

type IDecodeHandler interface {
	HandleDecode(bs []byte, data interface{})
}

type ICodingHandler interface {
	IEncodeHandler
	IDecodeHandler
}

//------------------------------------

type iBuffEncoder interface {
	bytex.IBuffByteReader
	bytex.IBuffDataWriter
	EncodeDataToBuff(data ...interface{})
}

type iBuffDecoder interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataReader
	DecodeDataFromBuff(data ...interface{})
}

type IBuffEncoder interface {
	iBuffEncoder
	bytex.IBuffReset
	bytex.IBuffLen
}

type IBuffDecoder interface {
	iBuffDecoder
	bytex.IBuffReset
	bytex.IBuffLen
}

type IBuffCodecs interface {
	iBuffEncoder
	iBuffDecoder
	bytex.IBuffReset
	bytex.IBuffLen
}
