//
//Created by xuzhuoxi
//on 2019-02-22.
//@author xuzhuoxi
//
package bytex

import (
	"bytes"
	"fmt"
	"testing"
)

func TestByteBuff(t *testing.T) {
	bs1 := []byte{1, 2, 3}
	bs2 := []byte{3, 2, 1}
	buff := bytes.NewBuffer(nil)
	buff.Write(bs1)
	buff.Write(bs2)
	fmt.Println(buff.Bytes())
	bs1[0] = 2
	fmt.Println(buff.Bytes())
}
