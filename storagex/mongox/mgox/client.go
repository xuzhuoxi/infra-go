package mgox

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// NewIMongoClient creates an IMongoClient backed by MongoClient.
// NewIMongoClient 创建由 MongoClient 实现的 IMongoClient。
func NewIMongoClient(clientName string, options *options.ClientOptions, context context.Context) IMongoClient {
	return NewMongoClient(clientName, options, context)
}

// NewMongoClient creates a named MongoClient with the given options and context.
// NewMongoClient 使用指定选项与 context 创建具名 MongoClient。
func NewMongoClient(clientName string, options *options.ClientOptions, context context.Context) *MongoClient {
	return &MongoClient{name: clientName, Options: options, CliContext: context}
}

// IMongoClient manages a single MongoDB client connection and database selection.
// IMongoClient 管理单个 MongoDB 客户端连接及数据库选择。
type IMongoClient interface {
	// Name returns the client name.
	// Name 返回客户端名称。
	Name() string
	// Connect establishes the client connection and loads database names.
	// Connect 建立客户端连接并加载数据库名称列表。
	Connect() error
	// Disconnect closes the client connection.
	// Disconnect 关闭客户端连接。
	Disconnect() error
	// Connecting reports whether the client is connected.
	// Connecting 报告客户端是否已连接。
	Connecting() bool
	// Reconnect disconnects and reconnects the client.
	// Reconnect 断开并重新连接客户端。
	Reconnect() error
	// Client returns the underlying mongo.Client.
	// Client 返回底层 mongo.Client。
	Client() *mongo.Client

	// GetDatabaseNames returns cached database names from the last Connect.
	// GetDatabaseNames 返回上次 Connect 时缓存的数据库名称列表。
	GetDatabaseNames() []string
	// LocateDatabase selects a database by name and optionally overrides its context.
	// LocateDatabase 按名称选择数据库，并可选择覆盖其 context。
	LocateDatabase(dbName string, newContext context.Context) (db IMongoDatabase, change bool)
	// LocatedDatabase returns the currently selected database wrapper.
	// LocatedDatabase 返回当前选中的数据库封装。
	LocatedDatabase() IMongoDatabase
}

// MongoClient wraps mongo.Client with connection state and database tracking.
// MongoClient 封装 mongo.Client，并跟踪连接状态与当前数据库。
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

// Name returns the client name.
// Name 返回客户端名称。
func (c *MongoClient) Name() string {
	return c.name
}

// Connect establishes the client connection and loads database names.
// Connect 建立客户端连接并加载数据库名称列表。
func (c *MongoClient) Connect() error {
	c.ConnectMutex.Lock()
	defer c.ConnectMutex.Unlock()
	err := c.connect()
	if nil != err {
		return err
	}
	names, err := c.Cli.ListDatabaseNames(c.CliContext, bson.D{})
	if nil != err {
		return err
	}
	c.DatabaseNames = names
	return nil
}

// Disconnect closes the client connection.
// Disconnect 关闭客户端连接。
func (c *MongoClient) Disconnect() error {
	c.ConnectMutex.Lock()
	defer c.ConnectMutex.Unlock()
	return c.disconnect()
}

// Connecting reports whether the client is connected.
// Connecting 报告客户端是否已连接。
func (c *MongoClient) Connecting() bool {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return nil != c.Cli
}

// Reconnect disconnects and reconnects the client.
// Reconnect 断开并重新连接客户端。
func (c *MongoClient) Reconnect() error {
	c.ConnectMutex.Lock()
	defer c.ConnectMutex.Unlock()
	c.disconnect()
	return c.connect()
}

// Client returns the underlying mongo.Client.
// Client 返回底层 mongo.Client。
func (c *MongoClient) Client() *mongo.Client {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return c.Cli
}

// GetDatabaseNames returns cached database names from the last Connect.
// GetDatabaseNames 返回上次 Connect 时缓存的数据库名称列表。
func (c *MongoClient) GetDatabaseNames() []string {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return c.DatabaseNames
}

// LocateDatabase selects a database by name and optionally overrides its context.
// LocateDatabase 按名称选择数据库，并可选择覆盖其 context。
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

// LocatedDatabase returns the currently selected database wrapper.
// LocatedDatabase 返回当前选中的数据库封装。
func (c *MongoClient) LocatedDatabase() IMongoDatabase {
	c.ConnectMutex.RLock()
	defer c.ConnectMutex.RUnlock()
	return c.Database
}

//---------------------------------

func (c *MongoClient) connect() error {
	client, err := mongo.Connect(c.Options)
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
