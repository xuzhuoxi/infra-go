// Package protocolx
// Created by xuzhuoxi
// on 2019-02-12.
// @author xuzhuoxi
//
package protocolx

//
//func NewDefaultProtocolEncoder(encodeHandler encodingx.IDataEncodeHandler) IEncodingProtocol {
//	return newProtocolCoder(DefaultOrder, DefaultData2BlockHandler, nil, encodeHandler, nil)
//}
//
//func NewDefaultProtocolDecoder(decodeHandler encodingx.IDataDecodeHandler) IDecodingProtocol {
//	return newProtocolCoder(DefaultOrder, nil, DefaultBlock2DataHandler, nil, decodeHandler)
//}
//
//func NewDefaultProtocolCoder(encodeHandler encodingx.IDataEncodeHandler, decodeHandler encodingx.IDataDecodeHandler) ICodingProtocol {
//	return newProtocolCoder(DefaultOrder, DefaultData2BlockHandler, DefaultBlock2DataHandler, encodeHandler, decodeHandler)
//}
//
//func NewProtocolEncoder(order binaryx.ByteOrder, data2blockHandler bytex.DataToBlockHandler, encodeHandler encodingx.IDataEncodeHandler) IEncodingProtocol {
//	return newProtocolCoder(order, data2blockHandler, nil, encodeHandler, nil)
//}
//
//func NewProtocolDecoder(order binaryx.ByteOrder, block2dataHandler bytex.BlockToDataHandler, decodeHandler encodingx.IDataDecodeHandler) IDecodingProtocol {
//	return newProtocolCoder(order, nil, block2dataHandler, nil, decodeHandler)
//}
//
//func NewProtocolCoder(order binaryx.ByteOrder, data2blockHandler bytex.DataToBlockHandler, block2dataHandler bytex.BlockToDataHandler,
//	encodeHandler encodingx.IDataEncodeHandler, decodeHandler encodingx.IDataDecodeHandler) ICodingProtocol {
//	return newProtocolCoder(order, data2blockHandler, block2dataHandler, encodeHandler, decodeHandler)
//}
//
//func newProtocolCoder(order binaryx.ByteOrder, data2blockHandler bytex.DataToBlockHandler, block2dataHandler bytex.BlockToDataHandler,
//	encodeHandler encodingx.IDataEncodeHandler, decodeHandler encodingx.IDataDecodeHandler) *protocolData {
//	return &protocolData{buff: bytes.NewBuffer(nil), order: order,
//		data2blockHandler: data2blockHandler, block2dataHandler: block2dataHandler,
//		encodeHandler: encodeHandler, decodeHandler: decodeHandler}
//}
//
//type ICodingProtocol interface {
//	encodingx.ICodingData
//	IProtocolData
//}
//
//type protocolData struct {
//	protocolId   string
//	protocolData []interface{}
//	buff         *bytes.Buffer
//
//	order             binaryx.ByteOrder
//	data2blockHandler bytex.DataToBlockHandler
//	block2dataHandler bytex.BlockToDataHandler
//	encodeHandler     encodingx.IDataEncodeHandler
//	decodeHandler     encodingx.IDataDecodeHandler
//}
//
//func (d *protocolData) ProtocolId() string {
//	return d.protocolId
//}
//
//func (d *protocolData) ProtocolData() []interface{} {
//	return d.protocolData
//}
//
//func (d *protocolData) EncodeToBytes() []byte {
//	d.buff.Reset()
//	binaryx.Write(d.buff, d.order, &d.protocolId)
//	if len(d.protocolData) > 0 {
//		for index := 0; index < len(d.protocolData); index++ {
//			var data []byte
//			if bytes, ok := d.protocolData[index].([]byte); ok { //数据为字节数据
//				data = bytes
//			} else if eData, ok := d.protocolData[index].(encodingx.IEncodingData); ok { //数据自带序列化方法
//				data = eData.EncodeToBytes()
//			} else if nil != d.encodeHandler { //序列化方法
//				data = d.encodeHandler.HandleEncode(d.protocolData[index])
//			} else {
//				panic("Encode Error!")
//			}
//			block := d.data2blockHandler(data, d.order)
//			d.buff.Write(block)
//		}
//	}
//	return slicex.CopyByte(d.buff.Bytes())
//}
//
//func (d *protocolData) DecodeFromBytes(data []byte) bool {
//	d.buff.Reset()
//	d.buff.Write(data)
//	err := binaryx.Read(d.buff, d.order, &d.protocolId)
//	if nil != err {
//		return false
//	}
//	var rs []interface{}
//	for d.buff.Len() > 0 {
//		dataByte, l, ok := d.block2dataHandler(d.buff.Bytes(), d.order)
//		if !ok {
//			return false
//		}
//		d.buff.Next(l)
//		rs = append(rs, d.decodeHandler.HandleDecode(dataByte))
//	}
//	d.protocolData = rs
//	return true
//}
