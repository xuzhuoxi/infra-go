// Package protox
// Created by xuzhuoxi
// on 2019-05-20.
// @author xuzhuoxi
//
package protox

import "fmt"

// IExtensionHeader
// Extension参数头接口
type IExtensionHeader interface {
	// ExtensionName
	// 请求Extension名称
	ExtensionName() string
	// ProtoId
	// 请求Extension中对应的协议标识
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
	GetHeaderInfo() IExtensionHeader
}

type ExtensionHeader struct {
	EName    string
	PId      string
	CId      string
	CAddress string
}

func (h *ExtensionHeader) String() string {
	return fmt.Sprintf("Header{ENmae='%s', PId='%s', CID='%s', CAddr='%s'}",
		h.EName, h.PId, h.CId, h.CAddress)
}

func (h *ExtensionHeader) GetHeaderInfo() IExtensionHeader {
	return h
}

func (h *ExtensionHeader) ExtensionName() string {
	return h.EName
}

func (h *ExtensionHeader) ProtoId() string {
	return h.PId
}

func (h *ExtensionHeader) ClientId() string {
	return h.CId
}

func (h *ExtensionHeader) ClientAddress() string {
	return h.CAddress
}

func (h *ExtensionHeader) SetHeader(extensionName string, protoId string, clientId string, clientAddress string) {
	h.EName, h.PId, h.CId, h.CAddress = extensionName, protoId, clientId, clientAddress
}
