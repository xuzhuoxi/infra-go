// Package lang
// Created by xuzhuoxi
// on 2019-03-23.
// @author xuzhuoxi
//
package lang

import "sync"

type PoolCheckFunc func(instance interface{}) bool

type IPoolInstanceChecker interface {
	SetCheckFunc(check PoolCheckFunc)
}

type IObjectPool interface {
	Register(newFunc func() interface{}, check PoolCheckFunc)
	GetInstance() interface{}
	Recycle(instance interface{}) bool
}

type ISizeObjectPool interface {
	SetMaxSize(size int)
	IObjectPool
}

func NewObjectPool(sync bool) IObjectPool {
	if sync {
		return NewObjectPoolSync()
	} else {
		return NewObjectPoolAsync()
	}
}

func NewSizeObjectPool(size int, sync bool) IObjectPool {
	if sync {
		return NewSizeObjectPoolSync(size)
	} else {
		return NewSizeObjectPoolAsync(size)
	}
}

func NewObjectPoolAsync() IObjectPool {
	return &objectPool{}
}

func NewObjectPoolSync() IObjectPool {
	return &objectPoolSync{pool: &objectPool{}}
}

func NewSizeObjectPoolAsync(size int) ISizeObjectPool {
	return &sizeObjectPool{objPool: NewObjectPoolAsync()}
}

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
