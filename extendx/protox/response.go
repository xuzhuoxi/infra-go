// Package protox
// Created by xuzhuoxi
// on 2019-03-22.
// @author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/netx"
)

type iExtResponse interface {
	IExtensionResponseSettings
	IExtensionResponse
}

type IExtensionResponseSettings interface {
	netx.IAddressProxySetter
	netx.ISockSenderSetter
}

// IExtensionResponse
// 响应对象参数集合接口
type IExtensionResponse interface {
	IExtensionHeader
	// SetParamInfo
	// 设置参数类型与处理器
	SetParamInfo(paramType ExtensionParamType, paramHandler IProtocolParamsHandler)
	// SetResultCode
	// 设置返回状态码
	SetResultCode(rsCode int32)
	// PrepareResponse
	// 准备设置回复参数
	PrepareResponse()
	// SendResponse
	// 根据设置好的参数响应
	SendResponse() error
	// SendResponseToClient
	// 根据设置好的参数响应到其它用户
	SendResponseToClient(interruptOnErr bool, clientIds ...string) error

	iExtRespNone
	iExtRespBase
	iExtRespBin
	iExtRespStr
	iExtResJson
	iExtRespObject
}

type iExtRespNone interface {
	// SendNoneResponse
	// 无额外参数响应
	SendNoneResponse() error
	// SendNoneResponseToClient
	// 无额外参数响到其它用户
	SendNoneResponseToClient(interruptOnErr bool, clientIds ...string) error
}

type iExtRespBin interface {
	// AppendBinary
	// 追加响应参数 - 字节数组
	AppendBinary(data ...[]byte) error
	// SendBinaryResponse
	// 响应客户端(二进制参数)
	SendBinaryResponse(data ...[]byte) error
}

type iExtRespBase interface {
	// AppendCommon
	// 追加响应参数 - 通用数据类型
	AppendCommon(data ...interface{}) error
	// SendCommonResponse
	// 响应客户端(基础数据参数)
	SendCommonResponse(data ...interface{}) error
}

type iExtRespStr interface {
	// AppendString
	// 追加响应返回- 字符串
	AppendString(data ...string) error
	// SendStringResponse
	// 响应客户端(字符串参数)
	SendStringResponse(data ...string) error
}

type iExtResJson interface {
	// AppendJson
	// 追加响应返回- Json字符串 或 可序列化的Struct
	AppendJson(data ...interface{}) error
	// SendJsonResponse
	// 响应客户端(Json字符串参数)
	SendJsonResponse(data ...interface{}) error
}

type iExtRespObject interface {
	// AppendObject
	// 追加响应参数
	AppendObject(data ...interface{}) error
	// SendObjectResponse
	// 响应客户端(具体结构体参数)
	SendObjectResponse(data ...interface{}) error
}

func NewSockResponse() *SockResponse {
	return &SockResponse{BuffToBlock: bytex.NewDefaultBuffToBlock()}
}

type SockResponse struct {
	ExtensionHeader
	RsCode int32

	SockSender   netx.ISockSender
	AddressProxy netx.IAddressProxy

	ParamType    ExtensionParamType
	ParamHandler IProtocolParamsHandler
	BuffToBlock  bytex.IBuffToBlock
}

func (resp *SockResponse) SetAddressProxy(proxy netx.IAddressProxy) {
	resp.AddressProxy = proxy
}

func (resp *SockResponse) SetSockSender(sockSender netx.ISockSender) {
	resp.SockSender = sockSender
}

func (resp *SockResponse) SetParamInfo(paramType ExtensionParamType, paramHandler IProtocolParamsHandler) {
	resp.ParamType, resp.ParamHandler = paramType, paramHandler
}

func (resp *SockResponse) SetResultCode(rsCode int32) {
	resp.RsCode = rsCode
}

func (resp *SockResponse) PrepareResponse() {
	resp.BuffToBlock.Reset()
	resp.BuffToBlock.WriteString(resp.EName)
	resp.BuffToBlock.WriteString(resp.PId)
	resp.BuffToBlock.WriteString(resp.CId)
	binaryx.Write(resp.BuffToBlock, resp.BuffToBlock.GetOrder(), resp.RsCode)
}

func (resp *SockResponse) SendResponse() error {
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}

func (resp *SockResponse) SendResponseToClient(interruptOnErr bool, clientIds ...string) error {
	if len(clientIds) == 0 {
		return nil
	}
	msg := resp.BuffToBlock.ReadBytes()
	for _, clientId := range clientIds {
		if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
			err := resp.SockSender.SendPackTo(msg, address)
			if nil != err && interruptOnErr {
				return err
			}
		}
	}
	return nil
}

func (resp *SockResponse) SendNoneResponse() error {
	resp.PrepareResponse()
	return resp.SendResponse()
}

func (resp *SockResponse) SendNoneResponseToClient(interruptOnErr bool, clientIds ...string) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.PrepareResponse()
	return resp.SendResponseToClient(interruptOnErr, clientIds...)
}
