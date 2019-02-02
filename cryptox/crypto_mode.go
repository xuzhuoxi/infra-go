//
//Created by xuzhuoxi
//on 2019-02-03.
//@author xuzhuoxi
//
package cryptox

import (
	"crypto/cipher"
	"strconv"
)

//ECB(electronic code book)是最简单的方式，它将明文分组加密后的结果直接成为密文分组。
//优缺点：模式操作简单；明文中的重复内容将在密文中表现出来，特别对于图像数据和明文变化较少的数据；适于短报文的加密传递。
//
//CBC(cipher block chaining)的原理是加密算法的输入是当前的明文分组和前一密文分组的异或，第一个明文分组和一个初始向量进行异或，这样同一个明文分组重复出现时会产生不同的密文分组。
//特点：同一个明文分组重复出现时产生不同的密文分组；加密函数的输入是当前的明文分组和前一个密文分组的异或；每个明文分组的加密函数的输入与明文分组之间不再有固定的关系；适合加密长消息。
type DESMode uint

const (
	ECB DESMode = 1 + iota
	CBC
	CTR
	OFB
	CFB
	maxMode
)

type newFunc func(block cipher.Block, iv []byte) cipher.BlockMode

var (
	encrypterHashes = make([]newFunc, maxMode)
	decrypterHashes = make([]newFunc, maxMode)
)

func (m DESMode) NewEncrypter(block cipher.Block, iv []byte) cipher.BlockMode {
	if m > 0 && m < maxMode {
		f := encrypterHashes[m]
		if f != nil {
			return f(block, iv)
		}
	}
	panic("des_mode: requested hash function #" + strconv.Itoa(int(m)) + " is unavailable")
}

func (m DESMode) NewDecrypter(block cipher.Block, iv []byte) cipher.BlockMode {
	if m > 0 && m < maxMode {
		f := decrypterHashes[m]
		if f != nil {
			return f(block, iv)
		}
	}
	panic("des_mode: requested hash function #" + strconv.Itoa(int(m)) + " is unavailable")
}

func (m DESMode) Available() bool {
	return m == ECB || m == CBC
}

func RegisterHash(m DESMode, ef newFunc, df newFunc) {
	if m >= maxMode {
		panic("des_mode: RegisterHash of unknown hash function")
	}
	encrypterHashes[m] = ef
	decrypterHashes[m] = df
}

func init() {
	RegisterHash(CBC, cipher.NewCBCEncrypter, cipher.NewCBCDecrypter)
}
