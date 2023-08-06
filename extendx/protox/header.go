// Package protox
// Created by xuzhuoxi
// on 2019-05-20.
// @author xuzhuoxi
//
package protox

import "fmt"

// IProtoHeader
// 协议参数头接口
type IProtoHeader interface {
	// ProtoGroup
	// 协议分组
	ProtoGroup() string
	// ProtoId
	// 协议标识
	ProtoId() string
	// ClientId
	// 客户端标识
	ClientId() string
	// ClientAddress
	// 客户端地址
	ClientAddress() string
	// SetHeader
	// 设置参数头信息
	SetHeader(extensionName string, protoId string, clientId string, clientAddress string)
	// GetHeaderInfo 取头信息
	GetHeaderInfo() IProtoHeader
}

type ProtoHeader struct {
	PGroup   string
	PId      string
	CId      string
	CAddress string
}

func (h *ProtoHeader) String() string {
	return fmt.Sprintf("Header{ENmae='%s', PId='%s', CID='%s', CAddr='%s'}",
		h.PGroup, h.PId, h.CId, h.CAddress)
}

func (h *ProtoHeader) GetHeaderInfo() IProtoHeader {
	return h
}

func (h *ProtoHeader) ProtoGroup() string {
	return h.PGroup
}

func (h *ProtoHeader) ProtoId() string {
	return h.PId
}

func (h *ProtoHeader) ClientId() string {
	return h.CId
}

func (h *ProtoHeader) ClientAddress() string {
	return h.CAddress
}

func (h *ProtoHeader) SetHeader(extensionName string, protoId string, clientId string, clientAddress string) {
	h.PGroup, h.PId, h.CId, h.CAddress = extensionName, protoId, clientId, clientAddress
}
