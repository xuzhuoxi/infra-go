// Package netx
// Create on 2025/2/15
// @author xuzhuoxi
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"sync"
)

// IUserConnMapper
// 用户Id与连接id的双向映射
type IUserConnMapper interface {
	eventx.IEventDispatcher
	// GetUserId
	// 通过 用户id 找 连接id
	GetUserId(connId string) (UserId string, ok bool)
	// GetConnId
	// 通过 连接id 找 用户id
	GetConnId(userId string) (connId string, ok bool)
	// MapInto
	// 把用户id和连接Id加入映射表
	MapInto(userId string, connId string)
	// RemoveByUserId
	// 移除用户id相关映射
	RemoveByUserId(userId string) (suc bool, connId string)
	// RemoveByConnId
	// 移除连接id相关映射
	RemoveByConnId(connId string) (suc bool, userId string)
	// Size
	// 获取映射表大小
	Size() int
	// Reset
	// 重置
	Reset()
}

type IUserConnMapperSetter interface {
	SetUserConnMapper(mapper IUserConnMapper)
}

type IUserConnMapperGetter interface {
	GetUserConnMapper() IUserConnMapper
}

func NewIUserConnMapper() IUserConnMapper {
	return NewUserConnMapper()
}

func NewUserConnMapper() *UserConnMapper {
	return NewUserConnMapperWithName("default")
}

func NewIUserConnMapperWithName(name string) IUserConnMapper {
	return NewUserConnMapperWithName(name)
}

func NewUserConnMapperWithName(name string) *UserConnMapper {
	return &UserConnMapper{
		name:      name,
		user2conn: make(map[string]string),
		conn2user: make(map[string]string)}
}

type UserConnMapper struct {
	eventx.EventDispatcher
	name      string
	user2conn map[string]string
	conn2user map[string]string
	lock      sync.RWMutex
}

func (o *UserConnMapper) GetUserId(connId string) (UserId string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	connId, ok = o.conn2user[connId]
	return
}

func (o *UserConnMapper) GetConnId(userId string) (connId string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	connId, ok = o.user2conn[userId]
	return
}

func (o *UserConnMapper) MapInto(userId string, connId string) {
	o.lock.Lock()
	if o.checkMapping(userId, connId) {
		o.lock.Unlock()
		return
	}
	_, _ = o.removeUserId(userId)
	_, _ = o.removeConnId(connId)
	o.mapInfo(userId, connId)
	o.lock.Unlock()
	o.DispatchEvent(EventUserConnMappingAdded, o, UserConnMapperEventInfo{UserID: userId, ConnId: connId})
}

func (o *UserConnMapper) RemoveByUserId(userId string) (suc bool, connId string) {
	o.lock.Lock()
	connId, suc = o.removeUserId(userId)
	o.lock.Unlock()
	if suc {
		o.DispatchEvent(EventUserConnMappingRemoved, o, UserConnMapperEventInfo{Key: userId, UserID: userId, ConnId: connId})
	}
	return
}

func (o *UserConnMapper) RemoveByConnId(connId string) (suc bool, userId string) {
	o.lock.Lock()
	userId, suc = o.removeConnId(connId)
	o.lock.Unlock()
	if suc {
		o.DispatchEvent(EventUserConnMappingRemoved, o, UserConnMapperEventInfo{Key: connId, UserID: userId, ConnId: connId})
	}
	return
}

func (o *UserConnMapper) Size() int {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return len(o.conn2user)
}

func (o *UserConnMapper) Reset() {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.user2conn = make(map[string]string)
	o.conn2user = make(map[string]string)
}

func (o *UserConnMapper) mapInfo(userId string, connId string) {
	o.user2conn[userId] = connId
	o.conn2user[connId] = userId
}

func (o *UserConnMapper) removeUserId(userId string) (connId string, ok bool) {
	if connId, ok = o.user2conn[userId]; ok {
		delete(o.user2conn, userId)
		delete(o.conn2user, connId)
		return
	}
	return "", false
}

func (o *UserConnMapper) removeConnId(connId string) (userId string, ok bool) {
	if userId, ok = o.conn2user[connId]; ok {
		delete(o.conn2user, connId)
		delete(o.user2conn, userId)
		return
	}
	return "", false
}

func (o *UserConnMapper) checkMapping(userId string, connId string) bool {
	address1, ok1 := o.user2conn[userId]
	id2, ok2 := o.conn2user[connId]
	if ok1 && ok2 && connId == address1 && userId == id2 {
		return true
	}
	return false
}

// ————————————————————————————————————————————————————————————————————————————————————————————

// IUserAddressMapper
// 用户id 与 通信地址 的双向映射
type IUserAddressMapper interface {
	eventx.IEventDispatcher
	// GetUserId
	// 通过过地址找id
	GetUserId(address string) (id string, ok bool)
	// GetAddress
	// 通过id找地址
	GetAddress(id string) (address string, ok bool)
	// MapInto
	// 把id和地址加入映射表
	MapInto(id string, address string)
	// RemoveByUserId
	// 移除id相关映射
	RemoveByUserId(id string)
	// RemoveByAddress
	// 移除地址相关映射
	RemoveByAddress(address string)
	// Size
	// 获取映射表大小
	Size() int
	// Reset
	// 重置
	Reset()
}

type IUserAddressMapperSetter interface {
	SetUserAddressMapper(proxy IUserAddressMapper)
}

type IUserAddressMapperGetter interface {
	GetUserAddressMapper() IUserAddressMapper
}

func NewIUserAddressMapper() IUserAddressMapper {
	return NewUserAddressMapper()
}

func NewUserAddressMapper() *UserAddressMapper {
	return NewUserAddressMapperWithName("default")
}

func NewIUserAddressMapperWithName(name string) IUserAddressMapper {
	return NewIUserAddressMapperWithName(name)
}

func NewUserAddressMapperWithName(name string) *UserAddressMapper {
	return &UserAddressMapper{
		name:         name,
		user2address: make(map[string]string),
		address2user: make(map[string]string)}
}

type UserAddressMapper struct {
	eventx.EventDispatcher
	name         string
	user2address map[string]string
	address2user map[string]string
	lock         sync.RWMutex
}

func (o *UserAddressMapper) GetUserId(address string) (id string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	id, ok = o.address2user[address]
	return
}

func (o *UserAddressMapper) GetAddress(id string) (address string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	address, ok = o.user2address[id]
	return
}

func (o *UserAddressMapper) MapInto(id string, address string) {
	//fmt.Println(fmt.Sprintf("UserAddressMapper[%s].MapInto:", p.name), id, address)
	o.lock.Lock()
	if o.checkMapping(id, address) {
		o.lock.Unlock()
		return
	}
	_, _ = o.removeUserId(id)
	_, _ = o.removeAddress(address)
	o.mapInfo(id, address)
	o.lock.Unlock()
	o.DispatchEvent(EventUserAddressMappingAdded, o, UserAddressMapperEventInfo{UserId: id, Address: address})
}

func (o *UserAddressMapper) RemoveByUserId(id string) {
	//fmt.Println(fmt.Sprintf("UserAddressMapper[%s].RemoveByUserId:", o.name), id)
	o.lock.Lock()
	address, ok := o.removeUserId(id)
	o.lock.Unlock()
	if ok {
		o.DispatchEvent(EventUserAddressMappingRemoved, o, UserAddressMapperEventInfo{Key: id, UserId: id, Address: address})
	}
}

func (o *UserAddressMapper) RemoveByAddress(address string) {
	//fmt.Println(fmt.Sprintf("UserAddressMapper[%s].RemoveByAddress:", o.name), address)
	o.lock.Lock()
	id, ok := o.removeAddress(address)
	o.lock.Unlock()
	if ok {
		o.DispatchEvent(EventUserAddressMappingRemoved, o, UserAddressMapperEventInfo{Key: address, UserId: id, Address: address})
	}
}

func (o *UserAddressMapper) Size() int {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return len(o.user2address)
}

func (o *UserAddressMapper) Reset() {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.user2address = make(map[string]string)
	o.address2user = make(map[string]string)
}

func (o *UserAddressMapper) mapInfo(userId string, address string) {
	o.user2address[userId] = address
	o.address2user[address] = userId
}

func (o *UserAddressMapper) removeUserId(userId string) (address string, ok bool) {
	if address, ok = o.user2address[userId]; ok {
		delete(o.address2user, address)
		delete(o.user2address, userId)
		return address, true
	}
	return "", false
}

func (o *UserAddressMapper) removeAddress(address string) (userId string, ok bool) {
	if userId, ok = o.address2user[address]; ok {
		delete(o.user2address, userId)
		delete(o.address2user, address)
		return userId, true
	}
	return "", false
}

func (o *UserAddressMapper) checkMapping(userId string, address string) bool {
	address1, ok1 := o.user2address[userId]
	id2, ok2 := o.address2user[address]
	if ok1 && ok2 && address == address1 && userId == id2 {
		return true
	}
	return false
}
