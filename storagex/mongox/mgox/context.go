package mgox

import "context"

// IContextSetter sets the context used by a MongoDB wrapper.
// IContextSetter 设置 MongoDB 封装层使用的 context。
type IContextSetter interface {
	// SetContext assigns the operation context.
	// SetContext 设置操作 context。
	SetContext(context context.Context) error
}

// IContextGetter returns the context used by a MongoDB wrapper.
// IContextGetter 返回 MongoDB 封装层使用的 context。
type IContextGetter interface {
	// GetContext returns the current operation context.
	// GetContext 返回当前操作 context。
	GetContext() context.Context
}
