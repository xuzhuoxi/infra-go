package encoding

import (
	"bytes"
	"encoding/gob"
)

type GobCodecs struct {
	buff bytes.Buffer
}

func (c *GobCodecs) Encode(e interface{}) []byte {
	c.buff.Reset()
	enc := gob.NewEncoder(&c.buff)
	enc.Encode(e)
	return c.buff.Bytes()
}

func (c *GobCodecs) Decoder(data []byte, rs interface{}) {
	c.buff.Reset()
	c.buff.Write(data)
	dec := gob.NewDecoder(&c.buff)
	dec.Decode(rs)
}

//----------------------------------------------------

var DefaultCodecs GobCodecs

func NewCodecs() *GobCodecs {
	return &GobCodecs{}
}

func Encode(e interface{}) []byte {
	return DefaultCodecs.Encode(e)
}

func Decoder(data []byte, e interface{}) {
	DefaultCodecs.Decoder(data, e)
}
