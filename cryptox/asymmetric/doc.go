// Package asymmetric provides RSA key management, PKCS/PEM encoding, hybrid encryption, and signing.
// Package asymmetric 提供 RSA 密钥管理、PKCS/PEM 编解码、混合加密与签名验签。
//
// Created by xuzhuoxi on 2019-02-03.
// @author xuzhuoxi
//
// PKCS standards (reference):
//   - PKCS#1: RSA encryption and signature (used by PKCS#7 envelopes)
//   - PKCS#3: Diffie-Hellman key exchange
//   - PKCS#5: password-based key derivation (DES-CBC)
//   - PKCS#6: X.509 certificate syntax extensions
//   - PKCS#7: general message syntax (signatures, encryption; PEM-compatible)
//   - PKCS#8: private key information format
//   - PKCS#9–#15: attributes, CSR, Cryptoki, PFX, EC, PRNG, token info
//
// Capabilities:
//   - Generate RSA key pairs and save/load PEM or OpenSSH key files
//   - IRSAPrivateCipher / IRSAPublicCipher / IRSACipher: encrypt, decrypt, hybrid (RSA+AES-CTR), sign, verify
//   - Encode/decode PKCS#1, PKCS#8, X.509, and OpenSSH public keys; format conversion helpers
//
// PKCS 标准（参考）：
//   - PKCS#1：RSA 加密与签名（PKCS#7 数字信封等会用到）
//   - PKCS#3：Diffie-Hellman 密钥交换
//   - PKCS#5：基于口令的密钥派生（DES-CBC）
//   - PKCS#6：X.509 证书语法扩展
//   - PKCS#7：通用消息语法（签名、加密；与 PEM 兼容）
//   - PKCS#8：私钥信息格式
//   - PKCS#9–#15：属性、CSR、Cryptoki、PFX、椭圆曲线、伪随机数、令牌信息等
//
// 主要职责：
//   - 生成 RSA 密钥对，保存/加载 PEM 或 OpenSSH 格式密钥文件
//   - IRSAPrivateCipher / IRSAPublicCipher / IRSACipher：加解密、混合加密（RSA+AES-CTR）、签名验签
//   - PKCS#1、PKCS#8、X.509、OpenSSH 公钥编解码及格式转换
package asymmetric
