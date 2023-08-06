// Package protox
// Created by xuzhuoxi
// on 2019-05-19.
// @author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/lang"
)

var (
	DefaultRequestPool  = NewPoolExtensionRequest()
	DefaultResponsePool = NewPoolExtensionResponse()
	DefaultNotifyPool   = NewPoolExtensionNotify()
)

// IPoolExtensionRequest
// 请求参数集的对象池接口
type IPoolExtensionRequest interface {
	// Register 注册创建方法
	Register(newFunc func() IExtensionRequest)
	// GetInstance 获取一个实例
	GetInstance() IExtensionRequest
	// Recycle 回收一个实例
	Recycle(instance IExtensionRequest) bool
}

// IPoolExtensionResponse
// 响应参数集的对象池接口
type IPoolExtensionResponse interface {
	// Register 注册创建方法
	Register(newFunc func() IExtensionResponse)
	// GetInstance 获取一个实例
	GetInstance() IExtensionResponse
	// Recycle 回收一个实例
	Recycle(instance IExtensionResponse) bool
}

// IPoolExtensionNotify
// 通知参数集的对象池接口
type IPoolExtensionNotify interface {
	// Register 注册创建方法
	Register(newFunc func() IExtensionNotify)
	// GetInstance 获取一个实例
	GetInstance() IExtensionNotify
	// Recycle 回收一个实例
	Recycle(instance IExtensionNotify) bool
}

func init() {
	DefaultRequestPool.Register(func() IExtensionRequest {
		return NewSockRequest()
	})
	DefaultResponsePool.Register(func() IExtensionResponse {
		return NewSockResponse()
	})
	DefaultNotifyPool.Register(func() IExtensionNotify {
		return NewSockNotify()
	})
}

//--------------------------------------------

func NewPoolExtensionRequest() IPoolExtensionRequest {
	return &reqPool{pool: lang.NewObjectPoolSync()}
}

func NewPoolExtensionResponse() IPoolExtensionResponse {
	return &respPool{pool: lang.NewObjectPoolSync()}
}

func NewPoolExtensionNotify() IPoolExtensionNotify {
	return &notifyPool{pool: lang.NewObjectPoolSync()}
}

type reqPool struct {
	pool lang.IObjectPool
}

func (p *reqPool) Register(newFunc func() IExtensionRequest) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		return nil != instance
	})
}

func (p *reqPool) GetInstance() IExtensionRequest {
	return p.pool.GetInstance().(IExtensionRequest)
}

func (p *reqPool) Recycle(instance IExtensionRequest) bool {
	return p.pool.Recycle(instance)
}

type respPool struct {
	pool lang.IObjectPool
}

func (p *respPool) Register(newFunc func() IExtensionResponse) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		return nil != instance
	})
}

func (p *respPool) GetInstance() IExtensionResponse {
	return p.pool.GetInstance().(IExtensionResponse)
}

func (p *respPool) Recycle(instance IExtensionResponse) bool {
	return p.pool.Recycle(instance)
}

type notifyPool struct {
	pool lang.IObjectPool
}

func (p *notifyPool) Register(newFunc func() IExtensionNotify) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		return nil != instance
	})
}

func (p *notifyPool) GetInstance() IExtensionNotify {
	return p.pool.GetInstance().(IExtensionNotify)
}

func (p *notifyPool) Recycle(instance IExtensionNotify) bool {
	return p.pool.Recycle(instance)
}
