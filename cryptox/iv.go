package cryptox

import (
	"crypto/rand"
	"io"
)

// RandomIV fills a cryptographically secure random IV of the given size.
// RandomIV 生成指定长度的密码学安全随机 IV。
func RandomIV(size int) ([]byte, error) {
	iv := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}
