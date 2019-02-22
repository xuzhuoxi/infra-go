//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package bytex

import (
	"bytes"
	"github.com/xuzhuoxi/infra-go/slicex"
	"io"
	"sync"
)

type IBuffByteReader interface {
	io.Reader
	//读取缓冲中全部字节
	//非数据安全
	ReadBytes() []byte
	//读取缓冲中全部字节
	//数据安全
	ReadCopyBytes() []byte
}

type IBuffByteWriter interface {
	io.Writer
	//把字节写入缓冲
	WriteBytes(bytes []byte)
}

type IBuffDataReader interface {
	//读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	//非数据安全
	ReadData() []byte
	//读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	//数据安全
	ReadCopyData() []byte
}

type IBuffDataWriter interface {
	//把数据包装为一个Block,写入到缓冲中，数据长度为0时不进行处理
	WriteData(bytes []byte)
}

type IBuffReset interface {
	//清空缓冲区
	Reset()
}

type IBuffLen interface {
	Len() int
}

//---------------------------------------------

type iBuffToBlock interface {
	IBuffDataWriter
	IBuffByteReader
}

type iBuffToData interface {
	IBuffDataReader
	IBuffByteWriter
}

type IBuffToBlock interface {
	iBuffToBlock
	IBuffReset
	IBuffLen
}

type IBuffToData interface {
	iBuffToData
	IBuffReset
	IBuffLen
}

type IBuffDataBlock interface {
	iBuffToBlock
	iBuffToData
	IBuffReset
	IBuffLen
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

//返回：buff缓存中的切片
//安全：共享数据，非安全，如果要保存使用，请先进行复制
func (b *buffDataBlock) ReadBytes() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	l := b.buff.Len()
	return b.buff.Next(l)
}

func (b *buffDataBlock) ReadCopyBytes() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	cp := slicex.CopyUint8(b.buff.Bytes())
	b.buff.Reset()
	return cp
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

func (b *buffDataBlock) Len() int {
	return b.buff.Len()
}

//把数据编码为[]byte
//注意：数据安全性视handler行为而定，如果返回的是共享切片，应该马上处理
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

func (b *buffDataBlock) ReadCopyData() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	rs, l, ok := b.handler.BlockToData(b.buff.Bytes())
	if !ok {
		return nil
	}
	b.buff.Next(l)
	return slicex.CopyUint8(rs)
}

func (b *buffDataBlock) WriteData(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	block := b.handler.DataToBlock(bytes)
	b.buff.Write(block)
}
