// Package protox
// Created by xuzhuoxi
// on 2019-05-19.
// @author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/encodingx"
)

// ParamFactory 通用参数构造器
type ParamFactory = func() interface{}

// IProtocolParamsHandler
// 协议参数处理器接口
// 要求：并发安全
type IProtocolParamsHandler interface {
	// SetCodingHandler
	// 设置编解码器
	SetCodingHandler(handler encodingx.ICodingHandler)
	// HandleRequestParam
	// 处理请求参数转换：二进制->具体数据
	HandleRequestParam(data []byte) interface{}
	// HandleRequestParams
	// 处理请求参数转换：二进制->具体数据(批量)
	HandleRequestParams(data [][]byte) []interface{}
	// HandleResponseParam
	// 处理响应参数转换：具体数据->二进制
	HandleResponseParam(data interface{}) []byte
	// HandleResponseParams
	// 处理响应参数转换：具体数据->二进制(批量)
	HandleResponseParams(data []interface{}) [][]byte
}

//----------------------------

func NewProtoObjectParamsHandler(factory ParamFactory, handler encodingx.ICodingHandler) IProtocolParamsHandler {
	return &ProtoObjectParamsHandler{ParamFactory: factory, Handler: handler}
}

type ProtoObjectParamsHandler struct {
	ParamFactory ParamFactory
	Handler      encodingx.ICodingHandler
}

func (o *ProtoObjectParamsHandler) SetCodingHandler(handler encodingx.ICodingHandler) {
	o.Handler = handler
}

func (o *ProtoObjectParamsHandler) HandleRequestParam(data []byte) interface{} {
	rs := o.ParamFactory()
	err := o.Handler.HandleDecode(data, rs)
	if nil != err {
		return nil
	}
	return rs
}

func (o *ProtoObjectParamsHandler) HandleRequestParams(data [][]byte) []interface{} {
	var objData []interface{}
	for index := range data {
		objData = append(objData, o.HandleRequestParam(data[index]))
	}
	return objData
}

func (o *ProtoObjectParamsHandler) HandleResponseParam(data interface{}) []byte {
	return o.Handler.HandleEncode(data)
}

func (o *ProtoObjectParamsHandler) HandleResponseParams(data []interface{}) [][]byte {
	var byteData [][]byte
	for index := range data {
		byteData = append(byteData, o.HandleResponseParam(data[index]))
	}
	return byteData
}
