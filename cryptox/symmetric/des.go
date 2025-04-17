// Package symmetric
// Create on 2025/4/6
// @author xuzhuoxi
package symmetric

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"github.com/xuzhuoxi/infra-go/cryptox"
	"io"
)

// IDESCipher
// DES：Data Encrytion Standard（数据加密标准），对应算法是DEA
// 特点：
//  1. 对称加密
//  2. 同一个SK
type IDESCipher interface {
	cryptox.ICipher
	// SetPaddingFunc 设置填充方式
	SetPaddingFunc(padding cryptox.FuncPadding, unPadding cryptox.FuncUnPadding)
	// EncryptMode 指定BlockMode加密
	EncryptMode(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// DecryptMode 指定BlockMode解密
	DecryptMode(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// EncryptECB 使用ECB模式加密
	EncryptECB(plaintext []byte) ([]byte, error)
	// DecryptECB 使用ECB模式解密
	DecryptECB(ciphertext []byte) ([]byte, error)
	// EncryptCBC 使用CBC模式加密
	EncryptCBC(plaintext []byte) ([]byte, error)
	// DecryptCBC 使用CBC模式解密
	DecryptCBC(ciphertext []byte) ([]byte, error)
	// EncryptCTR 使用CTR模式加密
	EncryptCTR(plaintext []byte) ([]byte, error)
	// DecryptCTR 使用CTR模式解密
	DecryptCTR(ciphertext []byte) ([]byte, error)
}

func NewDESCipher(key []byte) IDESCipher {
	rs := &desCipher{key: key}
	keyLen := len(key)
	if keyLen == des.BlockSize {
		rs.block, rs.err = des.NewCipher(key)
	} else if keyLen == 16 {
		newKey := append(key[:], key[:8]...)
		rs.block, rs.err = des.NewTripleDESCipher(newKey)
	} else if keyLen == 24 {
		rs.block, rs.err = des.NewTripleDESCipher(key)
	} else {
		rs.err = errors.New("key must be 8 (DES) or 24 (3DES) bytes")
		return rs
	}
	rs.padding, rs.unPadding = cryptox.PKCS7Padding, cryptox.PKCS7UnPadding
	return rs
}

type desCipher struct {
	key       []byte
	padding   cryptox.FuncPadding
	unPadding cryptox.FuncUnPadding
	block     cipher.Block
	err       error
}

func (o *desCipher) SetPaddingFunc(padding cryptox.FuncPadding, unPadding cryptox.FuncUnPadding) {
	o.padding, o.unPadding = padding, unPadding
}

func (o *desCipher) Encrypt(plaintext []byte) ([]byte, error) {
	return o.EncryptCBC(plaintext)
}

func (o *desCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	return o.DecryptCBC(ciphertext)
}

func (o *desCipher) EncryptMode(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.ECB:
		return o.EncryptECB(plaintext)
	case cryptox.CBC:
		return o.EncryptCBC(plaintext)
	case cryptox.CTR:
		return o.EncryptCTR(plaintext)
	default:
		return nil, errors.New("Unsupported DES block mode! ")
	}
}

func (o *desCipher) DecryptMode(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.ECB:
		return o.DecryptECB(ciphertext)
	case cryptox.CBC:
		return o.DecryptCBC(ciphertext)
	case cryptox.CTR:
		return o.DecryptCTR(ciphertext)
	default:
		return nil, errors.New("Unsupported DES block mode! ")
	}
}

func (o *desCipher) EncryptECB(plaintext []byte) ([]byte, error) {
	blockSize := o.block.BlockSize()
	padded, err := o.padding(plaintext, blockSize)
	if nil != err {
		return nil, err
	}
	ciphertext := make([]byte, len(padded))
	for bs, be := 0, blockSize; bs < len(padded); bs, be = bs+blockSize, be+blockSize {
		o.block.Encrypt(ciphertext[bs:be], padded[bs:be])
	}
	return ciphertext, nil
}

func (o *desCipher) DecryptECB(ciphertext []byte) ([]byte, error) {
	blockSize := o.block.BlockSize()
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("invalid ciphertext length")
	}
	plain := make([]byte, len(ciphertext))
	for bs, be := 0, blockSize; bs < len(ciphertext); bs, be = bs+blockSize, be+blockSize {
		o.block.Decrypt(plain[bs:be], ciphertext[bs:be])
	}
	return o.unPadding(plain)
}

func (o *desCipher) EncryptCBC(plaintext []byte) ([]byte, error) {
	blockSize := o.block.BlockSize()
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	padded, err := o.padding(plaintext, blockSize)
	if nil != err {
		return nil, err
	}
	ciphertext := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(o.block, iv)
	mode.CryptBlocks(ciphertext, padded)
	return append(iv, ciphertext...), nil
}

func (o *desCipher) DecryptCBC(ciphertext []byte) ([]byte, error) {
	blockSize := o.block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:blockSize]
	data := ciphertext[blockSize:]
	if len(data)%blockSize != 0 {
		return nil, errors.New("invalid CBC ciphertext length")
	}
	mode := cipher.NewCBCDecrypter(o.block, iv)
	plain := make([]byte, len(data))
	mode.CryptBlocks(plain, data)
	return o.unPadding(plain)
}

func (o *desCipher) EncryptCTR(plaintext []byte) ([]byte, error) {
	blockSize := o.block.BlockSize()
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(o.block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return append(iv, ciphertext...), nil
}

func (o *desCipher) DecryptCTR(ciphertext []byte) ([]byte, error) {
	blockSize := o.block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:blockSize]
	data := ciphertext[blockSize:]
	stream := cipher.NewCTR(o.block, iv)
	plain := make([]byte, len(data))
	stream.XORKeyStream(plain, data)
	return plain, nil
}
