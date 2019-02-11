//
//Created by xuzhuoxi
//on 2019-02-11.
//@author xuzhuoxi
//
package encodingx

import "encoding/binary"

type DataToBlockHandler func(data []byte, order binary.ByteOrder) (block []byte, ok bool)
type BlockToDataHandler func(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool)

func DefaultDataToBlockHandler(data []byte, order binary.ByteOrder) (block []byte, ok bool) {
	l := uint16(len(data))
	if 0 == l {
		return nil, false
	}
	rs := make([]byte, l+2)
	order.PutUint16(rs[:2], l)
	copy(rs[2:], data)
	return rs, true
}
func DefaultBlockToDataHandler(block []byte, order binary.ByteOrder) (data []byte, length int, ok bool) {
	var l = int(order.Uint16(block[:2]))
	if len(block) < l+2 {
		return nil, 0, false
	}
	return block[2 : 2+l], l + 2, true
}

type IDataToBlockHandler interface {
	//设置由数据到Block的处理方法
	SetDataToBlockHandler(handler DataToBlockHandler)
}

type IBlockToDataHandler interface {
	//设置同数据到Block的处理方法
	SetBlockToDataHandler(handler BlockToDataHandler)
}
