// Package mgox
// Created by xuzhuoxi
// on 2019-06-14.
// @author xuzhuoxi
//
package mgox

import (
	"context"
)

var (
	DefaultContext          = context.Background()
	ClientContext           = context.Background()
	DBReaderContext         = context.Background()
	DBWriterContext         = context.Background()
	CollectionReaderContext = context.Background()
	CollectionWriterContext = context.Background()
)
