//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package bytex

import "encoding/binary"

type DataToBlockHandler func(data []byte, order binary.ByteOrder) (block []byte)

//数据数组　+　Block长度 + 成功?
type BlockToDataHandler func(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool)

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
func DefaultBlockToDataHandler(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool) {
	var l = int(order.Uint16(block[:2]))
	if 0 == l && 2 == len(block) {
		return nil, 2, true
	}
	if len(block) < l+2 {
		return nil, 0, false
	}
	return block[2 : 2+l], l + 2, true
}

//----------------------------------------------------------

type IDataToBlockHandler interface {
	//对数据封装上长度
	DataToBlock(data []byte) (block []byte)
}

type IBlockToDataHandler interface {
	//拆分一个数据出来
	BlockToData(block []byte) (data []byte, length int, ok bool)
}

type IDataBlockHandler interface {
	IDataToBlockHandler
	IBlockToDataHandler
}

var DefaultDataBlockHandler = NewDefaultDataBlockHandler()

func NewDefaultDataBlockHandler() IDataBlockHandler {
	return newDataBlockHandler(DefaultOrder, DefaultDataToBlockHandler, DefaultBlockToDataHandler)
}

func NewDataBlockHandler(order binary.ByteOrder, data2block DataToBlockHandler, block2data BlockToDataHandler) IDataBlockHandler {
	return newDataBlockHandler(order, data2block, block2data)
}

func NewDataToBlockHandler(order binary.ByteOrder, data2block DataToBlockHandler) IDataToBlockHandler {
	return newDataBlockHandler(order, data2block, nil)
}

func NewBlockToDataHandler(order binary.ByteOrder, block2data BlockToDataHandler) IBlockToDataHandler {
	return newDataBlockHandler(order, nil, block2data)
}

//----------------------------------------------------------

func newDataBlockHandler(order binary.ByteOrder, data2block DataToBlockHandler, block2data BlockToDataHandler) *dataBlockHandler {
	return &dataBlockHandler{order: order, data2Block: data2block, block2Data: block2data}
}

type dataBlockHandler struct {
	order      binary.ByteOrder
	data2Block DataToBlockHandler
	block2Data BlockToDataHandler
}

func (h *dataBlockHandler) DataToBlock(data []byte) (block []byte) {
	return h.data2Block(data, h.order)
}

func (h *dataBlockHandler) BlockToData(block []byte) (data []byte, length int, ok bool) {
	return h.block2Data(block, h.order)
}
