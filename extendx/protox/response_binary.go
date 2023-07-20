// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
	"fmt"
)

type IExtensionBinaryResponse interface {
	IExtensionResponse
	// SendBinaryResponse
	// 响应客户端(二进制参数)
	SendBinaryResponse(data ...[]byte) error
	// SendBinaryResponseToClient
	// 响应指定客户端(二进制参数)
	SendBinaryResponseToClient(clientId string, data ...[]byte) error
	// SendBinaryResponseToClients
	// 响应指定客户端(二进制参数)
	SendBinaryResponseToClients(clientIds []string, data ...[]byte) error
}

func (resp *SockResponse) SendBinaryResponse(data ...[]byte) error {
	resp.setHeader()
	resp.setBinaryData(data...)
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}

func (resp *SockResponse) SendBinaryResponseToClient(clientId string, data ...[]byte) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		resp.setHeader()
		resp.setBinaryData(data...)
		msg := resp.BuffToBlock.ReadBytes()
		return resp.SockSender.SendPackTo(msg, address)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendBinaryResponseToClients(clientIds []string, data ...[]byte) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.setHeader()
	resp.setBinaryData(data...)
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
