// Package asymmetric
// Created by xuzhuoxi
// on 2019-02-04.
// @author xuzhuoxi
//
package asymmetric

import (
	"crypto"
	"github.com/xuzhuoxi/infra-go/filex"
	"testing"

	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	_ "golang.org/x/crypto/ripemd160"
)

var (
	rsaDir  string
	keyDir  string = "/key"
	fileDir string = "/file"
	bits    int    = 2048
)

var (
	privateSSHKeyName = "id_rsa"
	publicSSHKeyName  = "id_rsa.pub"

	privateRsaKeyName   = "rsa_private.pem"
	publicRsaKeyName    = "rsa_public.pem"
	privatePkcs8KeyName = "pkcs8_private.pem"
	publicX509KeyName   = "x509_public.pem"
)

var (
	contentShort = `abcdabcdabcd`
	contentLong  = `
abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,
abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,
abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,
abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,
abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,
abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,abcdabcdabcd,
`
)

var (
	hashlist = []crypto.Hash{
		crypto.MD5,       // import crypto/md5
		crypto.SHA1,      // import crypto/sha1
		crypto.SHA224,    // import crypto/sha256
		crypto.SHA256,    // import crypto/sha256
		crypto.SHA384,    // import crypto/sha512
		crypto.SHA512,    // import crypto/sha512
		crypto.RIPEMD160, // import golang.org/x/crypto/ripemd160
	}
)

func TestGenerateKeyFiles(t *testing.T) {
	privateKey, _, err := GenerateRSAKeyPair(bits)
	if nil != err {
		t.Fatal(err)
	}
	err = SaveOpenSSHKeyFiles(privateKey,
		filex.Combine(rsaDir, keyDir, privateSSHKeyName),
		filex.Combine(rsaDir, keyDir, publicSSHKeyName))
	if nil != err {
		t.Fatal(err)
	}
	err = SaveRSAKeyFiles(privateKey,
		filex.Combine(rsaDir, keyDir, privateRsaKeyName),
		filex.Combine(rsaDir, keyDir, publicRsaKeyName))
	if nil != err {
		t.Fatal(err)
	}
	err = SavePemKeyFiles(privateKey,
		filex.Combine(rsaDir, keyDir, privatePkcs8KeyName),
		filex.Combine(rsaDir, keyDir, publicX509KeyName))
	if nil != err {
		t.Fatal(err)
	}
}

func TestRSAKeyEncrypt(t *testing.T) {
	privateCipher, publicCipher, err := LoadCiphersRSA(
		filex.Combine(rsaDir, keyDir, privateRsaKeyName),
		filex.Combine(rsaDir, keyDir, publicRsaKeyName))

	if nil != err {
		t.Fatal(err)
	}
	testKeyEncrypt(privateCipher, publicCipher, contentShort, t)
	testKeyEncrypt(privateCipher, publicCipher, contentLong, t)
}

func TestSSHKeyEncrypt(t *testing.T) {
	privateCipher, publicCipher, err := LoadCiphersSSH(
		filex.Combine(rsaDir, keyDir, privateSSHKeyName),
		filex.Combine(rsaDir, keyDir, publicSSHKeyName))

	if nil != err {
		t.Fatal(err)
	}
	testKeyEncrypt(privateCipher, publicCipher, contentShort, t)
	testKeyEncrypt(privateCipher, publicCipher, contentLong, t)
}

func TestPemKeyEncrypt(t *testing.T) {
	privateCipher, publicCipher, err := LoadCiphersPEM(
		filex.Combine(rsaDir, keyDir, privatePkcs8KeyName),
		filex.Combine(rsaDir, keyDir, publicX509KeyName))

	if nil != err {
		t.Fatal(err)
	}
	testKeyEncrypt(privateCipher, publicCipher, contentShort, t)
	testKeyEncrypt(privateCipher, publicCipher, contentLong, t)
}

func TestRSAKeySign(t *testing.T) {
	privateCipher, publicCipher, err := LoadCiphersRSA(
		filex.Combine(rsaDir, keyDir, privateRsaKeyName),
		filex.Combine(rsaDir, keyDir, publicRsaKeyName))

	if nil != err {
		t.Fatal(err)
	}
	testKeySign(privateCipher, publicCipher, []byte(contentShort), t)
	testKeySign(privateCipher, publicCipher, []byte(contentLong), t)
	for _, hash := range hashlist {
		testKeySignHash(privateCipher, publicCipher, []byte(contentShort), hash, t)
		testKeySignHash(privateCipher, publicCipher, []byte(contentLong), hash, t)
	}
}
func TestSSHKeySign(t *testing.T) {
	privateCipher, publicCipher, err := LoadCiphersSSH(
		filex.Combine(rsaDir, keyDir, privateSSHKeyName),
		filex.Combine(rsaDir, keyDir, publicSSHKeyName))

	if nil != err {
		t.Fatal(err)
	}
	testKeySign(privateCipher, publicCipher, []byte(contentShort), t)
	testKeySign(privateCipher, publicCipher, []byte(contentLong), t)
	for _, hash := range hashlist {
		testKeySignHash(privateCipher, publicCipher, []byte(contentShort), hash, t)
		testKeySignHash(privateCipher, publicCipher, []byte(contentLong), hash, t)
	}
}

func TestPemKeySign(t *testing.T) {
	privateCipher, publicCipher, err := LoadCiphersPEM(
		filex.Combine(rsaDir, keyDir, privatePkcs8KeyName),
		filex.Combine(rsaDir, keyDir, publicX509KeyName))

	if nil != err {
		t.Fatal(err)
	}
	testKeySign(privateCipher, publicCipher, []byte(contentShort), t)
	testKeySign(privateCipher, publicCipher, []byte(contentLong), t)
	for _, hash := range hashlist {
		testKeySignHash(privateCipher, publicCipher, []byte(contentShort), hash, t)
		testKeySignHash(privateCipher, publicCipher, []byte(contentLong), hash, t)
	}
}

func testKeyEncrypt(private IRSAPrivateCipher, public IRSAPublicCipher, content string, t *testing.T) {
	ciphertext, err := public.Encrypt([]byte(content))
	if nil != err {
		t.Fatal(err)
	}
	plaintext, err := private.Decrypt(ciphertext)
	if nil != err {
		t.Fatal(err)
	}
	if string(plaintext) != content {
		t.Fatal("plaintext != content")
	}
}
func testKeySign(private IRSAPrivateCipher, public IRSAPublicCipher, content []byte, t *testing.T) {
	signed, err := private.Sign(content)
	if nil != err {
		t.Fatal(err)
	}
	b, err := public.VerifySign(content, signed)
	if nil != err {
		t.Fatal(err)
	}
	if !b {
		t.Fatal("sign fail")
	}
}
func testKeySignHash(private IRSAPrivateCipher, public IRSAPublicCipher, content []byte, hash crypto.Hash, t *testing.T) {
	signed, err := private.SignHash(content, hash)
	if nil != err {
		t.Fatal(err)
	}
	b, err := public.VerifySignHash(content, signed, hash)
	if nil != err {
		t.Fatal(hash, err)
	}
	if !b {
		t.Fatal("sign hash fail")
	}
}
