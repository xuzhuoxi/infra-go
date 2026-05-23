package encodingx

import (
	"sync"

	"github.com/xuzhuoxi/infra-go/bytex"
)

// NewBuffEncoder returns a buffer encoder using handler and encodeHandler.
// NewBuffEncoder 使用 handler 与 encodeHandler 创建缓冲编码器。
func NewBuffEncoder(handler bytex.IDataBlockHandler, encodeHandler IEncodeHandler) IBuffEncoder {
	return newBuffCodecs(handler, encodeHandler, nil)
}

// NewBuffDecoder returns a buffer decoder using handler and decodeHandler.
// NewBuffDecoder 使用 handler 与 decodeHandler 创建缓冲解码器。
func NewBuffDecoder(handler bytex.IDataBlockHandler, decodeHandler IDecodeHandler) IBuffDecoder {
	return newBuffCodecs(handler, nil, decodeHandler)
}

// NewBuffCodecs returns a combined buffer encoder and decoder.
// NewBuffCodecs 创建同时具备编码与解码能力的缓冲编解码器。
func NewBuffCodecs(handler bytex.IDataBlockHandler, codingHandler ICodingHandler) IBuffCodecs {
	return newBuffCodecs(handler, codingHandler, codingHandler)
}

//------------------------------------------

func newBuffCodecs(handler bytex.IDataBlockHandler, encodeHandler IEncodeHandler, decodeHandler IDecodeHandler) *buffCodecs {
	return &buffCodecs{IBuffDataBlock: bytex.NewBuffDataBlock(handler), encodeHandler: encodeHandler, decodeHandler: decodeHandler}
}

type buffCodecs struct {
	bytex.IBuffDataBlock
	encodeHandler IEncodeHandler
	decodeHandler IDecodeHandler
	codecsLock    sync.RWMutex
}

func (bc *buffCodecs) EncodeDataToBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		if cd, ok := data[index].(IEncodingData); ok {
			bs, _ := cd.EncodeToBytes()
			bc.WriteData(bs)
		} else {
			bs, _ := bc.encodeHandler.HandleEncode(data[index])
			bc.WriteData(bs)
		}
	}
}

func (bc *buffCodecs) DecodeDataFromBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	bc.codecsLock.Lock()
	defer bc.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		readData := bc.ReadData()
		if cd, ok := data[index].(IDecodingData); ok {
			cd.DecodeFromBytes(readData)
		} else {
			bc.decodeHandler.HandleDecode(readData, data[index])
		}
	}
}
