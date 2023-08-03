// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

func (resp *SockResponse) AppendBinary(data ...[]byte) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		resp.BuffToBlock.WriteData(data[index])
	}
	return nil
}

func (resp *SockResponse) SendBinaryResponse(data ...[]byte) error {
	resp.PrepareResponse()
	resp.AppendBinary(data...)
	return resp.SendResponse()
}
