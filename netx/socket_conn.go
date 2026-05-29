package netx

import (
	"fmt"
	"sync"
)

const (
	sep = "-"
)

var (
	connNum  = 0
	connLock sync.Mutex
)

// NewConnInfo creates connection info with a unique auto-increment id.
// NewConnInfo 创建带自增唯一 ID 的连接信息。
func NewConnInfo(localAddress string, removeAddress string) IConnInfo {
	rs := &_ConnInfo{localAddress: localAddress, remoteAddress: removeAddress}
	connLock.Lock()
	connNum++
	rs.connNum = connNum
	connLock.Unlock()
	// NAT网关可能出现端口复用，使得服务器中多个连接的四元信息一致，无法区别连接，所以需要加上一个自增ID
	rs.connId = fmt.Sprintf("[%d]%s%v", rs.connNum, sep, removeAddress)
	return rs
}

// NewRemoteOnlyConnInfo creates connection info keyed only by remote address.
// NewRemoteOnlyConnInfo 创建仅以远程地址为标识的连接信息。
func NewRemoteOnlyConnInfo(localAddress string, removeAddress string) IConnInfo {
	return &_ConnInfo{
		connNum:       -1,
		connId:        removeAddress,
		localAddress:  localAddress,
		remoteAddress: removeAddress}
}

// NewRemoteMuiltConnInfo creates connection info for one-to-many remote endpoints.
// NewRemoteMuiltConnInfo 创建面向多远程端点的一对多连接信息。
func NewRemoteMuiltConnInfo(localAddress string, removeAddress string) IConnInfo {
	return &_ConnInfo{
		connNum:       0,
		connId:        localAddress,
		localAddress:  localAddress,
		remoteAddress: removeAddress}
}

// IConnInfo describes connection identity and addresses.
// IConnInfo 描述连接标识与地址信息。
type IConnInfo interface {
	// One2One reports whether this is a one-to-one connection.
	// One2One 是否为一对一连接。
	One2One() bool
	// GetConnNum returns the connection sequence number.
	// GetConnNum 获取连接序号。
	GetConnNum() int
	// GetConnId returns the connection id.
	// GetConnId 获取连接 ID。
	GetConnId() string
	// GetLocalAddress returns the local address.
	// GetLocalAddress 获取连接的本地地址。
	GetLocalAddress() string
	// GetRemoteAddress returns the remote address.
	// GetRemoteAddress 获取连接的远程地址。
	GetRemoteAddress() string
}

// IConnInfoSetter injects IConnInfo.
// IConnInfoSetter 注入 IConnInfo。
type IConnInfoSetter interface {
	SetConnInfo(connInfo IConnInfo)
}

// IConnInfoGetter returns the bound IConnInfo.
// IConnInfoGetter 返回已绑定的 IConnInfo。
type IConnInfoGetter interface {
	GetConnInfo() IConnInfo
}

type _ConnInfo struct {
	connNum       int
	connId        string
	localAddress  string
	remoteAddress string
}

func (o *_ConnInfo) One2One() bool {
	return o.connNum != 0
}

func (o *_ConnInfo) String() string {
	if o.connNum > 0 {
		return o.connId
	}
	return fmt.Sprintf("[%d]%s->%s", o.connNum, o.localAddress, o.remoteAddress)
}

func (o *_ConnInfo) GetConnNum() int {
	return o.connNum
}

func (o *_ConnInfo) GetConnId() string {
	return o.connId
}

func (o *_ConnInfo) GetLocalAddress() string {
	return o.localAddress
}

func (o *_ConnInfo) GetRemoteAddress() string {
	return o.remoteAddress
}
