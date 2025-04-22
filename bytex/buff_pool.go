// Package bytex
// Created by xuzhuoxi
// on 2019-03-23.
// @author xuzhuoxi
//
package bytex

import (
	"sync"
)

var (
	DefaultPoolBuffDataBlock = NewPoolBuffDataBlock(NewDefaultBuffDataBlock)
	DefaultPoolBuffToData    = NewPoolBuffToData(NewDefaultBuffToData)
	DefaultPoolBuffToBlock   = NewPoolBuffToBlock(NewDefaultBuffToBlock)
)

// IPoolBuffDataBlock ---------- ---------- ---------- ---------- ----------

func NewPoolBuffDataBlock(newFunc func() IBuffDataBlock) IPoolBuffDataBlock {
	return &poolBuffDataBlock{pool: &sync.Pool{
		New: func() interface{} {
			return newFunc()
		},
	}}
}

type IPoolBuffDataBlock interface {
	GetInstance() IBuffDataBlock
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

// IPoolBuffToData ---------- ---------- ---------- ---------- ----------

func NewPoolBuffToData(newFunc func() IBuffToData) IPoolBuffToData {
	return &poolBuffToData{pool: &sync.Pool{
		New: func() interface{} {
			return newFunc()
		},
	}}
}

type IPoolBuffToData interface {
	GetInstance() IBuffToData
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

// IPoolBuffToBlock ---------- ---------- ---------- ---------- ----------

type IPoolBuffToBlock interface {
	GetInstance() IBuffToBlock
	Recycle(instance IBuffToBlock)
}

func NewPoolBuffToBlock(newFunc func() IBuffToBlock) IPoolBuffToBlock {
	return &poolBuffToBlock{pool: &sync.Pool{
		New: func() interface{} {
			return newFunc()
		},
	}}
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
