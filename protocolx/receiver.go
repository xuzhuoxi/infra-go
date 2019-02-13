//
//Created by xuzhuoxi
//on 2019-02-12.
//@author xuzhuoxi
//
package protocolx

import (
	"github.com/xuzhuoxi/util-go/bytex"
	"io"
)

type IProtocolReceiver interface {
	//接收数据
	io.Writer
	ReceiveBytes(bytes []byte)
}

//-------------------------------------------------

type protocolReceiver struct {
	handlerTable IExtensionTable
	buffToData   bytex.IBuffToData
	packData     IPackData
}

func (r *protocolReceiver) Write(p []byte) (n int, err error) {
	r.ReceiveBytes(p)
	return len(p), nil
}

func (r *protocolReceiver) ReceiveBytes(bytes []byte) {
	r.buffToData.WriteBytes(bytes)
	bytesData := r.buffToData.ReadData()
	if nil == bytes {
		return
	}
	succ := r.packData.DecodeFromBytes(bytesData)
	if !succ {
		return
	}
	pId := r.packData.ProtocolId()
	pData := r.packData.ProtocolData()
	handler := r.handlerTable.GetProtocolHandler(pId)
	if nil == handler {
		return
	}
	r.handleData(handler, pId, pData)
}

func (r *protocolReceiver) handleData(handler IProtocolExtension, pId string, data []interface{}) {
	l := len(data)
	if 0 == l {
		handler.OnRequest(pId, nil)
	} else if 1 == l {
		handler.OnRequest(pId, data[0])
	} else {
		handler.OnRequest(pId, data[0], data[1:]...)
	}
}
