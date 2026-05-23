package bytex

import (
	"encoding/binary"
)

// DefaultDataToBlockHandler encodes data as a uint16-length-prefixed block; the returned block is safe to retain.
// DefaultDataToBlockHandler 将数据编码为 uint16 长度前缀块；返回的 block 可安全持有。
func DefaultDataToBlockHandler(data []byte, order binary.ByteOrder) (block []byte) {
	l := uint16(len(data))
	if 0 == l {
		return []byte{0, 0}
	}
	rs := make([]byte, l+2)
	order.PutUint16(rs[:2], l)
	copy(rs[2:], data)
	return rs
}

// DefaultBlockToDataHandler decodes one block; data references a sub-slice of block and is not safe to retain.
// DefaultBlockToDataHandler 从块中解码一个数据包；data 为 block 的共享子切片，非安全持有。
func DefaultBlockToDataHandler(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool) {
	totalLen := len(block)
	if totalLen < BlockSizeLen {
		return nil, 0, false
	}
	var packLen = int(order.Uint16(block[:BlockSizeLen]))
	if 0 == packLen {
		return nil, BlockSizeLen, true
	}
	length = packLen + BlockSizeLen
	if totalLen < length {
		return nil, 0, false
	}
	return block[BlockSizeLen:length], length, true
}

//----------------------------------------------------------

// NewDefaultDataBlockHandler returns a handler using DefaultOrder and the default encode/decode functions.
// NewDefaultDataBlockHandler 使用 DefaultOrder 与默认编解码函数创建处理器。
func NewDefaultDataBlockHandler() IDataBlockHandler {
	return newDataBlockHandler(DefaultOrder, DefaultDataToBlockHandler, DefaultBlockToDataHandler)
}

// NewDataBlockHandler returns a handler with custom byte order and encode/decode functions.
// NewDataBlockHandler 使用自定义字节序与编解码函数创建处理器。
func NewDataBlockHandler(order binary.ByteOrder, data2block HandlerDataToBlock, block2data HandlerBlockToData) IDataBlockHandler {
	return newDataBlockHandler(order, data2block, block2data)
}

// NewDataToBlockHandler returns an encode-only handler.
// NewDataToBlockHandler 创建仅编码方向的处理器。
func NewDataToBlockHandler(order binary.ByteOrder, data2block HandlerDataToBlock) IDataToBlockHandler {
	return newDataBlockHandler(order, data2block, nil)
}

// NewBlockToDataHandler returns a decode-only handler.
// NewBlockToDataHandler 创建仅解码方向的处理器。
func NewBlockToDataHandler(order binary.ByteOrder, block2data HandlerBlockToData) IBlockToDataHandler {
	return newDataBlockHandler(order, nil, block2data)
}

//----------------------------------------------------------

func newDataBlockHandler(order binary.ByteOrder, data2block HandlerDataToBlock, block2data HandlerBlockToData) *dataBlockHandler {
	return &dataBlockHandler{order: order, data2Block: data2block, block2Data: block2data}
}

type dataBlockHandler struct {
	order      binary.ByteOrder
	data2Block HandlerDataToBlock
	block2Data HandlerBlockToData
}

func (h *dataBlockHandler) SetDataToBlockHandler(handler HandlerDataToBlock) {
	h.data2Block = handler
}

func (h *dataBlockHandler) SetBlockToDataHandler(handler HandlerBlockToData) {
	h.block2Data = handler
}

func (h *dataBlockHandler) SetOrder(order binary.ByteOrder) {
	h.order = order
}

func (h *dataBlockHandler) GetOrder() binary.ByteOrder {
	return h.order
}

func (h *dataBlockHandler) DataToBlock(data []byte) (block []byte) {
	return h.data2Block(data, h.order)
}

func (h *dataBlockHandler) BlockToData(block []byte) (data []byte, length int, ok bool) {
	return h.block2Data(block, h.order)
}
