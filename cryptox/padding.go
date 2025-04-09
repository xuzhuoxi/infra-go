package cryptox

import (
	"bytes"
	"crypto/rand"
	"errors"
)

type FuncPadding func(data []byte, blockSize int) ([]byte, error)
type FuncUnPadding func(data []byte) ([]byte, error)

// PKCS7Padding PKCS#7 填充算法步骤
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

// PKCS7UnPadding PKCS#7 填充算法的还原算法
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

// ZeroPadding Zero Padding 填充算法
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

// ZeroUnPadding Zero Padding 填充算法的还原算法
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

// Iso10126Padding ISO10126 填充算法
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

// Iso10126UnPadding ISO10126 填充算法的还原
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

// Ansix923Padding ANSI X9.23 填充算法
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

// Ansix923UnPadding ANSI X9.23 填充算法还原算法
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
