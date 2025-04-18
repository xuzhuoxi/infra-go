// Package poolx
// Create on 2025/4/19
// @author xuzhuoxi
package poolx

import "sync"

var (
	defaultBytePoolCap = 2048
	bytePoolLock       sync.RWMutex
	bytePools          = make(map[int]IByteSlicePool)
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
	// Get 获取一个切片引用
	Get() *[]byte
	// Put 归还一个切片引用
	Put(bs *[]byte)
}

type byteSlicePool struct {
	cap  int
	pool *sync.Pool
}

func (o *byteSlicePool) Get() *[]byte {
	return o.pool.Get().(*[]byte)
}

func (o *byteSlicePool) Put(bs *[]byte) {
	o.pool.Put(bs)
}
