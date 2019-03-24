//
//Created by xuzhuoxi
//on 2019-03-25.
//@author xuzhuoxi
//
package gobx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

var (
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

	DefaultPoolGobCodingHandler = encodingx.NewPoolCodingHandler()
	DefaultPoolGobBuffEncoder   = encodingx.NewPoolBuffEncoder()
	DefaultPoolGobBuffDecoder   = encodingx.NewPoolBuffDecoder()
	DefaultPoolGobBuffCodecs    = encodingx.NewPoolBuffCodecs()
)

func init() {
	DefaultPoolGobCodingHandler.Register(NewDefaultGobCodingHandler)
	DefaultPoolGobBuffEncoder.Register(NewDefaultGobBuffEncoder)
	DefaultPoolGobBuffDecoder.Register(NewDefaultGobBuffDecoder)
	DefaultPoolGobBuffCodecs.Register(NewDefaultGobBuffCodecs)
}

func NewDefaultGobCodingHandler() encodingx.ICodingHandler {
	return NewGobCodingHandlerAsync()
	//return NewGobCodingHandler()
}

func NewDefaultGobBuffEncoder() encodingx.IBuffEncoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffDecoder() encodingx.IBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffCodecs() encodingx.IBuffCodecs {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}
