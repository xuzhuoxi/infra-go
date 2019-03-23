//
//Created by xuzhuoxi
//on 2019-03-13.
//@author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"sync"
)

const (
	EventAddressRemoved = "EventAddressRemoved"
)

type IAddressProxySetter interface {
	SetAddressProxy(proxy IAddressProxy)
}

type IAddressProxyGetter interface {
	GetAddressProxy() IAddressProxy
}

//难住地址与id的双向映射
type IAddressProxy interface {
	eventx.IEventDispatcher
	//能过地址找id
	GetId(address string) (id string, ok bool)
	//能过id找地址
	GetAddress(id string) (address string, ok bool)
	//把id和地址加入映射表
	MapIdAddress(id string, address string)
	//移除id相关映射
	RemoveById(id string)
	//移除地址相关映射
	RemoveByAddress(address string)
	//重置
	Reset()
}

func NewIAddressProxy() IAddressProxy {
	return NewAddressProxy()
}

func NewAddressProxy() *AddressProxy {
	return &AddressProxy{idAddr: make(map[string]string), addrId: make(map[string]string)}
}

type AddressProxy struct {
	eventx.EventDispatcher
	idAddr map[string]string
	addrId map[string]string
	mu     sync.RWMutex
}

func (p *AddressProxy) Reset() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.idAddr = make(map[string]string)
	p.addrId = make(map[string]string)
}

func (p *AddressProxy) GetId(address string) (id string, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	id, ok = p.addrId[address]
	return
}

func (p *AddressProxy) GetAddress(id string) (address string, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	id, ok = p.idAddr[id]
	return
}

func (p *AddressProxy) MapIdAddress(id string, address string) {
	p.mu.Lock()
	if p.checkGroup(id, address) {
		p.mu.Unlock()
		return
	}
	var removeAddress string
	var ok bool
	defer func() {
		p.mu.Unlock()
		if ok {
			p.DispatchEvent(EventAddressRemoved, p, removeAddress)
		}
	}()
	removeAddress, ok = p.removeId(id)
	p.removeAddress(address)

	p.idAddr[id] = address
	p.addrId[address] = id
}

func (p *AddressProxy) RemoveById(id string) {
	p.mu.Lock()
	var address string
	var ok bool
	defer func() {
		p.mu.Unlock()
		if ok {
			p.DispatchEvent(EventAddressRemoved, p, address)
		}
	}()
	address, ok = p.removeId(id)
}

func (p *AddressProxy) RemoveByAddress(address string) {
	p.mu.Lock()
	var ok bool
	defer func() {
		p.mu.Unlock()
		if ok {
			p.DispatchEvent(EventAddressRemoved, p, address)
		}
	}()
	_, ok = p.removeAddress(address)
}

func (p *AddressProxy) removeId(id string) (address string, ok bool) {
	if address, ok := p.idAddr[id]; ok {
		delete(p.addrId, address)
		delete(p.idAddr, id)
		return address, true
	}
	return "", false
}

func (p *AddressProxy) removeAddress(address string) (id string, ok bool) {
	if id, ok := p.addrId[address]; ok {
		delete(p.idAddr, id)
		delete(p.addrId, address)
		return id, true
	}
	return "", false
}

func (p *AddressProxy) checkGroup(id string, address string) bool {
	address1, ok1 := p.idAddr[id]
	id2, ok2 := p.addrId[address]
	if ok1 && ok2 && address == address1 && id == id2 {
		return true
	}
	return false
}
