package netx

import (
	"bytes"
)

//ByteSplitter----------------------

func NewByteSplitter() IByteSplitter {
	return &ByteSplitter{splitHandler: DefaultByteSplitHandler, byteBuff: &bytes.Buffer{}}
}

type ByteSplitter struct {
	splitHandler func(buff []byte) ([]byte, []byte)

	byteBuff *bytes.Buffer
	frontLen []byte
	frontMsg []byte
}

func (b *ByteSplitter) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	b.splitHandler = handler
	return nil
}

func (b *ByteSplitter) AppendBytes(data []byte) {
	b.byteBuff.Write(data)
}

func (b *ByteSplitter) CheckSplit() bool {
	//log.Println("ByteSplitter.CheckSplit：", b, b.splitHandler == nil, b.frontLen, b.frontMsg, b.byteBuff.Bytes())
	if nil == b.splitHandler {
		return false
	}
	if nil != b.frontLen && nil != b.frontMsg {
		return true
	}
	l, msg := b.splitHandler(b.byteBuff.Bytes())
	if nil != l {
		b.frontLen = l
		b.frontMsg = msg
		b.byteBuff.Next(len(l) + len(msg))
		return true
	}
	return false
}

func (b *ByteSplitter) FrontBytes() []byte {
	if nil == b.frontMsg {
		return nil
	}
	rs := b.frontMsg
	b.frontLen = nil
	b.frontMsg = nil
	return rs
}

//第一个byte为长度
func DefaultByteSplitHandler(buff []byte) ([]byte, []byte) {
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
