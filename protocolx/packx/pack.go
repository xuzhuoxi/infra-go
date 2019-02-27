//
//Created by xuzhuoxi
//on 2019-02-13.
//@author xuzhuoxi
//
package packx

import (
	"bytes"
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
)

type IPackByte interface {
	encodingx.ICodingData
	ProtocolId() []byte
	ProtocolData() [][]byte
	SetId(id []byte)
	SetData(data ...[]byte)
	Set(id []byte, data ...[]byte)
}

func NewDefaultPackByte() IPackByte {
	return &PackByte{buff: bytes.NewBuffer(nil), dataBlockHandler: bytex.NewDefaultDataBlockHandler()}
}

func NewPackByte(dataBlockHandler bytex.IDataBlockHandler) IPackByte {
	return &PackByte{buff: bytes.NewBuffer(nil), dataBlockHandler: dataBlockHandler}
}

type PackByte struct {
	id   []byte
	data [][]byte
	buff *bytes.Buffer

	dataBlockHandler bytex.IDataBlockHandler
}

func (d *PackByte) ProtocolId() []byte {
	return d.id
}

func (d *PackByte) ProtocolData() [][]byte {
	return d.data
}

func (d *PackByte) SetId(id []byte) {
	d.id = id
}

func (d *PackByte) SetData(data ...[]byte) {
	d.data = data
}

func (d *PackByte) Set(id []byte, data ...[]byte) {
	d.SetId(id)
	d.SetData(data...)
}

func (d *PackByte) EncodeToBytes() []byte {
	d.buff.Reset()
	d.buff.Write(d.dataBlockHandler.DataToBlock(d.id))
	if len(d.data) > 0 {
		for index := 0; index < len(d.data); index++ {
			block := d.dataBlockHandler.DataToBlock(d.data[index])
			d.buff.Write(block)
		}
	}
	return d.dataBlockHandler.DataToBlock(d.buff.Bytes())
}

func (d *PackByte) DecodeFromBytes(bs []byte) bool {
	data, _, ok := d.dataBlockHandler.BlockToData(bs)
	if !ok {
		return false
	}
	d.buff.Reset()
	d.buff.Write(data)
	id, l, ok := d.dataBlockHandler.BlockToData(d.buff.Bytes())
	if !ok {
		return false
	}
	d.buff.Next(l)
	d.id = id
	d.data = nil
	for d.buff.Len() > 0 {
		sd, l, ok := d.dataBlockHandler.BlockToData(d.buff.Bytes())
		if !ok {
			return false
		}
		d.buff.Next(l)
		d.data = append(d.data, sd)
	}
	return true
}

//-----------------------------

type IPackData interface {
	encodingx.ICodingData
	ProtocolId() string
	ProtocolData() []interface{}
	SetId(id string)
	SetData(data ...interface{})
	Set(id string, data ...interface{})
}

func NewDefaultPackData(codeHandler encodingx.ICodingHandler) IPackData {
	return &PackData{packByte: NewDefaultPackByte(), codeHandler: codeHandler}
}

func NewPackData(packByte IPackByte, codeHandler encodingx.ICodingHandler) IPackData {
	return &PackData{packByte: packByte, codeHandler: codeHandler}
}

type PackData struct {
	id          string
	data        []interface{} //结构体数组
	tempData    interface{}   //结构体
	packByte    IPackByte
	codeHandler encodingx.ICodingHandler
}

func (d *PackData) ProtocolId() string {
	return d.id
}

func (d *PackData) ProtocolData() []interface{} {
	return d.data
}

func (d *PackData) SetId(id string) {
	d.id = id
}

func (d *PackData) SetData(data ...interface{}) {
	d.data = data
}

func (d *PackData) Set(id string, data ...interface{}) {
	d.SetId(id)
	d.SetData(data...)
}

func (d *PackData) EncodeToBytes() []byte {
	id := d.codeHandler.HandleEncode(d.id)
	data := [][]byte{}
	for index := 0; index < len(d.data); index++ {
		data = append(data, d.codeHandler.HandleEncode(d.data[index]))
	}
	d.packByte.Set(id, data...)
	return d.packByte.EncodeToBytes()
}

func (d *PackData) DecodeFromBytes(bs []byte) bool {
	ok := d.packByte.DecodeFromBytes(bs)
	if !ok {
		return false
	}
	d.codeHandler.HandleDecode(d.packByte.ProtocolId(), &d.id)
	d.data = nil
	for _, bs := range d.packByte.ProtocolData() {
		d.codeHandler.HandleDecode(bs, &d.tempData)
		temp := d.tempData
		d.data = append(d.data, temp)
	}
	return true
}
