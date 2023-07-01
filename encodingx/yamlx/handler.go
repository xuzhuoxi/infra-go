// Package yamlx
// Created by xuzhuoxi
// on 2019-02-24.
// @author xuzhuoxi
//
package yamlx

import (
	"github.com/xuzhuoxi/infra-go/encodingx"
	"gopkg.in/yaml.v2"
)

func NewYamlCodingHandlerAsync() encodingx.ICodingHandler {
	return yamlHandlerAsync{}
}

func NewYamlCodingHandlerSync() encodingx.ICodingHandler {
	return yamlHandlerAsync{}
}

type yamlHandlerAsync struct{}

func (c yamlHandlerAsync) HandleEncode(data interface{}) []byte {
	bs, err := yaml.Marshal(data)
	if nil != err {
		return nil
	}
	return bs
}

func (c yamlHandlerAsync) HandleDecode(bs []byte, data interface{}) error {
	return yaml.Unmarshal(bs, data)
}
