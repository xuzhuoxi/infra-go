package jsonx

import (
	"github.com/json-iterator/go"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

// NewJsonCodingHandlerAsync returns a JSON ICodingHandler (jsoniter, per-call safe).
// NewJsonCodingHandlerAsync 返回 JSON 编解码处理器（jsoniter，每次调用独立）。
func NewJsonCodingHandlerAsync() encodingx.ICodingHandler {
	return jsonHandlerAsync{}
}

// NewJsonCodingHandlerSync returns a JSON ICodingHandler (alias of async implementation).
// NewJsonCodingHandlerSync 返回 JSON 编解码处理器（与 Async 实现相同）。
func NewJsonCodingHandlerSync() encodingx.ICodingHandler {
	return jsonHandlerAsync{}
}

type jsonHandlerAsync struct{}

func (c jsonHandlerAsync) HandleEncode(data interface{}) (bs []byte, err error) {
	return jsoniter.Marshal(data)
}

func (c jsonHandlerAsync) HandleDecode(bs []byte, data interface{}) error {
	return jsoniter.Unmarshal(bs, data)
}
