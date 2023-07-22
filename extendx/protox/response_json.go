// Package protox
// Create on 2023/7/22
// @author xuzhuoxi
package protox

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	jsonMarshal func(v interface{}) ([]byte, error)
)

func SetJsonMarshalHandler(handler func(v interface{}) ([]byte, error)) {
	jsonMarshal = handler
}

func (resp *SockResponse) SendJsonResponse(data ...interface{}) error {
	return resp.sendJsonResp(resp.CAddress, data...)
}

func (resp *SockResponse) SendJsonResponseToClient(clientId string, data ...interface{}) error {
	if len(clientId) == 0 {
		return nil
	}
	if address, ok := resp.AddressProxy.GetAddress(clientId); ok {
		return resp.sendJsonResp(address, data...)
	}
	return errors.New(fmt.Sprintf("No clidnetId[%s] in AddressProxy! ", clientId))
}

func (resp *SockResponse) SendJsonResponseToClients(clientIds []string, data ...interface{}) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.writeHeader()
	resp.setJsonData(data...)
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

func (resp *SockResponse) sendJsonResp(address string, data ...interface{}) error {
	resp.writeHeader()
	err := resp.setJsonData(data...)
	if nil != err {
		return err
	}
	msg := resp.BuffToBlock.ReadBytes()
	return resp.SockSender.SendPackTo(msg, address)
}

func (resp *SockResponse) setJsonData(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		json, err1 := resp.toJson(data[index])
		if nil != err1 {
			return err1
		}
		resp.setStringData(json)
	}
	return nil
}

func (resp *SockResponse) toJson(o interface{}) (json string, err error) {
	switch d := o.(type) {
	case string:
		return d, nil
	default:
		bs, err1 := jsonMarshal(d)
		if nil != err1 {
			return "", err1
		}
		return string(bs), nil
	}
}

func (resp *SockResponse) jsonMarshal(obj interface{}) ([]byte, error) {
	if nil == jsonMarshal {
		return json.Marshal(obj)
	} else {
		return jsonMarshal(obj)
	}
}
