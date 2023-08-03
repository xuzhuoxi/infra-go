// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
	"fmt"
)

func (resp *SockResponse) SendObjectResponse(data ...interface{}) error {
	resp.writeHeader()
	err := resp.setObjectData(data...)
	if nil != err {
		return err
	}
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, resp.CAddress)
}

func (resp *SockResponse) SendObjectResponseToClient(clientId string, data ...interface{}) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		resp.writeHeader()
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
	resp.writeHeader()
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

func (resp *SockResponse) setObjectData(data ...interface{}) error {
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
