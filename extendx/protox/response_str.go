// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
	"fmt"
)

func (resp *SockResponse) SendStringResponse(data ...string) error {
	return resp.sendStringResp(resp.CAddress, data...)
}

func (resp *SockResponse) SendStringResponseToClient(clientId string, data ...string) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		return resp.sendStringResp(address, data...)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendStringResponseToClients(clientIds []string, data ...string) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.writeHeader()
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

func (resp *SockResponse) sendStringResp(address string, data ...string) error {
	resp.writeHeader()
	err := resp.setStringData(data...)
	if nil != err {
		return err
	}
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, address)
}

func (resp *SockResponse) setStringData(data ...string) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		resp.BuffToBlock.WriteString(data[index])
	}
	return nil
}
