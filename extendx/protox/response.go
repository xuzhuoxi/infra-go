//
//Created by xuzhuoxi
//on 2019-03-22.
//@author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/netx"
)

// 响应对象参数集合接口
type IExtensionResponse interface {
	IExtensionHeader
	netx.IAddressProxySetter
	netx.ISockServerSetter
	// 设置参数类型与处理器
	SetParamInfo(paramType ExtensionParamType, paramHandler IProtocolParamsHandler)
}

type IExtensionBinaryResponse interface {
	IExtensionResponse
	// 响应客户端(二进制参数)
	SendBinaryResponse(data ...[]byte)
	// 响应指定客户端(二进制参数)
	SendBinaryResponseToClient(clientId string, data ...[]byte)
}

type IExtensionJsonResponse interface {
	IExtensionResponse
	// 响应客户端(Json字符串参数)
	SendJsonResponse(data ...string)
	// 响应指定客户端(Json字符串参数)
	SendJsonResponseToClient(clientId string, data ...string)
}

type IExtensionObjectResponse interface {
	IExtensionResponse
	// 响应客户端(具体结构体参数)
	SendObjectResponse(data ...interface{})
	// 响应指定客户端(具体结构体参数)
	SendObjectResponseToClient(clientId string, data ...interface{})
}

//-----------------

func NewSockResponse() *SockResponse {
	return &SockResponse{BuffToBlock: bytex.NewDefaultBuffToBlock()}
}

type SockResponse struct {
	ExtensionHeader

	SockServer   netx.ISockServer
	AddressProxy netx.IAddressProxy

	ParamType    ExtensionParamType
	ParamHandler IProtocolParamsHandler
	BuffToBlock  bytex.IBuffToBlock
}

func (resp *SockResponse) SetAddressProxy(proxy netx.IAddressProxy) {
	resp.AddressProxy = proxy
}
func (resp *SockResponse) SetSockServer(server netx.ISockServer) {
	resp.SockServer = server
}
func (resp *SockResponse) SetParamInfo(paramType ExtensionParamType, paramHandler IProtocolParamsHandler) {
	resp.ParamType, resp.ParamHandler = paramType, paramHandler
}

func (resp *SockResponse) SendBinaryResponse(data ...[]byte) {
	msg := resp.makePackMessage(data)
	resp.SockServer.SendPackTo(msg, resp.CAddress)
}
func (resp *SockResponse) SendBinaryResponseToClient(clientId string, data ...[]byte) {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		msg := resp.makePackMessage(data)
		resp.SockServer.SendPackTo(msg, address)
	}
}

func (resp *SockResponse) SendJsonResponse(data ...string) {
	msg := resp.makePackMessage(data)
	resp.SockServer.SendPackTo(msg, resp.CAddress)
}
func (resp *SockResponse) SendJsonResponseToClient(clientId string, data ...string) {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		msg := resp.makePackMessage(data)
		resp.SockServer.SendPackTo(msg, address)
	}
}

func (resp *SockResponse) SendObjectResponse(data ...interface{}) {
	msg := resp.makePackMessage(data)
	resp.SockServer.SendPackTo(msg, resp.CAddress)
}
func (resp *SockResponse) SendObjectResponseToClient(clientId string, data ...interface{}) {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		msg := resp.makePackMessage(data)
		resp.SockServer.SendPackTo(msg, address)
	}
}

func (resp *SockResponse) makePackMessage(data interface{}) []byte {
	resp.BuffToBlock.Reset()
	resp.BuffToBlock.WriteData([]byte(resp.EName))
	resp.BuffToBlock.WriteData([]byte(resp.PId))
	resp.BuffToBlock.WriteData([]byte(resp.CId))
	switch t := data.(type) {
	case [][]byte:
		for index, _ := range t {
			resp.BuffToBlock.WriteData(resp.ParamHandler.HandleResponseParam(t[index]))
		}
	case []string:
		for index, _ := range t {
			resp.BuffToBlock.WriteData(resp.ParamHandler.HandleResponseParam(t[index]))
		}
	case []interface{}:
		for index, _ := range t {
			resp.BuffToBlock.WriteData(resp.ParamHandler.HandleResponseParam(t[index]))
		}
	}
	return resp.BuffToBlock.ReadBytes()
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
