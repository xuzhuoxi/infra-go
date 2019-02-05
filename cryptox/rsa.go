//
//Created by xuzhuoxi
//on 2019-02-03.
//@author xuzhuoxi
//
package cryptox

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

type IRSAPublicCipher interface {
	Encrypt(origData []byte) ([]byte, error)
	VerySign(origData, signData []byte, hash crypto.Hash) error
}

type IRSAPrivateCipher interface {
	Decrypt(crypted []byte) ([]byte, error)
	Sign(origData []byte, hash crypto.Hash) ([]byte, error)
}

type IRSACipher interface {
	IRSAPublicCipher
	IRSAPrivateCipher
}

type rsaBase struct {
	publicKey      *rsa.PublicKey
	privateKey     *rsa.PrivateKey
	encryptPartLen int
	decryptPartLen int
}

func newRsa(public *rsa.PublicKey, private *rsa.PrivateKey) *rsaBase {
	if nil == public {
		public = &private.PublicKey
	}
	encryptPartLen := public.N.BitLen()/8 - 11
	decryptPartLen := public.N.BitLen() / 8
	return &rsaBase{publicKey: public, privateKey: private, encryptPartLen: encryptPartLen, decryptPartLen: decryptPartLen}
}

//加密
func (b *rsaBase) Encrypt(origData []byte) ([]byte, error) {
	if len(origData) < b.encryptPartLen {
		return rsa.EncryptPKCS1v15(rand.Reader, b.publicKey, origData)
	}
	chunks := splitGroup(origData, b.encryptPartLen)
	buff := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, b.publicKey, chunk)
		if nil != err {
			return nil, err
		}
		buff.Write(bytes)
	}
	return buff.Bytes(), nil
}

//解密
func (b *rsaBase) Decrypt(crypted []byte) ([]byte, error) {
	if len(crypted) < b.decryptPartLen {
		return rsa.EncryptPKCS1v15(rand.Reader, b.publicKey, crypted)
	}
	chunks := splitGroup(crypted, b.decryptPartLen)
	buff := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		bytes, err := rsa.DecryptPKCS1v15(rand.Reader, b.privateKey, chunk)
		if nil != err {
			return nil, err
		}
		buff.Write(bytes)
	}
	return buff.Bytes(), nil
}

//签名
func (b *rsaBase) Sign(origData []byte, hash crypto.Hash) ([]byte, error) {
	hashed := Hash(hash, origData)
	fmt.Println(1111, hashed, len(hashed), hash.Size())
	signature, err := rsa.SignPKCS1v15(rand.Reader, b.privateKey, hash, hashed)
	return signature, err
}

//验签
func (b *rsaBase) VerySign(origData, signData []byte, hash crypto.Hash) error {
	hashed := Hash(hash, origData)
	return rsa.VerifyPKCS1v15(b.publicKey, hash, hashed, signData)
}

func splitGroup(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

//Public-------------------------------------------

func NewRsaPublicCipher(pemPublic []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParseRSAPublicKey(pemPublic)
	if nil != err {
		return nil, err
	}
	return newRsa(publicKey, nil), nil
}

func NewRsaPrivateCipher(pemPrivate []byte, pkcs8 bool) (IRSAPrivateCipher, error) {
	var privateKey *rsa.PrivateKey
	var err error
	if pkcs8 {
		privateKey, err = ParseRSAPrivateKeyPKCS8(pemPrivate)
	} else {
		privateKey, err = ParseRSAPrivateKeyPKCS1(pemPrivate)
	}
	if nil != err {
		return nil, err
	}
	return newRsa(nil, privateKey), nil
}
