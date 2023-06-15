// Package cryptox
// Created by xuzhuoxi
// on 2019-02-05.
// @author xuzhuoxi
//
package cryptox

// PKCS#1：定义RSA公开密钥算法加密和签名机制，主要用于组织PKCS#7中所描述的数字签名和数字信封[22]。
// PKCS#3：定义Diffie-Hellman密钥交换协议[23]。
// PKCS#5：描述一种利用从口令派生出来的安全密钥加密字符串的方法。使用MD2或MD5 从口令中派生密钥，并采用DES-CBC模式加密。主要用于加密从一个计算机传送到另一个计算机的私人密钥，不能用于加密消息[24]。
// PKCS#6：描述了公钥证书的标准语法，主要描述X.509证书的扩展格式[25]。
// PKCS#7：定义一种通用的消息语法，包括数字签名和加密等用于增强的加密机制，PKCS#7与PEM兼容，所以不需其他密码操作，就可以将加密的消息转换成PEM消息[26]。
// PKCS#8：描述私有密钥信息格式，该信息包括公开密钥算法的私有密钥以及可选的属性集等[27]。
// PKCS#9：定义一些用于PKCS#6证书扩展、PKCS#7数字签名和PKCS#8私钥加密信息的属性类型[28]。
// PKCS#10：描述证书请求语法[29]。
// PKCS#11：称为Cyptoki，定义了一套独立于技术的程序设计接口，用于智能卡和PCMCIA卡之类的加密设备[30]。
// PKCS#12：描述个人信息交换语法标准。描述了将用户公钥、私钥、证书和其他相关信息打包的语法[31]。
// PKCS#13：椭圆曲线密码体制标准[32]。
// PKCS#14：伪随机数生成标准。
// PKCS#15：密码令牌信息格式标准[33]。
