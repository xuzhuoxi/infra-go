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

type IBuffByteReader interface {
	//读取缓冲中全部字节
	ReadBytes() []byte
}

type IBuffByteWriter interface {
	//把字节写入缓冲
	WriteBytes(bytes []byte)
}

type IBuffDataReader interface {
	IBlockToDataHandler
	//读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	ReadData() []byte
}

type IBuffDataWriter interface {
	IDataToBlockHandler
	//把数据包装为一个Block,写入到缓冲中，数据长度为0时不进行处理
	WriteData(bytes []byte)
}

type IBuffReset interface {
	//清空缓冲区
	Reset()
}

//---------------------------------------------

type IBuffBase interface {
	IBuffByteReader
	IBuffByteWriter
	IBuffDataReader
	IBuffDataWriter
	IBuffReset
}

func NewBuffBase(order binary.ByteOrder) IBuffBase {
	rs := newBuffBase(order)
	return &rs
}

func newBuffBase(order binary.ByteOrder) buffBase {
	return buffBase{order: order, buff: bytes.NewBuffer(nil), handlerBlock2Byte: DefaultBlockToDataHandler, handlerByte2Block: DefaultDataToBlockHandler}
}

type buffBase struct {
	handlerBlock2Byte BlockToDataHandler
	handlerByte2Block DataToBlockHandler

	order binary.ByteOrder
	buff  *bytes.Buffer
	lock  sync.RWMutex
}

func (b *buffBase) SetBlockToDataHandler(handler BlockToDataHandler) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.handlerBlock2Byte = handler
}

func (b *buffBase) SetDataToBlockHandler(handler DataToBlockHandler) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.handlerByte2Block = handler
}

func (b *buffBase) ReadBytes() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	l := b.buff.Len()
	return b.buff.Next(l)
}

func (b *buffBase) WriteBytes(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buff.Write(bytes)
}

func (b *buffBase) ReadData() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	rs, l, ok := b.handlerBlock2Byte(b.buff.Bytes(), b.order)
	if !ok {
		return nil
	}
	b.buff.Next(l)
	return rs
}

func (b *buffBase) WriteData(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	block, ok := b.handlerByte2Block(bytes, b.order)
	if !ok {
		return
	}
	b.buff.Write(block)
}

func (b *buffBase) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buff.Reset()
}
