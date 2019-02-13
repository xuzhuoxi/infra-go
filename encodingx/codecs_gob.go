package encodingx

import (
	"bytes"
	"encoding/binary"
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
	return newGobBuffCodecs(DefaultOrder, nil, bytex.DefaultDataToBlockHandler)
}

func NewDefaultGobBuffDecoder() IGobBuffDecoder {
	return newGobBuffCodecs(DefaultOrder, bytex.DefaultBlockToDataHandler, nil)
}

func NewDefaultGobBuffCodecs() IGobBuffDecoder {
	return newGobBuffCodecs(DefaultOrder, bytex.DefaultBlockToDataHandler, bytex.DefaultDataToBlockHandler)
}

func NewGobBuffEncoder(order binary.ByteOrder, data2block bytex.DataToBlockHandler) IGobBuffEncoder {
	return newGobBuffCodecs(order, nil, data2block)
}

func NewGobBuffDecoder(order binary.ByteOrder, block2data bytex.BlockToDataHandler) IGobBuffDecoder {
	return newGobBuffCodecs(order, block2data, nil)
}

func NewGobBuffCodecs(order binary.ByteOrder, block2data bytex.BlockToDataHandler, data2block bytex.DataToBlockHandler) IGobBuffDecoder {
	return newGobBuffCodecs(order, block2data, data2block)
}

//-------------------------------------

func newGobBuffCodecs(order binary.ByteOrder, block2data bytex.BlockToDataHandler, data2block bytex.DataToBlockHandler) *gobBuffCodecs {
	return &gobBuffCodecs{IBuffDataBlock: bytex.NewBuffDataBlock(order, data2block, block2data), codecs: NewGobCodecs()}
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
