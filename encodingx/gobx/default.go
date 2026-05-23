package gobx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

var (
	// DefaultDataBlockHandler is the default bytex data block handler for Gob codecs.
	// DefaultDataBlockHandler 是 Gob 编解码使用的默认 bytex 数据块处理器。
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

	// DefaultPoolGobCodingHandler is the default pool for Gob ICodingHandler instances.
	// DefaultPoolGobCodingHandler 是 Gob 编解码处理器的默认对象池。
	DefaultPoolGobCodingHandler = encodingx.NewPoolCodingHandler()
	// DefaultPoolGobBuffEncoder is the default pool for Gob IBuffEncoder instances.
	// DefaultPoolGobBuffEncoder 是 Gob 缓冲编码器的默认对象池。
	DefaultPoolGobBuffEncoder = encodingx.NewPoolBuffEncoder()
	// DefaultPoolGobBuffDecoder is the default pool for Gob IBuffDecoder instances.
	// DefaultPoolGobBuffDecoder 是 Gob 缓冲解码器的默认对象池。
	DefaultPoolGobBuffDecoder = encodingx.NewPoolBuffDecoder()
	// DefaultPoolGobBuffCodecs is the default pool for Gob IBuffCodecs instances.
	// DefaultPoolGobBuffCodecs 是 Gob 缓冲编解码器的默认对象池。
	DefaultPoolGobBuffCodecs = encodingx.NewPoolBuffCodecs()
)

func init() {
	DefaultPoolGobCodingHandler.Register(NewDefaultGobCodingHandler)
	DefaultPoolGobBuffEncoder.Register(NewDefaultGobBuffEncoder)
	DefaultPoolGobBuffDecoder.Register(NewDefaultGobBuffDecoder)
	DefaultPoolGobBuffCodecs.Register(NewDefaultGobBuffCodecs)
}

// NewDefaultGobCodingHandler returns the default Gob ICodingHandler (async, full type descriptors).
// NewDefaultGobCodingHandler 返回默认 Gob 编解码处理器（Async，每次附带完整类型描述）。
func NewDefaultGobCodingHandler() encodingx.ICodingHandler {
	return NewGobCodingHandlerAsync()
	//return NewGobCodingHandler()
}

// NewDefaultGobBuffEncoder returns a Gob buffer encoder with DefaultDataBlockHandler.
// NewDefaultGobBuffEncoder 使用 DefaultDataBlockHandler 创建 Gob 缓冲编码器。
func NewDefaultGobBuffEncoder() encodingx.IBuffEncoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

// NewDefaultGobBuffDecoder returns a Gob buffer decoder with DefaultDataBlockHandler.
// NewDefaultGobBuffDecoder 使用 DefaultDataBlockHandler 创建 Gob 缓冲解码器。
func NewDefaultGobBuffDecoder() encodingx.IBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

// NewDefaultGobBuffCodecs returns a Gob buffer codec with DefaultDataBlockHandler.
// NewDefaultGobBuffCodecs 使用 DefaultDataBlockHandler 创建 Gob 缓冲编解码器。
func NewDefaultGobBuffCodecs() encodingx.IBuffCodecs {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}
