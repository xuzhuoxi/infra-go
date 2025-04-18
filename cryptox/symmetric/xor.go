// Package symmetric
// Create on 2025/4/18
// @author xuzhuoxi
package symmetric

import (
	"github.com/xuzhuoxi/infra-go/cryptox"
)

// IXORCipher
// 异常混淆器：Data Encrytion Standard（数据加密标准），对应算法是DEA
// 特点：
//  1. 快
//  2. 不安全
type IXORCipher interface {
	cryptox.ICipher
}

func NewXORCipher(key []byte) IXORCipher {
	rs := &xorCipher{key: key, keyLen: len(key)}
	return rs
}

type xorCipher struct {
	key    []byte
	keyLen int
}

func (o *xorCipher) Encrypt(plaintext []byte) ([]byte, error) {
	if o.keyLen == 0 {
		return plaintext, nil
	}
	return o.xor(plaintext), nil
}

func (o *xorCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	if o.keyLen == 0 {
		return ciphertext, nil
	}
	return o.xor(ciphertext), nil
}
func (o *xorCipher) xor(text []byte) []byte {
	result := make([]byte, len(text))

	keyIndex := 0
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ o.key[keyIndex]
		keyIndex++
		if keyIndex == o.keyLen {
			keyIndex = 0
		}
	}

	return result
}
