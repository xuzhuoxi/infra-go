//
//Created by xuzhuoxi
//on 2019-03-22.
//@author xuzhuoxi
//
package extendx

import "github.com/xuzhuoxi/infra-go/netx"

type IExtensionResponse interface {
	SenderAddress() string
	SendResponse(data []byte)
	SendResponseToUser(data []byte, uid string)
}

//-----------------

type SockServerResponse struct {
	SockServer   netx.ISockServer
	Address      string
	AddressProxy netx.IAddressProxy
}

func (resp *SockServerResponse) SenderAddress() string {
	return resp.Address
}

func (resp *SockServerResponse) SendResponse(data []byte) {
	resp.SockServer.SendPackTo(data, resp.Address)
}

func (resp *SockServerResponse) SendResponseToUser(data []byte, uid string) {
	if address, ok := resp.AddressProxy.GetAddress(uid); ok {
		resp.SockServer.SendPackTo(data, address)
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
