// Package packx
// Created by xuzhuoxi
// on 2019-02-12.
// @author xuzhuoxi
//
package packx

//
//type IProtocolReceiver interface {
//	//接收数据
//	io.Writer
//	Receive(bytes []byte)
//}
//
////-------------------------------------------------
//
//func NewProtocolReceiver(handlerContainer IProtocolContainer, packData IPackData) IProtocolReceiver {
//	return &protocolReceiver{container: handlerContainer, packData: packData}
//}
//
//type protocolReceiver struct {
//	container IProtocolContainer
//	packData  IPackData
//}
//
//func (r *protocolReceiver) Write(p []byte) (n int, err error) {
//	r.Receive(p)
//	return len(p), nil
//}
//
//func (r *protocolReceiver) Receive(protoData []byte) {
//	suc := r.packData.DecodeFromBytes(protoData)
//	if !suc {
//		return
//	}
//	pId := r.packData.ProtocolId()
//	pData := r.packData.ProtocolData()
//	handler := r.container.GetExtension(pId).(IProtocolExtension)
//	if nil == handler {
//		return
//	}
//	r.handleData(handler, pId, pData)
//}
//
//func (r *protocolReceiver) handleData(handler IProtocolExtension, pId string, data []interface{}) {
//	l := len(data)
//	if 0 == l {
//		handler.HandleRequest(pId, nil)
//	} else if 1 == l {
//		handler.HandleRequest(pId, data[0])
//	} else {
//		handler.HandleRequest(pId, data[0], data[1:]...)
//	}
//}
//
////-------------------------------------------------
//
//type byteReceiver struct {
//	container  IProtocolContainer
//	buffToData bytex.IBuffToData
//	packData   IPackData
//}
//
//func (r *byteReceiver) Write(p []byte) (n int, err error) {
//	r.ReceiveBytes(p)
//	return len(p), nil
//}
//
//func (r *byteReceiver) ReceiveBytes(bytes []byte) {
//	r.buffToData.WriteBytes(bytes)
//	bytesData := r.buffToData.ReadData()
//	if nil == bytes {
//		return
//	}
//	succ := r.packData.DecodeFromBytes(bytesData)
//	if !succ {
//		return
//	}
//	pId := r.packData.ProtocolId()
//	pData := r.packData.ProtocolData()
//	handler := r.container.GetExtension(pId).(IProtocolExtension)
//	if nil == handler {
//		return
//	}
//	r.handleData(handler, pId, pData)
//}
//
//func (r *byteReceiver) handleData(handler IProtocolExtension, pId string, data []interface{}) {
//	l := len(data)
//	if 0 == l {
//		handler.HandleRequest(pId, nil)
//	} else if 1 == l {
//		handler.HandleRequest(pId, data[0])
//	} else {
//		handler.HandleRequest(pId, data[0], data[1:]...)
//	}
//}
