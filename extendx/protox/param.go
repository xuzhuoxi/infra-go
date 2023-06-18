// Package protox
// Created by xuzhuoxi
// on 2019-05-19.
// @author xuzhuoxi
//
package protox

import "github.com/xuzhuoxi/infra-go/encodingx"

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

func NewProtoStringParamsHandler() IProtocolParamsHandler {
	return &ProtoStringParamsHandler{}
}

type ProtoStringParamsHandler struct{}

func (h *ProtoStringParamsHandler) SetCodingHandler(handler encodingx.ICodingHandler) {
	return
}

func (h *ProtoStringParamsHandler) HandleRequestParam(data []byte) interface{} {
	return string(data)
}
func (h *ProtoStringParamsHandler) HandleRequestParams(data [][]byte) []interface{} {
	var objData []interface{}
	for index := range data {
		objData = append(objData, h.HandleRequestParam(data[index]))
	}
	return objData
}

func (h *ProtoStringParamsHandler) HandleResponseParam(data interface{}) []byte {
	return []byte(data.(string))
}
func (h *ProtoStringParamsHandler) HandleResponseParams(data []interface{}) [][]byte {
	var byteData [][]byte
	for index := range data {
		byteData = append(byteData, h.HandleResponseParam(data[index]))
	}
	return byteData
}
