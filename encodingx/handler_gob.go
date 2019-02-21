//
//Created by xuzhuoxi
//on 2019-02-16.
//@author xuzhuoxi
//
package encodingx

import (
	"bytes"
	"encoding/gob"
	"sync"
)

func NewGobCodingHandler() ICodingHandler {
	buff := bytes.NewBuffer(nil)
	return &gobHandler{buff: buff, encoder: gob.NewEncoder(buff), decoder: gob.NewDecoder(buff)}
}

func NewGobCodingHandlerSync() ICodingHandler {
	buff := bytes.NewBuffer(nil)
	return &gobHandlerSync{buff: buff, encoder: gob.NewEncoder(buff), decoder: gob.NewDecoder(buff)}
}

func NewGobCodingHandlerAsync() ICodingHandler {
	return gobHandlerAsync{}
}

//------------------------------------------

type gobHandler struct {
	buff    *bytes.Buffer
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func (c *gobHandler) HandleEncode(data interface{}) []byte {
	c.buff.Reset()
	c.encoder.Encode(data)
	return c.buff.Bytes()
}

func (c *gobHandler) HandleDecode(bs []byte, data interface{}) {
	c.buff.Reset()
	c.buff.Write(bs)
	c.decoder.Decode(data)
}

//------------------------------------------

type gobHandlerSync struct {
	buff    *bytes.Buffer
	encoder *gob.Encoder
	decoder *gob.Decoder
	mu      sync.RWMutex
}

func (c *gobHandlerSync) HandleEncode(data interface{}) []byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.buff.Reset()
	c.encoder.Encode(data)
	return c.buff.Bytes()
}

func (c *gobHandlerSync) HandleDecode(bs []byte, data interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.buff.Reset()
	c.buff.Write(bs)
	c.decoder.Decode(data)
}

//------------------------------------------

type gobHandlerAsync struct {
}

func (c gobHandlerAsync) HandleEncode(data interface{}) []byte {
	buff := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buff)
	enc.Encode(data)
	return buff.Bytes()
}

func (c gobHandlerAsync) HandleDecode(bs []byte, data interface{}) {
	buff := bytes.NewBuffer(bs)
	dec := gob.NewDecoder(buff)
	dec.Decode(data)
}
