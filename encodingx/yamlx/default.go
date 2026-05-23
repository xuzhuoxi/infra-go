package yamlx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

var (
	// DefaultDataBlockHandler is the default bytex data block handler for YAML codecs.
	// DefaultDataBlockHandler 是 YAML 编解码使用的默认 bytex 数据块处理器。
	DefaultDataBlockHandler = bytex.NewDefaultDataBlockHandler()

	// DefaultPoolYamlCodingHandler is the default pool for YAML ICodingHandler instances.
	// DefaultPoolYamlCodingHandler 是 YAML 编解码处理器的默认对象池。
	DefaultPoolYamlCodingHandler = encodingx.NewPoolCodingHandler()
	// DefaultPoolYamlBuffEncoder is the default pool for YAML IBuffEncoder instances.
	// DefaultPoolYamlBuffEncoder 是 YAML 缓冲编码器的默认对象池。
	DefaultPoolYamlBuffEncoder = encodingx.NewPoolBuffEncoder()
	// DefaultPoolYamlBuffDecoder is the default pool for YAML IBuffDecoder instances.
	// DefaultPoolYamlBuffDecoder 是 YAML 缓冲解码器的默认对象池。
	DefaultPoolYamlBuffDecoder = encodingx.NewPoolBuffDecoder()
	// DefaultPoolYamlBuffCodecs is the default pool for YAML IBuffCodecs instances.
	// DefaultPoolYamlBuffCodecs 是 YAML 缓冲编解码器的默认对象池。
	DefaultPoolYamlBuffCodecs = encodingx.NewPoolBuffCodecs()
)

func init() {
	DefaultPoolYamlCodingHandler.Register(NewDefaultYamlCodingHandler)
	DefaultPoolYamlBuffEncoder.Register(NewDefaultYamlBuffEncoder)
	DefaultPoolYamlBuffDecoder.Register(NewDefaultYamlBuffDecoder)
	DefaultPoolYamlBuffCodecs.Register(NewDefaultYamlBuffCodecs)
}

// NewDefaultYamlCodingHandler returns the default YAML ICodingHandler.
// NewDefaultYamlCodingHandler 返回默认 YAML 编解码处理器。
func NewDefaultYamlCodingHandler() encodingx.ICodingHandler {
	return NewYamlCodingHandlerAsync()
}

// NewDefaultYamlBuffEncoder returns a YAML buffer encoder with DefaultDataBlockHandler.
// NewDefaultYamlBuffEncoder 使用 DefaultDataBlockHandler 创建 YAML 缓冲编码器。
func NewDefaultYamlBuffEncoder() encodingx.IBuffEncoder {
	return NewYamlBuffEncoder(DefaultDataBlockHandler)
}

// NewDefaultYamlBuffDecoder returns a YAML buffer decoder with DefaultDataBlockHandler.
// NewDefaultYamlBuffDecoder 使用 DefaultDataBlockHandler 创建 YAML 缓冲解码器。
func NewDefaultYamlBuffDecoder() encodingx.IBuffDecoder {
	return NewYamlBuffDecoder(DefaultDataBlockHandler)
}

// NewDefaultYamlBuffCodecs returns a YAML buffer codec with DefaultDataBlockHandler.
// NewDefaultYamlBuffCodecs 使用 DefaultDataBlockHandler 创建 YAML 缓冲编解码器。
func NewDefaultYamlBuffCodecs() encodingx.IBuffCodecs {
	return NewYamlBuffCodecs(DefaultDataBlockHandler)
}
