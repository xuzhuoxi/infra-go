package cryptox

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"os"
)

var ErrUnsupportedHash = errors.New("unsupported hash")

func Md5(data []byte) string {
	return hex.EncodeToString(_md5(data))
}
func Md5String(data string) string {
	return Md5([]byte(data))
}
func Md5File(filePath string) string {
	data, err := os.ReadFile(filePath)
	if nil != err {
		return ""
	}
	return Md5(data)
}

func Sha1(data []byte) string {
	return hex.EncodeToString(_sha1(data))
}
func Sha1String(data string) string {
	return Sha1([]byte(data))
}
func Sha1File(filePath string) string {
	data, err := os.ReadFile(filePath)
	if nil != err {
		return ""
	}
	return Sha1(data)
}

func Hash(hash crypto.Hash, data []byte) []byte {
	if hash.Available() {
		h := hash.New()
		h.Write(data)
		return h.Sum(nil)
	}
	return nil
}
func Hash2Hex(hash crypto.Hash, data []byte) string {
	rs := Hash(hash, data)
	if nil == rs {
		return ""
	}
	return hex.EncodeToString(rs)
}

func HashString(hash crypto.Hash, data string) []byte {
	return Hash(hash, []byte(data))
}

func HashString2Hex(hash crypto.Hash, data string) string {
	return Hash2Hex(hash, []byte(data))
}

func HashFile(hash crypto.Hash, filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if nil != err {
		return nil
	}
	return Hash(hash, data)
}

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
