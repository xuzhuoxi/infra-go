package gobx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
	"sync"
)

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
		bytes, err := b.gobHandler.HandleEncode(data[index])
		if nil != err {
			return
		}
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
		err := b.gobHandler.HandleDecode(bytes, data[index])
		if nil != err {
			return
		}
		//fmt.Println("DecodeDataFromBuff:", bytes, data[index])
	}
}
