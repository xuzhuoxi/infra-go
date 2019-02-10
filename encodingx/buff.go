//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package encodingx

import (
	"bytes"
	"encoding/binary"
	"sync"
)

type IBuffEncoded interface {
	EncodedBytes() []byte
}

type IBuffDecoded interface {
	DecodedBytes(bytes []byte)
}

//---------------------------------------------

type IBuffBase interface {
	AppendBytes(bytes []byte)
	ClearBytes() []byte
	Reset()

	WriteData(bytes []byte)
	ReadData() []byte
}

func NewBuffBase(order binary.ByteOrder) IBuffBase {
	rs := newBuffBase(order)
	return &rs
}

func newBuffBase(order binary.ByteOrder) buffBase {
	return buffBase{order: order, buff: bytes.NewBuffer(nil)}
}

type buffBase struct {
	order binary.ByteOrder
	buff  *bytes.Buffer
	lock  sync.RWMutex
}

func (b *buffBase) AppendBytes(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buff.Write(bytes)
}

func (b *buffBase) ClearBytes() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	l := b.buff.Len()
	return b.buff.Next(l)
}

func (b *buffBase) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buff.Reset()
}

func (b *buffBase) WriteData(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	buffWriteBytes(b.buff, b.order, bytes)
}

func (b *buffBase) ReadData() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	return buffReadBytes(b.buff, b.order)
}

func buffWriteBytes(buff *bytes.Buffer, order binary.ByteOrder, bytes []byte) {
	l := uint16(len(bytes))
	if 0 == l {
		return
	}
	binary.Write(buff, order, l)
	buff.Write(bytes)
}

func buffReadBytes(buff *bytes.Buffer, order binary.ByteOrder) []byte {
	var l uint16
	binary.Read(buff, order, &l)
	return buff.Next(int(l))
}
