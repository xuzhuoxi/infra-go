// Package symmetric
// Create on 2025/4/6
// @author xuzhuoxi
package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/xuzhuoxi/infra-go/cryptox"
)

// IAESCipher
// AES：Advanced Encrytion Standard（高级加密标准），对应算法Rijndael
// 特点：
//  1. 对称加密
//  2. 一个SK扩展成多个子SK，多轮加密
type IAESCipher interface {
	cryptox.ICipher
	// Key 获取密钥,只读
	Key() []byte
	// SetPaddingFunc 设置填充方式
	SetPaddingFunc(padding cryptox.FuncPadding, unPadding cryptox.FuncUnPadding)

	// EncryptMode 指定BlockMode加密, 使用随机IV，并把IV数据放在密文头部
	EncryptMode(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// EncryptIVMode 指定BlockMode加密, 指定IV
	EncryptIVMode(plaintext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// DecryptMode Decrypt 指定BlockMode解密, 从密文头部提取IV
	DecryptMode(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// DecryptIVMode Decrypt 指定BlockMode解密, 指定IV
	DecryptIVMode(ciphertext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error)

	// EncryptCBC 使用CBC模式加密, 使用随机IV，并把IV数据放在密文头部
	EncryptCBC(plaintext []byte) ([]byte, error)
	// EncryptIVCBC 使用CBC模式加密, 指定IV
	EncryptIVCBC(plaintext []byte, iv []byte) ([]byte, error)
	// DecryptCBC 使用CBC模式解密，从密文头部提取IV
	DecryptCBC(ciphertext []byte) ([]byte, error)
	// DecryptIVCBC 使用CBC模式解密, 指定IV
	DecryptIVCBC(ciphertext []byte, iv []byte) ([]byte, error)

	// EncryptCTR 使用CTR模式加密, 使用随机IV(nonce+counter)，并把IV数据放在密文头部
	EncryptCTR(plaintext []byte) ([]byte, error)
	// EncryptNonceCTR 使用CTR模式加密, 指定Nonce，并使用默认counter，返回结果不包含IV数据
	EncryptNonceCTR(plaintext []byte, nonce []byte, counter []byte) ([]byte, error)
	// EncryptIVCTR 使用CTR模式加密, 指定IV(nonce+counter)，返回结果不包含IV数据
	EncryptIVCTR(plaintext []byte, iv []byte) ([]byte, error)
	// DecryptCTR 使用CTR模式解密，从密文头部提取IV
	DecryptCTR(ciphertext []byte) ([]byte, error)
	// DecryptNonceCTR 使用CTR模式解密, 密文内不包含IV数据, 需要指定Nonce数据，并使用默认的counter
	DecryptNonceCTR(ciphertext []byte, nonce []byte, counter []byte) ([]byte, error)
	// DecryptIVCTR  使用CTR模式解密, 密文内不包含IV数据, 需要指定IV(nonce+counter)数据
	DecryptIVCTR(ciphertext []byte, iv []byte) ([]byte, error)

	// EncryptGCM 使用GCM模式加密, 使用随机nonce，并把nonce数据放在密文头部
	EncryptGCM(plaintext []byte) ([]byte, error)
	// EncryptNonceGCM 使用GCM模式加密, 指定nonce
	EncryptNonceGCM(plaintext []byte, nonce []byte) ([]byte, error)
	// DecryptGCM 使用GCM模式解密，从密文头部提取nonce
	DecryptGCM(ciphertext []byte) ([]byte, error)
	// DecryptNonceGCM 使用GCM模式解密, 指定nonce
	DecryptNonceGCM(ciphertext []byte, nonce []byte) ([]byte, error)
}

func NewAESCipher(key []byte) IAESCipher {
	rs := &aesCipher{key: key, keyStr: string(key)}
	rs.block, rs.err = aes.NewCipher(key)
	rs.padding, rs.unPadding = cryptox.PKCS7Padding, cryptox.PKCS7UnPadding
	return rs
}

type aesCipher struct {
	key       []byte
	keyStr    string
	padding   cryptox.FuncPadding
	unPadding cryptox.FuncUnPadding
	block     cipher.Block
	err       error
}

func (o *aesCipher) Key() []byte {
	return []byte(o.keyStr)
}

func (o *aesCipher) SetPaddingFunc(padding cryptox.FuncPadding, unPadding cryptox.FuncUnPadding) {
	o.padding, o.unPadding = padding, unPadding
}

func (o *aesCipher) Encrypt(plaintext []byte) ([]byte, error) {
	return o.EncryptGCM(plaintext)
	//return o.EncryptCTR(plaintext)
}

func (o *aesCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	return o.DecryptGCM(ciphertext)
	//return o.DecryptCTR(ciphertext)
}

func (o *aesCipher) EncryptMode(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error) {
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

func (o *aesCipher) EncryptIVMode(plaintext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.CBC:
		return o.EncryptIVCBC(plaintext, iv)
	case cryptox.CTR:
		return o.EncryptIVCTR(plaintext, iv)
	case cryptox.GCM:
		return o.EncryptNonceGCM(plaintext, iv)
	default:
		return nil, errors.New("Unsupported AES block mode! ")
	}
}

func (o *aesCipher) DecryptMode(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error) {
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

func (o *aesCipher) DecryptIVMode(ciphertext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.CBC:
		return o.DecryptIVCBC(ciphertext, iv)
	case cryptox.CTR:
		return o.DecryptIVCTR(ciphertext, iv)
	case cryptox.GCM:
		return o.DecryptNonceGCM(ciphertext, iv)
	default:
		return nil, errors.New("Unsupported AES block mode! ")
	}
}

// CBC ---------- ---------- ---------- ---------- ----------

func (o *aesCipher) EncryptCBC(plaintext []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	ciphertext, err := o.EncryptIVCBC(plaintext, iv)
	if nil != err {
		return nil, err
	}
	return append(iv, ciphertext...), nil
}
func (o *aesCipher) EncryptIVCBC(plaintext []byte, iv []byte) ([]byte, error) {
	if len(iv) != aes.BlockSize {
		return nil, errors.New("iv length must be equal to block size! ")
	}
	padded, err := o.padding(plaintext, aes.BlockSize)
	if nil != err {
		return nil, err
	}
	ciphertext := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(o.block, iv)
	mode.CryptBlocks(ciphertext, padded)
	return ciphertext, nil
}

func (o *aesCipher) DecryptCBC(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	data := ciphertext[aes.BlockSize:]
	return o.DecryptIVCBC(data, iv)
}

func (o *aesCipher) DecryptIVCBC(ciphertext []byte, iv []byte) ([]byte, error) {
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext not a multiple of block size")
	}
	mode := cipher.NewCBCDecrypter(o.block, iv)
	plainPadded := make([]byte, len(ciphertext))
	mode.CryptBlocks(plainPadded, ciphertext)
	return o.unPadding(plainPadded)
}

// CTR ---------- ---------- ---------- ---------- ----------

var ctrCounter = []byte{0, 0, 0, 1}

func (o *aesCipher) EncryptCTR(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, aes.BlockSize-4)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	iv := append(nonce, ctrCounter...)
	ciphertext, _ := o.EncryptIVCTR(plaintext, iv)
	return append(iv, ciphertext...), nil
}

func (o *aesCipher) EncryptNonceCTR(plaintext []byte, nonce []byte, counter []byte) ([]byte, error) {
	iv := append(nonce, counter...)
	return o.EncryptIVCTR(plaintext, iv)
}

func (o *aesCipher) EncryptIVCTR(plaintext []byte, iv []byte) ([]byte, error) {
	if len(iv) != aes.BlockSize {
		return nil, errors.New("iv length must be equal to block size! ")
	}
	stream := cipher.NewCTR(o.block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

func (o *aesCipher) DecryptCTR(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	data := ciphertext[aes.BlockSize:]
	plaintext, _ := o.DecryptIVCTR(data, iv)
	return plaintext, nil
}

func (o *aesCipher) DecryptNonceCTR(ciphertext []byte, nonce []byte, counter []byte) ([]byte, error) {
	iv := append(nonce, counter...)
	return o.DecryptIVCTR(ciphertext, iv)
}

func (o *aesCipher) DecryptIVCTR(ciphertext []byte, iv []byte) ([]byte, error) {
	if len(iv) != aes.BlockSize {
		return nil, errors.New("iv length must be equal to block size! ")
	}
	stream := cipher.NewCTR(o.block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// GCM ---------- ---------- ---------- ---------- ----------

const (
	GcmStdNonceSize = 12
)

func (o *aesCipher) EncryptGCM(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, GcmStdNonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext, err := o.EncryptNonceGCM(plaintext, nonce)
	if nil != err {
		return nil, err
	}
	return append(nonce, ciphertext...), nil
}
func (o *aesCipher) EncryptNonceGCM(plaintext []byte, nonce []byte) ([]byte, error) {
	aead, err := cipher.NewGCM(o.block)
	if err != nil {
		return nil, err
	}
	if len(nonce) != aead.NonceSize() {
		return nil, errors.New("nonce length must be equal to GcmStdNonceSize! ")
	}
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func (o *aesCipher) DecryptGCM(ciphertext []byte) ([]byte, error) {
	nonceSize := GcmStdNonceSize
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce := ciphertext[:nonceSize]
	data := ciphertext[nonceSize:]
	return o.DecryptNonceGCM(data, nonce)
}
func (o *aesCipher) DecryptNonceGCM(ciphertext []byte, nonce []byte) ([]byte, error) {
	if len(nonce) != GcmStdNonceSize {
		return nil, errors.New("nonce length must be equal to GcmStdNonceSize! ")
	}
	aead, err := cipher.NewGCM(o.block)
	if err != nil {
		return nil, err
	}
	return aead.Open(nil, nonce, ciphertext, nil)
}
