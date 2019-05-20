//
//Created by xuzhuoxi
//on 2019-05-19.
//@author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/lang"
)

var (
	DefaultRequestPool  = NewPoolExtensionRequest()
	DefaultResponsePool = NewPoolExtensionResponse()
)

// 请求参数集的对象池接口
type IPoolExtensionRequest interface {
	// 注册创建方法
	Register(newFunc func() IExtensionRequest)
	// 获取一个实例
	GetInstance() IExtensionRequest
	// 回收一个实例
	Recycle(instance IExtensionRequest) bool
}

//响应参数集的对象池接口
type IPoolExtensionResponse interface {
	// 注册创建方法
	Register(newFunc func() IExtensionResponse)
	// 获取一个实例
	GetInstance() IExtensionResponse
	// 回收一个实例
	Recycle(instance IExtensionResponse) bool
}

func init() {
	DefaultRequestPool.Register(func() IExtensionRequest {
		return NewSockRequest()
	})
	DefaultResponsePool.Register(func() IExtensionResponse {
		return NewSockResponse()
	})
}

//--------------------------------------------

func NewPoolExtensionRequest() IPoolExtensionRequest {
	return &reqPool{pool: lang.NewObjectPoolSync()}
}

func NewPoolExtensionResponse() IPoolExtensionResponse {
	return &respPool{pool: lang.NewObjectPoolSync()}
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
