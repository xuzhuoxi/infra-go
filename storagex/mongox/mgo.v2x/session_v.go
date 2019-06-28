package mgo_v2x

import (
	"gopkg.in/mgo.v2"
	"io"
)

type IMongoSessionV interface {
	io.Closer
	cloneSession(newName string) IMongoSessionV

	GetSessionName() string
	GetDialInfo() *mgo.DialInfo
	GetMode() mgo.Mode

	LocateDB(dbName string) (db *mgo.Database, change bool)
	LocateCollection(dbName string, cName string) (c *mgo.Collection, change bool)

	Session() *mgo.Session
	CurrentDataBase() *mgo.Database
	CurrentCollection() *mgo.Collection
}

type MongoSessionV struct {
	Key             string
	SessionDialInfo *mgo.DialInfo

	MgoSession *mgo.Session
	MgoDB      *mgo.Database
	MgoC       *mgo.Collection
	DBName     string
	CName      string
}

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

func (ses *MongoSessionV) GetSessionName() string {
	return ses.Key
}

func (ses *MongoSessionV) GetDialInfo() *mgo.DialInfo {
	return ses.SessionDialInfo
}

func (ses *MongoSessionV) GetMode() mgo.Mode {
	return ses.MgoSession.Mode()
}

func (ses *MongoSessionV) LocateDB(dbName string) (db *mgo.Database, change bool) {
	if dbName == ses.DBName && nil != ses.MgoDB {
		return ses.MgoDB, false
	}
	db = ses.MgoSession.DB(dbName)
	ses.MgoDB = db
	ses.DBName = dbName
	return db, true
}

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

func (ses *MongoSessionV) Session() *mgo.Session {
	return ses.MgoSession
}

func (ses *MongoSessionV) CurrentDataBase() *mgo.Database {
	return ses.MgoDB
}

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
