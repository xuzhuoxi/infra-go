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

type IEncoder interface {
	//序列化
	Encode() []byte
}

type IDecoder interface {
	//反序列化更新
	Decode([]byte)
}

type ICodecs interface {
	IEncoder
	IDecoder
}

//------------------------------------------------

type IBuffEncoder interface {
	IBuffByteReader
	IBuffDataWriter
	IBuffReset
	EncodeToBuff(encoders ...IEncoder)
}

type IBuffDecoder interface {
	IBuffByteWriter
	IBuffDataReader
	IBuffReset
	DecodeFromBuff(decoders ...IDecoder)
}

type IBuffCodecs interface {
	IBuffByteWriter
	IBuffDataWriter
	IBuffByteReader
	IBuffDataReader
	IBuffReset
	EncodeToBuff(encoders ...IEncoder)
	DecodeFromBuff(decoders ...IDecoder)
}

func NewBuffEncoder(order binary.ByteOrder) IBuffEncoder {
	return newBuffCodecs(order)
}

func NewBuffDecoder(order binary.ByteOrder) IBuffDecoder {
	return newBuffCodecs(order)
}

func NewBuffCodecs(order binary.ByteOrder) IBuffCodecs {
	return newBuffCodecs(order)
}

func newBuffCodecs(order binary.ByteOrder) *buffCodecs {
	return &buffCodecs{buffBase: newBuffBase(order)}
}

type buffCodecs struct {
	buffBase
	codecsLock sync.RWMutex
}

func (bc *buffCodecs) EncodeToBuff(encoders ...IEncoder) {
	if len(encoders) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for _, encoder := range encoders {
		bc.WriteData(encoder.Encode())
	}
}

func (bc *buffCodecs) DecodeFromBuff(decoders ...IDecoder) {
	if len(decoders) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for _, decoder := range decoders {
		decoder.Decode(bc.ReadData())
	}
}
