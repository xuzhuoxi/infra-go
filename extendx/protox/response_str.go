// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
	"fmt"
)

type IExtensionStringResponse interface {
	IExtensionResponse
	// SendStringResponse
	// 响应客户端(Json字符串参数)
	SendStringResponse(data ...string) error
	// SendStringResponseToClient
	// 响应指定客户端(Json字符串参数)
	SendStringResponseToClient(clientId string, data ...string) error
	// SendStringResponseToClients
	// 响应指定客户端(Json字符串参数)
	SendStringResponseToClients(clientIds []string, data ...string) error
}

func (resp *SockResponse) SendStringResponse(data ...string) error {
	resp.setHeader()
	err := resp.setStringData(data...)
	if nil != err {
		return err
	}
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}

func (resp *SockResponse) SendStringResponseToClient(clientId string, data ...string) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		resp.setHeader()
		err := resp.setStringData(data...)
		if nil != err {
			return err
		}
		msg := resp.BuffToBlock.ReadBytes()
		return resp.SockSender.SendPackTo(msg, address)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendStringResponseToClients(clientIds []string, data ...string) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.setHeader()
	resp.setStringData(data...)
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
