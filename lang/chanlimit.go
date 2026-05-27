package lang

import "sync"

// LChannelLimitMaxSetter configures the maximum concurrent acquisitions.
// LChannelLimitMaxSetter 配置最大并发占用数量。
type LChannelLimitMaxSetter interface {
	// SetMax sets the maximum limit; only effective before StartLimit.
	// SetMax 设置上限；仅在 StartLimit 之前生效。
	SetMax(max int)
}

// IChannelLimitSwitch starts or stops limiting.
// IChannelLimitSwitch 启动或停止并发限制。
type IChannelLimitSwitch interface {
	// StartLimit enables limiting and returns whether it was started.
	// StartLimit 启动限制，返回是否成功启动。
	StartLimit() bool
	// StopLimit disables limiting and returns whether it was stopped.
	// StopLimit 停止限制，返回是否成功停止。
	StopLimit() bool
}

// IChannelLimitHandler acquires and releases limit slots.
// IChannelLimitHandler 获取与释放限制槽位。
type IChannelLimitHandler interface {
	// Add acquires one slot (blocks when the limit is active).
	// Add 占用一个槽位（限制生效时会阻塞）。
	Add()
	// Done releases one acquired slot.
	// Done 释放一个已占用的槽位。
	Done()
}

// IChannelLimit combines max configuration, switch control, and acquire/release.
// IChannelLimit 组合上限配置、开关控制与占用/释放能力。
type IChannelLimit interface {
	LChannelLimitMaxSetter
	IChannelLimitSwitch
	IChannelLimitHandler
}

// ChannelLimitNone is a no-op implementation of IChannelLimit.
// ChannelLimitNone 是 IChannelLimit 的空实现。
type ChannelLimitNone struct{}

// SetMax is a no-op on ChannelLimitNone.
// SetMax 在 ChannelLimitNone 上为空操作。
func (l *ChannelLimitNone) SetMax(max int) {
	return
}

// StartLimit always returns false on ChannelLimitNone.
// StartLimit 在 ChannelLimitNone 上恒返回 false。
func (l *ChannelLimitNone) StartLimit() bool {
	return false
}

// StopLimit always returns false on ChannelLimitNone.
// StopLimit 在 ChannelLimitNone 上恒返回 false。
func (l *ChannelLimitNone) StopLimit() bool {
	return false
}

// Add is a no-op on ChannelLimitNone.
// Add 在 ChannelLimitNone 上为空操作。
func (l *ChannelLimitNone) Add() {
	return
}

// Done is a no-op on ChannelLimitNone.
// Done 在 ChannelLimitNone 上为空操作。
func (l *ChannelLimitNone) Done() {
	return
}

// ChannelLimit limits concurrency using a buffered channel semaphore.
// ChannelLimit 使用带缓冲 channel 作为信号量限制并发。
type ChannelLimit struct {
	useLimit bool
	chanMax  int

	isStart bool
	limitMu sync.Mutex
	chanSem chan struct{}
}

// SetMax sets the maximum limit; only effective before StartLimit.
// SetMax 设置上限；仅在 StartLimit 之前生效。
func (l *ChannelLimit) SetMax(max int) {
	l.limitMu.Lock()
	defer l.limitMu.Unlock()
	if l.isStart {
		return
	}
	l.chanMax = max
	if max > 0 {
		l.useLimit = true
	} else {
		l.useLimit = false
	}
}

// StartLimit enables limiting and returns whether it was started.
// StartLimit 启动限制，返回是否成功启动。
func (l *ChannelLimit) StartLimit() bool {
	l.limitMu.Lock()
	defer l.limitMu.Unlock()
	if l.isStart {
		return false
	}
	if l.useLimit {
		l.chanSem = make(chan struct{}, l.chanMax)
		return true
	}
	return false
}

// StopLimit disables limiting and returns whether it was stopped.
// StopLimit 停止限制，返回是否成功停止。
func (l *ChannelLimit) StopLimit() bool {
	l.limitMu.Lock()
	defer l.limitMu.Unlock()
	if !l.isStart {
		return false
	}
	if l.useLimit {
		close(l.chanSem)
		l.chanSem = nil
		return true
	}
	return false
}

// Add acquires one slot (blocks when the limit is active).
// Add 占用一个槽位（限制生效时会阻塞）。
func (l *ChannelLimit) Add() {
	if l.useLimit {
		l.chanSem <- struct{}{}
	}
}

// Done releases one acquired slot.
// Done 释放一个已占用的槽位。
func (l *ChannelLimit) Done() {
	if l.useLimit {
		<-l.chanSem
	}
}
