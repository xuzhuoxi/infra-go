// Package yamlx
// Created by xuzhuoxi
// on 2019-02-24.
// @author xuzhuoxi
//
package yamlx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

func NewYamlBuffEncoder(handler bytex.IDataBlockHandler) encodingx.IBuffEncoder {
	return encodingx.NewBuffEncoder(handler, NewYamlCodingHandlerAsync())
}

func NewYamlBuffDecoder(handler bytex.IDataBlockHandler) encodingx.IBuffDecoder {
	return encodingx.NewBuffDecoder(handler, NewYamlCodingHandlerAsync())
}

func NewYamlBuffCodecs(handler bytex.IDataBlockHandler) encodingx.IBuffCodecs {
	return encodingx.NewBuffCodecs(handler, NewYamlCodingHandlerAsync())
}
