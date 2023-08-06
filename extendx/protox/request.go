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
	IProtoHeader
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
	// BinaryData
	// RequestBinaryData
	// 请求的参数数据(二进制)
	BinaryData() [][]byte
}

// IExtensionStringRequest
// 数据参数为Json的请求对象参数集合接口
type IExtensionStringRequest interface {
	IExtensionRequest
	// StringData
	// 请求的参数数据(String)
	StringData() []string
}

// IExtensionObjectRequest
// 数据参数为结构体的请求对象参数集合接口
type IExtensionObjectRequest interface {
	IExtensionRequest
	// ObjectData
	// 请求的参数数据(具体数据)
	ObjectData() []interface{}
}

//---------------------------------------

func NewSockRequest() *SockRequest {
	return &SockRequest{}
}

type SockRequest struct {
	ProtoHeader
	ParamType ExtensionParamType
	binData   [][]byte
	strData   []string
	objData   []interface{}
}

func (req *SockRequest) String() string {
	return fmt.Sprintf("{Request: %v, %v, %v, %v}",
		req.ProtoHeader, req.ParamType, req.binData, req.objData)
}

func (req *SockRequest) DataSize() int {
	switch req.ParamType {
	case Binary:
		return len(req.binData)
	case String:
		return len(req.strData)
	case Object:
		return len(req.objData)
	}
	return 0
}

func (req *SockRequest) SetRequestData(paramType ExtensionParamType, paramHandler IProtocolParamsHandler, data [][]byte) {
	req.ParamType = paramType
	req.binData = data
	switch paramType {
	case None:
		req.strData, req.objData = nil, nil
	case Binary:
		req.strData, req.objData = nil, nil
	case String:
		req.strData, req.objData = req.toStringArray(data), nil
	case Object:
		objData := paramHandler.HandleRequestParams(data)
		req.strData, req.objData = nil, objData
	}
}

func (req *SockRequest) BinaryData() [][]byte {
	return req.binData
}

func (req *SockRequest) StringData() []string {
	return req.strData
}

func (req *SockRequest) ObjectData() []interface{} {
	return req.objData
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
