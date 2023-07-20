// Package bytex
// Create on 2023/6/18
// @author xuzhuoxi
package bytex

import "io"

// IBuffByteReader
// 字节处理，不用读取长度信息
type IBuffByteReader interface {
	io.Reader
	// Bytes
	// 缓冲中全部字节
	Bytes() []byte
	// ReadBytes
	// 读取缓冲中全部字节
	// 非数据安全
	ReadBytes() []byte
	// ReadBytesTo
	// 读取缓冲中全部字节，并写入到dst中
	ReadBytesTo(dst []byte) (n int, rs []byte)
	// ReadBytesCopy
	// 读取缓冲中全部字节
	// 数据安全
	ReadBytesCopy() []byte
	// ReadBinary
	// 读取一个二进制数据到out
	// out只支持binary.Write中支持的类型
	ReadBinary(out interface{}) error
}

// IBuffByteWriter
// 字节处理，不用写入长度信息
type IBuffByteWriter interface {
	io.Writer
	// WriteBytes
	// 把字节写入缓冲
	WriteBytes(bytes []byte)
	// WriteBinary
	// 把in写入数据
	WriteBinary(in interface{}) error
}

// IBuffDataReader
// 数据级处理，先读取长度信息，然后再根据长度读取数据
type IBuffDataReader interface {
	io.Reader
	// ReadData
	// 读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	// 非数据安全
	ReadData() []byte
	// ReadString
	// 读取一个Block字节数据，并拆分出字符串部分返回，数据不足返回nil
	// 非数据安全
	ReadString() string
	// ReadDataTo
	// 读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	// 如果不是nil,把数据写入到dst中，返回dst写入的数据长度
	ReadDataTo(dst []byte) (n int, rs []byte)
	// ReadDataCopy
	// 读取一个Block字节数据，并拆分出数据部分返回，数据不足返回nil
	// 数据安全
	ReadDataCopy() []byte
}

// IBuffDataWriter
// 数据级处理，先写入长度信息，然后再写入数据
type IBuffDataWriter interface {
	io.Writer
	// WriteData
	// 把数据包装为一个Block,写入到缓冲中，数据长度为0时不进行处理
	WriteData(bytes []byte)
	// WriteString
	// 把字符串装为一个Block,写入到缓冲中，数据长度为0时不进行处理
	WriteString(str string)
}

type IBuffReset interface {
	// Reset
	// 清空缓冲区
	Reset()
}

type IBuffLen interface {
	Len() int
}

type IBuffToBlock interface {
	iOrderGetter
	IBuffByteReader
	IBuffDataWriter
	IBuffReset
	IBuffLen
}

type IBuffToData interface {
	iOrderGetter
	IBuffDataReader
	IBuffByteWriter
	IBuffReset
	IBuffLen
}

type IBuffDataBlock interface {
	iOrderGetter
	IBuffDataWriter
	IBuffByteReader
	IBuffDataReader
	IBuffByteWriter
	IBuffReset
	IBuffLen
}
