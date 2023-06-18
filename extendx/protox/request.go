// Package protox
// Created by xuzhuoxi
// on 2019-05-18.
// @author xuzhuoxi
//
package protox

// IExtensionRequest
// 请求对象参数集合接口
type IExtensionRequest interface {
	IExtensionHeader
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

	BinaryData [][]byte
	StringData []string
	ObjectData []interface{}
}

func (req *SockRequest) SetRequestData(paramType ExtensionParamType, paramHandler IProtocolParamsHandler, data [][]byte) {
	switch paramType {
	case None:
		req.BinaryData, req.StringData, req.ObjectData = nil, nil, nil
	case Binary:
		req.BinaryData, req.StringData, req.ObjectData = data, nil, nil
	case String:
		objData := paramHandler.HandleRequestParams(data)
		req.BinaryData, req.StringData, req.ObjectData = nil, req.toStringArray(objData), nil
	case Object:
		objData := paramHandler.HandleRequestParams(data)
		req.BinaryData, req.StringData, req.ObjectData = nil, nil, objData
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

func (req *SockRequest) toStringArray(objArr []interface{}) []string {
	if nil == objArr || len(objArr) == 0 {
		return nil
	}
	rs := make([]string, len(objArr))
	for index := range objArr {
		rs[index] = objArr[index].(string)
	}
	return rs
}
