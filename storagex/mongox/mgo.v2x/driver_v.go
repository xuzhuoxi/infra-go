package mgo_v2x

import (
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/storagex/mongox"
	"gopkg.in/mgo.v2"
	"io"
	"sync"
)

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

type IMongoDriverV interface {
	io.Closer

	OpenSession(name string, info *mgo.DialInfo, mode mgo.Mode) (session IMongoSessionV, err error)
	OpenSessionByURL(name string, url string, mode mgo.Mode) (session IMongoSessionV, err error)
	CloseSession(name string) error
	CloneSession(session IMongoSessionV, newName string) IMongoSessionV
	NumSessions() int

	GetSession(name string) (session IMongoSessionV, ok bool)
	SetMasterSession(sessionName string) (master IMongoSessionV, change bool)
	MasterSession() IMongoSessionV
}

type MongoDriverV struct {
	Sessions    []IMongoSessionV
	MasterIndex int

	ChanLimit lang.IChannelLimit
	Lock      sync.RWMutex
}

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

func (d *MongoDriverV) OpenSessionByURL(name string, url string, mode mgo.Mode) (session IMongoSessionV, err error) {
	info := &mgo.DialInfo{Addrs: []string{url}}
	return d.OpenSession(name, info, mode)
}

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

func (d *MongoDriverV) NumSessions() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	return len(d.Sessions)
}

func (d *MongoDriverV) MasterSession() IMongoSessionV {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	if -1 == d.MasterIndex {
		return nil
	}
	return d.Sessions[d.MasterIndex]
}

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
