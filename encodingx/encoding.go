//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package encodingx

import (
	"encoding/binary"
	"github.com/xuzhuoxi/infra-go/bytex"
)

var DefaultOrder = binary.BigEndian
var DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

//-------------------------------

type IEncodingData interface {
	//序列化
	EncodeToBytes() []byte
}

type IDecodingData interface {
	//反序列化更新
	DecodeFromBytes(bs []byte) bool
}

type ICodingData interface {
	IEncodingData
	IDecodingData
}

//------------------------------------

type IDataEncodeHandler interface {
	HandleEncode(data interface{}) []byte
}

type IDataDecodeHandler interface {
	HandleDecode(bs []byte) interface{}
}

type IDataCodeHandler interface {
	IDataEncodeHandler
	IDataDecodeHandler
}
