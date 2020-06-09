//
//Created by xuzhuoxi
//on 2019-02-05.
//@author xuzhuoxi
//
package cryptox

import (
	"fmt"
	"testing"
)

var (
	winDir = "D:/log/rsa/key/"
	macDir = "/Users/xuzhuoxi/key/"
)

func TestGenRSAKeyFile(t *testing.T) {
	dir := macDir
	err := GenRSAKeyFile(1024, dir+"private.pem", dir+"public.pem")
	fmt.Println(err)
}

func TestGenSSHKeyFile(t *testing.T) {
	dir := macDir
	err := GenSSHKeyFile(1024, dir+"id_rsa", dir+"id_rsa.pub")
	fmt.Println(err)
}
