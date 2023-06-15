package bytex

import (
	"encoding/binary"
	"errors"
)

// BytesHandlerType 字节数组处理类型
type BytesHandlerType int

// BytesHandler 字节数组处理函数
type BytesHandler func(in []byte) (out []byte, err error)

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
	ErrTypeUnregister = errors.New("Type Unregister! ")

	ErrLittleUnblock = errors.New("LittleUnblock Error! ")
	ErrBigUnblock    = errors.New("BigUnblock Error! ")
)

var t2handler = make([]BytesHandler, 128, 128)

// IBytesProcessor
// 字节数组处理器
type IBytesProcessor interface {
	// AppendHandler
	// 追加处理函数
	AppendHandler(process BytesHandler)
	// AppendHandlerType
	// 追加处理类型
	AppendHandlerType(t BytesHandlerType) error
	// ClearHandlers
	// 清除全部处理
	ClearHandlers()

	// Handle
	// 数据处理
	Handle(in []byte) (out []byte, errIdx int, err error)
}

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

func RegisterBytesProcess(t BytesHandlerType, process BytesHandler) {
	t2handler[t] = process
}

func LittleBlockHandler(in []byte) (out []byte, err error) {
	return blockHandler(in, binary.LittleEndian)
}

func LittleUnblockProcess(in []byte) (out []byte, err error) {
	o, _, ok := unblockHandler(in, binary.LittleEndian)
	if ok {
		return o, nil
	}
	return nil, ErrLittleUnblock
}

func BigBlockHandler(in []byte) (out []byte, err error) {
	return blockHandler(in, binary.BigEndian)
}

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
