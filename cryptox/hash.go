package cryptox

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"os"
)

// ErrUnsupportedHash is returned when the requested hash algorithm is unavailable.
// ErrUnsupportedHash 在请求的哈希算法不可用时返回。
var ErrUnsupportedHash = errors.New("unsupported hash")

// Md5 returns the MD5 digest of data as a lowercase hex string.
// Md5 返回 data 的 MD5 摘要（小写十六进制字符串）。
func Md5(data []byte) string {
	return hex.EncodeToString(_md5(data))
}

// Md5String returns the MD5 digest of a string as lowercase hex.
// Md5String 返回字符串的 MD5 摘要（小写十六进制）。
func Md5String(data string) string {
	return Md5([]byte(data))
}

// Md5File returns the MD5 digest of a file as lowercase hex, or empty string on read error.
// Md5File 返回文件的 MD5 摘要（小写十六进制），读取失败时返回空字符串。
func Md5File(filePath string) string {
	data, err := os.ReadFile(filePath)
	if nil != err {
		return ""
	}
	return Md5(data)
}

// Sha1 returns the SHA-1 digest of data as a lowercase hex string.
// Sha1 返回 data 的 SHA-1 摘要（小写十六进制字符串）。
func Sha1(data []byte) string {
	return hex.EncodeToString(_sha1(data))
}

// Sha1String returns the SHA-1 digest of a string as lowercase hex.
// Sha1String 返回字符串的 SHA-1 摘要（小写十六进制）。
func Sha1String(data string) string {
	return Sha1([]byte(data))
}

// Sha1File returns the SHA-1 digest of a file as lowercase hex, or empty string on read error.
// Sha1File 返回文件的 SHA-1 摘要（小写十六进制），读取失败时返回空字符串。
func Sha1File(filePath string) string {
	data, err := os.ReadFile(filePath)
	if nil != err {
		return ""
	}
	return Sha1(data)
}

// Hash computes the digest using hash; returns nil if the algorithm is unavailable.
// Hash 使用 hash 计算摘要；算法不可用时返回 nil。
func Hash(hash crypto.Hash, data []byte) []byte {
	if hash.Available() {
		h := hash.New()
		h.Write(data)
		return h.Sum(nil)
	}
	return nil
}

// Hash2Hex computes the digest and encodes it as lowercase hex; returns empty string if unavailable.
// Hash2Hex 计算摘要并编码为小写十六进制；不可用时返回空字符串。
func Hash2Hex(hash crypto.Hash, data []byte) string {
	rs := Hash(hash, data)
	if nil == rs {
		return ""
	}
	return hex.EncodeToString(rs)
}

// HashString computes the digest of a string.
// HashString 计算字符串的摘要。
func HashString(hash crypto.Hash, data string) []byte {
	return Hash(hash, []byte(data))
}

// HashString2Hex computes the digest of a string and returns lowercase hex.
// HashString2Hex 计算字符串摘要并返回小写十六进制。
func HashString2Hex(hash crypto.Hash, data string) string {
	return Hash2Hex(hash, []byte(data))
}

// HashFile reads a file and returns its digest, or nil on read error.
// HashFile 读取文件并返回摘要，读取失败时返回 nil。
func HashFile(hash crypto.Hash, filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if nil != err {
		return nil
	}
	return Hash(hash, data)
}

// HashFile2Hex reads a file and returns its digest as lowercase hex, or empty string on error.
// HashFile2Hex 读取文件并返回小写十六进制摘要，失败时返回空字符串。
func HashFile2Hex(hash crypto.Hash, filePath string) string {
	data, err := os.ReadFile(filePath)
	if nil != err {
		return ""
	}
	return Hash2Hex(hash, data)
}

func _md5(data []byte) []byte {
	arr := md5.Sum(data)
	return arr[:]
}
func _sha1(data []byte) []byte {
	arr := sha1.Sum(data)
	return arr[:]
}
