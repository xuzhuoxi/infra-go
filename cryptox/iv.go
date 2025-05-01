// Package cryptox
// Create on 2025/4/29
// @author xuzhuoxi
package cryptox

import (
	"crypto/rand"
	"io"
)

func RandomIV(size int) ([]byte, error) {
	iv := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}
