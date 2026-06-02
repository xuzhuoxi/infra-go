package mgox

import (
	"context"
)

var (
	// DefaultContext is the default context for MongoDB operations.
	// DefaultContext 是 MongoDB 操作的默认 context。
	DefaultContext = context.Background()
	// ClientContext is the default context for client operations.
	// ClientContext 是客户端操作的默认 context。
	ClientContext = context.Background()
	// DBReaderContext is the default context for database read operations.
	// DBReaderContext 是数据库读操作的默认 context。
	DBReaderContext = context.Background()
	// DBWriterContext is the default context for database write operations.
	// DBWriterContext 是数据库写操作的默认 context。
	DBWriterContext = context.Background()
	// CollectionReaderContext is the default context for collection read operations.
	// CollectionReaderContext 是集合读操作的默认 context。
	CollectionReaderContext = context.Background()
	// CollectionWriterContext is the default context for collection write operations.
	// CollectionWriterContext 是集合写操作的默认 context。
	CollectionWriterContext = context.Background()
)
