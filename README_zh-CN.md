# infra-go

<div align=center>

<br/>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.16-9cf)
[![Tags](https://img.shields.io/badge/tags-1.0.4-green.svg)](https://github.com/xuzhuoxi/infra-go/tags)
[![GoDoc](https://godoc.org/github.com/xuzhuoxi/infra-go?status.svg)](https://pkg.go.dev/github.com/xuzhuoxi/infra-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuzhuoxi/infra-go)](https://goreportcard.com/report/github.com/xuzhuoxi/infra-go)
[![test](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xuzhuoxi/infra-go/blob/master/LICENSE)

</div>

<div STYLE="page-break-after: always;"></div>
<p style="font-size: 20px"> 
  A go library.
</p>

[English](./README.md) | 简体中文

## 功能特性
+ **核心依赖**：主要依赖 go 标准库 及 golang.org/x 扩展库。
+ **业务依赖**：具有业务性质的功能可能会依赖第三方库，例如：读取yaml配置、使用quic通信等。
+ **偶合度**：核心依赖相关包的偶合度一般，业务性质相关的包偶合度低高。
+ **可复用**
+ **可扩展**

## 安装

1. 要求使用 go1.16有以上版本。  
2. `go get github.com/xuzhuoxi/infra-go`  

## 用法
+ infra-go 以包结构组织，导入即可使用。
+ 例如： 使用TCP服务器功能
```go
import "github.com/xuzhuoxi/infra-go/netx/tcpx"
```

## 示例
```go
```

## 文档
<details>
<summary>展开代码说明(中文)</summary>
<pre><code>.
├── alg: 通用算法
│   ├── astar: 一个支持二维与三维的A星寻路算法
├── binaryx: 二进制数据的序列化与反序列化
├── bytex: 字节切片及缓存的序列化与反序列化
├── cmdx: 控制台命令行监听，解释与处理
├── cryptox: 加解密
├── encodingx: 编码与解码
│   ├── gobx: Gob编码与解码
│   ├── jsonx: Json编码与解码
├── graphicx: 图像处理库，图像色彩相关功能
│   ├── blendx: 混合模式支持
├── imagex: 图片处理库，包含各种图片格式的加载保存处理等功能
│   ├── formatx: 图片格式支持
│   │   ├── jpegx: jpg,jpeg,jps支持
│   │   ├── pngx: png支持
│   ├── resizex: 缩放支持
├── errorsx: 错误异常
├── eventx: 一个简单的事件处理模块
├── extendx: 通用扩展模块
│   ├── protox: 通用协议扩展模块
├── lang: 一些通用常用的功能函数
│   ├── listx: 列表，包含数据实现和链表实现
├── logx:  日志模块
├── mathx: 数学函数集
├── netx:  网络库，包含http,rpc,quic,tcp,udp,websocket等实现，同步支持服务端与客户端
├── osxu:  操作系统级的常用函数
├── regexpx: 常用正则表达式
├── slicex: 关于基本类型切片的常用函数
├── stringx: 关于字符串处理的常用函数
├── timex: 关于计时器的常用函数
</code></pre>
</details>  

## Contact
xuzhuoxi  
<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>  

## License
"infra-go" source code is available under the MIT [License](/LICENSE).
