//
//Created by xuzhuoxi
//on 2019-02-03.
//@author xuzhuoxi
//
package cryptox

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
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
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func newRsa(public *rsa.PublicKey, private *rsa.PrivateKey) *rsaBase {
	return &rsaBase{publicKey: public, privateKey: private}
}

//加密
func (b *rsaBase) Encrypt(origData []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, b.publicKey, origData)
}

//解密
func (b *rsaBase) Decrypt(crypted []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, b.privateKey, crypted)
}

//签名
func (b *rsaBase) Sign(origData []byte, hash crypto.Hash) ([]byte, error) {
	hashData := hash.New().Sum(origData)
	signature, err := rsa.SignPKCS1v15(rand.Reader, b.privateKey, hash, hashData[:])
	return signature, err
}

//验签
func (b *rsaBase) VerySign(origData, signData []byte, hash crypto.Hash) error {
	hashData := hash.New().Sum(origData)
	return rsa.VerifyPKCS1v15(b.publicKey, hash, hashData, signData)
}

//Public-------------------------------------------

func NewRsaPublicCipher(pemPublic []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParsePublicKey(pemPublic)
	if nil != err {
		return nil, err
	}
	return newRsa(publicKey, nil), nil
}

func NewRsaPrivateCipher(pemPrivate []byte, pkcs8 bool) (IRSAPrivateCipher, error) {
	var privateKey *rsa.PrivateKey
	var err error
	if pkcs8 {
		privateKey, err = ParsePrivateKeyPKCS8(pemPrivate)
	} else {
		privateKey, err = ParsePrivateKeyPKCS1(pemPrivate)
	}
	if nil != err {
		return nil, err
	}
	return newRsa(nil, privateKey), nil
}

func ParsePublicKey(pemPublicKey []byte) (*rsa.PublicKey, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(pemPublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	return pub, nil
}
func ParsePrivateKeyPKCS1(pemPrivateKey []byte) (*rsa.PrivateKey, error) {
	//解密
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return priv, err
}
func ParsePrivateKeyPKCS8(pemPrivateKey []byte) (*rsa.PrivateKey, error) {
	//解密
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS8格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	return priv.(*rsa.PrivateKey), err
}

//PKCS#1：定义RSA公开密钥算法加密和签名机制，主要用于组织PKCS#7中所描述的数字签名和数字信封[22]。
//PKCS#3：定义Diffie-Hellman密钥交换协议[23]。
//PKCS#5：描述一种利用从口令派生出来的安全密钥加密字符串的方法。使用MD2或MD5 从口令中派生密钥，并采用DES-CBC模式加密。主要用于加密从一个计算机传送到另一个计算机的私人密钥，不能用于加密消息[24]。
//PKCS#6：描述了公钥证书的标准语法，主要描述X.509证书的扩展格式[25]。
//PKCS#7：定义一种通用的消息语法，包括数字签名和加密等用于增强的加密机制，PKCS#7与PEM兼容，所以不需其他密码操作，就可以将加密的消息转换成PEM消息[26]。
//PKCS#8：描述私有密钥信息格式，该信息包括公开密钥算法的私有密钥以及可选的属性集等[27]。
//PKCS#9：定义一些用于PKCS#6证书扩展、PKCS#7数字签名和PKCS#8私钥加密信息的属性类型[28]。
//PKCS#10：描述证书请求语法[29]。
//PKCS#11：称为Cyptoki，定义了一套独立于技术的程序设计接口，用于智能卡和PCMCIA卡之类的加密设备[30]。
//PKCS#12：描述个人信息交换语法标准。描述了将用户公钥、私钥、证书和其他相关信息打包的语法[31]。
//PKCS#13：椭圆曲线密码体制标准[32]。
//PKCS#14：伪随机数生成标准。
//PKCS#15：密码令牌信息格式标准[33]。
