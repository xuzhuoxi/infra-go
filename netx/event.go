package netx

const (
	ServerEventStart = "netx.ServerEventStart"
	ServerEventStop  = "netx.ServerEventStop"

	// ServerEventConnOpened is dispatched when a connection opens; event data: IConnInfo.
	// ServerEventConnOpened 连接打开事件；事件数据格式：IConnInfo。
	ServerEventConnOpened = "netx.ServerEventConnOpened"
	// ServerEventConnClosed is dispatched when a connection closes; event data: IConnInfo.
	// ServerEventConnClosed 连接关闭事件；事件数据格式：IConnInfo。
	ServerEventConnClosed = "netx.ServerEventConnClosed"
)

const (
	// EventUserConnMappingAdded is dispatched when a user-connection mapping is added; event data: UserConnMapperEventInfo.
	// EventUserConnMappingAdded 用户连接映射添加事件；事件数据格式：UserConnMapperEventInfo。
	EventUserConnMappingAdded = "netx.EventUserConnMappingAdded"
	// EventUserConnMappingRemoved is dispatched when a user-connection mapping is removed; event data: UserConnMapperEventInfo.
	// EventUserConnMappingRemoved 用户连接映射移除事件；事件数据格式：UserConnMapperEventInfo。
	EventUserConnMappingRemoved = "netx.EventUserConnMappingRemoved"
)

// UserConnMapperEventInfo carries user-connection mapping event payload.
// UserConnMapperEventInfo 承载用户与连接映射事件数据。
type UserConnMapperEventInfo struct {
	Key    string
	UserID string
	ConnId string
}

const (
	// EventUserAddressMappingAdded is dispatched when a user-address mapping is added; event data: UserAddressMapperEventInfo.
	// EventUserAddressMappingAdded 用户地址映射添加事件；事件数据格式：UserAddressMapperEventInfo。
	EventUserAddressMappingAdded = "netx.EventUserAddressMappingAdded"
	// EventUserAddressMappingRemoved is dispatched when a user-address mapping is removed; event data: UserAddressMapperEventInfo.
	// EventUserAddressMappingRemoved 用户地址映射移除事件；事件数据格式：UserAddressMapperEventInfo。
	EventUserAddressMappingRemoved = "netx.EventUserAddressMappingRemoved"
)

// UserAddressMapperEventInfo carries user-address mapping event payload.
// UserAddressMapperEventInfo 承载用户与地址映射事件数据。
type UserAddressMapperEventInfo struct {
	Key     string
	UserId  string
	Address string
}
