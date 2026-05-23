package bytex

import "io"

// IBuffByteReader reads raw bytes from a buffer without length-prefixed framing.
// IBuffByteReader 字节处理，不用读取长度信息。
type IBuffByteReader interface {
	io.Reader
	// Bytes returns all bytes currently held in the buffer.
	// Bytes 缓冲中全部字节。
	Bytes() []byte
	// ReadBytes drains and returns all bytes in the buffer; the slice shares backing storage and is not safe to retain.
	// ReadBytes 读取缓冲中全部字节；返回切片共享底层存储，非数据安全。
	ReadBytes() []byte
	// ReadBytesTo drains the buffer into dst and returns the number of bytes copied and the written portion of dst.
	// ReadBytesTo 读取缓冲中全部字节，并写入到 dst 中。
	ReadBytesTo(dst []byte) (n int, rs []byte)
	// ReadBytesCopy drains the buffer and returns a copy of all bytes; triggers Reset after read.
	// ReadBytesCopy 读取缓冲中全部字节的克隆；读取后触发 Reset，数据安全。
	ReadBytesCopy() []byte
	// ReadBinary reads one binary value into out using the buffer byte order; out must be a type supported by binaryx.Write.
	// ReadBinary 读取一个二进制数据到 out；out 只支持 binaryx.Write 中支持的类型。
	ReadBinary(out interface{}) error
}

// IBuffByteWriter writes raw bytes to a buffer without length-prefixed framing.
// IBuffByteWriter 字节处理，不用写入长度信息。
type IBuffByteWriter interface {
	io.Writer
	// WriteBytes appends bytes to the buffer.
	// WriteBytes 把字节写入缓冲。
	WriteBytes(bytes []byte)
	// WriteBinary writes one binary value using the buffer byte order.
	// WriteBinary 把 in 写入缓冲。
	WriteBinary(in interface{}) error
}

// IBuffDataReader reads length-prefixed data blocks from a buffer.
// IBuffDataReader 数据级处理，先读取长度信息，然后再根据长度读取数据。
type IBuffDataReader interface {
	io.Reader
	// ReadData reads one block and returns the payload; returns nil if data is insufficient; not safe to retain the slice.
	// ReadData 读取一个 Block 字节数据，并拆分出数据部分返回，数据不足返回 nil；非数据安全。
	ReadData() []byte
	// ReadString reads one block and returns the payload as string; returns empty if data is insufficient; not safe to retain.
	// ReadString 读取一个 Block 字节数据，并拆分出字符串部分返回，数据不足返回空；非数据安全。
	ReadString() string
	// ReadDataTo reads one block into dst when dst is non-nil; returns copied length and the written portion of dst.
	// ReadDataTo 读取一个 Block 字节数据并写入 dst；数据不足返回 nil。
	ReadDataTo(dst []byte) (n int, rs []byte)
	// ReadDataCopy reads one block and returns a copy of the payload; safe to retain.
	// ReadDataCopy 读取一个 Block 字节数据并返回数据部分的克隆；数据安全。
	ReadDataCopy() []byte
}

// IBuffDataWriter writes length-prefixed data blocks to a buffer.
// IBuffDataWriter 数据级处理，先写入长度信息，然后再写入数据。
type IBuffDataWriter interface {
	io.Writer
	// WriteData wraps bytes as one block and appends it; no-op when len(bytes) is 0.
	// WriteData 把数据包装为一个 Block 写入缓冲；数据长度为 0 时不进行处理。
	WriteData(bytes []byte)
	// WriteString wraps str as one block and appends it; no-op when len(str) is 0.
	// WriteString 把字符串包装为一个 Block 写入缓冲；数据长度为 0 时不进行处理。
	WriteString(str string)
}

// IBuffReset clears the underlying buffer.
// IBuffReset 提供清空缓冲区的能力。
type IBuffReset interface {
	// Reset discards all buffered data.
	// Reset 清空缓冲区。
	Reset()
}

// IBuffLen reports the number of unread bytes in the buffer.
// IBuffLen 报告缓冲中未读字节数。
type IBuffLen interface {
	Len() int
}

// IBuffToBlock is a buffer that writes framed blocks and reads raw bytes.
// IBuffToBlock 写入分帧块、读取原始字节的缓冲。
type IBuffToBlock interface {
	iOrderGetter
	IBuffByteReader
	IBuffDataWriter
	IBuffReset
	IBuffLen
}

// IBuffToData is a buffer that reads framed blocks and writes raw bytes.
// IBuffToData 读取分帧块、写入原始字节的缓冲。
type IBuffToData interface {
	iOrderGetter
	IBuffDataReader
	IBuffByteWriter
	IBuffReset
	IBuffLen
}

// IBuffDataBlock is a full-duplex buffer for both raw bytes and framed blocks.
// IBuffDataBlock 同时支持原始字节与分帧块读写的全双工缓冲。
type IBuffDataBlock interface {
	iOrderGetter
	IBuffDataWriter
	IBuffByteReader
	IBuffDataReader
	IBuffByteWriter
	IBuffReset
	IBuffLen
}
