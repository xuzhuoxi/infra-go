// Package mgox
// Created by xuzhuoxi
// on 2019-06-16.
// @author xuzhuoxi
//
package mgox

import "context"

type IContextSetter interface {
	SetContext(context context.Context) error
}

type IContextGetter interface {
	GetContext() context.Context
}
