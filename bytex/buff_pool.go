// Package bytex
// Created by xuzhuoxi
// on 2019-03-23.
// @author xuzhuoxi
//
package bytex

import (
	"github.com/xuzhuoxi/infra-go/lang"
)

var (
	DefaultPoolBuffDataBlock = NewPoolBuffDataBlock()
	DefaultPoolBuffToData    = NewPoolBuffToData()
	DefaultPoolBuffToBlock   = NewPoolBuffToBlock()
)

func init() {
	DefaultPoolBuffDataBlock.Register(NewDefaultBuffDataBlock)
	DefaultPoolBuffToData.Register(NewDefaultBuffToData)
	DefaultPoolBuffToBlock.Register(NewDefaultBuffToBlock)
}

type IPoolBuffDataBlock interface {
	Register(newFunc func() IBuffDataBlock)
	GetInstance() IBuffDataBlock
	Recycle(instance IBuffDataBlock) bool
}

type IPoolBuffToData interface {
	Register(newFunc func() IBuffToData)
	GetInstance() IBuffToData
	Recycle(instance IBuffToData) bool
}

type IPoolBuffToBlock interface {
	Register(newFunc func() IBuffToBlock)
	GetInstance() IBuffToBlock
	Recycle(instance IBuffToBlock) bool
}

func NewPoolBuffDataBlock() IPoolBuffDataBlock {
	return &poolBuffDataBlock{pool: lang.NewObjectPoolSync()}
}

func NewPoolBuffToData() IPoolBuffToData {
	return &poolBuffToData{pool: lang.NewObjectPoolSync()}
}

func NewPoolBuffToBlock() IPoolBuffToBlock {
	return &poolBuffToBlock{pool: lang.NewObjectPoolSync()}
}

//--------------------------------------------

type poolBuffDataBlock struct {
	pool lang.IObjectPool
}

func (p *poolBuffDataBlock) Register(newFunc func() IBuffDataBlock) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		return nil != instance
	})
}

func (p *poolBuffDataBlock) GetInstance() IBuffDataBlock {
	rs := p.pool.GetInstance().(IBuffDataBlock)
	rs.Reset()
	return rs
}

func (p *poolBuffDataBlock) Recycle(instance IBuffDataBlock) bool {
	return p.pool.Recycle(instance)
}

type poolBuffToData struct {
	pool lang.IObjectPool
}

func (p *poolBuffToData) Register(newFunc func() IBuffToData) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		return nil != instance
	})
}

func (p *poolBuffToData) GetInstance() IBuffToData {
	rs := p.pool.GetInstance().(IBuffToData)
	rs.Reset()
	return rs
}

func (p *poolBuffToData) Recycle(instance IBuffToData) bool {
	return p.pool.Recycle(instance)
}

type poolBuffToBlock struct {
	pool lang.IObjectPool
}

func (p *poolBuffToBlock) Register(newFunc func() IBuffToBlock) {
	p.pool.Register(func() interface{} {
		return newFunc()
	}, func(instance interface{}) bool {
		return nil != instance
	})
}

func (p *poolBuffToBlock) GetInstance() IBuffToBlock {
	rs := p.pool.GetInstance().(IBuffToBlock)
	rs.Reset()
	return rs
}

func (p *poolBuffToBlock) Recycle(instance IBuffToBlock) bool {
	return p.pool.Recycle(instance)
}
