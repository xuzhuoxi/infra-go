// Package symmetric
// Create on 2025/7/15
// @author xuzhuoxi
package symmetric

import (
	"bytes"
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
	ctrNoNonce = aesCtrCaseItem{
		Plaintext: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Key: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}

	fixedKey   = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
	fixedNonce = []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66}
	useCases   = []aesCtrCaseItem{
		// 1️⃣ 标准短文本
		{
			Plaintext:  []byte("Hello AES-CTR"),
			Key:        fixedKey,
			Nonce:      fixedNonce,
			Ciphertext: []byte{223, 169, 30, 237, 220, 68, 46, 212, 201, 210, 102, 198, 157},
		},

		// 2️⃣ 空明文
		{
			Plaintext:  []byte{},
			Key:        fixedKey,
			Nonce:      fixedNonce,
			Ciphertext: []byte{},
		},

		// 3️⃣ 多块数据（128个字节）
		{
			Plaintext: bytes.Repeat([]byte{'A'}, 128),
			Key:       fixedKey,
			Nonce:     fixedNonce,
			Ciphertext: []byte{
				214, 141, 51, 192, 242, 37, 46, 208, 219, 190, 100, 211, 142, 124, 7, 3,
				173, 1, 69, 212, 78, 129, 39, 155, 237, 103, 37, 144, 20, 130, 69, 241,
				224, 246, 73, 10, 118, 202, 39, 129, 220, 151, 16, 95, 199, 78, 129, 101,
				100, 189, 48, 254, 16, 103, 92, 60, 221, 246, 214, 72, 131, 175, 233, 255,
				18, 91, 159, 182, 63, 55, 125, 247, 140, 249, 129, 160, 35, 34, 132, 198,
				51, 208, 131, 12, 39, 162, 214, 78, 235, 30, 201, 172, 223, 237, 156, 138,
				141, 90, 68, 179, 147, 93, 80, 57, 149, 65, 204, 40, 90, 45, 211, 83,
				247, 81, 188, 73, 31, 201, 124, 237, 202, 251, 13, 191, 191, 65, 191, 4,
			},
		},

		// 4️⃣ 边界长度：31个字节非对齐长度
		{
			Plaintext:  []byte("This message is 31 bytes long!!!"),
			Key:        fixedKey,
			Nonce:      fixedNonce,
			Ciphertext: []byte{195, 164, 27, 242, 147, 9, 10, 226, 233, 158, 66, 247, 239, 84, 53, 98, 223, 113, 36, 247, 118, 180, 3, 169, 140, 74, 11, 191, 50, 226, 37, 145},
		},
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

func TestAesCtrNoNonce(t *testing.T) {
	cipher := NewAESCipher(ctrNoNonce.Key)
	c, err := cipher.EncryptCTR(ctrNoNonce.Plaintext)
	if nil != err {
		t.Error(err)
		return
	}
	p, err := cipher.DecryptCTR(c)
	if nil != err {
		t.Error(err)
		return
	}
	if !slicex.EqualUint8(p, ctrNoNonce.Plaintext) {
		t.Log(ctrNoNonce.Plaintext)
		t.Log(c)
		t.Log(p)
		t.Fatal("plaintext != ciphertext")
	}
}

func TestAesCtrUseCases(t *testing.T) {
	for index, useCase := range useCases {
		cipher := NewAESCipher(useCase.Key)
		decrypted, err := cipher.DecryptNonceCTR(useCase.Ciphertext, useCase.Nonce, ctrCounter)
		if nil != err {
			t.Error(err)
			return
		}
		if !slicex.EqualUint8(decrypted, useCase.Plaintext) {
			t.Log(ctrNoNonce.Decrypted)
			t.Fatal(index, "plaintext != ciphertext")
		}
	}
}
