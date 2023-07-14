// Package bytex
// Create on 2023/6/18
// @author xuzhuoxi
package bytex

import "encoding/binary"

var BlockSizeLen = 2

type HandlerDataToBlock func(data []byte, order binary.ByteOrder) (block []byte)

// HandlerBlockToData
// 数据数组 + Block长度 + 成功?
type HandlerBlockToData func(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool)

type iOrderSetter interface {
	SetOrder(order binary.ByteOrder)
}

type iOrderGetter interface {
	GetOrder() binary.ByteOrder
}

type iToBlockHandler interface {
	// DataToBlock
	// 对数据封装上长度
	DataToBlock(data []byte) (block []byte)
	SetDataToBlockHandler(handler HandlerDataToBlock)
}

type iToDataHandler interface {
	// BlockToData
	// 拆分一个数据出来
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
