// Package bytex
// Created by xuzhuoxi
// on 2019-02-11.
// @author xuzhuoxi
//
package bytex

import (
	"encoding/binary"
)

// DefaultDataToBlockHandler
// block是安全的
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

// DefaultBlockToDataHandler
// data为共享切片，非安全
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

func (h *dataBlockHandler) DataToBlock(data []byte) (block []byte) {
	return h.data2Block(data, h.order)
}

func (h *dataBlockHandler) BlockToData(block []byte) (data []byte, length int, ok bool) {
	return h.block2Data(block, h.order)
}
