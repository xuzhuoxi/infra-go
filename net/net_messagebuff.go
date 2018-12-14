package net

import (
	"bytes"
)

func NewMessageBuff() *MessageBuff {
	buff := &bytes.Buffer{}
	rs := &MessageBuff{buff: buff}
	rs.msgSplitHandler = DefaultSplitHandler
	return rs
}

type MessageBuff struct {
	buff *bytes.Buffer

	msgSplitHandler func(buff []byte) ([]byte, []byte)
	frontLen        []byte
	frontMsg        []byte
}

func (b *MessageBuff) SetCheckMessageHandler(handler func(buff []byte) ([]byte, []byte)) {
	b.msgSplitHandler = handler
}

func (b *MessageBuff) AppendBytes(data []byte) {
	b.buff.Write(data)
}

func (b *MessageBuff) CheckMessage() bool {
	if nil == b.msgSplitHandler {
		return false
	}
	if nil != b.frontLen && nil != b.frontMsg {
		return true
	}
	len, msg := b.msgSplitHandler(b.buff.Bytes())
	if nil != len {
		b.frontLen = len
		b.frontMsg = msg
		return true
	}
	return false
}

func (b *MessageBuff) FrontMessage() []byte {
	rs := b.frontMsg
	b.buff.Next(len(b.frontLen) + len(b.frontMsg))
	b.frontLen = nil
	b.frontMsg = nil
	return rs
}

//第一个byte为长度
func DefaultSplitHandler(buff []byte) ([]byte, []byte) {
	if nil == buff {
		return nil, nil
	}
	bLen := len(buff)
	if len(buff) == 0 {
		return nil, nil
	}
	lVal := int(buff[0])
	if lVal+1 > bLen {
		return nil, nil
	}
	if lVal == 0 {
		return []byte{0}, []byte{}
	}
	rs1 := buff[:1]
	rs2 := buff[1 : lVal+1]
	return rs1, rs2
}
