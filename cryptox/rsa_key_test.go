//
//Created by xuzhuoxi
//on 2019-02-05.
//@author xuzhuoxi
//
package cryptox

import "testing"

func TestGenRSAKeyFile(t *testing.T) {
	dir := "D:/log/rsa/key/"
	GenRSAKeyFile(1024, dir+"private.pem", dir+"public.pem")
}

func TestGenSSHKeyFile(t *testing.T) {
	dir := "D:/log/rsa/key/"
	GenSSHKeyFile(1024, dir+"id_rsa", dir+"id_rsa.pub")
}
