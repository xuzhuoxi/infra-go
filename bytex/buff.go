// Package bytex
// Created by xuzhuoxi
// on 2019-02-11.
// @author xuzhuoxi
//
package bytex

import (
	"bytes"
	"fmt"
	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/slicex"
	"io"
	"sync"
)

type IBuffByteReader interface {
	io.Reader
	// Bytes
	// 缓冲中全部字节
	Bytes() []byte
	// ReadBytes
	// 读取缓冲中全部字节
	// 非数据安全
	ReadBytes() []byte
	// ReadBytesTo 读取缓冲中全部字节，并写入到dst中
	ReadBytesTo(dst []byte) (n int, rs []byte)
	// ReadBytesCopy
	// 读取缓冲中全部字节
	// 数据安全
	ReadBytesCopy() []byte
	// ReadBinary
	// 读取一个二进制数据到out
	// out只支持binary.Write中支持的类型
	ReadBinary(out interface{})
}

type IBuffByteWriter interface {
	io.Writer
	// WriteBytes
	// 把字节写入缓冲
	WriteBytes(bytes []byte)
	// WriteBinary
	// 把in写入数据
	WriteBinary(in interface{})
}

type IBuffDataReader interface {
	// ReadData
	// 读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	// 非数据安全
	ReadData() []byte
	// ReadDataTo
	// 读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	// 如果不是nil,把数据写入到dst中，返回dst写入的数据长度
	ReadDataTo(dst []byte) (n int, rs []byte)
	// ReadDataCopy
	// 读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	// 数据安全
	ReadDataCopy() []byte
}

type IBuffDataWriter interface {
	// WriteData
	// 把数据包装为一个Block,写入到缓冲中，数据长度为0时不进行处理
	WriteData(bytes []byte)
}

type IBuffReset interface {
	// Reset
	// 清空缓冲区
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

func (b *buffDataBlock) Bytes() []byte {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.buff.Bytes()
}

// ReadBytes
// 返回：buff缓存中的切片
// 安全：共享数据，非安全，如果要保存使用，请先进行复制
func (b *buffDataBlock) ReadBytes() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.readBytes()
}

func (b *buffDataBlock) ReadBytesTo(dst []byte) (n int, rs []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if nil == dst || len(dst) == 0 {
		return
	}
	bs := b.readBytes()
	n = copy(dst, bs)
	rs = bs[:n]
	return
}

// ReadBytesCopy
// 返回：buff缓存中的切片的克隆
// 触发Reset
func (b *buffDataBlock) ReadBytesCopy() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	cp := slicex.CopyUint8(b.readBytes())
	b.buff.Reset()
	return cp
}

func (b *buffDataBlock) readBytes() []byte {
	return b.buff.Next(b.buff.Len())
}

func (b *buffDataBlock) ReadBinary(out interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	binaryx.Read(b.buff, b.handler.GetOrder(), out)
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

func (b *buffDataBlock) WriteBinary(in interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	err := binaryx.Write(b.buff, b.handler.GetOrder(), in)
	if nil != err {
		fmt.Println("buffDataBlock.WriteBinary:", err)
	}
}

func (b *buffDataBlock) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buff.Reset()
}

func (b *buffDataBlock) Len() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.buff.Len()
}

// ReadData
// 把数据编码为[]byte
// 注意：数据安全性视handler行为而定，如果返回的是共享切片，应该马上处理
func (b *buffDataBlock) ReadData() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.readData()
}

func (b *buffDataBlock) ReadDataTo(dst []byte) (n int, rs []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if nil == dst || len(dst) == 0 {
		return
	}
	data := b.readData()
	if nil == data {
		return
	}
	n = copy(dst, data)
	rs = dst[:n]
	return
}

func (b *buffDataBlock) ReadDataCopy() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	data := b.readData()
	if nil == data {
		return nil
	}
	return slicex.CopyUint8(data)
}

func (b *buffDataBlock) readData() []byte {
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
