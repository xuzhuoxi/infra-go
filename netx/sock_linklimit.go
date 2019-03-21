//
//Created by xuzhuoxi
//on 2019-03-21.
//@author xuzhuoxi
//
package netx

import "sync"

type LLinkLimitMaxSetter interface {
	SetLinkMax(max int)
}

type ILinkLimitSwitch interface {
	StartLimit() bool
	StopLimit() bool
}

type ILinkLimitHandler interface {
	Add()
	Done()
}

type ILinkLimit interface {
	LLinkLimitMaxSetter
	ILinkLimitSwitch
	ILinkLimitHandler
}

//-------------------------------

type NoLinkLimit struct{}

func (l *NoLinkLimit) SetLinkMax(max int) {
	return
}

func (l *NoLinkLimit) StartLimit() bool {
	return false
}

func (l *NoLinkLimit) StopLimit() bool {
	return false
}

func (l *NoLinkLimit) Add() {
	return
}

func (l *NoLinkLimit) Done() {
	return
}

//-----------------------------------

type LinkLimit struct {
	openLimit bool
	maxLink   int

	isLimitStart bool
	limitMu      sync.Mutex
	linkSem      chan struct{}
}

func (l *LinkLimit) SetLinkMax(max int) {
	l.limitMu.Lock()
	defer l.limitMu.Unlock()
	if l.isLimitStart {
		return
	}
	l.maxLink = max
	if max > 0 {
		l.openLimit = true
	} else {
		l.openLimit = false
	}
}

func (l *LinkLimit) StartLimit() bool {
	l.limitMu.Lock()
	defer l.limitMu.Unlock()
	if l.isLimitStart {
		return false
	}
	if l.openLimit {
		l.linkSem = make(chan struct{}, l.maxLink)
		return true
	}
	return false
}

func (l *LinkLimit) StopLimit() bool {
	l.limitMu.Lock()
	defer l.limitMu.Unlock()
	if !l.isLimitStart {
		return false
	}
	if l.openLimit {
		close(l.linkSem)
		l.linkSem = nil
		return true
	}
	return false
}

func (l *LinkLimit) Add() {
	if l.openLimit {
		l.linkSem <- struct{}{}
	}
}

func (l *LinkLimit) Done() {
	if l.openLimit {
		<-l.linkSem
	}
}
