package cryptox

import "bytes"

type FuncPadding func(ciphertext []byte, blockSize int) []byte
type FuncUnPadding func(origData []byte) []byte

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	//padding := blockSize - len(ciphertext)%blockSize
	//padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//return append(ciphertext, padtext...)
	return nil
}
func ZeroUnPadding(origData []byte) []byte {
	//length := len(origData)
	//unpadding := int(origData[length-1])
	//return origData[:(length - unpadding)]
	return nil
}

func Iso10126Padding(ciphertext []byte, blockSize int) []byte {
	//padding := blockSize - len(ciphertext)%blockSize
	//padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//return append(ciphertext, padtext...)
	return nil
}
func Iso10126UnPadding(origData []byte) []byte {
	//length := len(origData)
	//unpadding := int(origData[length-1])
	//return origData[:(length - unpadding)]
	return nil
}

func Ansix923Padding(ciphertext []byte, blockSize int) []byte {
	//padding := blockSize - len(ciphertext)%blockSize
	//padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//return append(ciphertext, padtext...)
	return nil
}
func Ansix923UnPadding(origData []byte) []byte {
	//length := len(origData)
	//unpadding := int(origData[length-1])
	//return origData[:(length - unpadding)]
	return nil
}
