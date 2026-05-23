package jsonx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

var (
	// DefaultDataBlockHandler is the default bytex data block handler for JSON codecs.
	// DefaultDataBlockHandler 是 JSON 编解码使用的默认 bytex 数据块处理器。
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

	// DefaultPoolJsonCodingHandler is the default pool for JSON ICodingHandler instances.
	// DefaultPoolJsonCodingHandler 是 JSON 编解码处理器的默认对象池。
	DefaultPoolJsonCodingHandler = encodingx.NewPoolCodingHandler()
	// DefaultPoolJsonBuffEncoder is the default pool for JSON IBuffEncoder instances.
	// DefaultPoolJsonBuffEncoder 是 JSON 缓冲编码器的默认对象池。
	DefaultPoolJsonBuffEncoder = encodingx.NewPoolBuffEncoder()
	// DefaultPoolJsonBuffDecoder is the default pool for JSON IBuffDecoder instances.
	// DefaultPoolJsonBuffDecoder 是 JSON 缓冲解码器的默认对象池。
	DefaultPoolJsonBuffDecoder = encodingx.NewPoolBuffDecoder()
	// DefaultPoolJsonBuffCodecs is the default pool for JSON IBuffCodecs instances.
	// DefaultPoolJsonBuffCodecs 是 JSON 缓冲编解码器的默认对象池。
	DefaultPoolJsonBuffCodecs = encodingx.NewPoolBuffCodecs()
)

func init() {
	DefaultPoolJsonCodingHandler.Register(NewDefaultJsonCodingHandler)
	DefaultPoolJsonBuffEncoder.Register(NewDefaultJsonBuffEncoder)
	DefaultPoolJsonBuffDecoder.Register(NewDefaultJsonBuffDecoder)
	DefaultPoolJsonBuffCodecs.Register(NewDefaultJsonBuffCodecs)
}

// NewDefaultJsonCodingHandler returns the default JSON ICodingHandler from the pool factory.
// NewDefaultJsonCodingHandler 返回默认 JSON 编解码处理器。
func NewDefaultJsonCodingHandler() encodingx.ICodingHandler {
	return NewJsonCodingHandlerAsync()
}

// NewDefaultJsonBuffEncoder returns a JSON buffer encoder with DefaultDataBlockHandler.
// NewDefaultJsonBuffEncoder 使用 DefaultDataBlockHandler 创建 JSON 缓冲编码器。
func NewDefaultJsonBuffEncoder() encodingx.IBuffEncoder {
	return NewJsonBuffEncoder(DefaultDataBlockHandler)
}

// NewDefaultJsonBuffDecoder returns a JSON buffer decoder with DefaultDataBlockHandler.
// NewDefaultJsonBuffDecoder 使用 DefaultDataBlockHandler 创建 JSON 缓冲解码器。
func NewDefaultJsonBuffDecoder() encodingx.IBuffDecoder {
	return NewJsonBuffDecoder(DefaultDataBlockHandler)
}

// NewDefaultJsonBuffCodecs returns a JSON buffer codec with DefaultDataBlockHandler.
// NewDefaultJsonBuffCodecs 使用 DefaultDataBlockHandler 创建 JSON 缓冲编解码器。
func NewDefaultJsonBuffCodecs() encodingx.IBuffCodecs {
	return NewJsonBuffCodecs(DefaultDataBlockHandler)
}
