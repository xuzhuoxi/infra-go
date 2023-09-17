// Package netx
// Created by xuzhuoxi
// on 2019-03-13.
// @author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"sync"
)

const (
	// EventAddressAdded
	// EventData: AddrProxyEventInfo
	EventAddressAdded = "EventAddressAdded"
	// EventAddressRemoved
	// EventData: AddrProxyEventInfo
	EventAddressRemoved = "EventAddressRemoved"
)

type AddrProxyEventInfo struct {
	Id   string
	Addr string
}

type IAddressProxySetter interface {
	SetAddressProxy(proxy IAddressProxy)
}

type IAddressProxyGetter interface {
	GetAddressProxy() IAddressProxy
}

// IAddressProxy
// 难住地址与id的双向映射
type IAddressProxy interface {
	eventx.IEventDispatcher
	// GetId
	// 能过地址找id
	GetId(address string) (id string, ok bool)
	// GetAddress
	// 能过id找地址
	GetAddress(id string) (address string, ok bool)
	// MapIdAddress
	// 把id和地址加入映射表
	MapIdAddress(id string, address string)
	// RemoveById
	// 移除id相关映射
	RemoveById(id string)
	// RemoveByAddress
	// 移除地址相关映射
	RemoveByAddress(address string)
	// Reset
	// 重置
	Reset()
}

func NewIAddressProxy() IAddressProxy {
	return NewAddressProxy()
}

func NewAddressProxy() *AddressProxy {
	return NewAddressProxyWithName("default")
}

func NewIAddressProxyWithName(name string) IAddressProxy {
	return NewAddressProxyWithName(name)
}

func NewAddressProxyWithName(name string) *AddressProxy {
	return &AddressProxy{
		name:   name,
		idAddr: make(map[string]string),
		addrId: make(map[string]string)}
}

type AddressProxy struct {
	eventx.EventDispatcher
	name   string
	idAddr map[string]string
	addrId map[string]string
	lock   sync.RWMutex
}

func (o *AddressProxy) Reset() {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.idAddr = make(map[string]string)
	o.addrId = make(map[string]string)
}

func (o *AddressProxy) GetId(address string) (id string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	id, ok = o.addrId[address]
	return
}

func (o *AddressProxy) GetAddress(id string) (address string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	address, ok = o.idAddr[id]
	return
}

func (o *AddressProxy) MapIdAddress(id string, address string) {
	//fmt.Println(fmt.Sprintf("AddressProxy[%s].MapIdAddress:", p.name), id, address)
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.checkGroup(id, address) {
		return
	}
	_, _ = o.removeId(id)
	_, _ = o.removeAddress(address)
	o.mapIdAddr(id, address)
	o.DispatchEvent(EventAddressAdded, o, AddrProxyEventInfo{Id: id, Addr: address})
}

func (o *AddressProxy) RemoveById(id string) {
	//fmt.Println(fmt.Sprintf("AddressProxy[%s].RemoveById:", o.name), id)
	o.lock.Lock()
	defer o.lock.Unlock()
	address, ok := o.removeId(id)
	if ok {
		o.DispatchEvent(EventAddressRemoved, o, AddrProxyEventInfo{Id: id, Addr: address})
	}
}

func (o *AddressProxy) RemoveByAddress(address string) {
	//fmt.Println(fmt.Sprintf("AddressProxy[%s].RemoveByAddress:", o.name), address)
	o.lock.Lock()
	defer o.lock.Unlock()
	id, ok := o.removeAddress(address)
	if ok {
		o.DispatchEvent(EventAddressRemoved, o, AddrProxyEventInfo{Id: id, Addr: address})
	}
}

func (o *AddressProxy) mapIdAddr(id string, address string) {
	o.idAddr[id] = address
	o.addrId[address] = id
}

func (o *AddressProxy) removeId(id string) (address string, ok bool) {
	if address, ok = o.idAddr[id]; ok {
		delete(o.addrId, address)
		delete(o.idAddr, id)
		return address, true
	}
	return "", false
}

func (o *AddressProxy) removeAddress(address string) (id string, ok bool) {
	if id, ok = o.addrId[address]; ok {
		delete(o.idAddr, id)
		delete(o.addrId, address)
		return id, true
	}
	return "", false
}

func (o *AddressProxy) checkGroup(id string, address string) bool {
	address1, ok1 := o.idAddr[id]
	id2, ok2 := o.addrId[address]
	if ok1 && ok2 && address == address1 && id == id2 {
		return true
	}
	return false
}

//func (p *AddressProxy) traceLen() {
//	fmt.Println("AddressProxy Len:", len(p.idAddr), len(p.addrId))
//}
