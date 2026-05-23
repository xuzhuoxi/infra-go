// Package key provides shared-secret derivation (SHA-256, PBKDF2) and Diffie-Hellman key exchange (RFC 3526 group 14).
// Package key 提供共享密钥派生（SHA-256、PBKDF2）及 Diffie-Hellman 密钥交换（RFC 3526 group 14）。
//
// Created by xuzhuoxi on 2025-04-06.
// @author xuzhuoxi
//
// Capabilities:
//   - SharedKeySha256*: derive 32-byte keys from passphrases (AES-256, HMAC)
//   - DeriveKeyPbkdf2*: PBKDF2-HMAC-SHA512 key derivation with configurable or default salt/iterations
//   - GenerateDHKeyPair / ComputeDHSharedK: 2048-bit MODP DH per RFC 3526
//
// 主要职责：
//   - SharedKeySha256*：从口令派生 32 字节密钥（适用于 AES-256、HMAC）
//   - DeriveKeyPbkdf2*：PBKDF2-HMAC-SHA512 派生，支持默认或可配置 salt/迭代次数
//   - GenerateDHKeyPair / ComputeDHSharedK：RFC 3526 2048 位 MODP DH 密钥交换
package key
