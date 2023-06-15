// Package bytex
// Created by xuzhuoxi
// on 2019-02-12.
// @author xuzhuoxi
//
package bytex

import "encoding/binary"

var DefaultOrder = binary.BigEndian
var DefaultDataBlockHandler = NewDefaultDataBlockHandler()
