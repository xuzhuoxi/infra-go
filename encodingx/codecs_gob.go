package encodingx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"sync"
)

func NewDefaultGobBuffEncoder() IBuffEncoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffDecoder() IBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffCodecs() IBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewGobBuffEncoder(handler bytex.IDataBlockHandler) IBuffEncoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffDecoder(handler bytex.IDataBlockHandler) IBuffDecoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffCodecs(handler bytex.IDataBlockHandler) IBuffDecoder {
	return newGobBuffCodecs(handler)
}

//-------------------------------------

func newGobBuffCodecs(handler bytex.IDataBlockHandler) *gobBuffCodecs {
	return &gobBuffCodecs{IBuffDataBlock: bytex.NewBuffDataBlock(handler), handler: NewGobCodingHandler()}
}

type gobBuffCodecs struct {
	bytex.IBuffDataBlock
	handler    ICodingHandler
	codecsLock sync.RWMutex
}

func (b *gobBuffCodecs) EncodeToBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	b.codecsLock.Lock()
	defer b.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.handler.HandleEncode(data[index])
		b.WriteData(bytes)
	}
}

func (b *gobBuffCodecs) DecodeFromBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	b.codecsLock.Lock()
	defer b.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.ReadData()
		b.handler.HandleDecode(bytes, data[index])
	}
}
