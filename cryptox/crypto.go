package cryptox

type BlockNilError string

func (e BlockNilError) Error() string   { return "Block Nil Error At: " + string(e) }
func (e BlockNilError) Timeout() bool   { return false }
func (e BlockNilError) Temporary() bool { return false }

// 加密方案一：
// 使用对称密钥（比如AES/DES等加解密方法）加密数据，然后使用非对称密钥（RSA加解密密钥）加密对称密钥；
//
// 加密方案二：
// 直接使用非对称密钥加密数据
