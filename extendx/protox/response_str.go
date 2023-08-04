// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

func (resp *SockResponse) AppendString(data ...string) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		resp.BuffToBlock.WriteString(data[index])
	}
	return nil
}

func (resp *SockResponse) SendStringResponse(data ...string) error {
	resp.PrepareResponse()
	err := resp.AppendString(data...)
	if nil != err {
		return err
	}
	return resp.SendPreparedResponse()
}
