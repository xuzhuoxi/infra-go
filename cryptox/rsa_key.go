//
//Created by xuzhuoxi
//on 2019-02-04.
//@author xuzhuoxi
//
package cryptox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/ssh"
	"os"
)

//解释Key-------------------------------

func ParseRSAPublicKey(pemPublicKey []byte) (*rsa.PublicKey, error) {
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
func ParseRSAPrivateKeyPKCS1(pemPrivateKey []byte) (*rsa.PrivateKey, error) {
	//解密
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return priv, err
}
func ParseRSAPrivateKeyPKCS8(pemPrivateKey []byte) (*rsa.PrivateKey, error) {
	//解密
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS8格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	return priv.(*rsa.PrivateKey), err
}
func ParseSSHPublicKey(publicKey []byte) (ssh.PublicKey, error) {
	return ssh.ParsePublicKey(publicKey)
}
func RSAPrivatePkcs1ToPkcs8(pkcs1PemFileData []byte) ([]byte, error) {
	privateKey, err := ParseRSAPrivateKeyPKCS1(pkcs1PemFileData)
	if nil != err {
		return nil, err
	}
	pkcs8, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if nil != err {
		return nil, err
	}
	return pkcs8, nil
}
func RSAPrivatePkcs8ToPkcs1(pkcs8PemFileData []byte) ([]byte, error) {
	privateKey, err := ParseRSAPrivateKeyPKCS8(pkcs8PemFileData)
	if nil != err {
		return nil, err
	}
	return x509.MarshalPKCS1PrivateKey(privateKey), nil
}

//生成Key----------------------------

//bits:密钥长度
func GenRSAKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if nil != err {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}
func GenRSAKeyBlock(bits int) (private *pem.Block, public *pem.Block, err error) {
	privateKey, publicKey, e := GenRSAKey(bits)
	if nil != e {
		err = e
		return
	}
	privateDer := x509.MarshalPKCS1PrivateKey(privateKey)
	publicDer, e := x509.MarshalPKIXPublicKey(publicKey)
	if nil != e {
		err = e
		return
	}
	private = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateDer}
	public = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicDer}
	return
}
func EncodeRSAKey(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (private []byte, public []byte, err error) {
	if nil != privateKey {
		private, _ = EncodeRSAPrivateKeyPKCS1(privateKey)
	}
	if nil != publicKey {
		public, err = EncodeRSAPublicKey(publicKey)
	}
	return
}
func EncodeRSAPublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: publicBytes,
		Type:  "PUBLIC KEY",
	}), nil
}
func EncodeSSHPublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}
func EncodeRSAPrivateKeyPKCS1(private *rsa.PrivateKey) ([]byte, error) {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY"}), nil
}
func EncodeRSAPrivateKeyPKCS8(private *rsa.PrivateKey) ([]byte, error) {
	bytes, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
func GenRSAKeyFile(bits int, privateFilePath, publicFilePath string) error {
	privateBlock, publicBlock, err := GenRSAKeyBlock(bits)
	if nil != err {
		return err
	}
	writeBlock := func(block *pem.Block, filePath string) error {
		file, err := os.Create(filePath)
		if nil != err {
			return err
		}
		err = pem.Encode(file, block)
		if nil != err {
			return err
		}
		return nil
	}
	err = writeBlock(privateBlock, privateFilePath)
	if nil != err {
		return err
	}
	err = writeBlock(publicBlock, publicFilePath)
	if nil != err {
		return err
	}
	return nil
}
func GenSSHKeyFile(bits int, privateFilePath, publicFilePath string) error {
	private, public, err := GenRSAKey(bits)
	if nil != err {
		return err
	}
	privateStr, _ := EncodeRSAPrivateKeyPKCS1(private)
	publicStr, err := EncodeSSHPublicKey(public)
	if nil != err {
		return err
	}
	os.WriteFile(privateFilePath, privateStr, os.ModePerm)
	os.WriteFile(publicFilePath, publicStr, os.ModePerm)
	return nil
}
