// Package protox
// Created by xuzhuoxi
// on 2019-02-26.
// @author xuzhuoxi
//
package protox

// ExtensionProtoInfo
// 协议定义
type ExtensionProtoInfo struct {
	ProtoId          string
	ParamType        ExtensionParamType
	ExtensionHandler interface{}

	ParamHandler IProtocolParamsHandler
	ParamOrigin  interface{}
}

//---------------------------------------

func NewProtocolExtensionSupport(Name string) ProtocolExtensionSupport {
	return ProtocolExtensionSupport{
		Name: Name, ProtoIdToInfo: make(map[string]*ExtensionProtoInfo),
	}
}

type ProtocolExtensionSupport struct {
	Name          string
	ProtoIdToInfo map[string]*ExtensionProtoInfo
}

func (s *ProtocolExtensionSupport) ExtensionName() string {
	return s.Name
}
func (s *ProtocolExtensionSupport) CheckProtocolId(protoId string) bool {
	_, ok := s.ProtoIdToInfo[protoId]
	return ok
}
func (s *ProtocolExtensionSupport) GetParamInfo(protoId string) (paramType ExtensionParamType, handler IProtocolParamsHandler) {
	info, _ := s.ProtoIdToInfo[protoId]
	return info.ParamType, info.ParamHandler
}

func (s *ProtocolExtensionSupport) SetRequestHandler(protoId string, handler ExtensionHandlerNoneParam) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: None, ExtensionHandler: handler}
}
func (s *ProtocolExtensionSupport) SetRequestHandlerBinary(protoId string, handler ExtensionHandlerBinaryParam) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: Binary, ExtensionHandler: handler}
}
func (s *ProtocolExtensionSupport) SetRequestHandlerString(protoId string, handler ExtensionHandlerStringParam) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: String, ExtensionHandler: handler, ParamHandler: NewProtoStringParamsHandler()}
}
func (s *ProtocolExtensionSupport) SetRequestHandlerObject(protoId string, handler ExtensionHandlerObjectParam, ObjectOrigin interface{}, paramHandler IProtocolParamsHandler) {
	s.ProtoIdToInfo[protoId] = &ExtensionProtoInfo{ProtoId: protoId, ParamType: Object, ExtensionHandler: handler, ParamOrigin: ObjectOrigin, ParamHandler: paramHandler}
}
func (s *ProtocolExtensionSupport) ClearRequestHandler(protoId string) {
	delete(s.ProtoIdToInfo, protoId)
}

func (s *ProtocolExtensionSupport) OnRequest(resp IExtensionResponse, req IExtensionRequest) {
	info, _ := s.ProtoIdToInfo[req.ProtoId()]
	switch info.ParamType {
	case None:
		handler := info.ExtensionHandler.(ExtensionHandlerNoneParam)
		handler(resp.(IExtensionResponse), req.(IExtensionRequest))
	case Binary:
		handler := info.ExtensionHandler.(ExtensionHandlerBinaryParam)
		handler(resp.(IExtensionBinaryResponse), req.(IExtensionBinaryRequest))
	case String:
		handler := info.ExtensionHandler.(ExtensionHandlerStringParam)
		handler(resp.(IExtensionStringResponse), req.(IExtensionStringRequest))
	case Object:
		handler := info.ExtensionHandler.(ExtensionHandlerObjectParam)
		handler(resp.(IExtensionObjectResponse), req.(IExtensionObjectRequest))
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
