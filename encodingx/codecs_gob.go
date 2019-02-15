package encodingx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"sync"
)

type IGobBuffEncoder interface {
	bytex.IBuffByteReader
	bytex.IBuffDataWriter
	bytex.IBuffReset
	EncodeToBuff(data ...interface{})
}

type IGobBuffDecoder interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataReader
	bytex.IBuffReset
	DecodeFromBuff(data ...interface{})
}

type IGobBuffCodecs interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataWriter
	bytex.IBuffByteReader
	bytex.IBuffDataReader
	bytex.IBuffReset
	EncodeToBuff(data ...interface{})
	DecodeFromBuff(data ...interface{})
}

func NewDefaultGobBuffEncoder() IGobBuffEncoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffDecoder() IGobBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffCodecs() IGobBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewGobBuffEncoder(handler bytex.IDataBlockHandler) IGobBuffEncoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffDecoder(handler bytex.IDataBlockHandler) IGobBuffDecoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffCodecs(handler bytex.IDataBlockHandler) IGobBuffDecoder {
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
