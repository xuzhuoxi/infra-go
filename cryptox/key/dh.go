// Package key
// Create on 2025/4/6
// @author xuzhuoxi
package key

import (
	"crypto/rand"
	"math/big"
)

// RFC 3526 Group 14 (2048-bit MODP)
const primeHex = `
FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E08
8A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B
302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9
A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE6
49286651ECE65381FFFFFFFFFFFFFFFF
`

var (
	p, _ = new(big.Int).SetString(removeSpaces(primeHex), 16)
	g    = big.NewInt(2)
)

// DHKeyPair 表示 DH 密钥对
type DHKeyPair struct {
	Private *big.Int // 私钥
	Public  *big.Int // 公钥 A = g^a mod p
}

// GenerateDHKeyPair 生成一对 DH 密钥
func GenerateDHKeyPair() (*DHKeyPair, error) {
	private, err := rand.Int(rand.Reader, p)
	if err != nil {
		return nil, err
	}
	public := new(big.Int).Exp(g, private, p)
	return &DHKeyPair{Private: private, Public: public}, nil
}

// ComputeDHSharedK 计算共享密钥 K = B^a mod p
func ComputeDHSharedK(theirPublic *big.Int, myPrivate *big.Int) *big.Int {
	return new(big.Int).Exp(theirPublic, myPrivate, p)
}

// 移除 primeHex 中的空格和换行
func removeSpaces(s string) string {
	var out []rune
	for _, r := range s {
		if r != '\n' && r != ' ' && r != '\t' {
			out = append(out, r)
		}
	}
	return string(out)
}
