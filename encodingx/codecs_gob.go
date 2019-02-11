package encodingx

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"sync"
)

type IGobBuffEncoder interface {
	IBuffByteReader
	IBuffDataWriter
	IBuffReset
	EncodeToBuff(data ...interface{})
}

type IGobBuffDecoder interface {
	IBuffByteWriter
	IBuffDataReader
	IBuffReset
	DecodeFromBuff(data ...interface{})
}

type IGobBuffCodecs interface {
	IBuffByteWriter
	IBuffDataWriter
	IBuffByteReader
	IBuffDataReader
	IBuffReset
	EncodeToBuff(data ...interface{})
	DecodeFromBuff(data ...interface{})
}

func NewGobBuffEncoder(order binary.ByteOrder) IGobBuffEncoder {
	return newGobBuffCodecs(order)
}

func NewGobBuffDecoder(order binary.ByteOrder) IGobBuffDecoder {
	return newGobBuffCodecs(order)
}

func NewGobBuffCodecs(order binary.ByteOrder) IGobBuffDecoder {
	return newGobBuffCodecs(order)
}

func newGobBuffCodecs(order binary.ByteOrder) *gobBuffCodecs {
	return &gobBuffCodecs{buffBase: newBuffBase(order), codecs: NewGobCodecs()}
}

type gobBuffCodecs struct {
	buffBase
	codecs     *gobCodecs
	codecsLock sync.RWMutex
}

func (b *gobBuffCodecs) EncodeToBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	b.codecsLock.Lock()
	defer b.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.codecs.Encode(data[index])
		b.WriteData(bytes)
	}
}

func (b *gobBuffCodecs) DecodeFromBuff(data ...interface{}) {
	if len(data) == 0 {
		return
	}
	b.codecsLock.Lock()
	defer b.codecsLock.Unlock()
	for index := 0; index < len(data); index++ {
		bytes := b.ReadData()
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

func NewGobCodecs() *gobCodecs {
	return &gobCodecs{}
}

func GobEncode(e interface{}) []byte {
	return DefaultGobCodecs.Encode(e)
}

func GobDecoder(data []byte, e interface{}) {
	DefaultGobCodecs.Decode(data, e)
}
