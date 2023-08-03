// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/binaryx"
)

func (resp *SockResponse) SendCommonResponse(data ...interface{}) error {
	return resp.sendBaseResp(resp.CAddress, data...)
}

func (resp *SockResponse) SendCommonResponseToClient(clientId string, data ...interface{}) error {
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		return resp.sendBaseResp(address, data...)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendCommonResponseToClients(clientIds []string, data ...interface{}) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.writeHeader()
	resp.setCommonData(data...)
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

func (resp *SockResponse) sendBaseResp(address string, data ...interface{}) error {
	resp.writeHeader()
	err := resp.setCommonData(data...)
	if nil != err {
		return err
	}
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, address)
}

func (resp *SockResponse) setCommonData(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	order := resp.BuffToBlock.GetOrder()
	for index := range data {
		binaryx.Write(resp.BuffToBlock, order, data[index])
	}
	return nil
}
