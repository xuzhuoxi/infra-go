//
//Created by xuzhuoxi
//on 2019-02-03.
//@author xuzhuoxi
//
package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

type IESCipher interface {
	Encrypt(origData []byte) ([]byte, error)
	Decrypt(crypted []byte) ([]byte, error)
}

func NewDESCipher(key, iv [des.BlockSize]byte, mode DESMode, padding FuncPadding, unPadding FuncUnPadding) IESCipher {
	rs := &cipherBase{KEY: key[:], IV: iv[:], Mode: mode,
		Padding: padding, UnPadding: unPadding,
		blockNew: des.NewCipher}
	rs.InitBlock(key[:])
	return rs
}

func NewTripleDESCipher(key [des.BlockSize * 3]byte, iv [des.BlockSize]byte, mode DESMode, padding FuncPadding, unPadding FuncUnPadding) IESCipher {
	rs := &cipherBase{KEY: key[:], IV: iv[:], Mode: mode,
		Padding: padding, UnPadding: unPadding,
		blockNew: des.NewTripleDESCipher}
	rs.InitBlock(key[:])
	return rs
}

func NewAESCipher(key, iv [aes.BlockSize]byte, mode DESMode, padding FuncPadding, unPadding FuncUnPadding) IESCipher {
	rs := &cipherBase{KEY: key[:], IV: iv[:], Mode: mode,
		Padding: padding, UnPadding: unPadding,
		blockNew: aes.NewCipher}
	rs.InitBlock(key[:])
	return rs
}

type cipherBase struct {
	KEY       []byte
	IV        []byte
	Mode      DESMode
	Padding   FuncPadding
	UnPadding FuncUnPadding

	block    cipher.Block
	blockNew func(key []byte) (cipher.Block, error)
}

func (c *cipherBase) InitBlock(key []byte) error {
	block, err := c.blockNew(key)
	if err != nil {
		return err
	}
	c.block = block
	return nil
}

func (c *cipherBase) Encrypt(origData []byte) ([]byte, error) {
	if nil != c.Padding {
		origData = c.Padding(origData, c.block.BlockSize())
	}
	mode := c.Mode.NewEncrypter(c.block, c.IV[:])
	crypted := make([]byte, len(origData))
	mode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (c *cipherBase) Decrypt(crypted []byte) ([]byte, error) {
	blockMode := c.Mode.NewDecrypter(c.block, c.IV[:])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	if nil != c.UnPadding {
		origData = c.UnPadding(origData)
	}
	return origData, nil
}
