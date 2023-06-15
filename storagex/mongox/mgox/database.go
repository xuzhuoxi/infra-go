// Package mgox
// Created by xuzhuoxi
// on 2019-06-16.
// @author xuzhuoxi
//
package mgox

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func NewIMongoDatabase(db *mongo.Database, context context.Context) IMongoDatabase {
	return NewMongoDatabase(db, context)
}

func NewMongoDatabase(db *mongo.Database, context context.Context) *MongoDatabase {
	return &MongoDatabase{DB: db, DBContext: context}
}

type IMongoDatabase interface {
	Connect() error
	Disconnect() error
	DatabaseName() string
	Database() *mongo.Database

	GetCollectionNames() []string
	GetCollection(cName string) (c *mongo.Collection, ok bool)
	LocateCollection(cName string) (c *mongo.Collection, change bool)
	LocatedCollection() *mongo.Collection
}

type MongoDatabase struct {
	DB        *mongo.Database
	DBContext context.Context

	CollectionNames []string
	Collections     map[string]*mongo.Collection

	Collection *mongo.Collection
	CName      string
}

func (db *MongoDatabase) Connect() error {
	names, err := db.DB.ListCollectionNames(db.DBContext, bsonx.Doc{})
	if nil != err {
		return err
	}
	db.CollectionNames = names
	db.Collections = make(map[string]*mongo.Collection, len(names))
	for _, n := range names {
		db.Collections[n] = db.DB.Collection(n)
	}
	return nil
}

func (db *MongoDatabase) Disconnect() error {
	return nil
}

func (db *MongoDatabase) DatabaseName() string {
	return db.DB.Name()
}

func (db *MongoDatabase) Database() *mongo.Database {
	return db.DB
}

func (db *MongoDatabase) GetCollectionNames() []string {
	return db.CollectionNames
}

func (db *MongoDatabase) GetCollection(cName string) (c *mongo.Collection, ok bool) {
	c, ok = db.Collections[cName]
	return
}

func (db *MongoDatabase) LocateCollection(cName string) (c *mongo.Collection, change bool) {
	if cName == db.CName && nil != db.Collection {
		return db.Collection, false
	}
	if col, ok := db.Collections[cName]; ok {
		db.Collection, db.CName = col, cName
		return col, true
	}
	return nil, false
}

func (db *MongoDatabase) LocatedCollection() *mongo.Collection {
	return db.Collection
}
