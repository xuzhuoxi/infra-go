package mgox

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// NewIMongoDatabase creates an IMongoDatabase backed by MongoDatabase.
// NewIMongoDatabase 创建由 MongoDatabase 实现的 IMongoDatabase。
func NewIMongoDatabase(db *mongo.Database, context context.Context) IMongoDatabase {
	return NewMongoDatabase(db, context)
}

// NewMongoDatabase wraps a mongo.Database with collection tracking.
// NewMongoDatabase 封装 mongo.Database 并跟踪集合选择状态。
func NewMongoDatabase(db *mongo.Database, context context.Context) *MongoDatabase {
	return &MongoDatabase{DB: db, DBContext: context}
}

// IMongoDatabase manages collection listing and selection on a MongoDB database.
// IMongoDatabase 管理 MongoDB 数据库上的集合列表与选择。
type IMongoDatabase interface {
	// Connect loads collection names and builds the collection map.
	// Connect 加载集合名称并构建集合映射表。
	Connect() error
	// Disconnect releases database resources (no-op for the official driver).
	// Disconnect 释放数据库资源（官方驱动下为空操作）。
	Disconnect() error
	// DatabaseName returns the database name.
	// DatabaseName 返回数据库名称。
	DatabaseName() string
	// Database returns the underlying mongo.Database.
	// Database 返回底层 mongo.Database。
	Database() *mongo.Database

	// GetCollectionNames returns cached collection names from the last Connect.
	// GetCollectionNames 返回上次 Connect 时缓存的集合名称列表。
	GetCollectionNames() []string
	// GetCollection returns a collection by name from the cached map.
	// GetCollection 从缓存映射中按名称返回集合。
	GetCollection(cName string) (c *mongo.Collection, ok bool)
	// LocateCollection selects a collection by name as the current collection.
	// LocateCollection 按名称选择集合并设为当前集合。
	LocateCollection(cName string) (c *mongo.Collection, change bool)
	// LocatedCollection returns the currently selected collection.
	// LocatedCollection 返回当前选中的集合。
	LocatedCollection() *mongo.Collection
}

// MongoDatabase wraps mongo.Database with collection name and selection state.
// MongoDatabase 封装 mongo.Database，并跟踪集合名称与当前选中集合。
type MongoDatabase struct {
	DB        *mongo.Database
	DBContext context.Context

	CollectionNames []string
	Collections     map[string]*mongo.Collection

	Collection *mongo.Collection
	CName      string
}

// Connect loads collection names and builds the collection map.
// Connect 加载集合名称并构建集合映射表。
func (db *MongoDatabase) Connect() error {
	names, err := db.DB.ListCollectionNames(db.DBContext, bson.D{})
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

// Disconnect releases database resources (no-op for the official driver).
// Disconnect 释放数据库资源（官方驱动下为空操作）。
func (db *MongoDatabase) Disconnect() error {
	return nil
}

// DatabaseName returns the database name.
// DatabaseName 返回数据库名称。
func (db *MongoDatabase) DatabaseName() string {
	return db.DB.Name()
}

// Database returns the underlying mongo.Database.
// Database 返回底层 mongo.Database。
func (db *MongoDatabase) Database() *mongo.Database {
	return db.DB
}

// GetCollectionNames returns cached collection names from the last Connect.
// GetCollectionNames 返回上次 Connect 时缓存的集合名称列表。
func (db *MongoDatabase) GetCollectionNames() []string {
	return db.CollectionNames
}

// GetCollection returns a collection by name from the cached map.
// GetCollection 从缓存映射中按名称返回集合。
func (db *MongoDatabase) GetCollection(cName string) (c *mongo.Collection, ok bool) {
	c, ok = db.Collections[cName]
	return
}

// LocateCollection selects a collection by name as the current collection.
// LocateCollection 按名称选择集合并设为当前集合。
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

// LocatedCollection returns the currently selected collection.
// LocatedCollection 返回当前选中的集合。
func (db *MongoDatabase) LocatedCollection() *mongo.Collection {
	return db.Collection
}
