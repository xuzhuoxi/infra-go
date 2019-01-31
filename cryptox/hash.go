package cryptox

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
)

//
//type FuncHash func([]byte) []byte
//
//var (
//	hashMap map[crypto.Hash]FuncHash
//)

func init() {
	//hashMap = make(map[crypto.Hash]FuncHash)
	//hashMap[crypto.MD5] = _md5
	//hashMap[crypto.Sha1] = _sha1
	//hashMap[crypto.SHA224] = _sha224
	//hashMap[crypto.SHA256] = _sha256
	//hashMap[crypto.SHA384] = _sha384
	//hashMap[crypto.SHA512] = _sha512
	//hashMap[crypto.SHA512_224] = _sha512_224
	//hashMap[crypto.SHA512_256] = _sha512_256
	//hashMap[crypto.RIPEMD160] = _ripemd160
	//hashMap[crypto.SHA3_224] = _sha3_224
	//hashMap[crypto.SHA3_256] = _sha3_256
	//hashMap[crypto.SHA3_384] = _sha3_384
	//hashMap[crypto.SHA3_512] = _sha3_512
	//hashMap[crypto.BLAKE2s_256] = _blake2s_256
	//hashMap[crypto.BLAKE2b_256] = _blake2b_256
	//hashMap[crypto.BLAKE2b_384] = _blake2b_384
	//hashMap[crypto.BLAKE2b_512] = _blake2b_512
}

func Md5(data []byte) string {
	return hex.EncodeToString(_md5(data))
}
func Md5String(data string) string {
	return Md5([]byte(data))
}
func Md5File(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
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
	data, err := ioutil.ReadFile(filePath)
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
	data, err := ioutil.ReadFile(filePath)
	if nil != err {
		return nil
	}
	return Hash(hash, data)
}

func HashFile2Hex(hash crypto.Hash, filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if nil != err {
		return ""
	}
	return Hash2Hex(hash, data)
}

func _md5(data []byte) []byte {
	return md5.Sum(data)[:]
}
func _sha1(data []byte) []byte {
	return sha1.Sum(data)[:]
}
func _sha224(data []byte) []byte {
	return sha256.Sum224(data)[:]
}
func _sha256(data []byte) []byte {
	return sha256.Sum256(data)[:]
}
func _sha384(data []byte) []byte {
	return sha512.Sum384(data)[:]
}
func _sha512(data []byte) []byte {
	return sha512.Sum512(data)[:]
}
func _sha512_224(data []byte) []byte {
	return sha512.Sum512_224(data)[:]
}
func _sha512_256(data []byte) []byte {
	return sha512.Sum512_256(data)[:]
}
func _ripemd160(data []byte) []byte {
	digest := ripemd160.New()
	digest.Write(data)
	return digest.Sum(nil)
}
func _sha3_224(data []byte) []byte {
	return sha3.Sum224(data)[:]
}
func _sha3_256(data []byte) []byte {
	return sha3.Sum256(data)[:]
}
func _sha3_384(data []byte) []byte {
	return sha3.Sum384(data)[:]
}
func _sha3_512(data []byte) []byte {
	return sha3.Sum512(data)[:]
}
func _blake2s_256(data []byte) []byte {
	return blake2s.Sum256(data)[:]
}
func _blake2b_256(data []byte) []byte {
	return blake2b.Sum256(data)[:]
}
func _blake2b_384(data []byte) []byte {
	return blake2b.Sum384(data)[:]
}
func _blake2b_512(data []byte) []byte {
	return blake2b.Sum512(data)[:]
}
