package gobx

import (
	"sync"

	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

// NewGobBuffEncoder returns a Gob buffer encoder.
// NewGobBuffEncoder 创建 Gob 缓冲编码器。
func NewGobBuffEncoder(handler bytex.IDataBlockHandler) encodingx.IBuffEncoder {
	return newGobBuffCodecs(handler)
}

// NewGobBuffDecoder returns a Gob buffer decoder.
// NewGobBuffDecoder 创建 Gob 缓冲解码器。
func NewGobBuffDecoder(handler bytex.IDataBlockHandler) encodingx.IBuffDecoder {
	return newGobBuffCodecs(handler)
}

// NewGobBuffCodecs returns a combined Gob buffer codec.
// NewGobBuffCodecs 创建 Gob 缓冲编解码器。
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
