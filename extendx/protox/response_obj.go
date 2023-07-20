// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
	"fmt"
)

type IExtensionObjectResponse interface {
	IExtensionResponse
	// SendObjectResponse
	// 响应客户端(具体结构体参数)
	SendObjectResponse(data ...interface{}) error
	// SendObjectResponseToClient
	// 响应指定客户端(具体结构体参数)
	SendObjectResponseToClient(clientId string, data ...interface{}) error
	// SendObjectResponseToClients
	// 响应指定客户端(具体结构体参数)
	SendObjectResponseToClients(clientIds []string, data ...interface{}) error
}

func (resp *SockResponse) SendObjectResponse(data ...interface{}) error {
	resp.setHeader()
	err := resp.setObjectData(data...)
	if nil != err {
		return err
	}
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}

func (resp *SockResponse) SendObjectResponseToClient(clientId string, data ...interface{}) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		resp.setHeader()
		err := resp.setObjectData(data...)
		if nil != err {
			return err
		}
		msg := resp.BuffToBlock.ReadBytes()
		return resp.SockSender.SendPackTo(msg, address)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendObjectResponseToClients(clientIds []string, data ...interface{}) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.setHeader()
	resp.setObjectData(data...)
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
