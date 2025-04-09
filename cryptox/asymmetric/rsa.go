// Package asymmetric
// Created by xuzhuoxi
// on 2019-02-03.
// @author xuzhuoxi
//
package asymmetric

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/xuzhuoxi/infra-go/cryptox"
	"github.com/xuzhuoxi/infra-go/filex"
	"golang.org/x/crypto/ssh"
	"os"
)

type IRSAPrivateCipher interface {
	PrivateKey() *rsa.PrivateKey
	// Decrypt 解密
	Decrypt(ciphertext []byte) ([]byte, error)
	// DecryptHybrid 混合解密
	DecryptHybrid(encryptedKey []byte, ciphertext []byte) (plaintext []byte, err error)
	// Sign 签名，使用sha256
	Sign(origData []byte) ([]byte, error)
	// SignHash 指定Hash算法进行签名
	// 注意：MD5SHA1不可用
	// 各个Hash算法，可能要引用指定的包，详细请看crypto/crypto.go中的常量定义说明
	SignHash(origData []byte, hash crypto.Hash) ([]byte, error)
	// SignBase64 Base64编码签名
	SignBase64(origData []byte, hash crypto.Hash) (string, error)
}

type IRSAPublicCipher interface {
	PublicKey() *rsa.PublicKey
	// Encrypt 加密(支持分组)
	Encrypt(plaintext []byte) ([]byte, error)
	// EncryptHybrid 混合加密
	EncryptHybrid(plaintext []byte) (encryptedKey []byte, ciphertext []byte, err error)
	// VerySign 验签，使用sha256
	VerySign(origData, signedData []byte) (bool, error)
	// VerySignHash 指定Hash算法验签
	// 注意：MD5SHA1不可用
	// 各个Hash算法，可能要引用指定的包，详细请看crypto/crypto.go中的常量定义说明
	VerySignHash(origData, signedData []byte, hash crypto.Hash) (bool, error)
	// VerifyBase64 Base64编译验签
	VerifyBase64(origData []byte, signedBase64Data string, hash crypto.Hash) (bool, error)
}

type IRSACipher interface {
	IRSAPrivateCipher
	IRSAPublicCipher
}

func newRsa(private *rsa.PrivateKey, public *rsa.PublicKey) *rsaCipher {
	return &rsaCipher{
		rsaPrivateCipher: *newPrivateRsa(private),
		rsaPublicCipher:  *newPublicRsa(public),
	}
}

func newPrivateRsa(private *rsa.PrivateKey) *rsaPrivateCipher {
	decryptPartLen := private.N.BitLen() / 8
	return &rsaPrivateCipher{PriKey: private, DecryptPartLen: decryptPartLen}
}

func newPublicRsa(public *rsa.PublicKey) *rsaPublicCipher {
	//RSA使用Pkcs进行padding，所以有11个byte被用于padding信息
	encryptPartLen := public.N.BitLen()/8 - 11
	return &rsaPublicCipher{PubKey: public, EncryptPartLen: encryptPartLen}
}

const (
	hybridAesKeyLen = 32
)

type rsaCipher struct {
	rsaPrivateCipher
	rsaPublicCipher
}

type rsaPrivateCipher struct {
	PriKey         *rsa.PrivateKey
	DecryptPartLen int
}

func (o *rsaPrivateCipher) PrivateKey() *rsa.PrivateKey {
	return o.PriKey
}

func (o *rsaPrivateCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < o.DecryptPartLen {
		return rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, ciphertext)
	}
	chunks := splitGroup(ciphertext, o.DecryptPartLen)
	buff := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		bytes, err := rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, chunk)
		if nil != err {
			return nil, err
		}
		buff.Write(bytes)
	}
	return buff.Bytes(), nil
}

// DecryptHybrid
// 步骤:
//  1. 不足分组长度，直接使用RSA解密, 忽略随机密钥
//  2. 长于分组长度的执行以下步骤
//   2.1 对AES密钥密文RSA解密，截取前32位得到AES密钥,余下为IV
//   2.2 使用AES密钥与IV，使用AES-CTR算法解密明文
//   2.3 返回明文
func (o *rsaPrivateCipher) DecryptHybrid(encryptedKey []byte, ciphertext []byte) (plaintext []byte, err error) {
	if len(ciphertext) < o.DecryptPartLen {
		return rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, ciphertext)
	}
	keyData, err := rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, encryptedKey)
	if err != nil {
		return nil, err
	}
	if len(keyData) < hybridAesKeyLen+aes.BlockSize {
		return nil, errors.New("invalid hybrid key data")
	}
	aesKey := keyData[:hybridAesKeyLen]
	iv := keyData[hybridAesKeyLen:]

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	plaintext = make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}
func (o *rsaPrivateCipher) Sign(origData []byte) ([]byte, error) {
	hashed := sha256.Sum256(origData)
	return rsa.SignPKCS1v15(rand.Reader, o.PriKey, crypto.SHA256, hashed[:])
}

func (o *rsaPrivateCipher) SignHash(origData []byte, hash crypto.Hash) ([]byte, error) {
	hashed := cryptox.Hash(hash, origData)
	if nil == hashed {
		return nil, cryptox.ErrUnsupportedHash
	}
	return rsa.SignPKCS1v15(rand.Reader, o.PriKey, hash, hashed)
}

func (o *rsaPrivateCipher) SignBase64(origData []byte, hash crypto.Hash) (string, error) {
	signed, err := o.SignHash(origData, hash)
	if nil != err {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signed), nil
}

func splitGroup(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

type rsaPublicCipher struct {
	PubKey         *rsa.PublicKey
	EncryptPartLen int
}

func (o *rsaPublicCipher) PublicKey() *rsa.PublicKey {
	return o.PubKey
}

func (o *rsaPublicCipher) Encrypt(plaintext []byte) ([]byte, error) {
	if len(plaintext) < o.EncryptPartLen {
		return rsa.EncryptPKCS1v15(rand.Reader, o.PubKey, plaintext)
	}
	chunks := splitGroup(plaintext, o.EncryptPartLen)
	buff := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, o.PubKey, chunk)
		if nil != err {
			return nil, err
		}
		buff.Write(bytes)
	}
	return buff.Bytes(), nil
}

// EncryptHybrid
// 步骤:
// 1. 不足分组长度，不生成随机AES密钥，直接使用RSA加密
// 2. 长于分组长度的执行以下步骤
//   2.1 生成32位随机AES密钥与16位随机IV，
//   2.2 使用RSA加密AES密钥与IV，得到AES密钥密文
//   2.3 使用AES密钥与IV，使用AES-CTR算法加密明文
//   2.4 返回(AES密钥密文, 密文)
func (o *rsaPublicCipher) EncryptHybrid(plaintext []byte) (encryptedKey []byte, ciphertext []byte, err error) {
	if len(plaintext) < o.EncryptPartLen {
		ciphertext, err = rsa.EncryptPKCS1v15(rand.Reader, o.PubKey, plaintext)
		return
	}
	aesKey := make([]byte, hybridAesKeyLen) // AES-256
	_, err = rand.Read(aesKey)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, nil, err
	}

	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return
	}

	ciphertext = make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	encryptedKey, err = rsa.EncryptPKCS1v15(rand.Reader, o.PubKey, append(aesKey, iv...))
	if err != nil {
		return nil, nil, err
	}

	return encryptedKey, ciphertext, nil
}

func (o *rsaPublicCipher) VerySign(origData, signedData []byte) (bool, error) {
	hashed := sha256.Sum256(origData)
	err := rsa.VerifyPKCS1v15(o.PubKey, crypto.SHA256, hashed[:], signedData)
	return nil == err, err
}

func (o *rsaPublicCipher) VerySignHash(origData, signedData []byte, hash crypto.Hash) (bool, error) {
	hashed := cryptox.Hash(hash, origData)
	if nil == hashed {
		return false, cryptox.ErrUnsupportedHash
	}
	err := rsa.VerifyPKCS1v15(o.PubKey, hash, hashed, signedData)
	return nil == err, err
}

func (o *rsaPublicCipher) VerifyBase64(origData []byte, signedBase64Data string, hash crypto.Hash) (bool, error) {
	sig, err := base64.StdEncoding.DecodeString(signedBase64Data)
	if err != nil {
		return false, err
	}
	return o.VerySignHash(origData, sig, hash)
}

// 生成Key ---------- ---------- ---------- ---------- ----------

// GenerateRSAKeyPair
// 生成RSA密钥对
// bits:密钥长度
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if nil != err {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// 编码为不懂格式的密钥数据 ---------- ---------- ---------- ---------- ----------

// Encode2RSAPrivatePKCS1
// 编码为 RSA专用私钥
// 执行 PKCS1 标准
// PEM类型为 PEMTypeRSAPrivateKey "RSA PRIVATE KEY"
func Encode2RSAPrivatePKCS1(private *rsa.PrivateKey) ([]byte, error) {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  cryptox.PEMTypeRSAPrivateKey}), nil
}

// Encode2RSAPublicPKCS1
// 编码为 RSA专用公钥
// 执行 PKCS1 标准
// PEM类型为 PEMTypeRSAPublicKey "RSA PUBLIC KEY"
func Encode2RSAPublicPKCS1(public *rsa.PublicKey) ([]byte, error) {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PublicKey(public),
		Type:  cryptox.PEMTypeRSAPublicKey}), nil
}

// Encode2PrivatePKCS8
// 编码为 PKCS8 私钥
// PEM类型为 PEMTypePKCS8PrivateKey "PRIVATE KEY"
func Encode2PrivatePKCS8(private *rsa.PrivateKey) ([]byte, error) {
	bytes, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: bytes,
		Type:  cryptox.PEMTypePKCS8PrivateKey,
	}), nil
}

// Encode2PublicX509
// 编码为 X.509 公钥
// 使用PEM类型为 PEMTypePKCS8PublicKey "PUBLIC KEY"
func Encode2PublicX509(public *rsa.PublicKey) ([]byte, error) {
	bytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: bytes,
		Type:  cryptox.PEMTypeX509PublicKey,
	}), nil
}

// Encode2PublicOpenSSH
// 编码为 OpenSSL 公钥
// 并非Pem标准
func Encode2PublicOpenSSH(public *rsa.PublicKey) ([]byte, error) {
	sshPublicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(sshPublicKey), nil
}

// Encode2PrivateOpenSSH
// go 1.16 官方库不支持直接导出
func Encode2PrivateOpenSSH(private *rsa.PrivateKey) ([]byte, error) {
	return nil, nil
}

type funcEncode2KeyPair = func(privateKey *rsa.PrivateKey) (private []byte, public []byte, err error)

// Encode2OpenSSHPair
// 编码为 OpenSSH适用的公私钥数据
// 私钥格式: RSA私钥
// 公钥格式: OpenSSH格式
func Encode2OpenSSHPair(privateKey *rsa.PrivateKey) (private []byte, public []byte, err error) {
	private, err = Encode2RSAPrivatePKCS1(privateKey)
	if err != nil {
		return nil, nil, err
	}
	public, err = Encode2PublicOpenSSH(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return
}

// Encode2RSAPair
// 编码为 RSA专用的公私钥数据
// 私钥格式: RSA私钥
// 公钥格式: RSA公钥
func Encode2RSAPair(privateKey *rsa.PrivateKey) (private []byte, public []byte, err error) {
	private, err = Encode2RSAPrivatePKCS1(privateKey)
	if err != nil {
		return nil, nil, err
	}
	public, err = Encode2RSAPublicPKCS1(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return
}

// Encode2PKCS8Pair
// 编码为 PEM通用的公私钥数据
// 私钥格式: PKCS8私钥
// 公钥格式: X509公钥
func Encode2PKCS8Pair(privateKey *rsa.PrivateKey) (private []byte, public []byte, err error) {
	private, err = Encode2PrivatePKCS8(privateKey)
	if err != nil {
		return nil, nil, err
	}
	public, err = Encode2PublicX509(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return
}

// 生成不同格式的密钥文件 ---------- ---------- ---------- ---------- ----------

// GenerateOpenSSHKeyFiles
// 生成OpenSSH格式的密钥对文件
// bits:密钥长度
// privateFilePath, publicFilePath: 私钥、公钥文件路径
// 私钥文件格式: PKCS1格式
// 公钥文件格式: OpenSSH格式
func GenerateOpenSSHKeyFiles(bits int, privateFilePath, publicFilePath string) error {
	return generateKeyFiles(bits, Encode2OpenSSHPair, privateFilePath, publicFilePath)
}

// SaveOpenSSHKeyFiles
// 保存OpenSSH格式的密钥对文件
// privateKey: 私钥
// privateFilePath, publicFilePath: 私钥、公钥文件路径
// 私钥文件格式: PKCS1格式
// 公钥文件格式: OpenSSH格式
func SaveOpenSSHKeyFiles(privateKey *rsa.PrivateKey, privateFilePath, publicFilePath string) error {
	return saveKeyFiles(privateKey, Encode2OpenSSHPair, privateFilePath, publicFilePath)
}

// GenerateRSAKeyFiles
// 生成RSA专用的密钥对文件
// bits:密钥长度
// privateFilePath, publicFilePath: 私钥、公钥文件路径
// 私钥文件格式: PKCS1格式
// 公钥文件格式: PKCS1格式
func GenerateRSAKeyFiles(bits int, privateFilePath, publicFilePath string) error {
	return generateKeyFiles(bits, Encode2RSAPair, privateFilePath, publicFilePath)
}

// SaveRSAKeyFiles
// 保存RSA专用的密钥对文件
// privateKey: 私钥
// privateFilePath, publicFilePath: 私钥、公钥文件路径
// 私钥文件格式: PKCS1格式
// 公钥文件格式: PKCS1格式
func SaveRSAKeyFiles(privateKey *rsa.PrivateKey, privateFilePath, publicFilePath string) error {
	return saveKeyFiles(privateKey, Encode2RSAPair, privateFilePath, publicFilePath)
}

// GeneratePemKeyFiles
// 生成PEM通用的密钥对文件
// bits:密钥长度
// privateFilePath, publicFilePath: 私钥、公钥文件路径
// 私钥文件格式: PKCS8格式
// 公钥文件格式: X509格式
func GeneratePemKeyFiles(bits int, privateFilePath, publicFilePath string) error {
	return generateKeyFiles(bits, Encode2PKCS8Pair, privateFilePath, publicFilePath)
}

// SavePemKeyFiles
// 保存PEM通用的密钥对文件
// privateKey: 私钥
// privateFilePath, publicFilePath: 私钥、公钥文件路径
// 私钥文件格式: PKCS8格式
// 公钥文件格式: X509格式
func SavePemKeyFiles(privateKey *rsa.PrivateKey, privateFilePath, publicFilePath string) error {
	return saveKeyFiles(privateKey, Encode2PKCS8Pair, privateFilePath, publicFilePath)
}

func generateKeyFiles(bits int, funcEncode funcEncode2KeyPair, privateFilePath, publicFilePath string) error {
	privateKey, _, err := GenerateRSAKeyPair(bits)
	if nil != err {
		return err
	}
	return saveKeyFiles(privateKey, funcEncode, privateFilePath, publicFilePath)
}

func saveKeyFiles(private *rsa.PrivateKey, funcEncode funcEncode2KeyPair, privateFilePath, publicFilePath string) error {
	privateDer, publicDer, err := funcEncode(private)
	if nil != err {
		return err
	}
	filex.WriteFile(privateFilePath, privateDer, os.ModePerm)
	filex.WriteFile(publicFilePath, publicDer, os.ModePerm)
	return nil
}

// load public keys ---------- ---------- ---------- ---------- ----------

var (
	ErrKeyPath   = errors.New("path is not exist")
	ErrKeyFormat = errors.New("key content format error")
)

// LoadPublicCipherSSH
// 加载SSH专用公钥
// 公钥格式: SSH格式
func LoadPublicCipherSSH(path string) (IRSAPublicCipher, error) {
	if !filex.IsExist(path) {
		return nil, ErrKeyPath
	}
	data, err := os.ReadFile(path)
	if nil != err {
		return nil, err
	}
	return NewPublicCipherBySSH(data)
}

// LoadPrivateCipherRSA
// 加载RSA专用私钥
// 私钥格式: PKCS1格式
func LoadPrivateCipherRSA(path string) (IRSAPrivateCipher, error) {
	if !filex.IsExist(path) {
		return nil, ErrKeyPath
	}
	data, err := os.ReadFile(path)
	if nil != err {
		return nil, err
	}
	return NewPrivateCipherByRSA(data)
}

// LoadPublicCipherRSA
// 加载RSA专用公钥
// 公钥格式: PKCS1格式
func LoadPublicCipherRSA(path string) (IRSAPublicCipher, error) {
	if !filex.IsExist(path) {
		return nil, ErrKeyPath
	}
	data, err := os.ReadFile(path)
	if nil != err {
		return nil, err
	}
	return NewPublicCipherByRSA(data)
}

// LoadPrivateCipherPEM
// 加载PEM通用私钥
// 私钥格式: PKCS8格式
func LoadPrivateCipherPEM(path string) (IRSAPrivateCipher, error) {
	if !filex.IsExist(path) {
		return nil, ErrKeyPath
	}
	data, err := os.ReadFile(path)
	if nil != err {
		return nil, err
	}
	return NewPrivateCipherByPEM(data)
}

// LoadPublicCipherPEM
// 加载PEM通用公钥
// 公钥格式: X509格式
func LoadPublicCipherPEM(path string) (IRSAPublicCipher, error) {
	if !filex.IsExist(path) {
		return nil, ErrKeyPath
	}
	data, err := os.ReadFile(path)
	if nil != err {
		return nil, err
	}
	return NewPublicCipherByPEM(data)
}

// LoadCiphersSSH
// 加载SSH密钥对
// 私钥格式: PKCS1格式
// 公钥格式: SSH格式
func LoadCiphersSSH(privatePath string, publicPath string) (IRSAPrivateCipher, IRSAPublicCipher, error) {
	if !filex.IsExist(privatePath) || !filex.IsExist(publicPath) {
		return nil, nil, ErrKeyPath
	}
	privateData, err1 := os.ReadFile(privatePath)
	if nil != err1 {
		return nil, nil, err1
	}
	publicData, err2 := os.ReadFile(publicPath)
	if nil != err2 {
		return nil, nil, err2
	}
	return NewCiphersBySSH(privateData, publicData)
}

// LoadCiphersRSA
// 加载RSA专用密钥对
// 私钥格式: PKCS1格式
// 公钥格式: PKCS1格式
func LoadCiphersRSA(privatePath string, publicPath string) (IRSAPrivateCipher, IRSAPublicCipher, error) {
	if !filex.IsExist(privatePath) || !filex.IsExist(publicPath) {
		return nil, nil, ErrKeyPath
	}
	privateData, err1 := os.ReadFile(privatePath)
	if nil != err1 {
		return nil, nil, err1
	}
	publicData, err2 := os.ReadFile(publicPath)
	if nil != err2 {
		return nil, nil, err2
	}
	return NewCiphersByRSA(privateData, publicData)
}

// LoadCiphersPEM
// 加载PEM封装的通用密钥对
// 私钥格式: PKCS8格式
// 公钥格式: x.509格式
func LoadCiphersPEM(privatePath string, publicPath string) (IRSAPrivateCipher, IRSAPublicCipher, error) {
	if !filex.IsExist(privatePath) || !filex.IsExist(publicPath) {
		return nil, nil, ErrKeyPath
	}
	privateData, err1 := os.ReadFile(privatePath)
	if nil != err1 {
		return nil, nil, err1
	}
	publicData, err2 := os.ReadFile(publicPath)
	if nil != err2 {
		return nil, nil, err2
	}
	return NewCiphersByPEM(privateData, publicData)
}

// NewPrivateCipherByRSA
// 根据pkcs1(RSA专用)私钥数据生成私钥处理器
// publicPkcs1DerData PKCS1格式的私钥内容
func NewPrivateCipherByRSA(privatePkcs1DerData []byte) (IRSAPrivateCipher, error) {
	privateKey, err := ParsePrivateKeyFromPKCS1(privatePkcs1DerData)
	if nil != err {
		return nil, err
	}
	return newPrivateRsa(privateKey), nil
}

// NewPublicCipherByRSA
// 根据pkcs1(RSA专用)公钥数据生成公钥处理器
func NewPublicCipherByRSA(publicPkcs1DerData []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParsePublicKeyFromPKCS1(publicPkcs1DerData)
	if nil != err {
		return nil, err
	}
	return newPublicRsa(publicKey), nil
}

// NewPublicCipherBySSH
// 根据SSH公钥数据生成公钥处理器
func NewPublicCipherBySSH(publicSSHData []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParsePublicKeyFromSSH(publicSSHData)
	if nil != err {
		return nil, err
	}
	return newPublicRsa(publicKey), nil
}

// NewPrivateCipherByPEM
// 根据pkcs8私钥数据生成私钥处理器
func NewPrivateCipherByPEM(pkcs8DerData []byte) (IRSAPrivateCipher, error) {
	privateKey, err := ParsePrivateKeyFromPKCS8(pkcs8DerData)
	if nil != err {
		return nil, err
	}
	return newPrivateRsa(privateKey), nil
}

// NewPublicCipherByPEM
// 根据x.509公钥数据生成公钥处理器
func NewPublicCipherByPEM(x509DerData []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParsePublicKeyFromX509(x509DerData)
	if nil != err {
		return nil, err
	}
	return newPublicRsa(publicKey), nil
}

// NewCiphersBySSH
// 根据密钥对数据生成密钥对处理器
// 私钥：PKCS1格式
// 公钥：SSH格式
func NewCiphersBySSH(privatePkcs1DerData []byte, publicSSHData []byte) (IRSAPrivateCipher, IRSAPublicCipher, error) {
	privateKey, err := ParsePrivateKeyFromPKCS1(privatePkcs1DerData)
	if nil != err {
		return nil, nil, err
	}
	publicKey, err := ParsePublicKeyFromSSH(publicSSHData)
	if nil != err {
		return nil, nil, err
	}
	return newPrivateRsa(privateKey), newPublicRsa(publicKey), nil
}

// NewCiphersByRSA
// 根据密钥对数据生成密钥对处理器
// 私钥：PKCS1格式
// 公钥：PKCS1格式
func NewCiphersByRSA(privatePkcs1DerData []byte, publicPkcs1DerData []byte) (IRSAPrivateCipher, IRSAPublicCipher, error) {
	privateKey, err := ParsePrivateKeyFromPKCS1(privatePkcs1DerData)
	if nil != err {
		return nil, nil, err
	}
	publicKey, err := ParsePublicKeyFromPKCS1(publicPkcs1DerData)
	if nil != err {
		return nil, nil, err
	}
	return newPrivateRsa(privateKey), newPublicRsa(publicKey), nil
}

// NewCiphersByPEM
// 根据密钥对数据生成密钥对处理器
// 私钥：PKCS8格式
// 公钥：X509格式
func NewCiphersByPEM(pkcs8DerData []byte, x509DerData []byte) (IRSAPrivateCipher, IRSAPublicCipher, error) {
	privateKey, err := ParsePrivateKeyFromPKCS8(pkcs8DerData)
	if nil != err {
		return nil, nil, err
	}
	publicKey, err := ParsePublicKeyFromX509(x509DerData)
	if nil != err {
		return nil, nil, err
	}
	return newPrivateRsa(privateKey), newPublicRsa(publicKey), nil
}

// 解释Key ---------- ---------- ---------- ---------- ----------

// ParsePrivateKeyFromPKCS1
// 解释PKCS1(RSA专用)格式的私钥
// return *rsa.PrivateKey
func ParsePrivateKeyFromPKCS1(privatePkcs1DerData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privatePkcs1DerData)
	if block == nil {
		return nil, ErrKeyFormat
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// ParsePublicKeyFromPKCS1
// 解释PKCS1(RSA专用)格式的公钥
// return *rsa.PublicKey
func ParsePublicKeyFromPKCS1(publicPkcs1DerData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicPkcs1DerData)
	if block == nil {
		return nil, ErrKeyFormat
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// ParsePublicKeyFromSSH
// 把OpenSSH格式的公钥解释为rsa.PublicKey
// return *rsa.PublicKey
func ParsePublicKeyFromSSH(sshPublicData []byte) (*rsa.PublicKey, error) {
	// 注意，使用ParsePublicKey来解释的话，要手动去掉ssh-rsa后，使用base64进行解码后传入
	sshPublicKey, _, _, _, err := ssh.ParseAuthorizedKey(sshPublicData)
	if nil != err {
		return nil, err
	}
	rsaPublicKey, ok := sshPublicKey.(ssh.CryptoPublicKey).CryptoPublicKey().(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("ssh public key is not rsa")
	}
	return rsaPublicKey, nil
}

// ParsePrivateKeyFromPKCS8
// 把PKCS8格式的私钥解释为rsa.PrivateKey
func ParsePrivateKeyFromPKCS8(pkcs8DerData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pkcs8DerData)
	if block == nil {
		return nil, ErrKeyFormat
	}
	//解析PKCS8格式的私钥
	privateInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	private, ok := privateInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, ErrKeyFormat
	}
	return private, err
}

// ParsePublicKeyFromX509
// 把X509格式的公钥解释为rsa.PublicKey
func ParsePublicKeyFromX509(x509DerData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(x509DerData)
	if block == nil {
		return nil, ErrKeyFormat
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, ErrKeyFormat
	}
	return pub, nil
}

// ConvertPublicPkcs1ToX509
// 把 PKCS1格式的公钥内容 转换为 X509格式的公钥内容
func ConvertPublicPkcs1ToX509(pkcs1DerData []byte) ([]byte, error) {
	publicKey, err := ParsePublicKeyFromPKCS1(pkcs1DerData)
	if nil != err {
		return nil, err
	}
	return Encode2PublicX509(publicKey)
}

// ConvertPublicX509ToPkcs1
// 把 X509格式的公钥内容 转换为 PKCS1格式的公钥内容
func ConvertPublicX509ToPkcs1(x509DerData []byte) ([]byte, error) {
	publicKey, err := ParsePublicKeyFromX509(x509DerData)
	if nil != err {
		return nil, err
	}
	return Encode2PublicX509(publicKey)
}

// ConvertPrivatePkcs1ToPkcs8
// 把 PKCS1格式的私钥内容 转换为 PKCS8格式的私钥内容
func ConvertPrivatePkcs1ToPkcs8(pkcs1DerData []byte) ([]byte, error) {
	privateKey, err := ParsePrivateKeyFromPKCS1(pkcs1DerData)
	if nil != err {
		return nil, ErrKeyFormat
	}
	return Encode2PrivatePKCS8(privateKey)
}

// ConvertPrivatePkcs8ToPkcs1
// 把 PKCS8格式的私钥内容 转换为 PKCS1格式的私钥内容
func ConvertPrivatePkcs8ToPkcs1(pkcs8DerData []byte) ([]byte, error) {
	privateKey, err := ParsePrivateKeyFromPKCS8(pkcs8DerData)
	if nil != err {
		return nil, ErrKeyFormat
	}
	return Encode2RSAPrivatePKCS1(privateKey)
}
