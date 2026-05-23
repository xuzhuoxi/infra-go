package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/xuzhuoxi/infra-go/cryptox"
)

// IAESCipher is the Advanced Encryption Standard (Rijndael) cipher API.
// IAESCipher 是高级加密标准 AES（Rijndael）的加解密接口。
// Features: symmetric; key schedule expands into round subkeys.
// 特点：1. 对称加密；2. 一个 SK 扩展成多个子 SK，多轮加密。
type IAESCipher interface {
	cryptox.ICipher
	// Key returns a copy of the key (read-only).
	// Key 获取密钥，只读。
	Key() []byte
	// SetPaddingFunc sets padding and unpadding functions.
	// SetPaddingFunc 设置填充方式。
	SetPaddingFunc(padding cryptox.FuncPadding, unPadding cryptox.FuncUnPadding)

	// EncryptMode encrypts with blockMode using a random IV prepended to ciphertext.
	// EncryptMode 指定 BlockMode 加密，使用随机 IV，并把 IV 放在密文头部。
	EncryptMode(plaintext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// EncryptIVMode encrypts with blockMode using the given IV.
	// EncryptIVMode 指定 BlockMode 加密，使用指定 IV。
	EncryptIVMode(plaintext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// DecryptMode decrypts with blockMode; IV is taken from the ciphertext prefix.
	// DecryptMode 指定 BlockMode 解密，从密文头部提取 IV。
	DecryptMode(ciphertext []byte, blockMode cryptox.BlockMode) ([]byte, error)
	// DecryptIVMode decrypts with blockMode using the given IV.
	// DecryptIVMode 指定 BlockMode 解密，使用指定 IV。
	DecryptIVMode(ciphertext []byte, iv []byte, blockMode cryptox.BlockMode) ([]byte, error)

	// EncryptCBC encrypts in CBC mode with a random IV prepended.
	// EncryptCBC 使用 CBC 模式加密，随机 IV 放在密文头部。
	EncryptCBC(plaintext []byte) ([]byte, error)
	// EncryptIVCBC encrypts in CBC mode with the given IV.
	// EncryptIVCBC 使用 CBC 模式加密，指定 IV。
	EncryptIVCBC(plaintext []byte, iv []byte) ([]byte, error)
	// DecryptCBC decrypts in CBC mode; IV is taken from the ciphertext prefix.
	// DecryptCBC 使用 CBC 模式解密，从密文头部提取 IV。
	DecryptCBC(ciphertext []byte) ([]byte, error)
	// DecryptIVCBC decrypts in CBC mode with the given IV.
	// DecryptIVCBC 使用 CBC 模式解密，指定 IV。
	DecryptIVCBC(ciphertext []byte, iv []byte) ([]byte, error)

	// EncryptCTR encrypts in CTR mode with random nonce+counter IV prepended.
	// EncryptCTR 使用 CTR 模式加密，随机 IV（nonce+counter）放在密文头部。
	EncryptCTR(plaintext []byte) ([]byte, error)
	// EncryptNonceCTR encrypts in CTR with nonce and counter; ciphertext has no IV prefix.
	// EncryptNonceCTR 使用 CTR 模式加密，指定 Nonce 与 counter，结果不包含 IV。
	EncryptNonceCTR(plaintext []byte, nonce []byte, counter []byte) ([]byte, error)
	// EncryptIVCTR encrypts in CTR with the given IV; ciphertext has no IV prefix.
	// EncryptIVCTR 使用 CTR 模式加密，指定 IV（nonce+counter），结果不包含 IV。
	EncryptIVCTR(plaintext []byte, iv []byte) ([]byte, error)
	// DecryptCTR decrypts in CTR mode; IV is taken from the ciphertext prefix.
	// DecryptCTR 使用 CTR 模式解密，从密文头部提取 IV。
	DecryptCTR(ciphertext []byte) ([]byte, error)
	// DecryptNonceCTR decrypts in CTR with external nonce and counter (no IV in ciphertext).
	// DecryptNonceCTR 使用 CTR 模式解密，密文不含 IV，需指定 Nonce 与 counter。
	DecryptNonceCTR(ciphertext []byte, nonce []byte, counter []byte) ([]byte, error)
	// DecryptIVCTR decrypts in CTR with the given IV (no IV in ciphertext).
	// DecryptIVCTR 使用 CTR 模式解密，密文不含 IV，需指定 IV（nonce+counter）。
	DecryptIVCTR(ciphertext []byte, iv []byte) ([]byte, error)

	// EncryptGCM encrypts in GCM with a random nonce prepended.
	// EncryptGCM 使用 GCM 模式加密，随机 nonce 放在密文头部。
	EncryptGCM(plaintext []byte) ([]byte, error)
	// EncryptNonceGCM encrypts in GCM with the given nonce.
	// EncryptNonceGCM 使用 GCM 模式加密，指定 nonce。
	EncryptNonceGCM(plaintext []byte, nonce []byte) ([]byte, error)
	// DecryptGCM decrypts in GCM; nonce is taken from the ciphertext prefix.
	// DecryptGCM 使用 GCM 模式解密，从密文头部提取 nonce。
	DecryptGCM(ciphertext []byte) ([]byte, error)
	// DecryptNonceGCM decrypts in GCM with the given nonce.
	// DecryptNonceGCM 使用 GCM 模式解密，指定 nonce。
	DecryptNonceGCM(ciphertext []byte, nonce []byte) ([]byte, error)
}

// NewAESCipher creates an AES cipher; default padding is PKCS#7 and default mode is GCM.
// NewAESCipher 创建 AES 加解密器；默认填充为 PKCS#7，默认加解密模式为 GCM。
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
	// GcmStdNonceSize is the standard GCM nonce size in bytes (12).
	// GcmStdNonceSize 为标准 GCM nonce 长度（12 字节）。
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
