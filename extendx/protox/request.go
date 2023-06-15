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

// IExtensionJsonRequest
// 数据参数为Json的请求对象参数集合接口
type IExtensionJsonRequest interface {
	IExtensionRequest
	// RequestJsonData
	// 请求的参数数据(Json)
	RequestJsonData() []string
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
	JsonData   []string
	ObjectData []interface{}
}

func (req *SockRequest) SetRequestData(paramType ExtensionParamType, paramHandler IProtocolParamsHandler, data [][]byte) {
	switch paramType {
	case None:
		req.BinaryData, req.JsonData, req.ObjectData = nil, nil, nil
	case Binary:
		req.BinaryData, req.JsonData, req.ObjectData = data, nil, nil
	case Json:
		objData := paramHandler.HandleRequestParams(data)
		req.BinaryData, req.JsonData, req.ObjectData = nil, req.toJsonArray(objData), nil
	case Object:
		objData := paramHandler.HandleRequestParams(data)
		req.BinaryData, req.JsonData, req.ObjectData = nil, nil, objData
	}
}

func (req *SockRequest) RequestBinaryData() [][]byte {
	return req.BinaryData
}

func (req *SockRequest) RequestJsonData() []string {
	return req.JsonData
}

func (req *SockRequest) RequestObjectData() []interface{} {
	return req.ObjectData
}

func (req *SockRequest) toJsonArray(objArr []interface{}) []string {
	if nil == objArr || len(objArr) == 0 {
		return nil
	}
	rs := make([]string, len(objArr))
	for index := range objArr {
		rs[index] = objArr[index].(string)
	}
	return rs
}
