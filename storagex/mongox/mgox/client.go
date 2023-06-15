// Package mgox
// Created by xuzhuoxi
// on 2019-06-15.
// @author xuzhuoxi
//
package mgox

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"sync"
)

func NewIMongoClient(clientName string, options *options.ClientOptions, context context.Context) IMongoClient {
	return NewMongoClient(clientName, options, context)
}

func NewMongoClient(clientName string, options *options.ClientOptions, context context.Context) *MongoClient {
	return &MongoClient{name: clientName, Options: options, CliContext: context}
}

type IMongoClient interface {
	Name() string
	Connect() error
	Disconnect() error
	Connecting() bool
	Reconnect() error
	Client() *mongo.Client

	GetDatabaseNames() []string
	LocateDatabase(dbName string, newContext context.Context) (db IMongoDatabase, change bool)
	LocatedDatabase() IMongoDatabase
}

type MongoClient struct {
	name         string
	Options      *options.ClientOptions
	CliContext   context.Context
	Cli          *mongo.Client
	ConnectMutex sync.RWMutex

	DatabaseNames []string

	Database IMongoDatabase
	DBName   string
}

func (c *MongoClient) Name() string {
	return c.name
}

func (c *MongoClient) Connect() error {
	c.ConnectMutex.Lock()
	defer c.ConnectMutex.Unlock()
	err := c.connect()
	if nil != err {
		return err
	}
	names, err := c.Cli.ListDatabaseNames(c.CliContext, bsonx.Doc{})
	if nil != err {
		return err
	}
	c.DatabaseNames = names
	return c.connect()
}

func (c *MongoClient) Disconnect() error {
	c.ConnectMutex.Lock()
	defer c.ConnectMutex.Unlock()
	return c.disconnect()
}

func (c *MongoClient) Connecting() bool {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return nil != c.Cli
}

func (c *MongoClient) Reconnect() error {
	c.ConnectMutex.Lock()
	defer c.ConnectMutex.Unlock()
	c.disconnect()
	return c.connect()
}

func (c *MongoClient) Client() *mongo.Client {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return c.Cli
}

func (c *MongoClient) GetDatabaseNames() []string {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return c.DatabaseNames
}

func (c *MongoClient) LocateDatabase(dbName string, newContext context.Context) (db IMongoDatabase, change bool) {
	c.ConnectMutex.Lock()
	defer c.ConnectMutex.Unlock()
	if dbName == c.DBName && nil != c.Database {
		return c.Database, false
	}
	newDB := c.Cli.Database(dbName)
	if nil == newContext {
		newContext = c.CliContext
	}
	db = NewMongoDatabase(newDB, newContext)
	err := db.Connect()
	if nil != err {
		return nil, false
	}
	return db, true
}

func (c *MongoClient) LocatedDatabase() IMongoDatabase {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return c.Database
}

//---------------------------------

func (c *MongoClient) connect() error {
	client, err := mongo.NewClient(c.Options)
	if nil != err {
		return err
	}
	err = client.Connect(c.CliContext)
	if nil != err {
		return err
	}
	c.Cli = client
	return nil
}

func (c *MongoClient) disconnect() error {
	if nil != c.Cli {
		err := c.Cli.Disconnect(c.CliContext)
		c.Cli = nil
		if nil != err {
			return err
		}
	}
	return nil
}
