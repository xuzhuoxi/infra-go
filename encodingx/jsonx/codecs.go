package jsonx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

// NewJsonBuffEncoder returns a JSON buffer encoder.
// NewJsonBuffEncoder 创建 JSON 缓冲编码器。
func NewJsonBuffEncoder(handler bytex.IDataBlockHandler) encodingx.IBuffEncoder {
	return encodingx.NewBuffEncoder(handler, NewJsonCodingHandlerAsync())
}

// NewJsonBuffDecoder returns a JSON buffer decoder.
// NewJsonBuffDecoder 创建 JSON 缓冲解码器。
func NewJsonBuffDecoder(handler bytex.IDataBlockHandler) encodingx.IBuffDecoder {
	return encodingx.NewBuffDecoder(handler, NewJsonCodingHandlerAsync())
}

// NewJsonBuffCodecs returns a combined JSON buffer codec.
// NewJsonBuffCodecs 创建 JSON 缓冲编解码器。
func NewJsonBuffCodecs(handler bytex.IDataBlockHandler) encodingx.IBuffCodecs {
	return encodingx.NewBuffCodecs(handler, NewJsonCodingHandlerAsync())
}
