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
	"os"

	"github.com/xuzhuoxi/infra-go/cryptox"
	"github.com/xuzhuoxi/infra-go/filex"
	"golang.org/x/crypto/ssh"
)

// IRSAPrivateCipher defines RSA private-key decrypt and sign operations.
// IRSAPrivateCipher 定义 RSA 私钥解密与签名操作。
type IRSAPrivateCipher interface {
	// PrivateKey returns the underlying RSA private key.
	// PrivateKey 返回底层 RSA 私钥。
	PrivateKey() *rsa.PrivateKey
	// Decrypt decrypts ciphertext with PKCS#1 v15 (supports chunked input).
	// Decrypt 使用 PKCS#1 v15 解密（支持分块）。
	Decrypt(ciphertext []byte) ([]byte, error)
	// DecryptHybrid hybrid-decrypts: RSA for key material, AES-CTR for bulk data.
	// DecryptHybrid 混合解密：RSA 解密密钥材料，AES-CTR 解密主体数据。
	DecryptHybrid(ciphertext []byte) (plaintext []byte, err error)
	// Sign signs origData with SHA-256 and PKCS#1 v15.
	// Sign 使用 SHA-256 与 PKCS#1 v15 对 origData 签名。
	Sign(origData []byte) ([]byte, error)
	// SignBase64 signs with SHA-256 and returns a Base64-encoded signature.
	// SignBase64 使用 SHA-256 签名，并将结果编码为 Base64 字符串。
	SignBase64(origData []byte, base64Encoding *base64.Encoding) (string, error)
	// SignHash signs origData with the given hash (MD5SHA1 not supported).
	// SignHash 使用指定 Hash 算法签名（不支持 MD5SHA1）。
	// Some hash algorithms require importing their packages; see crypto.Hash constants.
	// 部分 Hash 算法需引入对应包，详见 crypto.Hash 常量说明。
	SignHash(origData []byte, hash crypto.Hash) ([]byte, error)
	// SignHashBase64 signs with the given hash and returns Base64-encoded signature.
	// SignHashBase64 使用指定 Hash 签名，并返回 Base64 编码结果。
	SignHashBase64(origData []byte, hash crypto.Hash, base64Encoding *base64.Encoding) (string, error)
}

// IRSAPublicCipher defines RSA public-key encrypt and verify operations.
// IRSAPublicCipher 定义 RSA 公钥加密与验签操作。
type IRSAPublicCipher interface {
	// PublicKey returns the underlying RSA public key.
	// PublicKey 返回底层 RSA 公钥。
	PublicKey() *rsa.PublicKey
	// Encrypt encrypts plaintext with PKCS#1 v15 (chunks when needed).
	// Encrypt 使用 PKCS#1 v15 加密（必要时按块加密）。
	Encrypt(plaintext []byte) ([]byte, error)
	// EncryptHybrid hybrid-encrypts: RSA for key material, AES-CTR for bulk data.
	// EncryptHybrid 混合加密：RSA 加密密钥材料，AES-CTR 加密主体数据。
	EncryptHybrid(plaintext []byte) (ciphertext []byte, err error)
	// VerifySign verifies a SHA-256 PKCS#1 v15 signature.
	// VerifySign 使用 SHA-256 验证 PKCS#1 v15 签名。
	VerifySign(origData, signature []byte) (bool, error)
	// VerifySignBase64 verifies a Base64-encoded SHA-256 signature.
	// VerifySignBase64 验签，签名为 Base64 编码字符串，使用 SHA-256。
	VerifySignBase64(origData []byte, base64Signature string, base64Encoding *base64.Encoding) (bool, error)
	// VerifySignHash verifies a signature for the given hash (MD5SHA1 not supported).
	// VerifySignHash 使用指定 Hash 算法验签（不支持 MD5SHA1）。
	VerifySignHash(origData, signature []byte, hash crypto.Hash) (bool, error)
	// VerifySignHashBase64 verifies a Base64-encoded signature for the given hash.
	// VerifySignHashBase64 验签，签名为 Base64 编码字符串。
	VerifySignHashBase64(origData []byte, base64Signature string, hash crypto.Hash, base64Encoding *base64.Encoding) (bool, error)
}

// IRSACipher combines private and public RSA cipher operations.
// IRSACipher 组合 RSA 私钥与公钥加解密、签名验签能力。
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
	//fmt.Println("Private: PartLen=", decryptPartLen)
	return &rsaPrivateCipher{PriKey: private, DecryptPartLen: decryptPartLen}
}

func newPublicRsa(public *rsa.PublicKey) *rsaPublicCipher {
	//RSA使用Pkcs进行padding，所以有11个byte被用于padding信息
	encryptPartLen := public.N.BitLen()/8 - 11
	//fmt.Println("Public: PartLen=", encryptPartLen)
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
	if len(ciphertext) == o.DecryptPartLen {
		return rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, ciphertext)
	}
	chunks := splitGroup(ciphertext, o.DecryptPartLen)
	buff := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		bs, err := rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, chunk)
		if nil != err {
			return nil, err
		}
		buff.Write(bs)
	}
	return buff.Bytes(), nil
}

// DecryptHybrid hybrid-decrypts per package scheme (see asymmetric doc).
// DecryptHybrid 混合解密。
// Steps:
//  1. If shorter than one RSA block, PKCS#1 decrypt only (no AES key).
//  2. Otherwise: RSA-decrypt leading block → 32-byte AES key + IV; AES-CTR decrypt remainder.
// 步骤:
//  1. 不足分组长度，直接使用 RSA 解密，忽略随机密钥。
//  2. 长于分组长度的执行以下步骤：
//   2.1 对 AES 密钥密文 RSA 解密，截取前 32 字节为 AES 密钥，余下为 IV。
//   2.2 使用 AES 密钥与 IV，AES-CTR 解密明文。
//   2.3 返回明文。
func (o *rsaPrivateCipher) DecryptHybrid(ciphertext []byte) (plaintext []byte, err error) {
	if len(ciphertext) < o.DecryptPartLen {
		return nil, errors.New("invalid rsa ciphertext data")
	}
	if len(ciphertext) == o.DecryptPartLen {
		return rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, ciphertext)
	}
	encryptedKey := ciphertext[:o.DecryptPartLen]
	ciphertext = ciphertext[o.DecryptPartLen:]
	encryptedKeyIV, err := rsa.DecryptPKCS1v15(rand.Reader, o.PriKey, encryptedKey)
	if err != nil {
		return nil, err
	}
	aesKey := encryptedKeyIV[:hybridAesKeyLen]
	iv := encryptedKeyIV[hybridAesKeyLen:]

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

func (o *rsaPrivateCipher) SignBase64(origData []byte, base64Encoding *base64.Encoding) (string, error) {
	signature, err := o.Sign(origData)
	if nil != err {
		return "", err
	}
	return base64Encoding.EncodeToString(signature), nil
}

func (o *rsaPrivateCipher) SignHash(origData []byte, hash crypto.Hash) ([]byte, error) {
	hashed := cryptox.Hash(hash, origData)
	if nil == hashed {
		return nil, cryptox.ErrUnsupportedHash
	}
	return rsa.SignPKCS1v15(rand.Reader, o.PriKey, hash, hashed)
}

func (o *rsaPrivateCipher) SignHashBase64(origData []byte, hash crypto.Hash, base64Encoding *base64.Encoding) (string, error) {
	signature, err := o.SignHash(origData, hash)
	if nil != err {
		return "", err
	}
	return base64Encoding.EncodeToString(signature), nil
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
		bs, err := rsa.EncryptPKCS1v15(rand.Reader, o.PubKey, chunk)
		if nil != err {
			return nil, err
		}
		buff.Write(bs)
	}
	return buff.Bytes(), nil
}

// EncryptHybrid hybrid-encrypts per package scheme (see asymmetric doc).
// EncryptHybrid 混合加密。
// Steps:
//  1. If plaintext fits one RSA block, PKCS#1 encrypt only.
//  2. Otherwise: random 32-byte AES key + 16-byte IV → RSA-encrypt key||IV → AES-CTR encrypt body → prepend RSA block.
// 步骤:
//  1. 不足分组长度，不生成随机 AES 密钥，直接使用 RSA 加密。
//  2. 长于分组长度的执行以下步骤：
//   2.1 生成 32 字节随机 AES 密钥与 16 字节随机 IV。
//   2.2 使用 RSA 加密 AES 密钥与 IV，得到 AES 密钥密文。
//   2.3 使用 AES 密钥与 IV，AES-CTR 加密明文。
//   2.4 返回（AES 密钥密文 || 密文）。
func (o *rsaPublicCipher) EncryptHybrid(plaintext []byte) (ciphertext []byte, err error) {
	if len(plaintext) <= o.EncryptPartLen {
		ciphertext, err = rsa.EncryptPKCS1v15(rand.Reader, o.PubKey, plaintext)
		return
	}
	aesKeyIV := make([]byte, hybridAesKeyLen+aes.BlockSize) // AES-Key + AES-IV
	_, err = rand.Read(aesKeyIV)
	if err != nil {
		return
	}
	aesKey := aesKeyIV[:hybridAesKeyLen]
	iv := aesKeyIV[hybridAesKeyLen:]

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	ciphertext = make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	encryptedKey, err := rsa.EncryptPKCS1v15(rand.Reader, o.PubKey, append(aesKey, iv...))
	if err != nil {
		return nil, err
	}

	return append(encryptedKey, ciphertext...), nil
}

func (o *rsaPublicCipher) VerifySign(origData, signature []byte) (bool, error) {
	hashed := sha256.Sum256(origData)
	err := rsa.VerifyPKCS1v15(o.PubKey, crypto.SHA256, hashed[:], signature)
	return nil == err, err
}

func (o *rsaPublicCipher) VerifySignBase64(origData []byte, base64Signature string, base64Encoding *base64.Encoding) (bool, error) {
	sig, err := base64Encoding.DecodeString(base64Signature)
	if err != nil {
		return false, err
	}
	return o.VerifySign(origData, sig)
}

func (o *rsaPublicCipher) VerifySignHash(origData, signature []byte, hash crypto.Hash) (bool, error) {
	hashed := cryptox.Hash(hash, origData)
	if nil == hashed {
		return false, cryptox.ErrUnsupportedHash
	}
	err := rsa.VerifyPKCS1v15(o.PubKey, hash, hashed, signature)
	return nil == err, err
}

func (o *rsaPublicCipher) VerifySignHashBase64(origData []byte, base64Signature string, hash crypto.Hash, base64Encoding *base64.Encoding) (bool, error) {
	sig, err := base64Encoding.DecodeString(base64Signature)
	if err != nil {
		return false, err
	}
	return o.VerifySignHash(origData, sig, hash)
}

// GenerateRSAKeyPair generates an RSA private/public key pair.
// GenerateRSAKeyPair 生成 RSA 密钥对。
// bits: key size in bits.
// bits：密钥长度（位）。
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if nil != err {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// Encode2RSAPrivatePKCS1 encodes a private key as PKCS#1 PEM (RSA PRIVATE KEY).
// Encode2RSAPrivatePKCS1 编码为 PKCS#1 RSA 专用私钥 PEM（PEMTypeRSAPrivateKey）。
func Encode2RSAPrivatePKCS1(private *rsa.PrivateKey) ([]byte, error) {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  cryptox.PEMTypeRSAPrivateKey}), nil
}

// Encode2RSAPublicPKCS1 encodes a public key as PKCS#1 PEM (RSA PUBLIC KEY).
// Encode2RSAPublicPKCS1 编码为 PKCS#1 RSA 专用公钥 PEM（PEMTypeRSAPublicKey）。
func Encode2RSAPublicPKCS1(public *rsa.PublicKey) ([]byte, error) {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PublicKey(public),
		Type:  cryptox.PEMTypeRSAPublicKey}), nil
}

// Encode2PrivatePKCS8 encodes a private key as PKCS#8 PEM (PRIVATE KEY).
// Encode2PrivatePKCS8 编码为 PKCS#8 私钥 PEM（PEMTypePKCS8PrivateKey）。
func Encode2PrivatePKCS8(private *rsa.PrivateKey) ([]byte, error) {
	bs, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: bs,
		Type:  cryptox.PEMTypePKCS8PrivateKey,
	}), nil
}

// Encode2PublicX509 encodes a public key as X.509 SubjectPublicKeyInfo PEM (PUBLIC KEY).
// Encode2PublicX509 编码为 X.509 公钥 PEM（PEMTypeX509PublicKey）。
func Encode2PublicX509(public *rsa.PublicKey) ([]byte, error) {
	bs, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{
		Bytes: bs,
		Type:  cryptox.PEMTypeX509PublicKey,
	}), nil
}

// Encode2PublicOpenSSH encodes a public key in OpenSSH authorized_keys format (not PEM).
// Encode2PublicOpenSSH 编码为 OpenSSH 公钥格式（非 PEM 标准）。
func Encode2PublicOpenSSH(public *rsa.PublicKey) ([]byte, error) {
	sshPublicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(sshPublicKey), nil
}

// Encode2PrivateOpenSSH is not supported by the Go standard library as of Go 1.16+.
// Encode2PrivateOpenSSH Go 1.16+ 标准库不支持直接导出 OpenSSH 私钥。
func Encode2PrivateOpenSSH(private *rsa.PrivateKey) ([]byte, error) {
	return nil, nil
}

type funcEncode2KeyPair = func(privateKey *rsa.PrivateKey) (private []byte, public []byte, err error)

// Encode2OpenSSHPair encodes PKCS#1 private key and OpenSSH public key bytes.
// Encode2OpenSSHPair 编码为 OpenSSH 适用的公私钥：私钥 PKCS#1，公钥 OpenSSH 格式。
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

// Encode2RSAPair encodes PKCS#1 RSA private and public key PEM bytes.
// Encode2RSAPair 编码为 RSA 专用 PKCS#1 公私钥 PEM。
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

// Encode2PKCS8Pair encodes PKCS#8 private and X.509 public key PEM bytes.
// Encode2PKCS8Pair 编码为通用 PEM：私钥 PKCS#8，公钥 X.509。
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

// GenerateOpenSSHKeyFiles generates and writes an OpenSSH-style key pair to disk.
// GenerateOpenSSHKeyFiles 生成 OpenSSH 格式密钥对文件。
// bits: key size; privateFilePath/publicFilePath: output paths; private PKCS#1, public OpenSSH.
// bits：密钥长度；privateFilePath/publicFilePath：输出路径；私钥 PKCS#1，公钥 OpenSSH。
func GenerateOpenSSHKeyFiles(bits int, privateFilePath, publicFilePath string) error {
	return generateKeyFiles(bits, Encode2OpenSSHPair, privateFilePath, publicFilePath)
}

// SaveOpenSSHKeyFiles writes an existing key as OpenSSH-style pair files.
// SaveOpenSSHKeyFiles 保存 OpenSSH 格式密钥对文件。
// privateKey: private key; paths as above; private PKCS#1, public OpenSSH.
// privateKey：私钥；路径同上；私钥 PKCS#1，公钥 OpenSSH。
func SaveOpenSSHKeyFiles(privateKey *rsa.PrivateKey, privateFilePath, publicFilePath string) error {
	return saveKeyFiles(privateKey, Encode2OpenSSHPair, privateFilePath, publicFilePath)
}

// GenerateRSAKeyFiles generates and writes a PKCS#1 RSA key pair to disk.
// GenerateRSAKeyFiles 生成 RSA 专用 PKCS#1 密钥对文件。
// bits: key size; both keys PKCS#1 PEM.
// bits：密钥长度；公私钥均为 PKCS#1 PEM。
func GenerateRSAKeyFiles(bits int, privateFilePath, publicFilePath string) error {
	return generateKeyFiles(bits, Encode2RSAPair, privateFilePath, publicFilePath)
}

// SaveRSAKeyFiles writes an existing key as PKCS#1 RSA pair files.
// SaveRSAKeyFiles 保存 RSA 专用 PKCS#1 密钥对文件。
func SaveRSAKeyFiles(privateKey *rsa.PrivateKey, privateFilePath, publicFilePath string) error {
	return saveKeyFiles(privateKey, Encode2RSAPair, privateFilePath, publicFilePath)
}

// GeneratePemKeyFiles generates and writes a generic PEM key pair (PKCS#8 + X.509).
// GeneratePemKeyFiles 生成通用 PEM 密钥对文件（私钥 PKCS#8，公钥 X.509）。
func GeneratePemKeyFiles(bits int, privateFilePath, publicFilePath string) error {
	return generateKeyFiles(bits, Encode2PKCS8Pair, privateFilePath, publicFilePath)
}

// SavePemKeyFiles writes an existing key as generic PEM pair files.
// SavePemKeyFiles 保存通用 PEM 密钥对文件（私钥 PKCS#8，公钥 X.509）。
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
	_ = filex.WriteFile(privateFilePath, privateDer, os.ModePerm)
	_ = filex.WriteFile(publicFilePath, publicDer, os.ModePerm)
	return nil
}

var (
	// ErrKeyPath is returned when a key file path does not exist.
	// ErrKeyPath 在密钥文件路径不存在时返回。
	ErrKeyPath = errors.New("path is not exist")
	// ErrKeyFormat is returned when key PEM/DER content cannot be parsed.
	// ErrKeyFormat 在密钥内容格式错误时返回。
	ErrKeyFormat = errors.New("key content format error")
)

// LoadPublicCipherSSH loads an SSH-format public key file into IRSAPublicCipher.
// LoadPublicCipherSSH 从文件加载 SSH 格式公钥。
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

// LoadPrivateCipherRSA loads a PKCS#1 RSA private key file.
// LoadPrivateCipherRSA 从文件加载 PKCS#1 RSA 私钥。
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

// LoadPublicCipherRSA loads a PKCS#1 RSA public key file.
// LoadPublicCipherRSA 从文件加载 PKCS#1 RSA 公钥。
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

// LoadPrivateCipherPEM loads a PKCS#8 private key file.
// LoadPrivateCipherPEM 从文件加载 PKCS#8 私钥。
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

// LoadPublicCipherPEM loads an X.509 public key file.
// LoadPublicCipherPEM 从文件加载 X.509 公钥。
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

// LoadCiphersSSH loads a key pair: private PKCS#1, public SSH.
// LoadCiphersSSH 加载密钥对：私钥 PKCS#1，公钥 SSH。
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

// LoadCiphersRSA loads a PKCS#1 RSA key pair from files.
// LoadCiphersRSA 从文件加载 PKCS#1 RSA 密钥对。
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

// LoadCiphersPEM loads a generic PEM pair: private PKCS#8, public X.509.
// LoadCiphersPEM 加载通用 PEM 密钥对：私钥 PKCS#8，公钥 X.509。
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

// NewPrivateCipherByRSA builds IRSAPrivateCipher from PKCS#1 private key PEM/DER bytes.
// NewPrivateCipherByRSA 根据 PKCS#1（RSA 专用）私钥数据生成私钥处理器。
// privatePkcs1DerData: PKCS#1 private key content.
// privatePkcs1DerData：PKCS#1 格式私钥内容。
func NewPrivateCipherByRSA(privatePkcs1DerData []byte) (IRSAPrivateCipher, error) {
	privateKey, err := ParsePrivateKeyFromPKCS1(privatePkcs1DerData)
	if nil != err {
		return nil, err
	}
	return newPrivateRsa(privateKey), nil
}

// NewPublicCipherByRSA builds IRSAPublicCipher from PKCS#1 public key PEM/DER bytes.
// NewPublicCipherByRSA 根据 PKCS#1（RSA 专用）公钥数据生成公钥处理器。
func NewPublicCipherByRSA(publicPkcs1DerData []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParsePublicKeyFromPKCS1(publicPkcs1DerData)
	if nil != err {
		return nil, err
	}
	return newPublicRsa(publicKey), nil
}

// NewPublicCipherBySSH builds IRSAPublicCipher from OpenSSH authorized_keys bytes.
// NewPublicCipherBySSH 根据 SSH 公钥数据生成公钥处理器。
func NewPublicCipherBySSH(publicSSHData []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParsePublicKeyFromSSH(publicSSHData)
	if nil != err {
		return nil, err
	}
	return newPublicRsa(publicKey), nil
}

// NewPrivateCipherByPEM builds IRSAPrivateCipher from PKCS#8 private key PEM/DER bytes.
// NewPrivateCipherByPEM 根据 PKCS#8 私钥数据生成私钥处理器。
func NewPrivateCipherByPEM(pkcs8DerData []byte) (IRSAPrivateCipher, error) {
	privateKey, err := ParsePrivateKeyFromPKCS8(pkcs8DerData)
	if nil != err {
		return nil, err
	}
	return newPrivateRsa(privateKey), nil
}

// NewPublicCipherByPEM builds IRSAPublicCipher from X.509 public key PEM/DER bytes.
// NewPublicCipherByPEM 根据 X.509 公钥数据生成公钥处理器。
func NewPublicCipherByPEM(x509DerData []byte) (IRSAPublicCipher, error) {
	publicKey, err := ParsePublicKeyFromX509(x509DerData)
	if nil != err {
		return nil, err
	}
	return newPublicRsa(publicKey), nil
}

// NewCiphersBySSH builds ciphers from PKCS#1 private and SSH public key bytes.
// NewCiphersBySSH 根据密钥对数据生成处理器：私钥 PKCS#1，公钥 SSH。
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

// NewCiphersByRSA builds ciphers from PKCS#1 private and public key bytes.
// NewCiphersByRSA 根据密钥对数据生成处理器：私钥与公钥均为 PKCS#1。
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

// NewCiphersByPEM builds ciphers from PKCS#8 private and X.509 public key bytes.
// NewCiphersByPEM 根据密钥对数据生成处理器：私钥 PKCS#8，公钥 X.509。
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

// ParsePrivateKeyFromPKCS1 parses a PKCS#1 RSA private key from PEM/DER bytes.
// ParsePrivateKeyFromPKCS1 解析 PKCS#1（RSA 专用）格式私钥，返回 *rsa.PrivateKey。
func ParsePrivateKeyFromPKCS1(privatePkcs1DerData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privatePkcs1DerData)
	if block == nil {
		return nil, ErrKeyFormat
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// ParsePublicKeyFromPKCS1 parses a PKCS#1 RSA public key from PEM/DER bytes.
// ParsePublicKeyFromPKCS1 解析 PKCS#1（RSA 专用）格式公钥，返回 *rsa.PublicKey。
func ParsePublicKeyFromPKCS1(publicPkcs1DerData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicPkcs1DerData)
	if block == nil {
		return nil, ErrKeyFormat
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

// ParsePublicKeyFromSSH parses an OpenSSH authorized_keys line into *rsa.PublicKey.
// ParsePublicKeyFromSSH 将 OpenSSH 格式公钥解析为 *rsa.PublicKey。
func ParsePublicKeyFromSSH(sshPublicData []byte) (*rsa.PublicKey, error) {
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

// ParsePrivateKeyFromPKCS8 parses a PKCS#8 private key into *rsa.PrivateKey.
// ParsePrivateKeyFromPKCS8 将 PKCS#8 格式私钥解析为 *rsa.PrivateKey。
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

// ParsePublicKeyFromX509 parses an X.509 SubjectPublicKeyInfo key into *rsa.PublicKey.
// ParsePublicKeyFromX509 将 X.509 格式公钥解析为 *rsa.PublicKey。
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

// ConvertPublicPkcs1ToX509 converts PKCS#1 public key PEM/DER to X.509 SPKI PEM.
// ConvertPublicPkcs1ToX509 将 PKCS#1 公钥内容转换为 X.509 公钥内容。
func ConvertPublicPkcs1ToX509(pkcs1DerData []byte) ([]byte, error) {
	publicKey, err := ParsePublicKeyFromPKCS1(pkcs1DerData)
	if nil != err {
		return nil, err
	}
	return Encode2PublicX509(publicKey)
}

// ConvertPublicX509ToPkcs1 converts X.509 public key PEM/DER to PKCS#1 PEM.
// ConvertPublicX509ToPkcs1 将 X.509 公钥内容转换为 PKCS#1 公钥内容。
func ConvertPublicX509ToPkcs1(x509DerData []byte) ([]byte, error) {
	publicKey, err := ParsePublicKeyFromX509(x509DerData)
	if nil != err {
		return nil, err
	}
	return Encode2PublicX509(publicKey)
}

// ConvertPrivatePkcs1ToPkcs8 converts PKCS#1 private key PEM/DER to PKCS#8 PEM.
// ConvertPrivatePkcs1ToPkcs8 将 PKCS#1 私钥内容转换为 PKCS#8 私钥内容。
func ConvertPrivatePkcs1ToPkcs8(pkcs1DerData []byte) ([]byte, error) {
	privateKey, err := ParsePrivateKeyFromPKCS1(pkcs1DerData)
	if nil != err {
		return nil, ErrKeyFormat
	}
	return Encode2PrivatePKCS8(privateKey)
}

// ConvertPrivatePkcs8ToPkcs1 converts PKCS#8 private key PEM/DER to PKCS#1 PEM.
// ConvertPrivatePkcs8ToPkcs1 将 PKCS#8 私钥内容转换为 PKCS#1 私钥内容。
func ConvertPrivatePkcs8ToPkcs1(pkcs8DerData []byte) ([]byte, error) {
	privateKey, err := ParsePrivateKeyFromPKCS8(pkcs8DerData)
	if nil != err {
		return nil, ErrKeyFormat
	}
	return Encode2RSAPrivatePKCS1(privateKey)
}
