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

// NewByteSlicePool creates or returns a shared byte slice pool for the given capacity.
// NewByteSlicePool 创建或返回指定容量的共享字节切片池。
//
// cap uses defaultBytePoolCap (2048) when cap is zero or negative.
// cap 为 0 或非正数时使用默认值 2048。
//
// Pools with the same cap share one instance; repeated calls return the same pool.
// 相同 cap 只维护一个池实例，重复创建返回同一对象。
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

// IByteSlicePool is a pool of reusable *[]byte backed by sync.Pool.
// IByteSlicePool 基于 sync.Pool 的可复用 *[]byte 字节切片池。
type IByteSlicePool interface {
	// Cap returns the configured slice capacity of the pool.
	// Cap 返回池配置的切片容量。
	Cap() int
	// Get acquires a *[]byte from the pool with length 0.
	// Get 从池中获取一个 len 为 0 的 *[]byte。
	Get() *[]byte
	// GetL acquires a *[]byte from the pool with length ln.
	// GetL 从池中获取长度为 ln 的 *[]byte。
	//
	// Returns an error when ln exceeds Cap.
	// ln 大于 Cap 时返回错误。
	GetL(ln int) (*[]byte, error)
	// Put returns a *[]byte to the pool.
	// Put 将 *[]byte 归还到池中。
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
