package bytex

import (
	"sync"
)

var (
	// DefaultPoolBuffDataBlock is the default pool for IBuffDataBlock instances.
	// DefaultPoolBuffDataBlock 是 IBuffDataBlock 实例的默认对象池。
	DefaultPoolBuffDataBlock = NewPoolBuffDataBlock(NewDefaultBuffDataBlock)
	// DefaultPoolBuffToData is the default pool for IBuffToData instances.
	// DefaultPoolBuffToData 是 IBuffToData 实例的默认对象池。
	DefaultPoolBuffToData = NewPoolBuffToData(NewDefaultBuffToData)
	// DefaultPoolBuffToBlock is the default pool for IBuffToBlock instances.
	// DefaultPoolBuffToBlock 是 IBuffToBlock 实例的默认对象池。
	DefaultPoolBuffToBlock = NewPoolBuffToBlock(NewDefaultBuffToBlock)
)

// NewPoolBuffDataBlock creates a sync.Pool for IBuffDataBlock using newFunc.
// NewPoolBuffDataBlock 使用 newFunc 创建 IBuffDataBlock 的对象池。
func NewPoolBuffDataBlock(newFunc func() IBuffDataBlock) IPoolBuffDataBlock {
	return &poolBuffDataBlock{pool: &sync.Pool{
		New: func() interface{} {
			return newFunc()
		},
	}}
}

// IPoolBuffDataBlock pools reusable IBuffDataBlock instances.
// IPoolBuffDataBlock 复用 IBuffDataBlock 实例的对象池。
type IPoolBuffDataBlock interface {
	// GetInstance acquires a buffer from the pool.
	// GetInstance 从池中获取一个缓冲实例。
	GetInstance() IBuffDataBlock
	// Recycle returns instance to the pool after Reset.
	// Recycle 在 Reset 后将实例归还到池中。
	Recycle(instance IBuffDataBlock)
}

type poolBuffDataBlock struct {
	pool *sync.Pool
}

func (p *poolBuffDataBlock) GetInstance() IBuffDataBlock {
	return p.pool.Get().(IBuffDataBlock)
}

func (p *poolBuffDataBlock) Recycle(instance IBuffDataBlock) {
	if nil == instance {
		return
	}
	instance.Reset()
	p.pool.Put(instance)
}

// NewPoolBuffToData creates a sync.Pool for IBuffToData using newFunc.
// NewPoolBuffToData 使用 newFunc 创建 IBuffToData 的对象池。
func NewPoolBuffToData(newFunc func() IBuffToData) IPoolBuffToData {
	return &poolBuffToData{pool: &sync.Pool{
		New: func() interface{} {
			return newFunc()
		},
	}}
}

// IPoolBuffToData pools reusable IBuffToData instances.
// IPoolBuffToData 复用 IBuffToData 实例的对象池。
type IPoolBuffToData interface {
	// GetInstance acquires a buffer from the pool.
	// GetInstance 从池中获取一个缓冲实例。
	GetInstance() IBuffToData
	// Recycle returns instance to the pool after Reset.
	// Recycle 在 Reset 后将实例归还到池中。
	Recycle(instance IBuffToData)
}

type poolBuffToData struct {
	pool *sync.Pool
}

func (p *poolBuffToData) GetInstance() IBuffToData {
	return p.pool.Get().(IBuffToData)
}

func (p *poolBuffToData) Recycle(instance IBuffToData) {
	if nil == instance {
		return
	}
	instance.Reset()
	p.pool.Put(instance)
}

// NewPoolBuffToBlock creates a sync.Pool for IBuffToBlock using newFunc.
// NewPoolBuffToBlock 使用 newFunc 创建 IBuffToBlock 的对象池。
func NewPoolBuffToBlock(newFunc func() IBuffToBlock) IPoolBuffToBlock {
	return &poolBuffToBlock{pool: &sync.Pool{
		New: func() interface{} {
			return newFunc()
		},
	}}
}

// IPoolBuffToBlock pools reusable IBuffToBlock instances.
// IPoolBuffToBlock 复用 IBuffToBlock 实例的对象池。
type IPoolBuffToBlock interface {
	// GetInstance acquires a buffer from the pool.
	// GetInstance 从池中获取一个缓冲实例。
	GetInstance() IBuffToBlock
	// Recycle returns instance to the pool after Reset.
	// Recycle 在 Reset 后将实例归还到池中。
	Recycle(instance IBuffToBlock)
}

type poolBuffToBlock struct {
	pool *sync.Pool
}

func (p *poolBuffToBlock) GetInstance() IBuffToBlock {
	return p.pool.Get().(IBuffToBlock)
}

func (p *poolBuffToBlock) Recycle(instance IBuffToBlock) {
	if nil == instance {
		return
	}
	instance.Reset()
	p.pool.Put(instance)
}
