package encodingx

import (
	"encoding/binary"

	"github.com/xuzhuoxi/infra-go/bytex"
)

var (
	// DefaultOrder is the default byte order for binary I/O.
	// DefaultOrder 是二进制读写的默认字节序。
	DefaultOrder binary.ByteOrder = bytex.DefaultOrder
	// DefaultDataBlockHandler is the default bytex data block handler.
	// DefaultDataBlockHandler 是默认的 bytex 数据块处理器。
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()
)

// NewDefaultBuffEncoder returns a buffer encoder with DefaultDataBlockHandler.
// NewDefaultBuffEncoder 使用 DefaultDataBlockHandler 创建缓冲编码器。
func NewDefaultBuffEncoder(encodeHandler IEncodeHandler) IBuffEncoder {
	return newBuffCodecs(DefaultDataBlockHandler, encodeHandler, nil)
}

// NewDefaultBuffDecoder returns a buffer decoder with DefaultDataBlockHandler.
// NewDefaultBuffDecoder 使用 DefaultDataBlockHandler 创建缓冲解码器。
func NewDefaultBuffDecoder(decodeHandler IDecodeHandler) IBuffDecoder {
	return newBuffCodecs(DefaultDataBlockHandler, nil, decodeHandler)
}

// NewDefaultBuffCodecs returns a combined buffer codec with DefaultDataBlockHandler.
// NewDefaultBuffCodecs 使用 DefaultDataBlockHandler 创建缓冲编解码器。
func NewDefaultBuffCodecs(codingHandler ICodingHandler) IBuffCodecs {
	return newBuffCodecs(DefaultDataBlockHandler, codingHandler, codingHandler)
}
