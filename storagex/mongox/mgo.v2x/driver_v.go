package mgo_v2x

import (
	"io"
	"sync"

	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/storagex/mongox"
	"gopkg.in/mgo.v2"
)

// NewMongoDriverV creates a driver that manages multiple named mgo sessions.
// NewMongoDriverV 创建管理多个具名 mgo 会话的驱动。
//
// maxSessions enables a concurrency limit when greater than zero; zero disables the limit.
// maxSessions 大于 0 时启用并发限制；为 0 时不限制。
func NewMongoDriverV(maxSessions int) IMongoDriverV {
	rs := &MongoDriverV{}
	if maxSessions > 0 {
		rs.ChanLimit = &lang.ChannelLimit{}
	} else {
		rs.ChanLimit = &lang.ChannelLimitNone{}
	}
	rs.ChanLimit.StartLimit()
	return rs
}

// IMongoDriverV manages the lifecycle of multiple named mgo sessions.
// IMongoDriverV 管理多个具名 mgo 会话的生命周期。
type IMongoDriverV interface {
	io.Closer

	// OpenSession dials a session with DialInfo and sets its consistency mode.
	// OpenSession 使用 DialInfo 拨号创建会话并设置一致性模式。
	OpenSession(name string, info *mgo.DialInfo, mode mgo.Mode) (session IMongoSessionV, err error)
	// OpenSessionByURL dials a session from a single address URL.
	// OpenSessionByURL 从单个地址 URL 拨号创建会话。
	OpenSessionByURL(name string, url string, mode mgo.Mode) (session IMongoSessionV, err error)
	// CloseSession closes and removes a session by name.
	// CloseSession 按名称关闭并移除会话。
	CloseSession(name string) error
	// CloneSession clones an existing session under a new name.
	// CloneSession 克隆已有会话并以新名称注册。
	CloneSession(session IMongoSessionV, newName string) IMongoSessionV
	// NumSessions returns the number of open sessions.
	// NumSessions 返回已打开会话数量。
	NumSessions() int

	// GetSession returns a session by name.
	// GetSession 按名称返回会话。
	GetSession(name string) (session IMongoSessionV, ok bool)
	// SetMasterSession sets or clears the master session by name.
	// SetMasterSession 按名称设置或清除主会话。
	SetMasterSession(sessionName string) (master IMongoSessionV, change bool)
	// MasterSession returns the current master session, or nil if unset.
	// MasterSession 返回当前主会话；未设置时返回 nil。
	MasterSession() IMongoSessionV
}

// MongoDriverV holds named mgo sessions with optional concurrency limiting.
// MongoDriverV 持有具名 mgo 会话，并支持可选的并发限制。
type MongoDriverV struct {
	Sessions    []IMongoSessionV
	MasterIndex int

	ChanLimit lang.IChannelLimit
	Lock      sync.RWMutex
}

// Close closes all sessions and stops the concurrency limiter.
// Close 关闭所有会话并停止并发限制器。
func (d *MongoDriverV) Close() error {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if 0 < len(d.Sessions) {
		for index := len(d.Sessions) - 1; index >= 0; index-- {
			d.Sessions[index].Close()
		}
	}
	d.ChanLimit.StopLimit()
	return nil
}

// OpenSession dials a session with DialInfo and sets its consistency mode.
// OpenSession 使用 DialInfo 拨号创建会话并设置一致性模式。
func (d *MongoDriverV) OpenSession(name string, info *mgo.DialInfo, mode mgo.Mode) (session IMongoSessionV, err error) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	//funcName := "MongoDriverV.OpenSession"
	d.ChanLimit.Add()
	sess, e := mgo.DialWithInfo(info)
	if nil != e {
		d.ChanLimit.Done()
		return nil, e
	}
	sess.SetMode(mode, true)
	session = &MongoSessionV{Key: name, SessionDialInfo: info, MgoSession: sess}
	session.LocateDB("")
	d.appendSession(session)
	return
}

// OpenSessionByURL dials a session from a single address URL.
// OpenSessionByURL 从单个地址 URL 拨号创建会话。
func (d *MongoDriverV) OpenSessionByURL(name string, url string, mode mgo.Mode) (session IMongoSessionV, err error) {
	info := &mgo.DialInfo{Addrs: []string{url}}
	return d.OpenSession(name, info, mode)
}

// CloseSession closes and removes a session by name.
// CloseSession 按名称关闭并移除会话。
func (d *MongoDriverV) CloseSession(name string) error {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if "" == name {
		return mongox.SessionNilError("MongoDriverV.CloseSession name=\"\"")
	}
	index := d.indexOfName(name)
	if -1 != index {
		defer func() {
			if index == d.MasterIndex {
				d.MasterIndex = -1
			}
			d.removeSession(d.Sessions[index])
			d.ChanLimit.Done()
		}()
		return d.Sessions[index].Close()
	}
	return mongox.SessionNilError("MongoDriverV.CloseSession name=\"" + name + "\"")
}

// CloneSession clones an existing session under a new name.
// CloneSession 克隆已有会话并以新名称注册。
func (d *MongoDriverV) CloneSession(session IMongoSessionV, newName string) IMongoSessionV {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if nil == session {
		return nil
	}
	d.ChanLimit.Add()
	sess := session.cloneSession(newName)
	d.appendSession(sess)
	return sess
}

// NumSessions returns the number of open sessions.
// NumSessions 返回已打开会话数量。
func (d *MongoDriverV) NumSessions() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	return len(d.Sessions)
}

// MasterSession returns the current master session, or nil if unset.
// MasterSession 返回当前主会话；未设置时返回 nil。
func (d *MongoDriverV) MasterSession() IMongoSessionV {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	if -1 == d.MasterIndex {
		return nil
	}
	return d.Sessions[d.MasterIndex]
}

// GetSession returns a session by name.
// GetSession 按名称返回会话。
func (d *MongoDriverV) GetSession(name string) (session IMongoSessionV, ok bool) {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	if "" == name {
		return nil, false
	}
	index := d.indexOfName(name)
	if -1 == index {
		return nil, false
	}
	return d.Sessions[index], true
}

// SetMasterSession sets or clears the master session by name.
// SetMasterSession 按名称设置或清除主会话。
func (d *MongoDriverV) SetMasterSession(sessionName string) (master IMongoSessionV, change bool) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if "" == sessionName {
		if -1 == d.MasterIndex {
			return nil, false
		} else {
			d.MasterIndex = -1
			return nil, true
		}
	}
	index := d.indexOfName(sessionName)
	change = index == d.MasterIndex
	d.MasterIndex = index
	return d.Sessions[index], change
}

//------------------------------------------------------------

func (d *MongoDriverV) indexOfSession(session IMongoSessionV) int {
	if nil == session {
		return -1
	}
	return d.indexOfName(session.GetSessionName())
}

func (d *MongoDriverV) indexOfName(name string) int {
	if "" == name {
		return -1
	}
	for index := range d.Sessions {
		if name == d.Sessions[index].GetSessionName() {
			return index
		}
	}
	return -1
}

func (d *MongoDriverV) appendSession(session IMongoSessionV) {
	if nil != session {
		d.Sessions = append(d.Sessions, session)
	}
}

func (d *MongoDriverV) removeSession(session IMongoSessionV) {
	index := d.indexOfSession(session)
	if -1 == index {
		return
	}
	d.Sessions = append(d.Sessions[:index], d.Sessions[index+1:]...)
}
