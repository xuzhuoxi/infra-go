// Package symmetric provides symmetric ciphers: AES (CBC/CTR/GCM), DES/3DES (ECB/CBC/CTR), and XOR obfuscation.
// Package symmetric 提供对称加密：AES（CBC/CTR/GCM）、DES/3DES（ECB/CBC/CTR）及 XOR 混淆。
//
// Created by xuzhuoxi on 2025-04-06.
// @author xuzhuoxi
//
// Capabilities:
//   - IAESCipher: AES with configurable padding; default encrypt/decrypt uses GCM
//   - IDESCipher: single DES or triple-DES (8/16/24-byte keys); default CBC
//   - IXORCipher: fast XOR stream (not cryptographically secure)
//   - Block modes integrate with cryptox.BlockMode and cryptox padding helpers
//
// 主要职责：
//   - IAESCipher：可配置填充的 AES，默认加解密使用 GCM
//   - IDESCipher：单 DES 或 3DES（8/16/24 字节密钥），默认 CBC
//   - IXORCipher：快速 XOR 流（不具备密码学安全性）
//   - 分组模式与 cryptox.BlockMode、填充函数配合使用
package symmetric
