// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

import (
	"github.com/xuzhuoxi/infra-go/binaryx"
)

func (resp *SockResponse) AppendCommon(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	order := resp.BuffToBlock.GetOrder()
	for index := range data {
		binaryx.Write(resp.BuffToBlock, order, data[index])
	}
	return nil
}

func (resp *SockResponse) SendCommonResponse(data ...interface{}) error {
	resp.PrepareResponse()
	err := resp.AppendCommon(data...)
	if nil != err {
		return err
	}
	return resp.SendPreparedResponse()
}
