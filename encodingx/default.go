//
//Created by xuzhuoxi
//on 2019-03-20.
//@author xuzhuoxi
//
package encodingx

import (
	"encoding/binary"
	"github.com/xuzhuoxi/infra-go/bytex"
)

var DataBlockHandler = bytex.NewDefaultDataBlockHandler()

var DefaultOrder = binary.BigEndian
