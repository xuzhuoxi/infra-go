//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package bytex

import "encoding/binary"

type HandlerDataToBlock func(data []byte, order binary.ByteOrder) (block []byte)

//数据数组　+　Block长度 + 成功?
type HandlerBlockToData func(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool)

//block是安全的
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

//data为共享切片，非安全
func DefaultBlockToDataHandler(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool) {
	blockLen := len(block)
	if blockLen < 2 {
		return nil, 0, false
	}
	var packLen = int(order.Uint16(block[:2]))
	if 0 == packLen {
		return nil, 2, true
	}
	if blockLen < packLen+2 {
		return nil, 0, false
	}
	return block[2 : 2+packLen], packLen + 2, true
}

//----------------------------------------------------------

type iOrderSetter interface {
	SetOrder(order binary.ByteOrder)
}

type iOrderGetter interface {
	GetOrder() binary.ByteOrder
}

type iToBlockHandler interface {
	//对数据封装上长度
	DataToBlock(data []byte) (block []byte)
	SetDataToBlockHandler(handler HandlerDataToBlock)
}

type iToDataHandler interface {
	//拆分一个数据出来
	BlockToData(block []byte) (data []byte, length int, ok bool)
	SetBlockToDataHandler(handler HandlerBlockToData)
}

type IDataToBlockHandler interface {
	iOrderSetter
	iOrderGetter
	iToBlockHandler
}

type IBlockToDataHandler interface {
	iOrderSetter
	iOrderGetter
	iToDataHandler
}

type IDataBlockHandler interface {
	iOrderSetter
	iOrderGetter
	iToBlockHandler
	iToDataHandler
}

func NewDefaultDataBlockHandler() IDataBlockHandler {
	return newDataBlockHandler(DefaultOrder, DefaultDataToBlockHandler, DefaultBlockToDataHandler)
}

func NewDataBlockHandler(order binary.ByteOrder, data2block HandlerDataToBlock, block2data HandlerBlockToData) IDataBlockHandler {
	return newDataBlockHandler(order, data2block, block2data)
}

func NewDataToBlockHandler(order binary.ByteOrder, data2block HandlerDataToBlock) IDataToBlockHandler {
	return newDataBlockHandler(order, data2block, nil)
}

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

func (h dataBlockHandler) DataToBlock(data []byte) (block []byte) {
	return h.data2Block(data, h.order)
}

func (h dataBlockHandler) BlockToData(block []byte) (data []byte, length int, ok bool) {
	return h.block2Data(block, h.order)
}
