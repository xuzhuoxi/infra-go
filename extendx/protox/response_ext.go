// Package protox
// Create on 2023/7/20
// @author xuzhuoxi
package protox

func (resp *SockResponse) SendNoneResponse() error {
	resp.PrepareData()
	return resp.SendResponse()
}

func (resp *SockResponse) SendNoneResponseToClient(interruptOnErr bool, clientIds ...string) error {
	if len(clientIds) == 0 {
		return nil
	}
	resp.PrepareData()
	return resp.SendResponseTo(interruptOnErr, clientIds...)
}

func (resp *SockResponse) SendBinaryResponse(data ...[]byte) error {
	resp.PrepareData()
	resp.AppendBinary(data...)
	return resp.SendResponse()
}

func (resp *SockResponse) SendCommonResponse(data ...interface{}) error {
	resp.PrepareData()
	err := resp.AppendCommon(data...)
	if nil != err {
		return err
	}
	return resp.SendResponse()
}

func (resp *SockResponse) SendStringResponse(data ...string) error {
	resp.PrepareData()
	err := resp.AppendString(data...)
	if nil != err {
		return err
	}
	return resp.SendResponse()
}

func (resp *SockResponse) SendJsonResponse(data ...interface{}) error {
	resp.PrepareData()
	err := resp.AppendJson(data...)
	if nil != err {
		return err
	}
	return resp.SendResponse()
}

func (resp *SockResponse) SendObjectResponse(data ...interface{}) error {
	resp.PrepareData()
	err := resp.AppendObject(data...)
	if nil != err {
		return err
	}
	return resp.SendResponse()
}
