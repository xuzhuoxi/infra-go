package mongox

import (
	"gopkg.in/mgo.v2"
	"io"
)

type IMongoSession interface {
	io.Closer
	LocateDB(dbName string) *mgo.Database
	LocateCollection(dbName string, cName string) *mgo.Collection
	CurrentSession() *mgo.Session
	CurrentDB() *mgo.Database
	CurrentCollection() *mgo.Collection
}

type MongoSession struct {
	Session *mgo.Session
	DB      *mgo.Database
	C       *mgo.Collection
	DBName  string
	CName   string

	CloseFunc func(sess IMongoSession)
}

func (ses *MongoSession) LocateDB(dbName string) *mgo.Database {
	if ses.DBName != dbName || ("" == dbName && nil == ses.DB) {
		db := ses.Session.DB(dbName)
		ses.DB = db
		ses.DBName = dbName
	}
	return ses.DB
}

func (ses *MongoSession) LocateCollection(dbName string, cName string) *mgo.Collection {
	db := ses.LocateDB(dbName)
	if ses.CName != cName || ("" == cName && nil == ses.C) {
		c := db.C(cName)
		ses.C = c
		ses.CName = cName

		c.DropCollection()
	}
	return ses.C
}

func (ses *MongoSession) CurrentSession() *mgo.Session {
	return ses.Session
}

func (ses *MongoSession) CurrentDB() *mgo.Database {
	return ses.DB
}

func (ses *MongoSession) CurrentCollection() *mgo.Collection {
	return ses.C
}

func (ses *MongoSession) Close() error {
	ses.close()
	ses.invokeFunc()
	return nil
}

func (ses *MongoSession) Clone() IMongoSession {
	rs := *ses
	rs.Session = ses.Session.Clone()
	if nil != rs.DB {
		rs.DB.Session = rs.Session
	}
	if nil != rs.C {
		rs.C.Database = rs.DB
	}
	return &rs
}

func (ses *MongoSession) close() {
	ses.CName = ""
	ses.DBName = ""
	ses.C = nil
	if nil != ses.DB {
		ses.DB.Logout()
		ses.DB = nil
	} else {
		ses.Session.DB("").Logout()
	}
	if nil != ses.Session {
		ses.Session.Close()
		ses.Session = nil
	}
}

func (ses *MongoSession) invokeFunc() {
	if nil != ses.CloseFunc {
		ses.CloseFunc(ses)
		ses.CloseFunc = nil
	}
}
