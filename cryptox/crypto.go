package cryptox

// IEncryptCipher encrypts plaintext to ciphertext.
// IEncryptCipher 将明文加密为密文。
type IEncryptCipher interface {
	// Encrypt encrypts plaintext.
	// Encrypt 加密明文。
	Encrypt(plaintext []byte) ([]byte, error)
}

// IDecryptCipher decrypts ciphertext to plaintext.
// IDecryptCipher 将密文解密为明文。
type IDecryptCipher interface {
	// Decrypt decrypts ciphertext.
	// Decrypt 解密密文。
	Decrypt(ciphertext []byte) ([]byte, error)
}

// ICipher combines encryption and decryption.
// ICipher 组合加密与解密能力。
type ICipher interface {
	IEncryptCipher
	IDecryptCipher
}
