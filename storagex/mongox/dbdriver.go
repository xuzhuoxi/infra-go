package mongox

import (
	"github.com/xuzhuoxi/util-go/logx"
	"gopkg.in/mgo.v2"
	"io"
	"sync"
)

type IMongoDriver interface {
	Open(info *mgo.DialInfo, mode mgo.Mode, maxSessions int) error
	io.Closer
	GetDialInfo() *mgo.DialInfo
	OriginSession() IMongoSession

	NewSession(sync bool) (IMongoSession, error)
	NumSessions() int
}

type MongoDriver struct {
	DialInfo     *mgo.DialInfo
	MongoSession IMongoSession

	maxSessions  int
	chanSessions chan bool
	sessions     []IMongoSession
	mu           sync.Mutex
}

func (d *MongoDriver) GetDialInfo() *mgo.DialInfo {
	return d.DialInfo
}

func (d *MongoDriver) Open(info *mgo.DialInfo, mode mgo.Mode, maxSessions int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	funcName := "MongoDriver.Open"
	session, err := mgo.DialWithInfo(info)
	if nil != err {
		return err
	}
	session.SetMode(mode, true)
	d.MongoSession = &MongoSession{Session: session, CloseFunc: d.removeSession}
	d.MongoSession.LocateDB("")
	if maxSessions < 1 {
		maxSessions = 1
	}
	d.maxSessions = maxSessions
	d.chanSessions = make(chan bool, maxSessions)
	d.chanSessions <- true
	d.sessions = nil
	logx.Infoln(funcName + "()")
	return nil
}

func (d *MongoDriver) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	funcName := "MongoDriver.Close"
	l := len(d.sessions)
	if l > 0 {
		for index := l - 1; index >= 0; index-- {
			d.sessions[index].(*MongoSession).close()
		}
	}
	d.sessions = nil
	if nil != d.MongoSession {
		d.MongoSession.(*MongoSession).close()
		d.MongoSession = nil
		close(d.chanSessions)
		logx.Infoln(funcName + "()")
		return nil
	}
	return SessionNilError(funcName)
}

func (d *MongoDriver) OriginSession() IMongoSession {
	return d.MongoSession
}

func (d *MongoDriver) NewSession(sync bool) (IMongoSession, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	funcName := "MongoDriver.NewSession"
	if d.maxSessions == 1 {
		return nil, SessionLimitError(funcName + ";1")
	}
	if !sync && len(d.chanSessions) == d.maxSessions {
		return nil, SessionLimitError(funcName + ":sync")
	}
	d.chanSessions <- true
	ses := d.MongoSession.(*MongoSession).Clone()
	d.sessions = append(d.sessions, ses)
	return ses, nil
}

func (d *MongoDriver) NumSessions() int {
	return len(d.chanSessions)
}

func (d *MongoDriver) removeSession(session IMongoSession) {
	if nil == session || nil == d.sessions || len(d.sessions) == 0 {
		return
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	for index, v := range d.sessions {
		if v == session {
			d.sessions = append(d.sessions[:index], d.sessions[index+1:]...)
			return
		}
	}
	<-d.chanSessions
}
