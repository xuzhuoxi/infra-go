// Package key
// Create on 2025/4/11
// @author xuzhuoxi
package key

import (
	"encoding/base64"
	"testing"
)

var (
	passphrase = "123456"
)

func TestSharedKeySha256(t *testing.T) {
	sk := SharedKeySha256Str(passphrase)
	t.Log(sk)
	t.Log(base64.StdEncoding.EncodeToString(sk))
}

func TestSharedKeyPbkdf2(t *testing.T) {
	sk := DeriveKeyPbkdf2Str(passphrase, salt, iterations, keyLen)
	t.Log(sk)
	t.Log(base64.StdEncoding.EncodeToString(sk))
}
