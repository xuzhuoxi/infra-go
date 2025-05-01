// Package symmetric
// Create on 2025/4/18
// @author xuzhuoxi
package symmetric

import (
	"bytes"
	"testing"
)

var (
	xorPlaintext = []byte("Hello, World!")
	xorKey       = []byte("MySecretKey123456")

	xorBigData = make([]byte, 10*1024*1024)
)

func TestXOR(t *testing.T) {
	cipher := NewXORCipher(xorKey)
	obfuscated, _ := cipher.Encrypt(xorPlaintext)
	deobfuscated, _ := cipher.Decrypt(obfuscated)
	if !bytes.Equal(xorPlaintext, deobfuscated) {
		t.Fatal(xorPlaintext, obfuscated, deobfuscated)
	}
}

func xorWithMod(data, key []byte) []byte {
	keyLen := len(key)
	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ key[i%keyLen]
	}
	return result
}

func xorWithCounter(data, key []byte) []byte {
	keyLen := len(key)
	result := make([]byte, len(data))
	k := 0
	for i := range data {
		result[i] = data[i] ^ key[k]
		k++
		if k == keyLen {
			k = 0
		}
	}
	return result
}

func BenchmarkXORWithMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xorWithMod(xorBigData, xorKey)
	}
}

func BenchmarkXORWithCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xorWithCounter(xorBigData, xorKey)
	}
}
