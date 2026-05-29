package netx

import (
	"sync"

	"github.com/xuzhuoxi/infra-go/eventx"
)

// IUserConnMapper maintains bidirectional mapping between user id and connection id.
// IUserConnMapper 维护用户 ID 与连接 ID 的双向映射。
type IUserConnMapper interface {
	eventx.IEventDispatcher
	// GetUserId returns the user id for the given connection id.
	// GetUserId 通过连接 id 查找用户 id。
	GetUserId(connId string) (userId string, ok bool)
	// GetConnId returns the connection id for the given user id.
	// GetConnId 通过用户 id 查找连接 id。
	GetConnId(userId string) (connId string, ok bool)
	// MapInto adds or updates a user id and connection id pair.
	// MapInto 把用户 id 和连接 id 加入映射表。
	MapInto(userId string, connId string)
	// RemoveByUserId removes mapping by user id.
	// RemoveByUserId 移除用户 id 相关映射。
	RemoveByUserId(userId string) (suc bool, connId string)
	// RemoveByConnId removes mapping by connection id.
	// RemoveByConnId 移除连接 id 相关映射。
	RemoveByConnId(connId string) (suc bool, userId string)
	// Size returns the number of mappings.
	// Size 获取映射表大小。
	Size() int
	// Reset clears all mappings.
	// Reset 重置映射表。
	Reset()
}

// IUserConnMapperSetter injects an IUserConnMapper.
// IUserConnMapperSetter 注入 IUserConnMapper。
type IUserConnMapperSetter interface {
	SetUserConnMapper(mapper IUserConnMapper)
}

// IUserConnMapperGetter returns the bound IUserConnMapper.
// IUserConnMapperGetter 返回已绑定的 IUserConnMapper。
type IUserConnMapperGetter interface {
	GetUserConnMapper() IUserConnMapper
}

// NewIUserConnMapper creates a default IUserConnMapper.
// NewIUserConnMapper 创建默认的 IUserConnMapper。
func NewIUserConnMapper() IUserConnMapper {
	return NewUserConnMapper()
}

// NewUserConnMapper creates a UserConnMapper named "default".
// NewUserConnMapper 创建名为 "default" 的 UserConnMapper。
func NewUserConnMapper() *UserConnMapper {
	return NewUserConnMapperWithName("default")
}

// NewIUserConnMapperWithName creates an IUserConnMapper with the given name.
// NewIUserConnMapperWithName 创建指定名称的 IUserConnMapper。
func NewIUserConnMapperWithName(name string) IUserConnMapper {
	return NewUserConnMapperWithName(name)
}

// NewUserConnMapperWithName creates a UserConnMapper with the given name.
// NewUserConnMapperWithName 创建指定名称的 UserConnMapper。
func NewUserConnMapperWithName(name string) *UserConnMapper {
	return &UserConnMapper{
		name:      name,
		user2conn: make(map[string]string),
		conn2user: make(map[string]string)}
}

// UserConnMapper is the default IUserConnMapper implementation.
// UserConnMapper 是 IUserConnMapper 的默认实现。
type UserConnMapper struct {
	eventx.EventDispatcher
	name      string
	user2conn map[string]string
	conn2user map[string]string
	lock      sync.RWMutex
}

// GetUserId returns the user id for the given connection id.
// GetUserId 通过连接 id 查找用户 id。
func (o *UserConnMapper) GetUserId(connId string) (userId string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	userId, ok = o.conn2user[connId]
	return
}

// GetConnId returns the connection id for the given user id.
// GetConnId 通过用户 id 查找连接 id。
func (o *UserConnMapper) GetConnId(userId string) (connId string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	connId, ok = o.user2conn[userId]
	return
}

// MapInto adds or updates a user id and connection id pair.
// MapInto 把用户 id 和连接 id 加入映射表。
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

// RemoveByUserId removes mapping by user id.
// RemoveByUserId 移除用户 id 相关映射。
func (o *UserConnMapper) RemoveByUserId(userId string) (suc bool, connId string) {
	o.lock.Lock()
	connId, suc = o.removeUserId(userId)
	o.lock.Unlock()
	if suc {
		o.DispatchEvent(EventUserConnMappingRemoved, o, UserConnMapperEventInfo{Key: userId, UserID: userId, ConnId: connId})
	}
	return
}

// RemoveByConnId removes mapping by connection id.
// RemoveByConnId 移除连接 id 相关映射。
func (o *UserConnMapper) RemoveByConnId(connId string) (suc bool, userId string) {
	o.lock.Lock()
	userId, suc = o.removeConnId(connId)
	o.lock.Unlock()
	if suc {
		o.DispatchEvent(EventUserConnMappingRemoved, o, UserConnMapperEventInfo{Key: connId, UserID: userId, ConnId: connId})
	}
	return
}

// Size returns the number of mappings.
// Size 获取映射表大小。
func (o *UserConnMapper) Size() int {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return len(o.conn2user)
}

// Reset clears all mappings.
// Reset 重置映射表。
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

// IUserAddressMapper maintains bidirectional mapping between user id and network address.
// IUserAddressMapper 维护用户 id 与通信地址的双向映射。
type IUserAddressMapper interface {
	eventx.IEventDispatcher
	// GetUserId returns the user id for the given address.
	// GetUserId 通过地址查找用户 id。
	GetUserId(address string) (id string, ok bool)
	// GetAddress returns the address for the given user id.
	// GetAddress 通过用户 id 查找地址。
	GetAddress(id string) (address string, ok bool)
	// MapInto adds or updates a user id and address pair.
	// MapInto 把 id 和地址加入映射表。
	MapInto(id string, address string)
	// RemoveByUserId removes mapping by user id.
	// RemoveByUserId 移除 id 相关映射。
	RemoveByUserId(id string)
	// RemoveByAddress removes mapping by address.
	// RemoveByAddress 移除地址相关映射。
	RemoveByAddress(address string)
	// Size returns the number of mappings.
	// Size 获取映射表大小。
	Size() int
	// Reset clears all mappings.
	// Reset 重置映射表。
	Reset()
}

// IUserAddressMapperSetter injects an IUserAddressMapper.
// IUserAddressMapperSetter 注入 IUserAddressMapper。
type IUserAddressMapperSetter interface {
	SetUserAddressMapper(proxy IUserAddressMapper)
}

// IUserAddressMapperGetter returns the bound IUserAddressMapper.
// IUserAddressMapperGetter 返回已绑定的 IUserAddressMapper。
type IUserAddressMapperGetter interface {
	GetUserAddressMapper() IUserAddressMapper
}

// NewIUserAddressMapper creates a default IUserAddressMapper.
// NewIUserAddressMapper 创建默认的 IUserAddressMapper。
func NewIUserAddressMapper() IUserAddressMapper {
	return NewUserAddressMapper()
}

// NewUserAddressMapper creates a UserAddressMapper named "default".
// NewUserAddressMapper 创建名为 "default" 的 UserAddressMapper。
func NewUserAddressMapper() *UserAddressMapper {
	return NewUserAddressMapperWithName("default")
}

// NewIUserAddressMapperWithName creates an IUserAddressMapper with the given name.
// NewIUserAddressMapperWithName 创建指定名称的 IUserAddressMapper。
func NewIUserAddressMapperWithName(name string) IUserAddressMapper {
	return NewUserAddressMapperWithName(name)
}

// NewUserAddressMapperWithName creates a UserAddressMapper with the given name.
// NewUserAddressMapperWithName 创建指定名称的 UserAddressMapper。
func NewUserAddressMapperWithName(name string) *UserAddressMapper {
	return &UserAddressMapper{
		name:         name,
		user2address: make(map[string]string),
		address2user: make(map[string]string)}
}

// UserAddressMapper is the default IUserAddressMapper implementation.
// UserAddressMapper 是 IUserAddressMapper 的默认实现。
type UserAddressMapper struct {
	eventx.EventDispatcher
	name         string
	user2address map[string]string
	address2user map[string]string
	lock         sync.RWMutex
}

// GetUserId returns the user id for the given address.
// GetUserId 通过地址查找用户 id。
func (o *UserAddressMapper) GetUserId(address string) (id string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	id, ok = o.address2user[address]
	return
}

// GetAddress returns the address for the given user id.
// GetAddress 通过用户 id 查找地址。
func (o *UserAddressMapper) GetAddress(id string) (address string, ok bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	address, ok = o.user2address[id]
	return
}

// MapInto adds or updates a user id and address pair.
// MapInto 把 id 和地址加入映射表。
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

// RemoveByUserId removes mapping by user id.
// RemoveByUserId 移除 id 相关映射。
func (o *UserAddressMapper) RemoveByUserId(id string) {
	//fmt.Println(fmt.Sprintf("UserAddressMapper[%s].RemoveByUserId:", o.name), id)
	o.lock.Lock()
	address, ok := o.removeUserId(id)
	o.lock.Unlock()
	if ok {
		o.DispatchEvent(EventUserAddressMappingRemoved, o, UserAddressMapperEventInfo{Key: id, UserId: id, Address: address})
	}
}

// RemoveByAddress removes mapping by address.
// RemoveByAddress 移除地址相关映射。
func (o *UserAddressMapper) RemoveByAddress(address string) {
	//fmt.Println(fmt.Sprintf("UserAddressMapper[%s].RemoveByAddress:", o.name), address)
	o.lock.Lock()
	id, ok := o.removeAddress(address)
	o.lock.Unlock()
	if ok {
		o.DispatchEvent(EventUserAddressMappingRemoved, o, UserAddressMapperEventInfo{Key: address, UserId: id, Address: address})
	}
}

// Size returns the number of mappings.
// Size 获取映射表大小。
func (o *UserAddressMapper) Size() int {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return len(o.user2address)
}

// Reset clears all mappings.
// Reset 重置映射表。
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
