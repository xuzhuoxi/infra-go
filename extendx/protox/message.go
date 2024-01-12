// Package protox
// Create on 2023/8/6
// @author xuzhuoxi
package protox

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/bytex"
)

type IProtoReturnMessage interface {
	IProtoHeader
	// PrepareData
	// 准备设置回复参数
	PrepareData()
	// AppendLen
	// 追加参数 - 长度值
	AppendLen(ln int) error
	// AppendBinary
	// 追加参数 - 字节数组, 自动补充长度数据
	AppendBinary(data ...[]byte) error
	// AppendCommon
	// 追加参数 - 通用数据类型
	AppendCommon(data ...interface{}) error
	// AppendString
	// 追加返回- 字符串
	AppendString(data ...string) error
	// AppendJson
	// 追加返回- Json字符串 或 可序列化的Struct
	AppendJson(data ...interface{}) error
	// AppendObject
	// 追加参数
	// data only supports pointer types
	// data 只支持指针类型
	AppendObject(data ...interface{}) error
	// GenMsgBytes
	// 生成消息
	GenMsgBytes() (msg []byte, err error)
}

func NewProtoReturnMessage() *ProtoReturnMessage {
	return &ProtoReturnMessage{
		DataBuff: bytex.NewDefaultBuffToBlock(),
		MsgBuff:  bytex.NewDefaultBuffToBlock(),
	}
}

type ProtoReturnMessage struct {
	ProtoHeader
	RsCode       int32
	ParamHandler IProtocolParamsHandler
	MsgBuff      bytex.IBuffToBlock
	DataBuff     bytex.IBuffToBlock
	dataBytes    []byte
}

func (o *ProtoReturnMessage) PrepareData() {
	o.dataBytes = nil
	o.DataBuff.Reset()
}

func (o *ProtoReturnMessage) AppendLen(ln int) error {
	order := o.DataBuff.GetOrder()
	return binaryx.Write(o.DataBuff, order, uint16(ln))
}

func (o *ProtoReturnMessage) AppendBinary(data ...[]byte) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		o.DataBuff.WriteData(data[index])
	}
	return nil
}

func (o *ProtoReturnMessage) AppendCommon(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	order := o.DataBuff.GetOrder()
	for index := range data {
		err := binaryx.Write(o.DataBuff, order, data[index])
		if nil != err {
			return err
		}
	}
	return nil
}

func (o *ProtoReturnMessage) AppendString(data ...string) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		o.DataBuff.WriteString(data[index])
	}
	return nil
}

func (o *ProtoReturnMessage) AppendJson(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	for index := range data {
		jsonStr, err1 := toJson(data[index])
		if nil != err1 {
			return err1
		}
		err2 := o.AppendString(jsonStr)
		if nil != err2 {
			return err2
		}
	}
	return nil
}

func (o *ProtoReturnMessage) AppendObject(data ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	if o.ParamHandler == nil {
		return errors.New("AppendObject Error: ParamHandler is nil! ")
	}
	for index := range data {
		bs := o.ParamHandler.HandleReturnParam(data[index])
		o.DataBuff.WriteData(bs)
	}
	return nil
}

func (o *ProtoReturnMessage) GenMsgBytes() (msg []byte, err error) {
	return o.genMsgBytes(o.PGroup, o.PId)
}

func (o *ProtoReturnMessage) genMsgBytes(eName string, pId string) (bytes []byte, err error) {
	err1 := o.writeHeaderToMsg(eName, pId)
	if nil != err1 {
		return nil, err1
	}
	err2 := o.writeDataToMsg()
	if nil != err2 {
		return nil, err2
	}
	return o.MsgBuff.ReadBytes(), nil
}

func (o *ProtoReturnMessage) writeHeaderToMsg(eName string, pId string) error {
	o.MsgBuff.Reset()
	o.MsgBuff.WriteString(eName)
	o.MsgBuff.WriteString(pId)
	o.MsgBuff.WriteString(o.CId)
	return binaryx.Write(o.MsgBuff, o.MsgBuff.GetOrder(), o.RsCode)
}

func (o *ProtoReturnMessage) writeDataToMsg() error {
	if nil == o.dataBytes {
		o.dataBytes = o.DataBuff.ReadBytesCopy()
		if nil == o.dataBytes {
			o.dataBytes = []byte{}
		}
	}
	_, err1 := o.MsgBuff.Write(o.dataBytes)
	return err1
}
