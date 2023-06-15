// Package encodingx
// Created by xuzhuoxi
// on 2019-03-25.
// @author xuzhuoxi
//
package encodingx

import "github.com/xuzhuoxi/infra-go/lang"

type IPoolCodingHandler interface {
	Register(newFunc func() ICodingHandler)
	GetInstance() ICodingHandler
	Recycle(instance ICodingHandler) bool
}

type IPoolBuffEncoder interface {
	Register(newFunc func() IBuffEncoder)
	GetInstance() IBuffEncoder
	Recycle(instance IBuffEncoder) bool
}

type IPoolBuffDecoder interface {
	Register(newFunc func() IBuffDecoder)
	GetInstance() IBuffDecoder
	Recycle(instance IBuffDecoder) bool
}

type IPoolBuffCodecs interface {
	Register(newFunc func() IBuffCodecs)
	GetInstance() IBuffCodecs
	Recycle(instance IBuffCodecs) bool
}

func NewPoolCodingHandler() IPoolCodingHandler {
	return &poolCodingHandler{pool: lang.NewObjectPoolSync()}
}

func NewPoolBuffEncoder() IPoolBuffEncoder {
	return &poolBuffEncoder{pool: lang.NewObjectPoolSync()}
}

func NewPoolBuffDecoder() IPoolBuffDecoder {
	return &poolBuffDecoder{pool: lang.NewObjectPoolSync()}
}

func NewPoolBuffCodecs() IPoolBuffCodecs {
	return &poolBuffCodecs{pool: lang.NewObjectPoolSync()}
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
