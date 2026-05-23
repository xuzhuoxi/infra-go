package encodingx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
)

//-------------------------------

// IEncodingData encodes the receiver to a byte slice.
// IEncodingData 将接收者编码为字节切片。
type IEncodingData interface {
	// EncodeToBytes serializes the receiver to bytes.
	// EncodeToBytes 将接收者序列化为字节。
	EncodeToBytes() (bs []byte, err error)
}

// IDecodingData decodes from a byte slice into the receiver.
// IDecodingData 从字节切片反序列化并更新接收者。
type IDecodingData interface {
	// DecodeFromBytes deserializes from bytes into the receiver.
	// DecodeFromBytes 从字节反序列化并更新接收者。
	DecodeFromBytes(bs []byte) error
}

// ICodingData combines encoding and decoding to byte slices.
// ICodingData 组合了与字节切片互转的编码与解码能力。
type ICodingData interface {
	IEncodingData
	IDecodingData
}

//------------------------------------

// IEncodeHandler encodes arbitrary values to byte slices.
// IEncodeHandler 将任意值编码为字节切片。
type IEncodeHandler interface {
	// HandleEncode encodes data to bytes.
	// HandleEncode 将 data 编码为字节。
	HandleEncode(data interface{}) (bs []byte, err error)
}

// IDecodeHandler decodes byte slices into target values.
// IDecodeHandler 将字节切片解码到目标值。
type IDecodeHandler interface {
	// HandleDecode decodes bytes into data.
	// HandleDecode 将字节解码到 data。
	HandleDecode(bs []byte, data interface{}) error
}

// ICodingHandler combines encode and decode handlers.
// ICodingHandler 组合编码与解码处理器。
type ICodingHandler interface {
	IEncodeHandler
	IDecodeHandler
}

//------------------------------------

type iBuffEncoder interface {
	bytex.IBuffByteReader
	bytex.IBuffDataWriter
	EncodeDataToBuff(data ...interface{})
}

type iBuffDecoder interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataReader
	DecodeDataFromBuff(data ...interface{})
}

// IBuffEncoder writes encoded values into a bytex buffer.
// IBuffEncoder 将编码后的值写入 bytex 缓冲区。
type IBuffEncoder interface {
	iBuffEncoder
	bytex.IBuffReset
	bytex.IBuffLen
}

// IBuffDecoder reads encoded values from a bytex buffer.
// IBuffDecoder 从 bytex 缓冲区读取并解码值。
type IBuffDecoder interface {
	iBuffDecoder
	bytex.IBuffReset
	bytex.IBuffLen
}

// IBuffCodecs combines buff encoding and decoding.
// IBuffCodecs 组合缓冲区的编码与解码能力。
type IBuffCodecs interface {
	iBuffEncoder
	iBuffDecoder
	bytex.IBuffReset
	bytex.IBuffLen
}
