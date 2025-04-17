package cryptox

// 加密方案一：
// 使用对称密钥（比如AES/DES等加解密方法）加密数据，然后使用非对称密钥（RSA加解密密钥）加密对称密钥；
//
// 加密方案二：
// 直接使用非对称密钥加密数据

// IEncryptCipher 加密处理器
type IEncryptCipher interface {
	// Encrypt 加密
	Encrypt(plaintext []byte) ([]byte, error)
}

// IDecryptCipher 解密处理器
type IDecryptCipher interface {
	// Decrypt 解密
	Decrypt(ciphertext []byte) ([]byte, error)
}

// ICipher 加密解密处理器
type ICipher interface {
	IEncryptCipher
	IDecryptCipher
}
