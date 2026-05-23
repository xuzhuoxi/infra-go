package symmetric

import (
	"github.com/xuzhuoxi/infra-go/cryptox"
)

// IXORCipher is a fast XOR stream obfuscator (not cryptographically secure).
// IXORCipher 是快速 XOR 流混淆器（不具备密码学安全性）。
// Features: fast; not secure for sensitive data.
// 特点：1. 快；2. 不安全。
type IXORCipher interface {
	cryptox.ICipher
}

// NewXORCipher creates an XOR cipher with the given repeating key.
// NewXORCipher 使用给定循环密钥创建 XOR 混淆器。
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
