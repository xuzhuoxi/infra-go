package gobx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
	"sync"
)

var DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

func NewDefaultGobBuffEncoder() encodingx.IBuffEncoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffDecoder() encodingx.IBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffCodecs() encodingx.IBuffCodecs {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewGobBuffEncoder(handler bytex.IDataBlockHandler) encodingx.IBuffEncoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffDecoder(handler bytex.IDataBlockHandler) encodingx.IBuffDecoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffCodecs(handler bytex.IDataBlockHandler) encodingx.IBuffCodecs {
	return newGobBuffCodecs(handler)
}

//-------------------------------------

func newGobBuffCodecs(handler bytex.IDataBlockHandler) *gobBuffCodecs {
	return &gobBuffCodecs{IBuffDataBlock: bytex.NewBuffDataBlock(handler), gobHandler: NewDefaultGobCodingHandler()}
}

type gobBuffCodecs struct {
	bytex.IBuffDataBlock
	gobHandler encodingx.ICodingHandler
	codecsLock sync.RWMutex
}

func (b *gobBuffCodecs) EncodeDataToBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	b.codecsLock.Lock()
	defer b.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.gobHandler.HandleEncode(data[index])
		b.WriteData(bytes)
		//fmt.Println("EncodeDataToBuff:", bytes, data[index])
	}
}

func (b *gobBuffCodecs) DecodeDataFromBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	b.codecsLock.Lock()
	defer b.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.ReadData()
		b.gobHandler.HandleDecode(bytes, data[index])
		//fmt.Println("DecodeDataFromBuff:", bytes, data[index])
	}
}
