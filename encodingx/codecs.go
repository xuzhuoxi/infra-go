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

func (bc *buffCodecs) EncodeToBuff(encoders ...interface{}) {
	if len(encoders) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for _, encoder := range encoders {
		bc.WriteData(encoder.(IEncodingData).EncodeToBytes())
	}
}

func (bc *buffCodecs) DecodeFromBuff(decoders ...interface{}) {
	if len(decoders) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for _, decoder := range decoders {
		decoder.(IDecodingData).DecodeFromBytes(bc.ReadData())
	}
}
