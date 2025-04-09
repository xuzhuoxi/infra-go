// Package cryptox
// Created by xuzhuoxi
// on 2019-02-03.
// @author xuzhuoxi
//
package cryptox

// BlockMode
//
// 模式	全称						并行性		是否需要IV	  是否自带认证	安全性 		常见用途
// ECB	Electronic Codebook		✅ 加密		❌			  ❌			❌低		不推荐，易泄漏模式结构
// CBC	Cipher Block Chaining	❌ 			✅			  ❌			⚠️中		常用于文件加密
// CFB	Cipher Feedback			❌ 			✅			  ❌			✅较高		加密串流数据
// OFB	Output Feedback			✅ 解密		✅			  ❌			✅较高		加密串流数据
// CTR	Counter Mode			✅ 全并行	✅（nonce）	  ❌			✅高		高性能通信流加密
// GCM	Galois/Counter Mode		✅ 全并行	✅			  ❌			✅✅ 很高	TLS、HTTPS、VPN、敏感数据传输
//
// 模式	用于DES？	用于AES？	说明
// ECB	✅ 是		✅ 是		可用于两者，但因不安全，基本只在演示或兼容场景使用
// CBC	✅ 常用		✅ 常用		经典通用，文件加密、旧协议广泛使用
// CFB	✅ 是		✅ 是		支持流式加密，适合字节数据传输场景
// OFB	✅ 是		✅ 是		少用，适用于需要同步加密流的地方
// CTR	✅ 是		✅ 是		非常高性能，在 AES 中尤其常见，适合并行加密
// GCM	❌ 否		✅ 是		专为 AES 设计的块加密模式
type BlockMode string

const (
	// ECB （电子密码本模式） ❌ 不安全
	// 每个分组独立加密，没有上下文。
	// 同一明文块总是会生成相同密文块 → 可被模式识别（结构泄漏）。
	// ✅ 优点：并行处理快
	// ❌ 缺点：明文块相同→密文块也相同，容易被攻击
	// 🚫 不推荐用于任何实际应用。
	ECB BlockMode = "ECB"
	// CBC （加密分组链接模式） ✅ 常用
	// 每个块会与上一个密文块异或后再加密。
	// 首个块需要一个随机 IV。
	// ✅ 安全性好
	// ❌ 加密不能并行（因为有依赖）
	// ✅ 解密可以并行
	CBC = "CBC"
	// CFB （加密反馈模式） 🔄 类似串流加密
	// 用前一密文作为输入流来生成密钥流。
	// 适合加密长度不定的串流数据。
	// ✅ 支持字节级加密
	// ❌ 不能并行
	// ⚠️ 一位错误会影响当前和下一个块
	CFB = "CFB"
	// OFB （输出反馈模式） ⏩ 可预计算密钥流
	// 类似 CFB，但只用之前的“加密输出”而不是密文。
	// ✅ 抗误差扩散（错误不会传播）
	// ✅ 可提前生成密钥流
	// ❌ 对 IV 非常敏感
	OFB = "OFB"
	// CTR （计数器模式） ⚡ 高性能之选
	// 把一个递增的“计数器”加密来生成密钥流。
	// ✅ 并行加密 & 解密
	// ✅ 快速、高性能
	// ✅ 常用于 TLS、VPN、SSH 等
	CTR = "CTR"
	// GCM （ Galois/Counter Mode ） ✅ 高安全
	// Authenticated Encryption（加密 + 验证）。
	// ✅ 高性能，硬件支持好（如 Intel AES-NI）。
	// ✅ 不需要填充，适合并行处理。
	// ✅ TLS（HTTPS）、VPN、JWT、数据库加密等高安全需求。
	GCM = "GCM"
)
