// Package key
// Create on 2025/4/6
// @author xuzhuoxi
package key

import (
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
)

var (
	salt       = []byte("infra-go:cryptox.key")
	iterations = 100000
	keyLen     = 32
)

// SharedKeySha256Str
// 将任意字符串（如密码、passphrase）转换为 32 字节密钥（适合 AES-256、HMAC）
func SharedKeySha256Str(passphrase string) []byte {
	hash := sha256.Sum256([]byte(passphrase))
	return hash[:]
}

// SharedKeySha256
// 将任意字符串（如密码、passphrase）转换为 32 字节密钥（适合 AES-256、HMAC）
func SharedKeySha256(passphrase []byte) []byte {
	hash := sha256.Sum256(passphrase)
	return hash[:]
}

// DeriveKeyPbkdf2StrDefault
// 使用 PBKDF2 派生出一个强密钥（推荐用于生产场景）
func DeriveKeyPbkdf2StrDefault(passphrase string) []byte {
	return DeriveKeyPbkdf2Str(passphrase, salt, iterations, keyLen)
}

// DeriveKeyPbkdf2Default
// 使用 PBKDF2 派生出一个强密钥（推荐用于生产场景）
func DeriveKeyPbkdf2Default(passphrase []byte) []byte {
	return DeriveKeyPbkdf2(passphrase, salt, iterations, keyLen)
}

// DeriveKeyPbkdf2Str
// 使用 PBKDF2 从密码 + salt 派生出一个强密钥（推荐用于生产场景）
// iterations 越大越安全，建议 10,000 以上
func DeriveKeyPbkdf2Str(passphrase string, salt []byte, iterations int, keyLen int) []byte {
	return pbkdf2.Key([]byte(passphrase), salt, iterations, keyLen, sha512.New)
}

// DeriveKeyPbkdf2
// 使用 PBKDF2 从密码 + salt 派生出一个强密钥（推荐用于生产场景）
// iterations 越大越安全，建议 10,000 以上
func DeriveKeyPbkdf2(passphrase []byte, salt []byte, iterations int, keyLen int) []byte {
	return pbkdf2.Key(passphrase, salt, iterations, keyLen, sha512.New)
}
