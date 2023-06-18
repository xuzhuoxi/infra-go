// Package protox
// Created by xuzhuoxi
// on 2019-03-22.
// @author xuzhuoxi
//
package protox

import (
	"errors"
	"fmt"
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
}

type IExtensionBinaryResponse interface {
	IExtensionResponse
	// SendBinaryResponse
	// 响应客户端(二进制参数)
	SendBinaryResponse(data ...[]byte) error
	// SendBinaryResponseToClient
	// 响应指定客户端(二进制参数)
	SendBinaryResponseToClient(clientId string, data ...[]byte) error
}

type IExtensionStringResponse interface {
	IExtensionResponse
	// SendStringResponse
	// 响应客户端(Json字符串参数)
	SendStringResponse(data ...string) error
	// SendStringResponseToClient
	// 响应指定客户端(Json字符串参数)
	SendStringResponseToClient(clientId string, data ...string) error
}

type IExtensionObjectResponse interface {
	IExtensionResponse
	// SendObjectResponse
	// 响应客户端(具体结构体参数)
	SendObjectResponse(data ...interface{}) error
	// SendObjectResponseToClient
	// 响应指定客户端(具体结构体参数)
	SendObjectResponseToClient(clientId string, data ...interface{}) error
}

//-----------------

func NewSockResponse() *SockResponse {
	return &SockResponse{BuffToBlock: bytex.NewDefaultBuffToBlock()}
}

type SockResponse struct {
	ExtensionHeader

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

func (resp *SockResponse) SendBinaryResponse(data ...[]byte) error {
	resp.setHeader()
	resp.writeBinary(data...)
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}
func (resp *SockResponse) SendBinaryResponseToClient(clientId string, data ...[]byte) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		resp.setHeader()
		resp.writeBinary(data...)
		msg := resp.BuffToBlock.ReadBytes()
		return resp.SockSender.SendPackTo(msg, address)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendStringResponse(data ...string) error {
	resp.setHeader()
	resp.writeString()
	msg, err := resp.makePackMessage(data)
	if nil != err {
		return err
	}
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}
func (resp *SockResponse) SendStringResponseToClient(clientId string, data ...string) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		msg, err := resp.makePackMessage(data)
		if nil != err {
			return err
		}
		resp.SockSender.SendPackTo(msg, address)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendObjectResponse(data ...interface{}) error {
	msg, err := resp.makePackMessage(data)
	if nil != err {
		return err
	}
	resp.SockSender.SendPackTo(msg, resp.CAddress)
}
func (resp *SockResponse) SendObjectResponseToClient(clientId string, data ...interface{}) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		msg, err := resp.makePackMessage(data)
		if nil != err {
			return err
		}
		resp.SockSender.SendPackTo(msg, address)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) makePackMessage(data interface{}) (bs []byte, err error) {
	resp.setHeader()
	switch t := data.(type) {
	case [][]byte:
		fmt.Println("666:", resp.ParamHandler)
		fmt.Println("777:", t)
		for index := range t {
			resp.BuffToBlock.WriteData(resp.ParamHandler.HandleResponseParam(t[index]))
		}
	case []string:
		for index := range t {
			resp.BuffToBlock.WriteData(resp.ParamHandler.HandleResponseParam(t[index]))
		}
	case []interface{}:
		for index := range t {
			resp.BuffToBlock.WriteData(resp.ParamHandler.HandleResponseParam(t[index]))
		}
	}
	return resp.BuffToBlock.ReadBytes(), nil
}

func (resp *SockResponse) setHeader() {
	resp.BuffToBlock.Reset()
	resp.BuffToBlock.WriteString(resp.EName)
	resp.BuffToBlock.WriteString(resp.PId)
	resp.BuffToBlock.WriteString(resp.CId)
}

func (resp *SockResponse) writeBinary(data ...[]byte) {
	if len(data) == 0 {
		return
	}
	for index := range data {
		resp.BuffToBlock.WriteData(data[index])
	}
}

func (resp *SockResponse) writeString(data ...string) error {
	if len(data) == 0 {
		return nil
	}
	if resp.ParamHandler == nil {
		return errors.New("SendStringResponse Error: ParamHandler is nil! ")
	}
	for index := range data {
		bs := resp.ParamHandler.HandleResponseParam(data[index])
		resp.BuffToBlock.WriteData(bs)
	}
	return nil
}

func (resp *SockResponse) writeObject(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	if resp.ParamHandler == nil {
		return errors.New("SendObjectResponse Error: ParamHandler is nil! ")
	}
	for index := range data {
		bs := resp.ParamHandler.HandleResponseParam(data[index])
		resp.BuffToBlock.WriteData(bs)
	}
	return nil
}

//-----------------

//type PackResponse struct {
//	PackSender   netx.IPackSender
//	AddressProxy netx.IAddressProxy
//
//	CId      string
//	CAddress string
//}
//
//func (resp *PackResponse) ClientId() string {
//	return resp.CId
//}
//
//func (resp *PackResponse) ClientAddress() string {
//	return resp.CAddress
//}
//
//func (resp *PackResponse) SendResponse(data []byte) {
//	resp.PackSender.SendPack(data, resp.CAddress)
//}
//
//func (resp *PackResponse) SendResponseToClient(clientId string, data []byte) {
//	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
//		resp.PackSender.SendPack(data, address)
//	}
//}
