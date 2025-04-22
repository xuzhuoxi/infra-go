// Package poolx
// Create on 2025/4/19
// @author xuzhuoxi
package poolx

import (
	"errors"
	"sync"
)

var (
	defaultBytePoolCap = 2048
	bytePoolLock       sync.RWMutex
	bytePools          = make(map[int]IByteSlicePool)
	errLenTooLarge     = errors.New("len too large")
)

// NewByteSlicePool 创建一个字节切片池
// 如果cap为0，则使用默认值2048
// 如果重复创建相同cap的池，只返回同一个
//
func NewByteSlicePool(cap int) IByteSlicePool {
	if cap <= 0 {
		cap = defaultBytePoolCap
	}
	bytePoolLock.Lock()
	defer bytePoolLock.Unlock()
	if pool, ok := bytePools[cap]; ok {
		return pool
	}
	pool := &byteSlicePool{
		cap: cap,
		pool: &sync.Pool{
			New: func() interface{} {
				buff := make([]byte, 0, cap)
				return &buff
			}},
	}
	bytePools[cap] = pool
	return pool
}

// IByteSlicePool 字节切片池
type IByteSlicePool interface {
	// Cap
	// 切片实例的设置容量
	Cap() int
	// Get
	// 获取一个切片引用
	// len=0
	Get() *[]byte
	// GetL
	// 获取一个切片引用
	// 切片的len值为l
	// 如果 ln>cap值，返回一个错误
	GetL(ln int) (*[]byte, error)
	// Put
	// 归还一个切片引用
	Put(bs *[]byte)
}

type byteSlicePool struct {
	cap  int
	pool *sync.Pool
}

func (o *byteSlicePool) Cap() int {
	return o.cap
}

func (o *byteSlicePool) GetL(ln int) (*[]byte, error) {
	if ln > o.cap {
		return nil, errLenTooLarge
	}
	rs := o.pool.Get().(*[]byte)
	*rs = (*rs)[:ln]
	return rs, nil
}

func (o *byteSlicePool) Get() *[]byte {
	return o.pool.Get().(*[]byte)
}

func (o *byteSlicePool) Put(bs *[]byte) {
	if nil == bs {
		return
	}
	o.pool.Put(bs)
}
