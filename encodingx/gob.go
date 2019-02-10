package encodingx

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"sync"
)

var GobOrder = binary.BigEndian

type IGobBuffEncoder interface {
	IBuffEncoded
	EncodeToBuff(data ...interface{})
}

type IGobBuffDecoder interface {
	IBuffDecoded
	DecodeFromBuff(data ...interface{})
}

func NewGobBuffEncoder() IGobBuffEncoder {
	return newGobBuff()
}

func NewGobBuffDecoder() IGobBuffDecoder {
	return newGobBuff()
}

func newGobBuff() *gobBuff {
	return &gobBuff{base: NewBuffBase(GobOrder), codecs: NewCodecs()}
}

type gobBuff struct {
	base     IBuffBase
	codecs   *gobCodecs
	buffLock sync.RWMutex
}

func (b *gobBuff) DecodedBytes(bytes []byte) {
	b.base.AppendBytes(bytes)
}

func (b *gobBuff) EncodedBytes() []byte {
	return b.base.ClearBytes()
}

func (b *gobBuff) EncodeToBuff(data ...interface{}) {
	b.buffLock.Lock()
	defer b.buffLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.codecs.Encode(data[index])
		b.base.WriteData(bytes)
	}
}

func (b *gobBuff) DecodeFromBuff(data ...interface{}) {
	b.buffLock.Lock()
	defer b.buffLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.base.ReadData()
		b.codecs.Decode(bytes, data[index])
	}
}

//----------------------------------------------------

type gobCodecs struct {
	buff bytes.Buffer
}

func (c *gobCodecs) Encode(e interface{}) []byte {
	c.buff.Reset()
	enc := gob.NewEncoder(&c.buff)
	enc.Encode(e)
	return c.buff.Bytes()
}

func (c *gobCodecs) Decode(data []byte, rs interface{}) {
	c.buff.Reset()
	c.buff.Write(data)
	dec := gob.NewDecoder(&c.buff)
	dec.Decode(rs)
}

//----------------------------------------------------

var DefaultGobCodecs gobCodecs

func NewCodecs() *gobCodecs {
	return &gobCodecs{}
}

func Encode(e interface{}) []byte {
	return DefaultGobCodecs.Encode(e)
}

func Decoder(data []byte, e interface{}) {
	DefaultGobCodecs.Decode(data, e)
}
