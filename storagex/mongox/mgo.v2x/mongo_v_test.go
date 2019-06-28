package mgo_v2x

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

//Using the driver "gopkg.in/mgo.v2"

func TestMongoDriver(t *testing.T) {
	dirver := NewMongoDriverV(0)
	sess, err := dirver.OpenSession("master", &mgo.DialInfo{Addrs: []string{"192.168.3.105"}}, mgo.Monotonic)
	if nil != err {
		t.Fatal(err)
	}
	sess.LocateCollection("TestDB", "TestC")
	err1 := sess.CurrentCollection().Insert(bson.M{"name": "哈哈"})
	if nil != err1 {
		t.Fatal(err1)
	}
	dirver.Close()
}

func TestMgoV2(t *testing.T) {
	sess, err := mgo.DialWithInfo(&mgo.DialInfo{Addrs: []string{"192.168.3.105"}})
	if nil != err {
		t.Fatal(err)
	}
	sess.SetMode(mgo.Monotonic, true)

	db := sess.DB("TestDB")
	c := db.C("TestC1")
	err1 := c.Insert(bson.M{"name": "ABC1", "id": 11})
	if nil != err1 {
		t.Fatal(err1)
	}
	sess.Close()
}
