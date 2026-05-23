package cryptox

import "errors"

const (
	// PEMTypeOpenSSHPrivateKey is the PEM label for OpenSSH private keys.
	// PEMTypeOpenSSHPrivateKey OpenSSH 特定私钥，常用于 OpenSSH。
	PEMTypeOpenSSHPrivateKey = "OPENSSH PRIVATE KEY"
	// PEMTypeSSH2PublicKey is the PEM label for SSH2 public keys (SecureCRT, Tectia, etc.).
	// PEMTypeSSH2PublicKey SSH2 公钥，常用于 SecureCRT、Tectia 等工具。
	PEMTypeSSH2PublicKey = "SSH2 PUBLIC KEY"

	// PEMTypeRSAPrivateKey is the PEM label for PKCS#1 RSA private keys (OpenSSL, Go).
	// PEMTypeRSAPrivateKey PKCS#1 私钥，常用于 OpenSSL、Go。
	PEMTypeRSAPrivateKey = "RSA PRIVATE KEY"
	// PEMTypeRSAPublicKey is the PEM label for PKCS#1 RSA public keys (rare; prefer SPKI).
	// PEMTypeRSAPublicKey PKCS#1 公钥（较少用），通常导出为 SubjectPublicKeyInfo。
	PEMTypeRSAPublicKey = "RSA PUBLIC KEY"

	// PEMTypePKCS8PrivateKey is the PEM label for PKCS#8 private keys (multi-algorithm).
	// PEMTypePKCS8PrivateKey PKCS#8 私钥，更通用，支持多算法。
	PEMTypePKCS8PrivateKey = "PRIVATE KEY"
	// PEMTypeEncryptedPKCS8PrivateKey is the PEM label for encrypted PKCS#8 private keys.
	// PEMTypeEncryptedPKCS8PrivateKey 加密的 PKCS#8 私钥，常用于 OpenSSL、Go。
	PEMTypeEncryptedPKCS8PrivateKey = "ENCRYPTED PRIVATE KEY"
	// PEMTypeX509PublicKey is the PEM label for X.509 SubjectPublicKeyInfo public keys.
	// PEMTypeX509PublicKey X.509 公钥（SubjectPublicKeyInfo），常用于 OpenSSL、Go。
	PEMTypeX509PublicKey = "PUBLIC KEY"

	// PEMTypeCertificate is the PEM label for X.509 certificates (SSL/TLS).
	// PEMTypeCertificate X.509 证书，常用于 SSL/TLS。
	PEMTypeCertificate = "CERTIFICATE"
	// PEMTypeCertRequest is the PEM label for PKCS#10 certificate signing requests.
	// PEMTypeCertRequest PKCS#10 证书签名请求（CSR），常用于 OpenSSL、Go。
	PEMTypeCertRequest = "CERTIFICATE REQUEST"
	// PEMTypeNewCertRequest is an OpenSSL alias for certificate signing requests.
	// PEMTypeNewCertRequest OpenSSL CSR 别名。
	PEMTypeNewCertRequest = "NEW CERTIFICATE REQUEST"

	// PEMTypeCRL is the PEM label for X.509 certificate revocation lists.
	// PEMTypeCRL X509 吊销列表（Certificate Revocation List）。
	PEMTypeCRL = "X509 CRL"

	// PEMTypeECPrivateKey is the PEM label for SEC1 ECDSA private keys.
	// PEMTypeECPrivateKey ECDSA 私钥（SEC1 格式），常用于 OpenSSL。
	PEMTypeECPrivateKey = "EC PRIVATE KEY"
	// PEMTypeDSAPrivateKey is the PEM label for DSA private keys (uncommon).
	// PEMTypeDSAPrivateKey DSA 私钥（少见）。
	PEMTypeDSAPrivateKey = "DSA PRIVATE KEY"
	// PEMTypeAttributeCertificate is the PEM label for attribute certificates (rare).
	// PEMTypeAttributeCertificate 属性证书（稀有）。
	PEMTypeAttributeCertificate = "ATTRIBUTE CERTIFICATE"

	// PEMTypePKCS7 is the PEM label for PKCS#7 signed/encrypted containers.
	// PEMTypePKCS7 PKCS#7 数据（签名/加密容器）。
	PEMTypePKCS7 = "PKCS7"
	// PEMTypeCMS is the PEM label for Cryptographic Message Syntax (PKCS#7 successor).
	// PEMTypeCMS Cryptographic Message Syntax（PKCS#7 的替代）。
	PEMTypeCMS = "CMS"
)

var (
	// ErrUnsupportedPemType is returned when a PEM block type is not recognized.
	// ErrUnsupportedPemType 在无法识别的 PEM 块类型时返回。
	ErrUnsupportedPemType = errors.New("unsupported pem type")
)
