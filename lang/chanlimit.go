//
//Created by xuzhuoxi
//on 2019-03-21.
//@author xuzhuoxi
//
package lang

import "sync"

type LChannelLimitMaxSetter interface {
	SetMax(max int)
}

type IChannelLimitSwitch interface {
	StartLimit() bool
	StopLimit() bool
}

type IChannelLimitHandler interface {
	Add()
	Done()
}

type IChannelLimit interface {
	LChannelLimitMaxSetter
	IChannelLimitSwitch
	IChannelLimitHandler
}

//-------------------------------

type ChannelLimitNone struct{}

func (l *ChannelLimitNone) SetMax(max int) {
	return
}

func (l *ChannelLimitNone) StartLimit() bool {
	return false
}

func (l *ChannelLimitNone) StopLimit() bool {
	return false
}

func (l *ChannelLimitNone) Add() {
	return
}

func (l *ChannelLimitNone) Done() {
	return
}

//-----------------------------------

type ChannelLimit struct {
	useLimit bool
	chanMax  int

	isStart bool
	limitMu sync.Mutex
	chanSem chan struct{}
}

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

func (l *ChannelLimit) Add() {
	if l.useLimit {
		l.chanSem <- struct{}{}
	}
}

func (l *ChannelLimit) Done() {
	if l.useLimit {
		<-l.chanSem
	}
}
