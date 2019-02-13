//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package encodingx

import (
	"encoding/binary"
	"github.com/xuzhuoxi/util-go/bytex"
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
	return newBuffCodecs(DefaultOrder, nil, bytex.DefaultDataToBlockHandler)
}

func NewDefaultBuffDecoder() IBuffDecoder {
	return newBuffCodecs(DefaultOrder, bytex.DefaultBlockToDataHandler, nil)
}

func NewDefaultBuffCodecs() IBuffCodecs {
	return newBuffCodecs(DefaultOrder, bytex.DefaultBlockToDataHandler, bytex.DefaultDataToBlockHandler)
}

func NewBuffEncoder(order binary.ByteOrder, data2block bytex.DataToBlockHandler) IBuffEncoder {
	return newBuffCodecs(order, nil, data2block)
}

func NewBuffDecoder(order binary.ByteOrder, block2data bytex.BlockToDataHandler) IBuffDecoder {
	return newBuffCodecs(order, block2data, nil)
}

func NewBuffCodecs(order binary.ByteOrder, block2data bytex.BlockToDataHandler, data2block bytex.DataToBlockHandler) IBuffCodecs {
	return newBuffCodecs(order, block2data, data2block)
}

//------------------------------------------

func newBuffCodecs(order binary.ByteOrder, block2data bytex.BlockToDataHandler, data2block bytex.DataToBlockHandler) *buffCodecs {
	return &buffCodecs{IBuffDataBlock: bytex.NewBuffDataBlock(order, data2block, block2data)}
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
