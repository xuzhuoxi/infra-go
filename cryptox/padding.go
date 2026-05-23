package cryptox

import (
	"bytes"
	"crypto/rand"
	"errors"
)

// FuncPadding pads plaintext to a multiple of blockSize.
// FuncPadding 将明文填充为 blockSize 的整数倍。
type FuncPadding func(data []byte, blockSize int) ([]byte, error)

// FuncUnPadding removes padding from ciphertext.
// FuncUnPadding 从密文中去除填充。
type FuncUnPadding func(data []byte) ([]byte, error)

// PKCS7Padding applies PKCS#7 padding.
// PKCS7Padding 应用 PKCS#7 填充。
// Steps:
// 1. padding length = blockSize - (len(data) % blockSize); if zero, use blockSize.
// 2. append paddingLen bytes each equal to paddingLen.
// 3. last byte value equals padding length (e.g. 0x04 for 4 padding bytes).
// 步骤：
// 1. 计算填充长度：填充长度是一个字节，表示填充的字节数。它等于块大小减去当前数据长度对块大小取余的结果。
// 2. 填充字节：所有填充字节的值都与填充长度相同。
// 3. 最后一个字节：最后一个字节的值表示填充的字节数（例如，如果填充了 4 字节，最后一个字节的值为 0x04）。
func PKCS7Padding(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("blockSize must be positive")
	}

	// 计算填充长度
	paddingLen := blockSize - len(data)%blockSize
	if paddingLen == 0 {
		paddingLen = blockSize
	}

	// 填充字节
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)

	// 返回填充后的数据
	return append(data, padding...), nil
}

// PKCS7UnPadding removes PKCS#7 padding.
// PKCS7UnPadding 去除 PKCS#7 填充。
// Steps:
// 1. read the last byte as padding length.
// 2. strip padding bytes to recover original data.
// 步骤：
// 1. 读取数据的最后一个字节，提取填充长度。
// 2. 删除填充字节，恢复原始数据。
func PKCS7UnPadding(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	// 获取最后一个字节的值，表示填充长度
	paddingLen := int(data[len(data)-1])
	if paddingLen > len(data) {
		return nil, errors.New("invalid padding")
	}

	// 去除填充并返回
	return data[:len(data)-paddingLen], nil
}

// ZeroPadding applies zero-byte padding to blockSize.
// ZeroPadding 使用零字节填充至 blockSize。
// Steps:
// 1. padding length = blockSize - (len(data) % blockSize); if zero, use blockSize.
// 2. append zero bytes (0x00) until length is a multiple of blockSize.
// 3. no length marker in padding; unpad by stripping trailing zeros.
// 步骤：
// 1. 计算填充长度：如果数据的长度没有达到块大小，则计算需要填充的字节数。
// 2. 填充字节：填充字节是 零字节 (0x00)，直到数据长度达到块大小。
// 3. 没有标记：Zero Padding 不会在填充内容中包含标记字节，所以解填充时需要注意是否使用块大小的倍数。
func ZeroPadding(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("blockSize must be positive")
	}

	// 计算需要填充的字节数
	paddingLen := blockSize - len(data)%blockSize
	if paddingLen == 0 {
		paddingLen = blockSize
	}

	// 填充零字节
	padding := make([]byte, paddingLen)

	// 返回填充后的数据
	return append(data, padding...), nil
}

// ZeroUnPadding removes zero-byte padding from the end.
// ZeroUnPadding 从末尾去除零字节填充。
// Steps:
// 1. strip trailing zero bytes until the first non-zero byte.
// 2. return the remaining original data.
// 步骤：
// 1. 从数据末尾开始，去除零字节直到遇到第一个非零字节。
// 2. 返回去除零字节后的原始数据。
func ZeroUnPadding(data []byte) ([]byte, error) {
	// 去除末尾的零字节
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != 0 {
			return data[:i+1], nil
		}
	}
	return nil, errors.New("invalid padding")
}

// Iso10126Padding applies ISO 10126 padding (random bytes + length byte).
// Iso10126Padding 应用 ISO 10126 填充（随机字节 + 长度字节）。
// Steps:
// 1. padding length = blockSize - (len(data) % blockSize); if zero, use blockSize.
// 2. fill with random bytes except the last length byte.
// 3. last byte equals padding length (e.g. 0x04 for 4 padding bytes).
// 步骤:
// 1. 计算填充长度：填充长度是一个字节，表示填充的字节数。它等于填充字节数的总长度。
// 2. 填充字节：填充内容由随机字节组成，直到填充字节数等于所需长度（一般填充到块大小）。
// 3. 最后一个字节：最后一个字节的值是填充字节的长度（例如，如果填充了 4 字节，最后一个字节的值为 0x04）。
func Iso10126Padding(ciphertext []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("blockSize must be positive")
	}
	// 计算填充长度
	paddingLen := blockSize - len(ciphertext)%blockSize
	if paddingLen == 0 {
		paddingLen = blockSize
	}

	// 填充随机字节
	padding := make([]byte, paddingLen-1)
	_, err := rand.Read(padding)
	if err != nil {
		return nil, err
	}

	// 添加最后一个字节，表示填充长度
	padding = append(padding, byte(paddingLen))

	// 返回填充后的数据
	return append(ciphertext, padding...), nil
}

// Iso10126UnPadding removes ISO 10126 padding.
// Iso10126UnPadding 去除 ISO 10126 填充。
// Steps:
// 1. read the last byte as padding length.
// 2. strip padding bytes to recover original data.
// 步骤：
// 1. 读取数据的最后一个字节，提取填充长度。
// 2. 删除填充字节，恢复原始数据。
func Iso10126UnPadding(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	// 获取最后一个字节的值，表示填充长度
	paddingLen := int(data[len(data)-1])
	if paddingLen > len(data) {
		return nil, errors.New("invalid padding")
	}

	// 去除填充并返回
	return data[:len(data)-paddingLen], nil
}

// Ansix923Padding applies ANSI X9.23 padding (zeros + length byte).
// Ansix923Padding 应用 ANSI X9.23 填充（零字节 + 长度字节）。
// Steps:
// 1. padding length = blockSize - (len(data) % blockSize); if zero, use blockSize.
// 2. fill with zero bytes except the last length byte.
// 3. last byte equals padding length (e.g. 0x04 for 4 padding bytes).
// 步骤：
// 1. 计算填充长度：填充长度是一个字节，表示填充的字节数。它等于填充字节数的总长度。
// 2. 填充字节：填充内容由 零 字节组成，直到填充字节数等于所需长度（一般填充到块大小）。
// 3. 最后一个字节：最后一个字节的值是填充的字节长度（例如，如果填充了 4 字节，最后一个字节的值为 0x04）。
func Ansix923Padding(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("blockSize must be positive")
	}

	// 计算填充长度
	paddingLen := blockSize - len(data)%blockSize
	if paddingLen == 0 {
		paddingLen = blockSize
	}

	// 填充零字节
	padding := make([]byte, paddingLen-1)
	// 最后一个字节是填充长度
	padding = append(padding, byte(paddingLen))

	// 返回填充后的数据
	return append(data, padding...), nil
}

// Ansix923UnPadding removes ANSI X9.23 padding.
// Ansix923UnPadding 去除 ANSI X9.23 填充。
// Steps:
// 1. read the last byte as padding length.
// 2. strip padding bytes to recover original data.
// 步骤：
// 1. 读取数据的最后一个字节，提取填充长度。
// 2. 删除填充字节，恢复原始数据。
func Ansix923UnPadding(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}

	// 获取最后一个字节的值，表示填充长度
	paddingLen := int(data[len(data)-1])
	if paddingLen > len(data) {
		return nil, errors.New("invalid padding")
	}

	// 去除填充并返回
	return data[:len(data)-paddingLen], nil
}
