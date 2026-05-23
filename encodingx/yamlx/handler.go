package yamlx

import (
	"github.com/xuzhuoxi/infra-go/encodingx"
	"gopkg.in/yaml.v2"
)

// NewYamlCodingHandlerAsync returns a YAML ICodingHandler.
// NewYamlCodingHandlerAsync 返回 YAML 编解码处理器。
func NewYamlCodingHandlerAsync() encodingx.ICodingHandler {
	return yamlHandlerAsync{}
}

// NewYamlCodingHandlerSync returns a YAML ICodingHandler (alias of async implementation).
// NewYamlCodingHandlerSync 返回 YAML 编解码处理器（与 Async 实现相同）。
func NewYamlCodingHandlerSync() encodingx.ICodingHandler {
	return yamlHandlerAsync{}
}

type yamlHandlerAsync struct{}

func (c yamlHandlerAsync) HandleEncode(data interface{}) (bs []byte, err error) {
	return yaml.Marshal(data)
}

func (c yamlHandlerAsync) HandleDecode(bs []byte, data interface{}) error {
	return yaml.Unmarshal(bs, data)
}
