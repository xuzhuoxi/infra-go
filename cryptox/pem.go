// Package cryptox
// Create on 2025/4/8
// @author xuzhuoxi
package cryptox

import "errors"

const (
	// PEMTypeOpenSSHPrivateKey
	// OpenSSH特定私钥
	// 常用于 OpenSSH
	PEMTypeOpenSSHPrivateKey = "OPENSSH PRIVATE KEY"
	// PEMTypeSSH2PublicKey
	// SSH2公钥
	// 常用于 SecureCRT、Tectia等工具
	PEMTypeSSH2PublicKey = "SSH2 PUBLIC KEY"

	// PEMTypeRSAPrivateKey
	// PKCS#1 私钥
	// 常用于 OpenSSL、Go
	PEMTypeRSAPrivateKey = "RSA PRIVATE KEY"
	// PEMTypeRSAPublicKey
	// PKCS#1 公钥（较少用）
	// 不常用，通常导出为 SubjectPublicKeyInfo
	PEMTypeRSAPublicKey = "RSA PUBLIC KEY"

	// PEMTypePKCS8PrivateKey
	// PKCS#8 私钥
	// 更通用，支持多算法
	PEMTypePKCS8PrivateKey = "PRIVATE KEY"
	// PEMTypeEncryptedPKCS8PrivateKey
	// 加密的 PKCS#8 私钥
	// 常用于 OpenSSL、Go
	PEMTypeEncryptedPKCS8PrivateKey = "ENCRYPTED PRIVATE KEY"
	// PEMTypeX509PublicKey
	// X.509 公钥（SubjectPublicKeyInfo）
	// 常用于 OpenSSL、Go
	PEMTypeX509PublicKey = "PUBLIC KEY"

	// PEMTypeCertificate
	// X.509 证书
	// 常用于 SSL/TLS
	PEMTypeCertificate = "CERTIFICATE"
	// PEMTypeCertRequest
	// PKCS#10 证书签名请求（CSR）
	// 常用于 OpenSSL、Go
	PEMTypeCertRequest = "CERTIFICATE REQUEST"
	// PEMTypeNewCertRequest
	// OpenSSL CSR 别名
	PEMTypeNewCertRequest = "NEW CERTIFICATE REQUEST"

	// PEMTypeCRL
	// X509 吊销列表（Certificate Revocation List）
	PEMTypeCRL = "X509 CRL"

	// PEMTypeECPrivateKey
	// ECDSA 私钥（SEC1格式）
	// 常用于 OpenSSL
	PEMTypeECPrivateKey = "EC PRIVATE KEY"
	// PEMTypeDSAPrivateKey
	// DSA 私钥（少见）
	PEMTypeDSAPrivateKey = "DSA PRIVATE KEY"
	// PEMTypeAttributeCertificate
	// 属性证书（稀有）
	PEMTypeAttributeCertificate = "ATTRIBUTE CERTIFICATE"

	// PEMTypePKCS7
	// PKCS#7数据（签名/加密容器）
	PEMTypePKCS7 = "PKCS7"
	// PEMTypeCMS
	// Cryptographic Message Syntax（PKCS#7 的替代）
	PEMTypeCMS = "CMS"
)

var (
	ErrUnsupportedPemType = errors.New("unsupported pem type")
)
