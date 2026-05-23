package encodingx

import "github.com/xuzhuoxi/infra-go/lang"

// IPoolKeyValue is an object pool for IKeyValue instances.
// IPoolKeyValue 是 IKeyValue 实例的对象池。
type IPoolKeyValue interface {
	// Register registers a factory for new IKeyValue instances.
	// Register 注册创建 IKeyValue 的工厂函数。
	Register(newFunc func() IKeyValue)
	// GetInstance obtains a cleared IKeyValue from the pool.
	// GetInstance 从池中获取已清空的 IKeyValue 实例。
	GetInstance() IKeyValue
	// Recycle returns an instance to the pool.
	// Recycle 将实例回收到池中。
	Recycle(instance IKeyValue) bool
}

// IPoolCodingHandler is an object pool for ICodingHandler instances.
// IPoolCodingHandler 是 ICodingHandler 实例的对象池。
type IPoolCodingHandler interface {
	// Register registers a factory for new ICodingHandler instances.
	// Register 注册创建 ICodingHandler 的工厂函数。
	Register(newFunc func() ICodingHandler)
	// GetInstance obtains an ICodingHandler from the pool.
	// GetInstance 从池中获取 ICodingHandler 实例。
	GetInstance() ICodingHandler
	// Recycle returns an instance to the pool.
	// Recycle 将实例回收到池中。
	Recycle(instance ICodingHandler) bool
}

// IPoolBuffEncoder is an object pool for IBuffEncoder instances.
// IPoolBuffEncoder 是 IBuffEncoder 实例的对象池。
type IPoolBuffEncoder interface {
	// Register registers a factory for new IBuffEncoder instances.
	// Register 注册创建 IBuffEncoder 的工厂函数。
	Register(newFunc func() IBuffEncoder)
	// GetInstance obtains a reset IBuffEncoder from the pool.
	// GetInstance 从池中获取已 Reset 的 IBuffEncoder 实例。
	GetInstance() IBuffEncoder
	// Recycle returns an instance to the pool.
	// Recycle 将实例回收到池中。
	Recycle(instance IBuffEncoder) bool
}

// IPoolBuffDecoder is an object pool for IBuffDecoder instances.
// IPoolBuffDecoder 是 IBuffDecoder 实例的对象池。
type IPoolBuffDecoder interface {
	// Register registers a factory for new IBuffDecoder instances.
	// Register 注册创建 IBuffDecoder 的工厂函数。
	Register(newFunc func() IBuffDecoder)
	// GetInstance obtains a reset IBuffDecoder from the pool.
	// GetInstance 从池中获取已 Reset 的 IBuffDecoder 实例。
	GetInstance() IBuffDecoder
	// Recycle returns an instance to the pool.
	// Recycle 将实例回收到池中。
	Recycle(instance IBuffDecoder) bool
}

// IPoolBuffCodecs is an object pool for IBuffCodecs instances.
// IPoolBuffCodecs 是 IBuffCodecs 实例的对象池。
type IPoolBuffCodecs interface {
	// Register registers a factory for new IBuffCodecs instances.
	// Register 注册创建 IBuffCodecs 的工厂函数。
	Register(newFunc func() IBuffCodecs)
	// GetInstance obtains a reset IBuffCodecs from the pool.
	// GetInstance 从池中获取已 Reset 的 IBuffCodecs 实例。
	GetInstance() IBuffCodecs
	// Recycle returns an instance to the pool.
	// Recycle 将实例回收到池中。
	Recycle(instance IBuffCodecs) bool
}

// NewPoolKeyValue creates a new IKeyValue object pool.
// NewPoolKeyValue 创建 IKeyValue 对象池。
func NewPoolKeyValue() IPoolKeyValue {
	return &poolKeyValue{pool: lang.NewObjectPoolSync()}
}

// NewPoolCodingHandler creates a new ICodingHandler object pool.
// NewPoolCodingHandler 创建 ICodingHandler 对象池。
func NewPoolCodingHandler() IPoolCodingHandler {
	return &poolCodingHandler{pool: lang.NewObjectPoolSync()}
}

// NewPoolBuffEncoder creates a new IBuffEncoder object pool.
// NewPoolBuffEncoder 创建 IBuffEncoder 对象池。
func NewPoolBuffEncoder() IPoolBuffEncoder {
	return &poolBuffEncoder{pool: lang.NewObjectPoolSync()}
}

// NewPoolBuffDecoder creates a new IBuffDecoder object pool.
// NewPoolBuffDecoder 创建 IBuffDecoder 对象池。
func NewPoolBuffDecoder() IPoolBuffDecoder {
	return &poolBuffDecoder{pool: lang.NewObjectPoolSync()}
}

// NewPoolBuffCodecs creates a new IBuffCodecs object pool.
// NewPoolBuffCodecs 创建 IBuffCodecs 对象池。
func NewPoolBuffCodecs() IPoolBuffCodecs {
	return &poolBuffCodecs{pool: lang.NewObjectPoolSync()}
}

//------------------------------
type poolKeyValue struct {
	pool lang.IObjectPool
}

func (p poolKeyValue) Register(newFunc func() IKeyValue) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		if nil == instance {
			return false
		}
		if _, ok := instance.(IKeyValue); ok {
			return ok
		}
		return false
	})
}

func (p poolKeyValue) GetInstance() IKeyValue {
	rs := p.pool.GetInstance().(IKeyValue)
	rs.Clear()
	return rs
}

func (p poolKeyValue) Recycle(instance IKeyValue) bool {
	return p.pool.Recycle(instance)
}

//------------------------------

type poolCodingHandler struct {
	pool lang.IObjectPool
}

func (p *poolCodingHandler) Register(newFunc func() ICodingHandler) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		if nil == instance {
			return false
		}
		if _, ok := instance.(ICodingHandler); ok {
			return ok
		}
		return false
	})
}

func (p *poolCodingHandler) GetInstance() ICodingHandler {
	return p.pool.GetInstance().(ICodingHandler)
}

func (p *poolCodingHandler) Recycle(instance ICodingHandler) bool {
	return p.pool.Recycle(instance)
}

//------------------------------

type poolBuffEncoder struct {
	pool lang.IObjectPool
}

func (p *poolBuffEncoder) Register(newFunc func() IBuffEncoder) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		if nil == instance {
			return false
		}
		if _, ok := instance.(IBuffEncoder); ok {
			return ok
		}
		return false
	})
}

func (p *poolBuffEncoder) GetInstance() IBuffEncoder {
	rs := p.pool.GetInstance().(IBuffEncoder)
	rs.Reset()
	return rs
}

func (p *poolBuffEncoder) Recycle(instance IBuffEncoder) bool {
	return p.pool.Recycle(instance)
}

//------------------------------

type poolBuffDecoder struct {
	pool lang.IObjectPool
}

func (p *poolBuffDecoder) Register(newFunc func() IBuffDecoder) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		if nil == instance {
			return false
		}
		if _, ok := instance.(IBuffDecoder); ok {
			return ok
		}
		return false
	})
}

func (p *poolBuffDecoder) GetInstance() IBuffDecoder {
	rs := p.pool.GetInstance().(IBuffDecoder)
	rs.Reset()
	return rs
}

func (p *poolBuffDecoder) Recycle(instance IBuffDecoder) bool {
	return p.pool.Recycle(instance)
}

//------------------------------

type poolBuffCodecs struct {
	pool lang.IObjectPool
}

func (p *poolBuffCodecs) Register(newFunc func() IBuffCodecs) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		if nil == instance {
			return false
		}
		if _, ok := instance.(IBuffCodecs); ok {
			return ok
		}
		return false
	})
}

func (p *poolBuffCodecs) GetInstance() IBuffCodecs {
	rs := p.pool.GetInstance().(IBuffCodecs)
	rs.Reset()
	return rs
}

func (p *poolBuffCodecs) Recycle(instance IBuffCodecs) bool {
	return p.pool.Recycle(instance)
}
