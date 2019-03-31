//
//Created by xuzhuoxi
//on 2019-03-22.
//@author xuzhuoxi
//
package extendx

import "github.com/xuzhuoxi/infra-go/netx"

type IExtensionResponse interface {
	SenderAddress() string
}

type IExtensionBinaryResponse interface {
	IExtensionResponse
	SendBinaryResponse(data []byte)
	SendBinaryResponseToUser(uid string, data []byte)
}

type IExtensionObjectResponse interface {
	IExtensionResponse
	SendObjectResponse(data ...interface{})
	SendObjectResponseToUser(uid string, data ...interface{})
}

//-----------------

type SockServerResponse struct {
	SockServer    netx.ISockServer
	Address       string
	AddressProxy  netx.IAddressProxy
	FuncObjToByte func(o ...interface{}) []byte
}

func (resp *SockServerResponse) SenderAddress() string {
	return resp.Address
}

func (resp *SockServerResponse) SendBinaryResponse(data []byte) {
	resp.SockServer.SendPackTo(data, resp.Address)
}

func (resp *SockServerResponse) SendBinaryResponseToUser(uid string, data []byte) {
	if address, ok := resp.AddressProxy.GetAddress(uid); ok {
		resp.SockServer.SendPackTo(data, address)
	}
}

func (resp *SockServerResponse) SendObjectResponse(data ...interface{}) {
	bs := resp.FuncObjToByte(data)
	resp.SockServer.SendPackTo(bs, resp.Address)
}

func (resp *SockServerResponse) SendObjectResponseToUser(uid string, data ...interface{}) {
	if address, ok := resp.AddressProxy.GetAddress(uid); ok {
		bs := resp.FuncObjToByte(data)
		resp.SockServer.SendPackTo(bs, address)
	}
}

//-----------------

type PackSenderResponse struct {
	PackSender   netx.IPackSender
	Address      string
	AddressProxy netx.IAddressProxy
}

func (resp *PackSenderResponse) SenderAddress() string {
	return resp.Address
}

func (resp *PackSenderResponse) SendResponse(data []byte) {
	resp.PackSender.SendPack(data, resp.Address)
}

func (resp *PackSenderResponse) SendResponseToUser(data []byte, uid string) {
	if address, ok := resp.AddressProxy.GetAddress(uid); ok {
		resp.PackSender.SendPack(data, address)
	}
}
