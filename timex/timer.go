//
//Created by xuzhuoxi
//on 2019-03-26.
//@author xuzhuoxi
//
package timex

import (
	"time"
)

func TimeFunc(f func()) (nano int64) {
	tn := time.Now().UnixNano()
	f()
	nano = time.Now().UnixNano() - tn
	return
}
