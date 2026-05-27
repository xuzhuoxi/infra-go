package lang

import "sync"

// PoolCheckFunc validates whether an instance can be recycled into the pool.
// PoolCheckFunc 校验实例是否可回收进对象池。
type PoolCheckFunc func(instance interface{}) bool

// IPoolInstanceChecker configures the recycle validation function.
// IPoolInstanceChecker 配置回收校验函数。
type IPoolInstanceChecker interface {
	// SetCheckFunc sets the function used to validate recycled instances.
	// SetCheckFunc 设置回收实例时的校验函数。
	SetCheckFunc(check PoolCheckFunc)
}

// IObjectPool manages reusable object instances.
// IObjectPool 管理可复用对象实例。
type IObjectPool interface {
	// Register sets the factory and optional recycle check function.
	// Register 设置工厂函数与可选的回收校验函数。
	Register(newFunc func() interface{}, check PoolCheckFunc)
	// GetInstance returns a pooled instance or creates a new one.
	// GetInstance 从池中取出实例，池空时创建新实例。
	GetInstance() interface{}
	// Recycle returns an instance to the pool when check passes.
	// Recycle 在校验通过时将实例放回池中。
	Recycle(instance interface{}) bool
}

// ISizeObjectPool is an object pool with a maximum outstanding instance count.
// ISizeObjectPool 是带最大借出数量的对象池。
type ISizeObjectPool interface {
	// SetMaxSize sets the maximum number of instances that may be checked out.
	// SetMaxSize 设置允许同时借出的最大实例数。
	SetMaxSize(size int)
	IObjectPool
}

// NewObjectPool returns an async or mutex-protected object pool.
// NewObjectPool 返回异步或带互斥锁保护的对象池。
func NewObjectPool(sync bool) IObjectPool {
	if sync {
		return NewObjectPoolSync()
	} else {
		return NewObjectPoolAsync()
	}
}

// NewSizeObjectPool returns a size-limited async or sync object pool.
// NewSizeObjectPool 返回带容量上限的异步或同步对象池。
func NewSizeObjectPool(size int, sync bool) IObjectPool {
	if sync {
		return NewSizeObjectPoolSync(size)
	} else {
		return NewSizeObjectPoolAsync(size)
	}
}

// NewObjectPoolAsync returns a non-mutex object pool.
// NewObjectPoolAsync 返回无互斥锁的对象池。
func NewObjectPoolAsync() IObjectPool {
	return &objectPool{}
}

// NewObjectPoolSync returns a mutex-protected object pool.
// NewObjectPoolSync 返回带互斥锁的对象池。
func NewObjectPoolSync() IObjectPool {
	return &objectPoolSync{pool: &objectPool{}}
}

// NewSizeObjectPoolAsync returns an async pool with a maximum checkout size.
// NewSizeObjectPoolAsync 返回异步的容量受限对象池。
func NewSizeObjectPoolAsync(size int) ISizeObjectPool {
	return &sizeObjectPool{objPool: NewObjectPoolAsync()}
}

// NewSizeObjectPoolSync returns a sync pool with a maximum checkout size.
// NewSizeObjectPoolSync 返回同步的容量受限对象池。
func NewSizeObjectPoolSync(size int) ISizeObjectPool {
	return &sizeObjectPool{objPool: NewObjectPoolSync()}
}

//--------------------------------

type sizeObjectPool struct {
	objPool IObjectPool
	size    chan struct{}
}

func (p *sizeObjectPool) SetMaxSize(size int) {
	if size <= 0 || nil != p.size {
		return
	}
	p.size = make(chan struct{}, size)
}

func (p *sizeObjectPool) Register(newFunc func() interface{}, check PoolCheckFunc) {
	p.objPool.Register(newFunc, check)
}

func (p *sizeObjectPool) GetInstance() interface{} {
	p.size <- struct{}{}
	return p.objPool.GetInstance()
}

func (p *sizeObjectPool) Recycle(instance interface{}) bool {
	if p.objPool.Recycle(instance) {
		defer func() {
			<-p.size
		}()
		return true
	}
	return false
}

//-------------------------
type objectPoolSync struct {
	pool *objectPool
	lock sync.Mutex
}

func (p *objectPoolSync) Register(newFunc func() interface{}, check PoolCheckFunc) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.pool.Register(newFunc, check)
}

func (p *objectPoolSync) GetInstance() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.pool.GetInstance()
}

func (p *objectPoolSync) Recycle(instance interface{}) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.pool.Recycle(instance)
}

//-------------------------

type objectPool struct {
	newFunc func() interface{}
	check   PoolCheckFunc
	objs    []interface{}
}

func (p *objectPool) Register(newFunc func() interface{}, check PoolCheckFunc) {
	p.newFunc = newFunc
	p.check = check
	list := make([]interface{}, 32)
	p.objs = list[0:0]
}

func (p *objectPool) GetInstance() interface{} {
	ln := len(p.objs)
	if ln > 0 {
		rs := p.objs[ln-1]
		p.objs = p.objs[:ln-1]
		return rs
	} else {
		return p.newFunc()
	}
}

func (p *objectPool) Recycle(instance interface{}) bool {
	if nil != instance && p.check(instance) {
		p.objs = append(p.objs, instance)
		return true
	}
	return false
}
