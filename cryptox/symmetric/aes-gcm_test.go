// Package symmetric
// Create on 2025/7/13
// @author xuzhuoxi
package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/xuzhuoxi/infra-go/slicex"
	"testing"
)

type aesGcmCaseItem struct {
	Plaintext  []byte
	Key        []byte
	Nonce      []byte
	Ciphertext []byte
	Decrypted  []byte
}

func (o aesGcmCaseItem) Tag() []byte {
	return o.Ciphertext[len(o.Plaintext):]
}

func (o aesGcmCaseItem) String() string {
	pStr := fmt.Sprintf("Plaintext: %v | %s", o.Plaintext, string(o.Plaintext))
	kStr := fmt.Sprintf("Key: %v", o.Key)
	//kStr := fmt.Sprintf("Key: %X", o.Key)
	nStr := fmt.Sprintf("Nonce: %v", o.Nonce)
	//nStr := fmt.Sprintf("Nonce: %X", o.Nonce)
	cStr := fmt.Sprintf("Ciphertext: %v | %s", o.Ciphertext, string(o.Ciphertext))
	tStr := fmt.Sprintf("Tag: %v", o.Tag())
	dStr := fmt.Sprintf("Decrypted: %v | %s", o.Decrypted, string(o.Decrypted))
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", pStr, kStr, nStr, cStr, tStr, dStr)
}

func TestAesGcmDemo(t *testing.T) {
	plaintext := []byte("Hello AES-GCM from Go!")
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	rand.Read(key)
	rand.Read(nonce)

	block, _ := aes.NewCipher(key)
	aesGcm, _ := cipher.NewGCM(block)
	ciphertext := aesGcm.Seal(nil, nonce, plaintext, nil)

	item := &aesGcmCaseItem{
		Plaintext:  plaintext,
		Key:        key,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}
	fmt.Println(item)
}

var (
	demoAesGcmCase = aesGcmCaseItem{
		Plaintext:  []byte("Hello AES-CTR from Go!"),
		Key:        []byte{221, 114, 67, 111, 131, 146, 36, 124, 118, 142, 179, 76, 214, 202, 1, 58, 40, 227, 121, 123, 131, 175, 210, 141, 26, 202, 201, 128, 239, 34, 252, 132},
		Nonce:      []byte{201, 61, 50, 23, 128, 203, 94, 229, 160, 193, 219, 108},
		Ciphertext: []byte{177, 245, 71, 30, 174, 147, 216, 112, 113, 47, 206, 149, 85, 239, 246, 157, 14, 221, 143, 66, 126, 245, 39, 38, 172, 93, 34, 15, 48, 157, 242, 52, 202, 236, 167, 231, 100, 219},
	}
	aesNonceCases = []aesGcmCaseItem{
		{
			Plaintext: []byte("Hello World!Hello World!Hello World!"),
			Key: []byte{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			},
			Nonce: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67},
		},
		{
			Plaintext: []byte("Hello World!Hello World!Hello World!666"),
			Key: []byte{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			},
			Nonce: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67},
		},
	}
)

func TestAesGcmDemoCase(t *testing.T) {
	cipher := NewAESCipher(demoAesGcmCase.Key)
	c, err := cipher.EncryptNonceGCM(demoAesGcmCase.Plaintext, demoAesGcmCase.Nonce)
	if err != nil {
		t.Error(err)
		return
	}
	p, err := cipher.DecryptNonceGCM(c, demoAesGcmCase.Nonce)
	if err != nil {
		t.Error(err)
		return
	}
	if !slicex.EqualUint8(p, demoAesCtrCase.Plaintext) {
		t.Log(demoAesCtrCase.Plaintext)
		t.Log(p)
		t.Fatal("plaintext != ciphertext")
	}
}

func TestGcmNonceCases(t *testing.T) {
	for _, value := range aesNonceCases {
		testGcmNonceDecrypt(t, value)
		fmt.Println("---------- ---------- ---------- ---------- ---------- ----------")
	}
}

func testGcmNonceDecrypt(t *testing.T, value aesGcmCaseItem) {
	cipher := NewAESCipher(value.Key)
	ciphertext, err1 := cipher.EncryptNonceGCM(value.Plaintext, value.Nonce)
	decrypted, err2 := cipher.DecryptNonceGCM(ciphertext, value.Nonce)
	value.Ciphertext, value.Decrypted = ciphertext, decrypted
	compareGcmResult(t, value, err1, err2)
}

func compareGcmResult(t *testing.T, value aesGcmCaseItem, err1, err2 error) {
	title := "Gcm"
	if err1 != nil || err2 != nil {
		t.Fatal(title, value, err1, err2)
		return
	}
	if !slicex.EqualUint8(value.Plaintext, value.Decrypted) {
		t.Fatal(title, value)
		return
	}
	t.Log(title, value)
}
