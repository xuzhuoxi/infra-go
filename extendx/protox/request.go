// Package protox
// Created by xuzhuoxi
// on 2019-05-18.
// @author xuzhuoxi
//
package protox

import "fmt"

// IExtensionRequest
// 请求对象参数集合接口
type IExtensionRequest interface {
	IExtensionHeader
	// DataSize 数据长度
	DataSize() int
	// SetRequestData
	// 设置集合数据信息
	SetRequestData(paramType ExtensionParamType, paramHandler IProtocolParamsHandler, data [][]byte)
}

// IExtensionBinaryRequest
// 数据参数为二进制的请求对象参数集合接口
type IExtensionBinaryRequest interface {
	IExtensionRequest
	// RequestBinaryData
	// 请求的参数数据(二进制)
	RequestBinaryData() [][]byte
}

// IExtensionStringRequest
// 数据参数为Json的请求对象参数集合接口
type IExtensionStringRequest interface {
	IExtensionRequest
	// RequestStringData
	// 请求的参数数据(String)
	RequestStringData() []string
}

// IExtensionObjectRequest
// 数据参数为结构体的请求对象参数集合接口
type IExtensionObjectRequest interface {
	IExtensionRequest
	// RequestObjectData
	// 请求的参数数据(具体数据)
	RequestObjectData() []interface{}
}

//---------------------------------------

func NewSockRequest() *SockRequest {
	return &SockRequest{}
}

type SockRequest struct {
	ExtensionHeader
	ParamType  ExtensionParamType
	BinaryData [][]byte
	StringData []string
	ObjectData []interface{}
}

func (req *SockRequest) String() string {
	return fmt.Sprintf("{Request: %v, %v, %v, %v}",
		req.ExtensionHeader, req.ParamType, req.BinaryData, req.ObjectData)
}

func (req *SockRequest) DataSize() int {
	switch req.ParamType {
	case Binary:
		return len(req.BinaryData)
	case String:
		return len(req.StringData)
	case Object:
		return len(req.ObjectData)
	}
	return 0
}

func (req *SockRequest) SetRequestData(paramType ExtensionParamType, paramHandler IProtocolParamsHandler, data [][]byte) {
	req.ParamType = paramType
	req.BinaryData = data
	switch paramType {
	case None:
		req.StringData, req.ObjectData = nil, nil
	case Binary:
		req.StringData, req.ObjectData = nil, nil
	case String:
		req.StringData, req.ObjectData = req.toStringArray(data), nil
	case Object:
		objData := paramHandler.HandleRequestParams(data)
		req.StringData, req.ObjectData = nil, objData
	}
}

func (req *SockRequest) RequestBinaryData() [][]byte {
	return req.BinaryData
}

func (req *SockRequest) RequestStringData() []string {
	return req.StringData
}

func (req *SockRequest) RequestObjectData() []interface{} {
	return req.ObjectData
}

func (req *SockRequest) toStringArray(data [][]byte) []string {
	if nil == data || len(data) == 0 {
		return nil
	}
	rs := make([]string, len(data))
	for index := range data {
		rs[index] = string(data[index])
	}
	return rs
}
