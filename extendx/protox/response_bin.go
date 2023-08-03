// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
	"fmt"
)

func (resp *SockResponse) SendBinaryResponse(data ...[]byte) error {
	resp.writeHeader()
	resp.setBinaryData(data...)
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}

func (resp *SockResponse) SendBinaryResponseToClient(clientId string, data ...[]byte) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		resp.writeHeader()
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
	resp.writeHeader()
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

func (resp *SockResponse) setBinaryData(data ...[]byte) {
	if len(data) == 0 {
		return
	}
	for index := range data {
		resp.BuffToBlock.WriteData(data[index])
	}
}
