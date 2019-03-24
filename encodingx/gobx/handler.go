//
//Created by xuzhuoxi
//on 2019-02-16.
//@author xuzhuoxi
//
package gobx

import (
	"bytes"
	"encoding/gob"
	"github.com/xuzhuoxi/infra-go/encodingx"
	"sync"
)

// 由于复用buff，相同结构体只有第一次才会追加TypeDescriptor
func NewGobCodingHandler() encodingx.ICodingHandler {
	encoderBuff := bytes.NewBuffer(nil)
	decoderBuff := bytes.NewBuffer(nil)
	return &gobHandler{encoderBuff: encoderBuff, decoderBuff: decoderBuff,
		encoder: gob.NewEncoder(encoderBuff), decoder: gob.NewDecoder(decoderBuff)}
}

// 由于复用buff，相同结构体只有第一次才会追加TypeDescriptor
func NewGobCodingHandlerSync() encodingx.ICodingHandler {
	encoderBuff := bytes.NewBuffer(nil)
	decoderBuff := bytes.NewBuffer(nil)
	return &gobHandlerSync{encoderBuff: encoderBuff, decoderBuff: decoderBuff,
		encoder: gob.NewEncoder(encoderBuff), decoder: gob.NewDecoder(decoderBuff)}
}

//每次都会追加TypeDescriptor到编码中去
func NewGobCodingHandlerAsync() encodingx.ICodingHandler {
	return gobHandlerAsync{}
}

//------------------------------------------

type gobHandler struct {
	encoderBuff *bytes.Buffer
	decoderBuff *bytes.Buffer
	encoder     *gob.Encoder
	decoder     *gob.Decoder
}

//把数据编码为[]byte
//注意：返回的数据应该马上使用
//因为：[]byte来源于buff的切片，再次执行会覆盖数据，导致上次的返回数据被修改
func (c *gobHandler) HandleEncode(data interface{}) []byte {
	c.encoderBuff.Reset()
	c.encoder.Encode(data)
	//return slicex.CopyUint8(c.encoderBuff.Bytes())
	return c.encoderBuff.Bytes()
}

func (c *gobHandler) HandleDecode(bs []byte, data interface{}) {
	c.decoderBuff.Reset()
	c.decoderBuff.Write(bs)
	c.decoder.Decode(data)
}

//------------------------------------------

type gobHandlerSync struct {
	encoderBuff *bytes.Buffer
	decoderBuff *bytes.Buffer
	encoder     *gob.Encoder
	decoder     *gob.Decoder
	eMu         sync.RWMutex
	dMu         sync.RWMutex
}

//把数据编码为[]byte
//注意：返回的数据应该马上使用
//因为：[]byte来源于buff的切片，再次执行会覆盖数据，导致上次的返回数据被修改
func (c *gobHandlerSync) HandleEncode(data interface{}) []byte {
	c.eMu.Lock()
	defer c.eMu.Unlock()
	c.encoderBuff.Reset()
	c.encoder.Encode(data)
	//return slicex.CopyUint8(c.encoderBuff.Bytes())
	return c.encoderBuff.Bytes()
}

func (c *gobHandlerSync) HandleDecode(bs []byte, data interface{}) {
	c.dMu.Lock()
	defer c.dMu.Unlock()
	c.decoderBuff.Reset()
	c.decoderBuff.Write(bs)
	c.decoder.Decode(data)
}

//------------------------------------------

type gobHandlerAsync struct{}

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
