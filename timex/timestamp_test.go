// Create on 2023/8/16
// @author xuzhuoxi
package timex

import (
	"fmt"
	"testing"
)

func TestTimestamp(t *testing.T) {
	fmt.Println("UTC:\t\t", ZeroUTC.UnixNano())
	fmt.Println("Local:\t\t", ZeroLocal.UnixNano())
	fmt.Println("UTC1970:\t", Zero1970UTC.UnixNano())
	fmt.Println("Local1970:\t", Zero1970Local.UnixNano())
}
