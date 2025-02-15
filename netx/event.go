package netx

const (
	ServerEventStart = "netx.ServerEventStart"
	ServerEventStop  = "netx.ServerEventStop"

	// ServerEventConnOpened
	// 事件数据格式: IConnInfo
	ServerEventConnOpened = "netx.ServerEventConnOpened"
	// ServerEventConnClosed
	// 事件数据格式: IConnInfo
	ServerEventConnClosed = "netx.ServerEventConnClosed"
)

const (
	// EventUserConnMappingAdded
	// 事件数据格式: UserConnMapperEventInfo
	EventUserConnMappingAdded = "netx.EventUserConnMappingAdded"
	// EventUserConnMappingRemoved
	// 事件数据格式: UserConnMapperEventInfo
	EventUserConnMappingRemoved = "netx.EventUserConnMappingRemoved"
)

type UserConnMapperEventInfo struct {
	Key    string
	UserID string
	ConnId string
}

const (
	// EventUserAddressMappingAdded
	// EventData: UserAddressMapperEventInfo
	EventUserAddressMappingAdded = "netx.EventUserAddressMappingAdded"
	// EventUserAddressMappingRemoved
	// EventData: UserAddressMapperEventInfo
	EventUserAddressMappingRemoved = "netx.EventUserAddressMappingRemoved"
)

type UserAddressMapperEventInfo struct {
	Key     string
	UserId  string
	Address string
}
