// Package mgox
// Created by xuzhuoxi
// on 2019-06-15.
// @author xuzhuoxi
//
package mgox

import "go.mongodb.org/mongo-driver/mongo/options"

func NewOptions(uri string) *options.ClientOptions {
	return options.Client().ApplyURI(uri)
}
