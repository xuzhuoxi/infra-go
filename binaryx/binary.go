package binaryx

import (
	"encoding/binary"
	"errors"
	"io"
)

// ReadLen reads a length from r as uint16 and returns it as int.
// ReadLen 从 Reader 中读取一个长度数据，按 uint16 格式读取并强制转换为 int 返回。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadLen(r io.Reader, order binary.ByteOrder) (ln int, err error) {
	return ReadLen2(r, order)
}

// ReadLen1 reads a length from r as uint8 and returns it as int.
// ReadLen1 从 Reader 中读取一个长度数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadLen1(r io.Reader, order binary.ByteOrder) (ln int, err error) {
	val := uint8(0)
	if err := binary.Read(r, order, &val); nil != err {
		return 0, err
	}
	return int(val), nil
}

// ReadLen2 reads a length from r as uint16 and returns it as int.
// ReadLen2 从 Reader 中读取一个长度数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadLen2(r io.Reader, order binary.ByteOrder) (ln int, err error) {
	val := uint16(0)
	if err := binary.Read(r, order, &val); nil != err {
		return 0, err
	}
	return int(val), nil
}

// ReadLen4 reads a length from r as uint32 and returns it as uint.
// ReadLen4 从 Reader 中读取一个长度数据。
// r: Reader.
// r: 读取源。
// order: byte order setting.
// order: 大小端设定。
func ReadLen4(r io.Reader, order binary.ByteOrder) (ln uint, err error) {
	val := uint32(0)
	if err := binary.Read(r, order, &val); nil != err {
		return 0, err
	}
	return uint(val), nil
}

// ReadLenTo reads a length from w into ln using the type of ln.
// ReadLenTo 从 Reader 中读取一个长度数据。
// w: Reader.
// w: 读取源。
// order: byte order setting.
// order: 大小端设定。
// ln: length value; supports int8, int16, int32, int64, uint8, uint16, uint32, uint64 and their pointer types.
// ln: 长度值，只支持 int8、int16、int32、int64、uint8、uint16、uint32、uint64 及其指针类型。
func ReadLenTo(w io.Reader, order binary.ByteOrder, ln interface{}) error {
	switch ln.(type) {
	case int8, int16, int32, int64, *int8, *int16, *int32, *int64:
		return binary.Read(w, order, ln)
	case uint8, uint16, uint32, uint64, *uint8, *uint16, *uint32, *uint64:
		return binary.Read(w, order, ln)
	}
	return errors.New("ln type error! ")
}

// WriteLen writes ln to w as uint16.
// WriteLen 把长度写入到 Writer 中。
// w: Writer.
// w: 写入目标。
// order: byte order setting.
// order: 大小端设定。
// ln: length value written as uint16.
// ln: 长度值，按 uint16 写入。
func WriteLen(w io.Writer, order binary.ByteOrder, ln int) error {
	return WriteLen2(w, order, ln)
}

// WriteLen1 writes ln to w as uint8.
// WriteLen1 把长度写入到 Writer 中。
// w: Writer.
// w: 写入目标。
// order: byte order setting.
// order: 大小端设定。
// ln: length value written as uint8.
// ln: 长度值，按 uint8 写入。
func WriteLen1(w io.Writer, order binary.ByteOrder, ln int) error {
	return binary.Write(w, order, uint8(ln))
}

// WriteLen2 writes ln to w as uint16.
// WriteLen2 把长度写入到 Writer 中。
// w: Writer.
// w: 写入目标。
// order: byte order setting.
// order: 大小端设定。
// ln: length value written as uint16.
// ln: 长度值，按 uint16 写入。
func WriteLen2(w io.Writer, order binary.ByteOrder, ln int) error {
	return binary.Write(w, order, uint16(ln))
}

// WriteLen4 writes ln to w as uint32.
// WriteLen4 把长度写入到 Writer 中。
// w: Writer.
// w: 写入目标。
// order: byte order setting.
// order: 大小端设定。
// ln: length value written as uint32.
// ln: 长度值，按 uint32 写入。
func WriteLen4(w io.Writer, order binary.ByteOrder, ln uint) error {
	return binary.Write(w, order, uint32(ln))
}

// WriteLenTo writes ln to w using the type of ln.
// WriteLenTo 把长度写入到 Writer 中。
// w: Writer.
// w: 写入目标。
// order: byte order setting.
// order: 大小端设定。
// ln: length value; supports int8, int16, int32, int64, uint8, uint16, uint32, uint64 and their pointer types.
// ln: 长度值，只支持 int8、int16、int32、int64、uint8、uint16、uint32、uint64 及其指针类型。
func WriteLenTo(w io.Writer, order binary.ByteOrder, ln interface{}) error {
	switch ln.(type) {
	case int8, int16, int32, int64, *int8, *int16, *int32, *int64:
		return binary.Write(w, order, ln)
	case uint8, uint16, uint32, uint64, *uint8, *uint16, *uint32, *uint64:
		return binary.Write(w, order, ln)
	}
	return errors.New("ln type error! ")
}
