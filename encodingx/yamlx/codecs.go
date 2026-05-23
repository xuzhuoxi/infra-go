package yamlx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

// NewYamlBuffEncoder returns a YAML buffer encoder.
// NewYamlBuffEncoder 创建 YAML 缓冲编码器。
func NewYamlBuffEncoder(handler bytex.IDataBlockHandler) encodingx.IBuffEncoder {
	return encodingx.NewBuffEncoder(handler, NewYamlCodingHandlerAsync())
}

// NewYamlBuffDecoder returns a YAML buffer decoder.
// NewYamlBuffDecoder 创建 YAML 缓冲解码器。
func NewYamlBuffDecoder(handler bytex.IDataBlockHandler) encodingx.IBuffDecoder {
	return encodingx.NewBuffDecoder(handler, NewYamlCodingHandlerAsync())
}

// NewYamlBuffCodecs returns a combined YAML buffer codec.
// NewYamlBuffCodecs 创建 YAML 缓冲编解码器。
func NewYamlBuffCodecs(handler bytex.IDataBlockHandler) encodingx.IBuffCodecs {
	return encodingx.NewBuffCodecs(handler, NewYamlCodingHandlerAsync())
}
