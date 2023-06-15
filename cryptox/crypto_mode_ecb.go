// Package cryptox
// Created by xuzhuoxi
// on 2019-02-03.
// @author xuzhuoxi
//
package cryptox

import (
	"crypto/cipher"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block, iv []byte) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

type ecbEncAble interface {
	NewECBEncrypter(iv []byte) cipher.BlockMode
}

func NewECBEncrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	if cbc, ok := b.(ecbEncAble); ok {
		return cbc.NewECBEncrypter(iv)
	}
	return (*ecbEncrypter)(newECB(b, iv))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	bs := x.blockSize
	block := x.b
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
}

type ecbDecrypter ecb

type ecbDecAble interface {
	NewECBDecrypter(iv []byte) cipher.BlockMode
}

func NewECBDecrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	if cbc, ok := b.(ecbDecAble); ok {
		return cbc.NewECBDecrypter(iv)
	}
	return (*ecbDecrypter)(newECB(b, iv))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	bs := x.blockSize
	block := x.b
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
}

func init() {
	RegisterHash(ECB, NewECBEncrypter, NewECBDecrypter)
}
