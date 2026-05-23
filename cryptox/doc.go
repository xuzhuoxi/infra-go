// Package cryptox provides cryptographic primitives: cipher interfaces, hashing, padding, block modes, and PEM type constants.
// Package cryptox 提供密码学基础能力：加解密接口、摘要哈希、填充算法、分组模式及 PEM 类型常量。
//
// Created by xuzhuoxi on 2019-02-03.
// @author xuzhuoxi
//
// Hybrid encryption schemes (for integrators):
//   - Scheme 1: encrypt data with a symmetric key (AES/DES, etc.), then encrypt the symmetric key with asymmetric RSA.
//   - Scheme 2: encrypt data directly with an asymmetric key.
//
// Capabilities:
//   - IEncryptCipher, IDecryptCipher, ICipher: encrypt/decrypt abstractions
//   - MD5/SHA1 and generic crypto.Hash helpers with hex encoding and file helpers
//   - BlockMode (ECB, CBC, CFB, OFB, CTR, GCM) with usage notes
//   - Padding: PKCS#7, Zero, ISO 10126, ANSI X9.23 and custom FuncPadding/FuncUnPadding hooks
//   - PEM block type string constants for RSA, PKCS#8, X.509, SSH, certificates, etc.
//   - RandomIV for block cipher initialization vectors
//
// 混合加密方案（供集成参考）：
//   - 方案一：先用对称密钥（如 AES/DES）加密数据，再用 RSA 等非对称密钥加密对称密钥。
//   - 方案二：直接使用非对称密钥加密数据。
//
// 主要职责：
//   - IEncryptCipher、IDecryptCipher、ICipher：加解密抽象
//   - MD5/SHA1 及通用 crypto.Hash 辅助（含十六进制与文件读写）
//   - BlockMode（ECB、CBC、CFB、OFB、CTR、GCM）及选用说明
//   - PKCS#7、Zero、ISO 10126、ANSI X9.23 填充及 FuncPadding/FuncUnPadding 钩子
//   - RSA、PKCS#8、X.509、SSH、证书等 PEM 块类型字符串常量
//   - RandomIV：生成分组密码初始化向量
package cryptox
