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

func NewBuffEncoder(handler bytex.IDataBlockHandler, encodeHandler IEncodeHandler) IBuffEncoder {
	return newBuffCodecs(handler, encodeHandler, nil)
}

func NewBuffDecoder(handler bytex.IDataBlockHandler, decodeHandler IDecodeHandler) IBuffDecoder {
	return newBuffCodecs(handler, nil, decodeHandler)
}

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
			bc.WriteData(cd.EncodeToBytes())
		} else {
			bc.WriteData(bc.encodeHandler.HandleEncode(data[index]))
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
