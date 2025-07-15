// Package symmetric
// Create on 2025/7/15
// @author xuzhuoxi
package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/xuzhuoxi/infra-go/slicex"
	"math/rand"
	"testing"
)

type aesCtrCaseItem struct {
	Plaintext  []byte
	Key        []byte
	Nonce      []byte
	Counter    []byte
	Ciphertext []byte
	Decrypted  []byte
}

func (o aesCtrCaseItem) String() string {
	pStr := fmt.Sprintf("Plaintext: %v | %s", o.Plaintext, string(o.Plaintext))
	kStr := fmt.Sprintf("Key: %v", o.Key)
	//kStr := fmt.Sprintf("Key: %X", o.Key)
	nStr := fmt.Sprintf("Nonce: %v", o.Nonce)
	//nStr := fmt.Sprintf("Nonce: %X", o.Nonce)
	cStr := fmt.Sprintf("Ciphertext: %v | %s", o.Ciphertext, string(o.Ciphertext))
	dStr := fmt.Sprintf("Decrypted: %v | %s", o.Decrypted, string(o.Decrypted))
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", pStr, kStr, nStr, cStr, dStr)
}

func TestAesCtrDemo(t *testing.T) {
	plaintext := []byte("Hello AES-CTR from Go!")
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	rand.Read(key)
	rand.Read(nonce)

	block, _ := aes.NewCipher(key)
	ctr := cipher.NewCTR(block, append(nonce, 0, 0, 0, 1)) // nonce + counter

	ciphertext := make([]byte, len(plaintext))
	ctr.XORKeyStream(ciphertext, plaintext)
	item := aesCtrCaseItem{Plaintext: plaintext, Key: key, Nonce: nonce, Ciphertext: ciphertext}
	fmt.Println(item)
}

var (
	demoAesCtrCase = aesCtrCaseItem{
		Plaintext:  []byte("Hello AES-CTR from Go!"),
		Key:        []byte{171, 145, 30, 202, 108, 201, 19, 35, 72, 50, 210, 129, 250, 80, 25, 96, 135, 61, 24, 141, 48, 27, 200, 107, 95, 4, 240, 93, 74, 127, 22, 208},
		Nonce:      []byte{188, 93, 26, 190, 198, 36, 25, 25, 115, 49, 208, 117},
		Ciphertext: []byte{25, 166, 2, 181, 129, 251, 167, 211, 152, 185, 167, 224, 84, 86, 161, 225, 65, 217, 34, 211, 44, 89},
	}
	demoAesCtrCaseMode = aesCtrCaseItem{
		Plaintext: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Key: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}
)

func TestAesCtrDemoCase(t *testing.T) {
	cipher := NewAESCipher(demoAesCtrCase.Key)
	c, err := cipher.EncryptNonceCTR(demoAesCtrCase.Plaintext, demoAesCtrCase.Nonce, ctrCounter)
	if nil != err {
		t.Error(err)
		return
	}
	p, err := cipher.DecryptNonceCTR(c, demoAesCtrCase.Nonce, ctrCounter)
	if nil != err {
		t.Error(err)
		return
	}
	if !slicex.EqualUint8(p, demoAesCtrCase.Plaintext) {
		t.Log(demoAesCtrCase.Plaintext)
		t.Log(p)
		t.Fatal("plaintext != ciphertext")
	}
}

func TestAesCtrDemoCaseMode(t *testing.T) {
	cipher := NewAESCipher(demoAesCtrCaseMode.Key)
	c, err := cipher.EncryptCTR(demoAesCtrCaseMode.Plaintext)
	if nil != err {
		t.Error(err)
		return
	}
	p, err := cipher.DecryptCTR(c)
	if nil != err {
		t.Error(err)
		return
	}
	if !slicex.EqualUint8(p, demoAesCtrCaseMode.Plaintext) {
		t.Log(demoAesCtrCaseMode.Plaintext)
		t.Log(c)
		t.Log(p)
		t.Fatal("plaintext != ciphertext")
	}
}
