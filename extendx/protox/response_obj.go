// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"errors"
)

func (resp *SockResponse) AppendObject(data ...interface{}) error {
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

func (resp *SockResponse) SendObjectResponse(data ...interface{}) error {
	resp.PrepareResponse()
	err := resp.AppendObject(data...)
	if nil != err {
		return err
	}
	return resp.SendResponse()
}
