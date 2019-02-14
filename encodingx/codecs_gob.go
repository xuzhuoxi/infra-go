package encodingx

import (
	"bytes"
	"encoding/gob"
	"github.com/xuzhuoxi/util-go/bytex"
	"sync"
)

type IGobBuffEncoder interface {
	bytex.IBuffByteReader
	bytex.IBuffDataWriter
	bytex.IBuffReset
	EncodeToBuff(data ...interface{})
}

type IGobBuffDecoder interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataReader
	bytex.IBuffReset
	DecodeFromBuff(data ...interface{})
}

type IGobBuffCodecs interface {
	bytex.IBuffByteWriter
	bytex.IBuffDataWriter
	bytex.IBuffByteReader
	bytex.IBuffDataReader
	bytex.IBuffReset
	EncodeToBuff(data ...interface{})
	DecodeFromBuff(data ...interface{})
}

func NewDefaultGobBuffEncoder() IGobBuffEncoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffDecoder() IGobBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewDefaultGobBuffCodecs() IGobBuffDecoder {
	return newGobBuffCodecs(DefaultDataBlockHandler)
}

func NewGobBuffEncoder(handler bytex.IDataBlockHandler) IGobBuffEncoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffDecoder(handler bytex.IDataBlockHandler) IGobBuffDecoder {
	return newGobBuffCodecs(handler)
}

func NewGobBuffCodecs(handler bytex.IDataBlockHandler) IGobBuffDecoder {
	return newGobBuffCodecs(handler)
}

//-------------------------------------

func newGobBuffCodecs(handler bytex.IDataBlockHandler) *gobBuffCodecs {
	return &gobBuffCodecs{IBuffDataBlock: bytex.NewBuffDataBlock(handler), codecs: NewGobCodecs()}
}

type gobBuffCodecs struct {
	bytex.IBuffDataBlock
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
