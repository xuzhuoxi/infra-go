package mgox

import "go.mongodb.org/mongo-driver/v2/mongo/options"

// NewOptions builds ClientOptions from a MongoDB connection URI.
// NewOptions 从 MongoDB 连接 URI 构建 ClientOptions。
func NewOptions(uri string) *options.ClientOptions {
	return options.Client().ApplyURI(uri)
}
