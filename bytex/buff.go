//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package bytex

import (
	"bytes"
	"io"
	"sync"
)

type IBuffByteReader interface {
	io.Reader
	//读取缓冲中全部字节
	ReadBytes() []byte
}

type IBuffByteWriter interface {
	io.Writer
	//把字节写入缓冲
	WriteBytes(bytes []byte)
}

type IBuffDataReader interface {
	//读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	ReadData() []byte
}

type IBuffDataWriter interface {
	//把数据包装为一个Block,写入到缓冲中，数据长度为0时不进行处理
	WriteData(bytes []byte)
}

type IBuffReset interface {
	//清空缓冲区
	Reset()
}

//---------------------------------------------

type IBuffToBlock interface {
	IBuffDataWriter
	IBuffByteReader
	IBuffReset
}

type IBuffToData interface {
	IBuffDataReader
	IBuffByteWriter
	IBuffReset
}

type IBuffDataBlock interface {
	IBuffByteReader
	IBuffByteWriter
	IBuffDataReader
	IBuffDataWriter
	IBuffReset
}

func NewDefaultBuffDataBlock() IBuffDataBlock {
	rs := newBuffDataBlock(DefaultDataBlockHandler)
	return rs
}

func NewDefaultBuffToBlock() IBuffToBlock {
	rs := newBuffDataBlock(DefaultDataBlockHandler)
	return rs
}

func NewDefaultBuffToData() IBuffToData {
	rs := newBuffDataBlock(DefaultDataBlockHandler)
	return rs
}

func NewBuffDataBlock(handler IDataBlockHandler) IBuffDataBlock {
	rs := newBuffDataBlock(handler)
	return rs
}

func NewBuffToBlock(handler IDataBlockHandler) IBuffToBlock {
	rs := newBuffDataBlock(handler)
	return rs
}

func NewBuffToData(handler IDataBlockHandler) IBuffToData {
	rs := newBuffDataBlock(handler)
	return rs
}

//----------------------------------------

func newBuffDataBlock(handler IDataBlockHandler) *buffDataBlock {
	return &buffDataBlock{buff: bytes.NewBuffer(nil), handler: handler}
}

type buffDataBlock struct {
	buff    *bytes.Buffer
	handler IDataBlockHandler
	lock    sync.RWMutex
}

func (b *buffDataBlock) Read(p []byte) (n int, err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.buff.Read(p)
}

func (b *buffDataBlock) ReadBytes() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	l := b.buff.Len()
	return b.buff.Next(l)
}

func (b *buffDataBlock) Write(p []byte) (n int, err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.buff.Write(p)
}

func (b *buffDataBlock) WriteBytes(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buff.Write(bytes)
}

func (b *buffDataBlock) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buff.Reset()
}

func (b *buffDataBlock) ReadData() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	rs, l, ok := b.handler.BlockToData(b.buff.Bytes())
	if !ok {
		return nil
	}
	b.buff.Next(l)
	return rs
}

func (b *buffDataBlock) WriteData(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	block := b.handler.DataToBlock(bytes)
	b.buff.Write(block)
}