package bytex

import (
	"encoding/binary"
	"errors"
)

// BytesHandlerType identifies byte array processing kind.
// BytesHandlerType 字节数组处理类型。
type BytesHandlerType int

// BytesHandler transforms an input byte slice into an output byte slice.
// BytesHandler 字节数组处理函数。
type BytesHandler func(in []byte) (out []byte, err error)

// Predefined BytesHandlerType values for block/unblock and base64 operations.
// 预定义的 BytesHandlerType 值，用于封包/解包及 base64 操作。
const (
	LittleBlock BytesHandlerType = iota + 1
	LittleUnblock
	BigBlock
	BigUnblock

	Base64StdEncode
	Base64StdDecode
	Base64RawStdEncode
	Base64RawStdDecode

	Base64UrlEncode
	Base64UrlDecode
	Base64RawUrlEncode
	Base64RawUrlDecode
)

var (
	// ErrTypeUnregister is returned when a handler type is not registered.
	// ErrTypeUnregister 在 handler 类型未注册时返回。
	ErrTypeUnregister = errors.New("Type Unregister! ")

	// ErrLittleUnblock is returned when little-endian unblock fails.
	// ErrLittleUnblock 在小端解包失败时返回。
	ErrLittleUnblock = errors.New("LittleUnblock Error! ")
	// ErrBigUnblock is returned when big-endian unblock fails.
	// ErrBigUnblock 在大端解包失败时返回。
	ErrBigUnblock = errors.New("BigUnblock Error! ")
)

var t2handler = make([]BytesHandler, 128, 128)

// IBytesProcessor runs a chain of byte handlers in sequence.
// IBytesProcessor 字节数组处理器。
type IBytesProcessor interface {
	// AppendHandler appends a custom handler to the pipeline.
	// AppendHandler 追加处理函数。
	AppendHandler(process BytesHandler)
	// AppendHandlerType appends a registered predefined handler type.
	// AppendHandlerType 追加处理类型。
	AppendHandlerType(t BytesHandlerType) error
	// ClearHandlers removes all handlers from the pipeline.
	// ClearHandlers 清除全部处理函数。
	ClearHandlers()

	// Handle runs all handlers on in and returns the result, failing index, and error.
	// Handle 依次处理数据。
	Handle(in []byte) (out []byte, errIdx int, err error)
}

// BytesProcessor is the default IBytesProcessor implementation.
// BytesProcessor 是 IBytesProcessor 的默认实现。
type BytesProcessor struct {
	handlers []BytesHandler
}

func (p *BytesProcessor) AppendHandler(process BytesHandler) {
	if nil != process {
		p.handlers = append(p.handlers, process)
	}
}

func (p *BytesProcessor) AppendHandlerType(t BytesHandlerType) error {
	if t2handler[t] != nil {
		p.handlers = append(p.handlers, t2handler[t])
		return nil
	}
	return nil
}

func (p *BytesProcessor) ClearHandlers() {
	p.handlers = nil
}

func (p *BytesProcessor) Handle(in []byte) (out []byte, errIdx int, err error) {
	if len(p.handlers) == 0 {
		return
	}
	tmp := out
	for idx := range p.handlers {
		tmp, err = p.handlers[idx](tmp)
		if nil != err {
			return nil, idx, err
		}
	}
	return tmp, -1, nil
}

// ----------------------------

// RegisterBytesProcess registers a handler for the given BytesHandlerType.
// RegisterBytesProcess 为指定 BytesHandlerType 注册处理函数。
func RegisterBytesProcess(t BytesHandlerType, process BytesHandler) {
	t2handler[t] = process
}

// LittleBlockHandler encodes in as a little-endian uint16-length-prefixed block.
// LittleBlockHandler 将小端 uint16 长度前缀块编码应用到输入数据。
func LittleBlockHandler(in []byte) (out []byte, err error) {
	return blockHandler(in, binary.LittleEndian)
}

// LittleUnblockProcess decodes one little-endian length-prefixed block from in.
// LittleUnblockProcess 从输入中解码一个小端长度前缀块。
func LittleUnblockProcess(in []byte) (out []byte, err error) {
	o, _, ok := unblockHandler(in, binary.LittleEndian)
	if ok {
		return o, nil
	}
	return nil, ErrLittleUnblock
}

// BigBlockHandler encodes in as a big-endian uint16-length-prefixed block.
// BigBlockHandler 将大端 uint16 长度前缀块编码应用到输入数据。
func BigBlockHandler(in []byte) (out []byte, err error) {
	return blockHandler(in, binary.BigEndian)
}

// BigUnblockProcess decodes one big-endian length-prefixed block from in.
// BigUnblockProcess 从输入中解码一个大端长度前缀块。
func BigUnblockProcess(in []byte) (out []byte, err error) {
	o, _, ok := unblockHandler(in, binary.BigEndian)
	if ok {
		return o, nil
	}
	return nil, ErrBigUnblock
}

//
//func Base64StdEncodeHandler(in []byte) (out []byte, err error) {
//	return base64Encode(in, base64.StdEncoding)
//}
//
//func Base64StdDecodeHandler(in []byte) (out []byte, err error) {
//	return base64Decode(in, base64.StdEncoding)
//}
//
//func Base64RawStdEncodeHandler(in []byte) (out []byte, err error) {
//	return base64Encode(in, base64.RawStdEncoding)
//}
//
//func Base64RawStdDecodeHandler(in []byte) (out []byte, err error) {
//	return base64Decode(in, base64.RawStdEncoding)
//}
//
//func Base64UrlEncodeHandler(in []byte) (out []byte, err error) {
//	return base64Encode(in, base64.URLEncoding)
//}
//
//func Base64UrlDecodeHandler(in []byte) (out []byte, err error) {
//	return base64Decode(in, base64.URLEncoding)
//}
//
//func Base64RawUrlEncodeHandler(in []byte) (out []byte, err error) {
//	return base64Encode(in, base64.RawURLEncoding)
//}
//
//func Base64RawUrlDecodeHandler(in []byte) (out []byte, err error) {
//	return base64Decode(in, base64.RawURLEncoding)
//}

// ----------------------------

func init() {
	RegisterBytesProcess(LittleBlock, LittleBlockHandler)
	RegisterBytesProcess(LittleUnblock, LittleUnblockProcess)
	RegisterBytesProcess(BigBlock, BigBlockHandler)
	RegisterBytesProcess(BigUnblock, BigUnblockProcess)
	//
	//RegisterBytesProcess(Base64StdEncode, Base64StdEncodeHandler)
	//RegisterBytesProcess(Base64StdDecode, Base64StdDecodeHandler)
	//RegisterBytesProcess(Base64RawStdEncode, Base64RawStdEncodeHandler)
	//RegisterBytesProcess(Base64RawStdDecode, Base64RawStdDecodeHandler)
	//
	//RegisterBytesProcess(Base64UrlEncode, Base64UrlEncodeHandler)
	//RegisterBytesProcess(Base64UrlDecode, Base64UrlDecodeHandler)
	//RegisterBytesProcess(Base64RawUrlEncode, Base64RawUrlEncodeHandler)
	//RegisterBytesProcess(Base64RawUrlDecode, Base64RawUrlDecodeHandler)
}

func blockHandler(in []byte, order binary.ByteOrder) (out []byte, err error) {
	l := uint16(len(in))
	if l == 0 {
		return []byte{0, 0}, nil
	}
	rs := make([]byte, len(in)+2)
	order.PutUint16(rs[:2], l)
	copy(rs[2:], in)
	return rs, nil
}

func unblockHandler(in []byte, order binary.ByteOrder) (out []byte, usedLen int, ok bool) {
	const lenLn = 2
	blockLen := len(in)
	if blockLen < lenLn {
		return nil, 0, false
	}
	var packLen = int(order.Uint16(in[:lenLn]))
	if 0 == packLen {
		return nil, lenLn, true
	}
	if blockLen < packLen+lenLn {
		return nil, 0, false
	}
	return in[lenLn : lenLn+packLen], packLen + lenLn, true
}

//
//func base64Encode(in []byte, enc *base64.Encoding) (out []byte, err error) {
//	out = make([]byte, enc.EncodedLen(len(in)))
//	enc.Encode(out, in)
//	return
//}
//
//func base64Decode(in []byte, enc *base64.Encoding) (out []byte, err error) {
//	dbuf := make([]byte, enc.DecodedLen(len(in)))
//	n, err := enc.Decode(dbuf, in)
//	return dbuf[:n], err
//}
