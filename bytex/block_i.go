package bytex

import "encoding/binary"

// BlockSizeLen is the byte length of the length prefix in a block (uint16).
// BlockSizeLen 是块长度前缀占用的字节数（uint16）。
var BlockSizeLen = 2

// HandlerDataToBlock encodes raw data into a length-prefixed block.
// HandlerDataToBlock 将原始数据编码为带长度前缀的块。
type HandlerDataToBlock func(data []byte, order binary.ByteOrder) (block []byte)

// HandlerBlockToData decodes one block and returns payload, consumed length, and success flag.
// HandlerBlockToData 从块中解码出数据，返回数据切片、消耗长度与是否成功。
type HandlerBlockToData func(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool)

type iOrderSetter interface {
	// SetOrder sets the byte order used for encoding and decoding.
	// SetOrder 设置编解码使用的字节序。
	SetOrder(order binary.ByteOrder)
}

type iOrderGetter interface {
	// GetOrder returns the current byte order.
	// GetOrder 返回当前字节序。
	GetOrder() binary.ByteOrder
}

type iToBlockHandler interface {
	// DataToBlock wraps data with a length prefix.
	// DataToBlock 对数据封装长度前缀。
	DataToBlock(data []byte) (block []byte)
	// SetDataToBlockHandler replaces the data-to-block handler.
	// SetDataToBlockHandler 设置数据到块的转换 handler。
	SetDataToBlockHandler(handler HandlerDataToBlock)
}

type iToDataHandler interface {
	// BlockToData extracts one framed payload from block.
	// BlockToData 从块中拆分出一个数据包。
	BlockToData(block []byte) (data []byte, length int, ok bool)
	// SetBlockToDataHandler replaces the block-to-data handler.
	// SetBlockToDataHandler 设置块到数据的转换 handler。
	SetBlockToDataHandler(handler HandlerBlockToData)
}

// IDataToBlockHandler encodes data to length-prefixed blocks.
// IDataToBlockHandler 将数据编码为带长度前缀的块。
type IDataToBlockHandler interface {
	iOrderSetter
	iOrderGetter
	iToBlockHandler
}

// IBlockToDataHandler decodes length-prefixed blocks to data.
// IBlockToDataHandler 将带长度前缀的块解码为数据。
type IBlockToDataHandler interface {
	iOrderSetter
	iOrderGetter
	iToDataHandler
}

// IDataBlockHandler combines block encoding and decoding.
// IDataBlockHandler 同时支持块的编码与解码。
type IDataBlockHandler interface {
	iOrderSetter
	iOrderGetter
	iToBlockHandler
	iToDataHandler
}
