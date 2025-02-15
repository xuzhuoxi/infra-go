// Package netx
// Create on 2025/2/14
// @author xuzhuoxi
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

func NewRemoteOnlyConnInfo(localAddress string, removeAddress string) IConnInfo {
	return &_ConnInfo{
		connNum:       -1,
		connId:        removeAddress,
		localAddress:  localAddress,
		remoteAddress: removeAddress}
}

func NewRemoteMuiltConnInfo(localAddress string, removeAddress string) IConnInfo {
	return &_ConnInfo{
		connNum:       0,
		connId:        localAddress,
		localAddress:  localAddress,
		remoteAddress: removeAddress}
}

type IConnInfo interface {
	// One2One 是否为一对一连接
	One2One() bool
	// GetConnNum 获取连接序号
	GetConnNum() int
	// GetConnId 获取连接ID
	GetConnId() string
	// GetLocalAddress 获取连接的本地地址
	GetLocalAddress() string
	// GetRemoteAddress 获取连接的远程地址
	GetRemoteAddress() string
}

type IConnInfoSetter interface {
	SetConnInfo(connInfo IConnInfo)
}

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
