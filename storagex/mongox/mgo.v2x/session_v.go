package mgo_v2x

import (
	"io"

	"gopkg.in/mgo.v2"
)

// IMongoSessionV wraps an mgo session with database and collection selection state.
// IMongoSessionV 封装 mgo 会话，并跟踪数据库与集合选择状态。
type IMongoSessionV interface {
	io.Closer
	cloneSession(newName string) IMongoSessionV

	// GetSessionName returns the session name.
	// GetSessionName 返回会话名称。
	GetSessionName() string
	// GetDialInfo returns the dial info used to create the session.
	// GetDialInfo 返回创建会话时使用的拨号信息。
	GetDialInfo() *mgo.DialInfo
	// GetMode returns the current session consistency mode.
	// GetMode 返回当前会话一致性模式。
	GetMode() mgo.Mode

	// LocateDB selects a database by name.
	// LocateDB 按名称选择数据库。
	LocateDB(dbName string) (db *mgo.Database, change bool)
	// LocateCollection selects a database and collection by name.
	// LocateCollection 按名称选择数据库与集合。
	LocateCollection(dbName string, cName string) (c *mgo.Collection, change bool)

	// Session returns the underlying mgo.Session.
	// Session 返回底层 mgo.Session。
	Session() *mgo.Session
	// CurrentDataBase returns the currently selected database.
	// CurrentDataBase 返回当前选中的数据库。
	CurrentDataBase() *mgo.Database
	// CurrentCollection returns the currently selected collection.
	// CurrentCollection 返回当前选中的集合。
	CurrentCollection() *mgo.Collection
}

// MongoSessionV is the default IMongoSessionV implementation backed by mgo.
// MongoSessionV 是基于 mgo 的默认 IMongoSessionV 实现。
type MongoSessionV struct {
	Key             string
	SessionDialInfo *mgo.DialInfo

	MgoSession *mgo.Session
	MgoDB      *mgo.Database
	MgoC       *mgo.Collection
	DBName     string
	CName      string
}

// Close closes the session and releases associated resources.
// Close 关闭会话并释放相关资源。
func (ses *MongoSessionV) Close() error {
	ses.close()
	return nil
}

func (ses *MongoSessionV) cloneSession(newKey string) IMongoSessionV {
	rs := *ses
	rs.Key = newKey
	dialInfo := *ses.SessionDialInfo
	session := *ses.MgoSession.Clone()
	rs.SessionDialInfo, rs.MgoSession = &dialInfo, &session
	if nil != rs.MgoDB {
		db := *rs.MgoDB
		rs.MgoDB = &db
		rs.MgoDB.Session = rs.MgoSession
	}
	if nil != rs.MgoC {
		c := *rs.MgoC
		rs.MgoC = &c
		rs.MgoC.Database = rs.MgoDB
	}
	return &rs
}

// GetSessionName returns the session name.
// GetSessionName 返回会话名称。
func (ses *MongoSessionV) GetSessionName() string {
	return ses.Key
}

// GetDialInfo returns the dial info used to create the session.
// GetDialInfo 返回创建会话时使用的拨号信息。
func (ses *MongoSessionV) GetDialInfo() *mgo.DialInfo {
	return ses.SessionDialInfo
}

// GetMode returns the current session consistency mode.
// GetMode 返回当前会话一致性模式。
func (ses *MongoSessionV) GetMode() mgo.Mode {
	return ses.MgoSession.Mode()
}

// LocateDB selects a database by name.
// LocateDB 按名称选择数据库。
func (ses *MongoSessionV) LocateDB(dbName string) (db *mgo.Database, change bool) {
	if dbName == ses.DBName && nil != ses.MgoDB {
		return ses.MgoDB, false
	}
	db = ses.MgoSession.DB(dbName)
	ses.MgoDB = db
	ses.DBName = dbName
	return db, true
}

// LocateCollection selects a database and collection by name.
// LocateCollection 按名称选择数据库与集合。
func (ses *MongoSessionV) LocateCollection(dbName string, cName string) (c *mgo.Collection, change bool) {
	db, change := ses.LocateDB(dbName)
	if cName == ses.CName && nil != ses.MgoC && !change {
		return ses.MgoC, false
	}
	c = db.C(cName)
	ses.MgoC = c
	ses.CName = cName
	return c, true
}

// Session returns the underlying mgo.Session.
// Session 返回底层 mgo.Session。
func (ses *MongoSessionV) Session() *mgo.Session {
	return ses.MgoSession
}

// CurrentDataBase returns the currently selected database.
// CurrentDataBase 返回当前选中的数据库。
func (ses *MongoSessionV) CurrentDataBase() *mgo.Database {
	return ses.MgoDB
}

// CurrentCollection returns the currently selected collection.
// CurrentCollection 返回当前选中的集合。
func (ses *MongoSessionV) CurrentCollection() *mgo.Collection {
	return ses.MgoC
}

func (ses *MongoSessionV) close() {
	ses.CName = ""
	ses.DBName = ""
	ses.MgoC = nil
	if nil != ses.MgoDB {
		ses.MgoDB.Logout()
		ses.MgoDB = nil
	} else {
		ses.MgoSession.DB("").Logout()
	}
	if nil != ses.MgoSession {
		ses.MgoSession.Close()
		ses.MgoSession = nil
	}
}
