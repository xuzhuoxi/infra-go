// Package symmetric
// Create on 2025/4/6
// @author xuzhuoxi
package symmetric

import (
	"crypto/cipher"
	"crypto/des"
	"errors"

	"github.com/xuzhuoxi/infra-go/cryptox"
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

	// EncryptMode 指定BlockMode加密, 使用随机IV，并把IV数据放在密文头部
	EncryptMode(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// EncryptIVMode 指定BlockMode加密, 指定IV
	EncryptIVMode(plaintext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// DecryptMode 指定BlockMode解密, 从密文头部提取IV
	DecryptMode(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// DecryptIVMode 指定BlockMode解密, 指定IV
	DecryptIVMode(ciphertext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error)

	// EncryptECB 使用ECB模式加密
	EncryptECB(plaintext []byte) ([]byte, error)
	// DecryptECB 使用ECB模式解密
	DecryptECB(ciphertext []byte) ([]byte, error)

	// EncryptCBC 使用CBC模式加密, 使用随机IV，并把IV数据放在密文头部
	EncryptCBC(plaintext []byte) ([]byte, error)
	// EncryptIVCBC 使用CBC模式加密,  指定IV
	EncryptIVCBC(plaintext []byte, iv []byte) ([]byte, error)
	// DecryptCBC 使用CBC模式解密, 从密文头部提取IV
	DecryptCBC(ciphertext []byte) ([]byte, error)
	// DecryptIVCBC 使用CBC模式解密, 指定IV
	DecryptIVCBC(ciphertext []byte, iv []byte) ([]byte, error)

	// EncryptCTR 使用CTR模式加密, 使用随机IV，并把IV数据放在密文头部
	EncryptCTR(plaintext []byte) ([]byte, error)
	// EncryptIVCTR 使用CTR模式解密, 指定IV
	EncryptIVCTR(plaintext []byte, iv []byte) ([]byte, error)
	// DecryptCTR 使用CTR模式解密, 从密文头部提取IV
	DecryptCTR(ciphertext []byte) ([]byte, error)
	// DecryptIVCTR 使用CTR模式解密, 指定IV
	DecryptIVCTR(ciphertext []byte, iv []byte) ([]byte, error)
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
func NewTripleDESCipher(key []byte) IDESCipher {
	rs := &desCipher{key: key}
	keyLen := len(key)
	var newKey []byte
	if keyLen == des.BlockSize {
		newKey = append(append(key, key...), key...)
	} else if keyLen == 16 {
		newKey = append(key[:], key[:8]...)
	} else if keyLen == 24 {
		newKey = key
	} else {
		rs.err = errors.New("key must be 8 (DES) or 24 (3DES) bytes")
		return rs
	}
	rs.block, rs.err = des.NewTripleDESCipher(newKey)
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
func (o *desCipher) EncryptIVMode(plaintext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.ECB:
		return o.EncryptECB(plaintext)
	case cryptox.CBC:
		return o.EncryptIVCBC(plaintext, iv)
	case cryptox.CTR:
		return o.EncryptIVCTR(plaintext, iv)
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

func (o *desCipher) DecryptIVMode(ciphertext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error) {
	if nil != o.err {
		return nil, o.err
	}
	switch blockMode {
	case cryptox.ECB:
		return o.DecryptECB(ciphertext)
	case cryptox.CBC:
		return o.DecryptIVCBC(ciphertext, iv)
	case cryptox.CTR:
		return o.DecryptIVCTR(ciphertext, iv)
	default:
		return nil, errors.New("Unsupported DES block mode! ")
	}
}

func (o *desCipher) EncryptECB(plaintext []byte) ([]byte, error) {
	blockSize := des.BlockSize
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
	blockSize := des.BlockSize
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("invalid ciphertext length")
	}
	plain := make([]byte, len(ciphertext))
	for bs, be := 0, blockSize; bs < len(ciphertext); bs, be = bs+blockSize, be+blockSize {
		o.block.Decrypt(plain[bs:be], ciphertext[bs:be])
	}
	return o.unPadding(plain)
}

// CBC ---------- ---------- ---------- ---------- ----------

func (o *desCipher) EncryptCBC(plaintext []byte) ([]byte, error) {
	iv, e := cryptox.RandomIV(des.BlockSize)
	if nil != e {
		return nil, e
	}
	ciphertext, err := o.EncryptIVCBC(plaintext, iv)
	if nil != err {
		return nil, err
	}
	return append(iv, ciphertext...), nil
}
func (o *desCipher) EncryptIVCBC(plaintext []byte, iv []byte) ([]byte, error) {
	padded, err := o.padding(plaintext, des.BlockSize)
	if nil != err {
		return nil, err
	}
	ciphertext := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(o.block, iv)
	mode.CryptBlocks(ciphertext, padded)
	return ciphertext, nil
}

func (o *desCipher) DecryptCBC(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	data := ciphertext[des.BlockSize:]
	return o.DecryptIVCBC(data, iv)
}

func (o *desCipher) DecryptIVCBC(ciphertext []byte, iv []byte) ([]byte, error) {
	if len(iv) < des.BlockSize {
		return nil, errors.New("iv length must be equal to BlockSize! ")
	}
	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("invalid CBC ciphertext length")
	}
	mode := cipher.NewCBCDecrypter(o.block, iv)
	plain := make([]byte, len(ciphertext))
	mode.CryptBlocks(plain, ciphertext)
	return o.unPadding(plain)
}

// CTR ---------- ---------- ---------- ---------- ----------

func (o *desCipher) EncryptCTR(plaintext []byte) ([]byte, error) {
	iv, e := cryptox.RandomIV(des.BlockSize)
	if nil != e {
		return nil, e
	}
	ciphertext, err := o.EncryptIVCTR(plaintext, iv)
	if nil != err {
		return nil, err
	}
	return append(iv, ciphertext...), nil
}
func (o *desCipher) EncryptIVCTR(plaintext []byte, iv []byte) ([]byte, error) {
	if len(iv) != des.BlockSize {
		return nil, errors.New("iv length must be equal to BlockSize! ")
	}
	stream := cipher.NewCTR(o.block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

func (o *desCipher) DecryptCTR(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	data := ciphertext[des.BlockSize:]
	return o.DecryptIVCTR(data, iv)
}
func (o *desCipher) DecryptIVCTR(ciphertext []byte, iv []byte) ([]byte, error) {
	if len(iv) != des.BlockSize {
		return nil, errors.New("iv length must be equal to BlockSize! ")
	}
	stream := cipher.NewCTR(o.block, iv)
	plain := make([]byte, len(ciphertext))
	stream.XORKeyStream(plain, ciphertext)
	return plain, nil
}
