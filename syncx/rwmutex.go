// Package syncx
// Create on 2023/9/20
// @author xuzhuoxi
package syncx

import "sync"

type RWMutex struct {
	lock sync.RWMutex
}

func (rw *RWMutex) Lock() {
	rw.lock.Lock()
}

func (rw *RWMutex) Unlock() {
	rw.lock.Unlock()
}

func (rw *RWMutex) RLock() {
	rw.lock.RLock()
}

func (rw *RWMutex) RUnlock() {
	rw.lock.RUnlock()
}

func (rw *RWMutex) RLocker() {
	rw.lock.RLocker()
}
