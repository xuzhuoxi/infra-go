package gobx

import (
	"bytes"
	"encoding/gob"
	"sync"

	"github.com/xuzhuoxi/infra-go/encodingx"
)

// NewGobCodingHandler returns a Gob handler that reuses internal buffers.
// NewGobCodingHandler 返回复用内部缓冲区的 Gob 编解码处理器。
//
// Because buffers are reused, the same struct type only appends a TypeDescriptor on the first encode.
// 由于复用 buff，相同结构体只有第一次编码才会追加 TypeDescriptor。
func NewGobCodingHandler() encodingx.ICodingHandler {
	encoderBuff := bytes.NewBuffer(nil)
	decoderBuff := bytes.NewBuffer(nil)
	return &gobHandler{encoderBuff: encoderBuff, decoderBuff: decoderBuff,
		encoder: gob.NewEncoder(encoderBuff), decoder: gob.NewDecoder(decoderBuff)}
}

// NewGobCodingHandlerSync returns a mutex-protected Gob handler that reuses internal buffers.
// NewGobCodingHandlerSync 返回带互斥锁、复用内部缓冲区的 Gob 编解码处理器。
//
// Because buffers are reused, the same struct type only appends a TypeDescriptor on the first encode.
// 由于复用 buff，相同结构体只有第一次编码才会追加 TypeDescriptor。
func NewGobCodingHandlerSync() encodingx.ICodingHandler {
	encoderBuff := bytes.NewBuffer(nil)
	decoderBuff := bytes.NewBuffer(nil)
	return &gobHandlerSync{encoderBuff: encoderBuff, decoderBuff: decoderBuff,
		encoder: gob.NewEncoder(encoderBuff), decoder: gob.NewDecoder(decoderBuff)}
}

// NewGobCodingHandlerAsync returns a Gob handler that allocates fresh buffers per call.
// NewGobCodingHandlerAsync 返回每次调用都使用新缓冲区的 Gob 编解码处理器。
//
// Each encode appends a full TypeDescriptor to the output.
// 每次编码都会向输出追加完整的 TypeDescriptor。
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

// HandleEncode encodes data to bytes backed by the internal buffer.
// HandleEncode 将 data 编码为 []byte。
//
// The returned slice must be used immediately: it aliases the internal buffer and a subsequent
// encode overwrites the buffer, invalidating prior results.
// 注意：返回的数据应立即使用；[]byte 来源于 buff 切片，再次编码会覆盖数据并修改上次结果。
func (c *gobHandler) HandleEncode(data interface{}) (bs []byte, err error) {
	c.encoderBuff.Reset()
	err = c.encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	//return slicex.CopyUint8(c.encoderBuff.Bytes())
	bs = c.encoderBuff.Bytes()
	return
}

func (c *gobHandler) HandleDecode(bs []byte, data interface{}) error {
	c.decoderBuff.Reset()
	c.decoderBuff.Write(bs)
	return c.decoder.Decode(data)
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

// HandleEncode encodes data to bytes backed by the internal buffer (mutex-protected).
// HandleEncode 将 data 编码为 []byte（带互斥保护）。
//
// The returned slice must be used immediately: it aliases the internal buffer and a subsequent
// encode overwrites the buffer, invalidating prior results.
// 注意：返回的数据应立即使用；[]byte 来源于 buff 切片，再次编码会覆盖数据并修改上次结果。
func (c *gobHandlerSync) HandleEncode(data interface{}) (bs []byte, err error) {
	c.eMu.Lock()
	defer c.eMu.Unlock()
	c.encoderBuff.Reset()
	err = c.encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	//return slicex.CopyUint8(c.encoderBuff.Bytes())
	bs = c.encoderBuff.Bytes()
	return
}

func (c *gobHandlerSync) HandleDecode(bs []byte, data interface{}) error {
	c.dMu.Lock()
	defer c.dMu.Unlock()
	c.decoderBuff.Reset()
	_, err := c.decoderBuff.Write(bs)
	if nil != err {
		return err
	}
	return c.decoder.Decode(data)
}

//------------------------------------------

type gobHandlerAsync struct{}

func (c gobHandlerAsync) HandleEncode(data interface{}) (bs []byte, err error) {
	buff := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buff)
	err = enc.Encode(data)
	if err != nil {
		return nil, err
	}
	bs = buff.Bytes()
	return
}

func (c gobHandlerAsync) HandleDecode(bs []byte, data interface{}) error {
	buff := bytes.NewBuffer(bs)
	dec := gob.NewDecoder(buff)
	return dec.Decode(data)
}
