// Package protox
// Created by xuzhuoxi
// on 2019-03-22.
// @author xuzhuoxi
//
package protox

import (
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
	IProtoReturnMessage
	// SetParamInfo
	// 设置参数类型与处理器
	SetParamInfo(paramType ExtensionParamType, paramHandler IProtocolParamsHandler)
	// SetResultCode
	// 设置返回状态码
	SetResultCode(rsCode int32)
	// SendResponse
	// 根据设置好的参数响应
	SendResponse() error
	// SendResponseTo
	// 根据设置好的参数响应到其它用户
	SendResponseTo(interruptOnErr bool, clientIds ...string) error
	// SendNotify
	// 根据设置好的参数响应
	SendNotify(eName string, notifyPId string) error
	// SendNotifyTo
	// 根据设置好的参数响应到其它用户
	SendNotifyTo(eName string, notifyPId string, interruptOnErr bool, clientIds ...string) error
	iExtResp
}

type iExtResp interface {
	// SendNoneResponse
	// 无额外参数响应
	SendNoneResponse() error
	// SendNoneResponseToClient
	// 无额外参数响到其它用户
	SendNoneResponseToClient(interruptOnErr bool, clientIds ...string) error
	// SendBinaryResponse
	// 响应客户端(二进制参数)
	SendBinaryResponse(data ...[]byte) error
	// SendCommonResponse
	// 响应客户端(基础数据参数)
	SendCommonResponse(data ...interface{}) error
	// SendStringResponse
	// 响应客户端(字符串参数)
	SendStringResponse(data ...string) error
	// SendJsonResponse
	// 响应客户端(Json字符串参数)
	SendJsonResponse(data ...interface{}) error
	// SendObjectResponse
	// 响应客户端(具体结构体参数)
	// data only supports pointer types
	// data 只支持指针类型
	SendObjectResponse(data ...interface{}) error
}

func NewSockResponse() *SockResponse {
	return &SockResponse{
		ProtoReturnMessage: *NewProtoReturnMessage(),
	}
}

type SockResponse struct {
	ProtoReturnMessage
	SockSender   netx.ISockSender
	AddressProxy netx.IAddressProxy
	ParamType    ExtensionParamType
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

func (resp *SockResponse) SendResponse() error {
	return resp.sendRedirectMsg(resp.PGroup, resp.PId)
}

func (resp *SockResponse) SendResponseTo(interruptOnErr bool, clientIds ...string) error {
	return resp.sendRedirectMsgTo(resp.PGroup, resp.PId, interruptOnErr, clientIds...)
}

func (resp *SockResponse) SendNotify(eName string, notifyPId string) error {
	return resp.sendRedirectMsg(eName, notifyPId)
}

func (resp *SockResponse) SendNotifyTo(eName string, notifyPId string, interruptOnErr bool, clientIds ...string) error {
	return resp.sendRedirectMsgTo(eName, notifyPId, interruptOnErr, clientIds...)
}

// private

func (resp *SockResponse) sendRedirectMsgTo(eName string, pId string,
	interruptOnErr bool, clientIds ...string) error {
	if len(clientIds) == 0 {
		return nil
	}
	msg, err1 := resp.genMsgBytes(eName, pId)
	if nil != err1 {
		return err1
	}
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

func (resp *SockResponse) sendRedirectMsg(eName string, pId string) error {
	msg, err := resp.genMsgBytes(eName, pId)
	if nil != err {
		return err
	}
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}
