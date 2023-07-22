// Package protox
// Created by xuzhuoxi
// on 2019-03-22.
// @author xuzhuoxi
//
package protox

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/netx"
)

// IExtensionResponse
// 响应对象参数集合接口
type IExtensionResponse interface {
	IExtensionHeader
	netx.IAddressProxySetter
	netx.ISockSenderSetter
	// SetParamInfo
	// 设置参数类型与处理器
	SetParamInfo(paramType ExtensionParamType, paramHandler IProtocolParamsHandler)
	// SetResultCode
	// 设置返回状态码
	SetResultCode(rsCode int32)
	// SendNoneResponse
	// 无额外参数响应
	SendNoneResponse() error
	// SendNoneResponseToClient
	// 无额外参数响到其它用户
	SendNoneResponseToClient(clientId string) error
	// SendNoneResponseToClients
	// 无额外参数响到其它用户
	SendNoneResponseToClients(clientIds []string) error
	// SendJsonResponse
	// 响应客户端(Json字符串参数)
	SendJsonResponse(data ...interface{}) error
	// SendJsonResponseToClient
	// 响应指定客户端(Json字符串参数)
	SendJsonResponseToClient(clientId string, data ...interface{}) error
	// SendJsonResponseToClients
	// 响应指定客户端(Json字符串参数)
	SendJsonResponseToClients(clientIds []string, data ...interface{}) error
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

func (resp *SockResponse) SendNoneResponse() error {
	resp.writeHeader()
	msg := resp.BuffToBlock.ReadBytes()
	fmt.Println("Response:", msg)
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}

func (resp *SockResponse) SendNoneResponseToClient(clientId string) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		resp.writeHeader()
		msg := resp.BuffToBlock.ReadBytes()
		return resp.SockSender.SendPackTo(msg, address)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendNoneResponseToClients(clientIds []string) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.writeHeader()
	msg := resp.BuffToBlock.ReadBytes()
	for _, clientId := range clientIds {
		if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
			err := resp.SockSender.SendPackTo(msg, address)
			if nil != err {
				return err
			}
		}
	}
	return nil
}

func (resp *SockResponse) writeHeader() {
	resp.BuffToBlock.Reset()
	resp.BuffToBlock.WriteString(resp.EName)
	resp.BuffToBlock.WriteString(resp.PId)
	resp.BuffToBlock.WriteString(resp.CId)
	binaryx.Write(resp.BuffToBlock, resp.BuffToBlock.GetOrder(), resp.RsCode)
}
