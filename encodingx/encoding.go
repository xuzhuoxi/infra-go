//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package encodingx

import (
	"encoding/binary"
	"sync"
)

//-------------------------------------------

type IBuffEncoder interface {
	IBuffEncoded
	EncodeToBuff(encoders ...IEncoder)
}

type IBuffDecoder interface {
	IBuffDecoded
	DecodeFromBuff(decoders ...IDecoder)
}

type IBuffCodec interface {
	IBuffEncoder
	IBuffDecoder
}

func NewBuffEncoder(order binary.ByteOrder) IBuffEncoder {
	return newBuffCodec(order)
}

func NewBuffDecoder(order binary.ByteOrder) IBuffDecoder {
	return newBuffCodec(order)
}

func NewBuffCodec(order binary.ByteOrder) IBuffCodec {
	return newBuffCodec(order)
}

func newBuffCodec(order binary.ByteOrder) *buffCodec {
	return &buffCodec{base: NewBuffBase(order)}
}

type buffCodec struct {
	base     IBuffBase
	buffLock sync.RWMutex
}

func (bc *buffCodec) EncodedBytes() []byte {
	return bc.base.ClearBytes()
}

func (bc *buffCodec) EncodeToBuff(encoders ...IEncoder) {
	if len(encoders) == 0 {
		return
	}
	bc.buffLock.Lock()
	defer bc.buffLock.Unlock()
	for _, encoder := range encoders {
		bc.base.WriteData(encoder.Encode())
	}
}

func (bc *buffCodec) DecodedBytes(bytes []byte) {
	bc.base.AppendBytes(bytes)
}

func (bc *buffCodec) DecodeFromBuff(decoders ...IDecoder) {
	if len(decoders) == 0 {
		return
	}
	bc.buffLock.Lock()
	defer bc.buffLock.Unlock()
	for _, decoder := range decoders {
		decoder.Decode(bc.base.ReadData())
	}
}

//---------------------------------

type IEncoder interface {
	//序列化
	Encode() []byte
}

type IDecoder interface {
	//反序列化更新
	Decode([]byte)
}

type ICodec interface {
	IEncoder
	IDecoder
}
