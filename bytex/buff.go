package bytex

import (
	"bytes"
	"encoding/binary"
	"sync"

	"github.com/xuzhuoxi/infra-go/binaryx"
	"github.com/xuzhuoxi/infra-go/slicex"
)

// NewDefaultBuffDataBlock returns a full-duplex buffer with DefaultDataBlockHandler.
// NewDefaultBuffDataBlock 使用 DefaultDataBlockHandler 创建全双工缓冲。
func NewDefaultBuffDataBlock() IBuffDataBlock {
	rs := newBuffDataBlock(DefaultDataBlockHandler)
	return rs
}

// NewDefaultBuffToBlock returns a buffer that writes framed blocks and reads raw bytes.
// NewDefaultBuffToBlock 创建写入分帧块、读取原始字节的缓冲。
func NewDefaultBuffToBlock() IBuffToBlock {
	rs := newBuffDataBlock(DefaultDataBlockHandler)
	return rs
}

// NewDefaultBuffToData returns a buffer that reads framed blocks and writes raw bytes.
// NewDefaultBuffToData 创建读取分帧块、写入原始字节的缓冲。
func NewDefaultBuffToData() IBuffToData {
	rs := newBuffDataBlock(DefaultDataBlockHandler)
	return rs
}

// NewBuffDataBlock returns a full-duplex buffer with the given handler.
// NewBuffDataBlock 使用指定 handler 创建全双工缓冲。
func NewBuffDataBlock(handler IDataBlockHandler) IBuffDataBlock {
	rs := newBuffDataBlock(handler)
	return rs
}

// NewBuffToBlock returns a to-block buffer with the given handler.
// NewBuffToBlock 使用指定 handler 创建写入分帧块的缓冲。
func NewBuffToBlock(handler IDataBlockHandler) IBuffToBlock {
	rs := newBuffDataBlock(handler)
	return rs
}

// NewBuffToData returns a to-data buffer with the given handler.
// NewBuffToData 使用指定 handler 创建读取分帧块的缓冲。
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

func (b *buffDataBlock) GetOrder() binary.ByteOrder {
	return b.handler.GetOrder()
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

func (b *buffDataBlock) ReadBytesCopy() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	cp := slicex.CopyUint8(b.readBytes())
	b.buff.Reset()
	return cp
}

func (b *buffDataBlock) ReadBinary(out interface{}) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	return binaryx.Read(b.buff, b.handler.GetOrder(), out)
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

func (b *buffDataBlock) WriteBinary(in interface{}) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	return binaryx.Write(b.buff, b.handler.GetOrder(), in)
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

func (b *buffDataBlock) ReadData() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.readData()
}

func (b *buffDataBlock) ReadString() string {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.readString()
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

func (b *buffDataBlock) WriteData(bytes []byte) {
	b.lock.Lock()
	defer b.lock.Unlock()
	block := b.handler.DataToBlock(bytes)
	b.buff.Write(block)
}

func (b *buffDataBlock) WriteString(str string) {
	b.WriteData([]byte(str))
}

// Private

func (b *buffDataBlock) readBytes() []byte {
	return b.buff.Next(b.buff.Len())
}

func (b *buffDataBlock) readData() []byte {
	if b.buff.Len() == 0 {
		return nil
	}
	rs, ln, ok := b.handler.BlockToData(b.buff.Bytes())
	if !ok {
		return nil
	}
	b.buff.Next(ln)
	return rs
}

func (b *buffDataBlock) readString() string {
	return string(b.readData())
}
