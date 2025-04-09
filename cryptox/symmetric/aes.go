// Package symmetric
// Create on 2025/4/6
// @author xuzhuoxi
package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/xuzhuoxi/infra-go/cryptox"
	"io"
)

// IAESCipher
// AES：Advanced Encrytion Standard（高级加密标准），对应算法Rijndael
// 特点：
//  1. 对称加密
//  2. 一个SK扩展成多个子SK，多轮加密
type IAESCipher interface {
	// SetPaddingFunc 设置填充方式
	SetPaddingFunc(padding cryptox.FuncPadding, unPadding cryptox.FuncUnPadding)
	// Encrypt 指定BlockMode加密
	Encrypt(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// Decrypt 指定BlockMode解密
	Decrypt(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// EncryptCBC 使用CBC模式加密
	EncryptCBC(plaintext []byte) ([]byte, error)
	// DecryptCBC 使用CBC模式解密
	DecryptCBC(ciphertext []byte) ([]byte, error)
	// EncryptCTR 使用CTR模式加密
	EncryptCTR(plaintext []byte) ([]byte, error)
	// DecryptCTR 使用CTR模式解密
	DecryptCTR(ciphertext []byte) ([]byte, error)
	// EncryptGCM 使用GCM模式加密
	EncryptGCM(plaintext []byte) ([]byte, error)
	// DecryptGCM 使用GCM模式解密
	DecryptGCM(ciphertext []byte) ([]byte, error)
}

func NewAESCipher(key []byte) IAESCipher {
	rs := &aesCipher{key: key}
	rs.block, rs.err = aes.NewCipher(key)
	rs.padding, rs.unPadding = cryptox.PKCS7Padding, cryptox.PKCS7UnPadding
	return rs
}

type aesCipher struct {
	key       []byte
	padding   cryptox.FuncPadding
	unPadding cryptox.FuncUnPadding
	block     cipher.Block
	err       error
}

func (o *aesCipher) SetPaddingFunc(padding cryptox.FuncPadding, unPadding cryptox.FuncUnPadding) {
	o.padding, o.unPadding = padding, unPadding
}

func (o *aesCipher) Encrypt(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.CBC:
		return o.EncryptCBC(plaintext)
	case cryptox.CTR:
		return o.EncryptCTR(plaintext)
	case cryptox.GCM:
		return o.EncryptGCM(plaintext)
	default:
		return nil, errors.New("Unsupported AES block mode! ")
	}
}

func (o *aesCipher) Decrypt(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.CBC:
		return o.DecryptCBC(ciphertext)
	case cryptox.CTR:
		return o.DecryptCTR(ciphertext)
	case cryptox.GCM:
		return o.DecryptGCM(ciphertext)
	default:
		return nil, errors.New("Unsupported AES block mode! ")
	}
}

func (o *aesCipher) EncryptCBC(plaintext []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	padded, err := o.padding(plaintext, aes.BlockSize)
	if nil != err {
		return nil, err
	}
	ciphertext := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(o.block, iv)
	mode.CryptBlocks(ciphertext, padded)
	return append(iv, ciphertext...), nil
}

func (o *aesCipher) DecryptCBC(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	data := ciphertext[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext not a multiple of block size")
	}
	mode := cipher.NewCBCDecrypter(o.block, iv)
	plainPadded := make([]byte, len(data))
	mode.CryptBlocks(plainPadded, data)
	return o.unPadding(plainPadded)
}

func (o *aesCipher) EncryptCTR(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(o.block, nonce)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return append(nonce, ciphertext...), nil
}

func (o *aesCipher) DecryptCTR(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce := ciphertext[:aes.BlockSize]
	data := ciphertext[aes.BlockSize:]
	stream := cipher.NewCTR(o.block, nonce)
	plaintext := make([]byte, len(data))
	stream.XORKeyStream(plaintext, data)
	return plaintext, nil
}

func (o *aesCipher) EncryptGCM(plaintext []byte) ([]byte, error) {
	aead, err := cipher.NewGCM(o.block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

func (o *aesCipher) DecryptGCM(ciphertext []byte) ([]byte, error) {
	aead, err := cipher.NewGCM(o.block)
	if err != nil {
		return nil, err
	}
	nonceSize := aead.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce := ciphertext[:nonceSize]
	data := ciphertext[nonceSize:]
	return aead.Open(nil, nonce, data, nil)
}
