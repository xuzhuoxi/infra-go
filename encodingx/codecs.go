//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package encodingx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"sync"
)

type IBuffEncoder interface {
	bytex.IBuffByteReader
	bytex.IBuffDataWriter
	bytex.IBuffReset
	EncodeToBuff(encoders ...IEncodingData)
}

type IBuffDecoder interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataReader
	bytex.IBuffReset
	DecodeFromBuff(decoders ...IDecodingData)
}

type IBuffCodecs interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataWriter
	bytex.IBuffByteReader
	bytex.IBuffDataReader
	bytex.IBuffReset
	EncodeToBuff(encoders ...IEncodingData)
	DecodeFromBuff(decoders ...IDecodingData)
}

func NewDefaultBuffEncoder() IBuffEncoder {
	return newBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultBuffDecoder() IBuffDecoder {
	return newBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultBuffCodecs() IBuffCodecs {
	return newBuffCodecs(DefaultDataBlockHandler)
}

func NewBuffEncoder(handler bytex.IDataBlockHandler) IBuffEncoder {
	return newBuffCodecs(handler)
}

func NewBuffDecoder(handler bytex.IDataBlockHandler) IBuffDecoder {
	return newBuffCodecs(handler)
}

func NewBuffCodecs(handler bytex.IDataBlockHandler) IBuffCodecs {
	return newBuffCodecs(handler)
}

//------------------------------------------

func newBuffCodecs(handler bytex.IDataBlockHandler) *buffCodecs {
	return &buffCodecs{IBuffDataBlock: bytex.NewBuffDataBlock(handler)}
}

type buffCodecs struct {
	bytex.IBuffDataBlock
	codecsLock sync.RWMutex
}

func (bc *buffCodecs) EncodeToBuff(encoders ...IEncodingData) {
	if len(encoders) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for _, encoder := range encoders {
		bc.WriteData(encoder.EncodeToBytes())
	}
}

func (bc *buffCodecs) DecodeFromBuff(decoders ...IDecodingData) {
	if len(decoders) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for _, decoder := range decoders {
		decoder.DecodeFromBytes(bc.ReadData())
	}
}
