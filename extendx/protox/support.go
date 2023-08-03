// Package protox
// Created by xuzhuoxi
// on 2019-02-26.
// @author xuzhuoxi
//
package protox

import "github.com/xuzhuoxi/infra-go/encodingx"

// ExtensionProtoInfo
// 协议定义
type ExtensionProtoInfo struct {
	ProtoId          string
	ParamType        ExtensionParamType
	ExtensionHandler interface{}

	ParamHandler    IProtocolParamsHandler
	ReqParamFactory ParamFactory
}

//---------------------------------------

func NewProtoExtensionSupport(Name string) ProtoExtensionSupport {
	return ProtoExtensionSupport{
		Name: Name, ProtoIdToInfo: make(map[string]*ExtensionProtoInfo),
	}
}

type ProtoExtensionSupport struct {
	Name          string
	ProtoIdToInfo map[string]*ExtensionProtoInfo
}

func (s *ProtoExtensionSupport) ExtensionName() string {
	return s.Name
}
func (s *ProtoExtensionSupport) CheckProtocolId(protoId string) bool {
	_, ok := s.ProtoIdToInfo[protoId]
	return ok
}
func (s *ProtoExtensionSupport) GetParamInfo(protoId string) (paramType ExtensionParamType, handler IProtocolParamsHandler) {
	info, _ := s.ProtoIdToInfo[protoId]
	return info.ParamType, info.ParamHandler
}

func (s *ProtoExtensionSupport) SetRequestHandler(protoId string, handler ExtensionHandlerNoneParam) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: None, ExtensionHandler: handler}
}
func (s *ProtoExtensionSupport) SetRequestHandlerBinary(protoId string, handler ExtensionHandlerBinaryParam) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: Binary, ExtensionHandler: handler}
}
func (s *ProtoExtensionSupport) SetRequestHandlerString(protoId string, handler ExtensionHandlerStringParam) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: String, ExtensionHandler: handler}
}
func (s *ProtoExtensionSupport) SetRequestHandlerObject(protoId string, handler ExtensionHandlerObjectParam,
	factory ParamFactory, codingHandler encodingx.ICodingHandler) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: Object, ExtensionHandler: handler,
		ReqParamFactory: factory, ParamHandler: NewProtoObjectParamsHandler(factory, codingHandler)}
}
func (s *ProtoExtensionSupport) ClearRequestHandler(protoId string) {
	delete(s.ProtoIdToInfo, protoId)
}

func (s *ProtoExtensionSupport) OnRequest(resp IExtensionResponse, req IExtensionRequest) {
	info, _ := s.ProtoIdToInfo[req.ProtoId()]
	switch info.ParamType {
	case None:
		handler := info.ExtensionHandler.(ExtensionHandlerNoneParam)
		handler(resp.(IExtensionResponse), req.(IExtensionRequest))
	case Binary:
		handler := info.ExtensionHandler.(ExtensionHandlerBinaryParam)
		handler(resp.(IExtensionResponse), req.(IExtensionBinaryRequest))
	case String:
		handler := info.ExtensionHandler.(ExtensionHandlerStringParam)
		handler(resp.(IExtensionResponse), req.(IExtensionStringRequest))
	case Object:
		handler := info.ExtensionHandler.(ExtensionHandlerObjectParam)
		handler(resp.(IExtensionResponse), req.(IExtensionObjectRequest))
	}
}

//---------------------------------------

func NewGoroutineExtensionSupport(MaxGo int) GoroutineExtensionSupport {
	return GoroutineExtensionSupport{MaxGoroutine: MaxGo}
}

type GoroutineExtensionSupport struct {
	MaxGoroutine int
}

func (s *GoroutineExtensionSupport) MaxGo() int {
	return s.MaxGoroutine
}
