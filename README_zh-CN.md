# infra-go

<div align="center">

![Go version](https://img.shields.io/badge/go-1.24-9cf)
[![Tags](https://img.shields.io/badge/tags-v2.0.0-green.svg)](https://github.com/xuzhuoxi/infra-go/tags)
[![GoDoc](https://godoc.org/github.com/xuzhuoxi/infra-go?status.svg)](https://pkg.go.dev/github.com/xuzhuoxi/infra-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuzhuoxi/infra-go)](https://goreportcard.com/report/github.com/xuzhuoxi/infra-go)
[![test](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xuzhuoxi/infra-go/blob/master/LICENSE)

**一个全面的 Go 基础设施库，为现代 Go 应用程序提供必要的工具和实用程序**

[English](./README.md) | [简体中文](./README_zh-CN.md)

</div>

## 📋 项目概述

`infra-go` 是一个模块化的 Go 基础设施库，用于构建健壮的应用程序。涵盖网络通信（TCP/UDP/QUIC/WebSocket/HTTP/RPC）、密码学、图像处理、数据序列化、文件 I/O、日志、集合等常见开发场景。各包可独立导入，耦合度低。

### ✨ 主要特性

- **模块化设计**：按职责划分包，按需导入
- **低耦合**：业务包主要依赖标准库和 `golang.org/x`
- **可重用**：各组件可独立使用
- **可扩展**：面向接口设计，便于自定义处理器与编解码器
- **高性能**：异步 JSON 编解码、对象池、高效二进制封包
- **可靠**：各包均有较完整的 `_test.go` 测试覆盖

## 🚀 快速开始

### 前置条件

- Go 1.24 或更高版本
- Git

### 安装

Go Modules 模式下，在代码中 `import` 即可，依赖会自动解析。当前最新版本标签为 **v2.0.0**。

```bash
go mod init your-project   # 已有 go.mod 可省略
go mod tidy                # 添加 import 后执行
```

### 基本使用

```go
package main

import (
    "github.com/xuzhuoxi/infra-go/logx"
    "github.com/xuzhuoxi/infra-go/netx"
    "github.com/xuzhuoxi/infra-go/netx/tcpx"
)

func main() {
    logger := logx.NewLogger()
    logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelInfo})
    logger.Infoln("启动服务器")

    server := tcpx.NewTCPServer()
    if err := server.StartServer(netx.SockParams{LocalAddress: ":8080"}); err != nil {
        logger.Errorln("启动服务器失败:", err)
    }
}
```

## 📦 包参考

| 包 | 说明 |
|----|------|
| `alg/astar` | 2D/3D A* 寻路与路径平滑 |
| `archivex` | ZIP、TAR 归档创建与解压 |
| `binaryx` | 基于 `io.Reader` / `io.Writer` 的二进制读写 |
| `bytex` | 长度前缀块、线程安全缓冲、字节流水线 |
| `cmdx` | 交互式 CLI 与扩展 flag/参数解析 |
| `cryptox` | 加解密接口、哈希、填充、分组模式 |
| `cryptox/asymmetric` | RSA 密钥管理、PKCS/PEM、混合加密 |
| `cryptox/key` | SHA-256/PBKDF2 密钥派生与 DH 密钥交换 |
| `cryptox/symmetric` | AES（CBC/CTR/GCM）、DES/3DES、XOR |
| `encodingx` | 序列化抽象、缓冲编解码、键值 Map |
| `encodingx/gobx` | Gob 编解码处理器 |
| `encodingx/jsonx` | JSON 编解码处理器（json-iterator） |
| `encodingx/yamlx` | YAML 编解码处理器 |
| `errorsx` | 领域专用错误类型 |
| `eventx` | 轻量级发布-订阅事件分发器 |
| `filex` | 跨平台路径、目录遍历、复制/移动/删除 |
| `graphicx` | 像素打包、色彩空间转换、灰度处理 |
| `graphicx/blendx` | Photoshop 风格图像混合模式 |
| `imagex` | 图像 I/O、像素缓冲、复制填充、邻域操作 |
| `imagex/featurex` | 特征描述子与直方图 |
| `imagex/filterx` | 卷积滤波、模糊、灰度、形态学 |
| `imagex/formatx` | 可插拔图像格式注册表 |
| `imagex/projectionx` | 水平/垂直像素投影统计 |
| `imagex/resizex` | Lanczos3 重采样缩放 |
| `lang` | 值相等、对象池、channel 限制、回调 |
| `lang/collectionx` | 按 ID 索引的有序集合 |
| `lang/listx` | `IList` 及切片/链表实现 |
| `logx` | 多输出、分级、文件滚动、命名 Logger 管理 |
| `mathx` | abs/min/max/clamp、位运算、随机、单位、进制 |
| `mathx/gaussx` | 二维高斯函数与卷积核 |
| `netx` | 统一套接字抽象与消息封包 |
| `netx/httpx` | HTTP 客户端辅助与 `ServeMux` 服务端封装 |
| `netx/quicx` | QUIC 客户端与服务端 |
| `netx/rpcx` | 轻量 `net/rpc` 封装 |
| `netx/tcpx` | TCP 客户端与服务端 |
| `netx/udpx` | UDP 客户端与服务端 |
| `netx/wsx` | WebSocket 客户端与服务端 |
| `osxu` | 进程检测与可执行文件路径工具 |
| `poolx` | 可复用对象池（含字节切片池） |
| `regexpx` | 预置正则表达式模式 |
| `serialx` | 顺序模块生命周期管理器 |
| `servicex` | 服务生命周期编排（init/data/save 流水线） |
| `slicex` | 多类型切片 merge/insert/remove/search 工具 |
| `storagex/mongox` | MongoDB 共享错误类型 |
| `storagex/mongox/mgox` | 官方 `mongo-driver` 多客户端生命周期 |
| `storagex/mongox/mgo.v2x` | `gopkg.in/mgo.v2` 多会话生命周期 |
| `storagex/redisx` | Redis 存储辅助 |
| `stringx` | Unicode 字符串工具与密码校验 |
| `timex` | 时间、时间戳、duration 解析与格式化 |

## 📁 项目结构

```
infra-go/
├── alg/
│   └── astar/                  # A* 寻路（2D/3D）
├── archivex/                   # 归档工具（tar, zip）
├── binaryx/                    # 二进制序列化
├── bytex/                      # 字节切片与缓冲操作
├── cmdx/                       # 命令行工具
├── cryptox/                    # 密码学核心
│   ├── asymmetric/             # RSA、PKCS
│   ├── key/                    # 密钥派生、DH
│   └── symmetric/              # AES、DES、XOR
├── encodingx/                  # 序列化抽象
│   ├── gobx/                   # Gob 编码
│   ├── jsonx/                  # JSON 编码
│   └── yamlx/                  # YAML 编码
├── errorsx/                    # 错误处理工具
├── eventx/                     # 事件系统
├── filex/                      # 文件操作
├── graphicx/                   # 图形与颜色处理
│   └── blendx/                 # 图像混合模式
├── imagex/                     # 图像处理
│   ├── featurex/               # 特征提取
│   ├── filterx/                # 滤镜与效果
│   ├── formatx/                # 格式注册表
│   │   ├── jpegx/
│   │   ├── pngx/
│   │   └── webpx/
│   ├── projectionx/            # 投影统计
│   └── resizex/                # 图像缩放
├── lang/                       # 语言工具
│   ├── collectionx/            # 有序集合
│   └── listx/                  # 列表实现
├── logx/                       # 日志系统
├── mathx/                      # 数学工具
│   └── gaussx/                 # 高斯函数
├── netx/                       # 网络核心（套接字抽象）
│   ├── httpx/                  # HTTP 客户端/服务端
│   ├── quicx/                  # QUIC 协议
│   ├── rpcx/                   # RPC 工具
│   ├── tcpx/                   # TCP 工具
│   ├── udpx/                   # UDP 工具
│   └── wsx/                    # WebSocket 工具
├── osxu/                       # 操作系统工具
├── poolx/                      # 对象池
├── regexpx/                    # 正则表达式工具
├── serialx/                    # 序列化生命周期
├── servicex/                   # 服务管理
├── slicex/                     # 切片工具
├── storagex/
│   ├── mongox/                 # MongoDB 共享类型
│   │   ├── mgox/               # mongo-driver 封装
│   │   └── mgo.v2x/            # mgo.v2 封装
│   └── redisx/                 # Redis 工具
├── stringx/                    # 字符串工具
└── timex/                      # 时间工具
```

## 📚 使用示例

### 日志

```go
import "github.com/xuzhuoxi/infra-go/logx"

logger := logx.NewLogger()
logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
logger.Traceln("hello", "world")

// 包级函数使用默认 Logger
logx.Infoln("使用默认 Logger")
```

### 网络 — TCP

```go
import (
    "github.com/xuzhuoxi/infra-go/netx"
    "github.com/xuzhuoxi/infra-go/netx/tcpx"
)

server := tcpx.NewTCPServer()
server.SetMaxConn(200)
go server.StartServer(netx.SockParams{LocalAddress: "127.0.0.1:9999"})

client := tcpx.NewTCPClient()
client.OpenClient(netx.SockParams{RemoteAddress: "127.0.0.1:9999"})
go client.StartReceiving()
client.SendPackTo([]byte("hello"))
client.CloseClient()
server.StopServer()
```

### 网络 — HTTP

```go
import (
    "fmt"
    "net/http"

    "github.com/xuzhuoxi/infra-go/netx/httpx"
)

// 包级辅助函数
httpx.HttpGet("https://api.example.com/data", func(res *http.Response, body *[]byte) {
    fmt.Println(res.StatusCode, string(*body))
})

// 或使用可配置超时的客户端实例
client := httpx.NewHttpClient()
client.HttpGet("https://api.example.com/data", func(res *http.Response, body *[]byte) {
    fmt.Println(string(*body))
})
```

### 图像处理

```go
import (
    "image"
    "image/color"

    "github.com/xuzhuoxi/infra-go/imagex"
    "github.com/xuzhuoxi/infra-go/imagex/filterx"
    "github.com/xuzhuoxi/infra-go/imagex/formatx"
    "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"

    _ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx" // 注册 JPEG 编解码器
)

// 创建并保存
rect := image.Rect(0, 0, 640, 480)
img := imagex.NewRGBA(rect, color.White)
imagex.SaveImage(img, "output.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)

// 加载并应用模糊
src, _, err := imagex.LoadImage("input.png", formatx.PNG)
dst := imagex.NewRGBA(src.Bounds(), color.Transparent)
filterx.FastBlur2(src, dst, 5)
```

### 密码学 — AES

```go
import (
    "github.com/xuzhuoxi/infra-go/cryptox"
    "github.com/xuzhuoxi/infra-go/cryptox/symmetric"
)

key := []byte("your-32-byte-secret-key-here!!")
plaintext := []byte("Hello, World!")

cipher := symmetric.NewAESCipher(key)
ciphertext, err := cipher.EncryptMode(plaintext, cryptox.CBC)
if err != nil {
    panic(err)
}
decrypted, err := cipher.DecryptMode(ciphertext, cryptox.CBC)
```

### 数据序列化 — JSON

```go
import "github.com/xuzhuoxi/infra-go/encodingx/jsonx"

handler := jsonx.NewJsonCodingHandlerAsync()

data := map[string]interface{}{"name": "John", "age": 30}
jsonBytes, err := handler.HandleEncode(data)

var decoded map[string]interface{}
err = handler.HandleDecode(jsonBytes, &decoded)
```

## 🛠️ 技术栈

### 核心

- **Go 标准库** — 主要基础
- **golang.org/x** — crypto、image、net 扩展

### 第三方依赖

| 库 | 用途 |
|----|------|
| [json-iterator/go](https://github.com/json-iterator/go) | 高性能 JSON |
| [quic-go/quic-go](https://github.com/quic-go/quic-go) | QUIC 协议 |
| [nfnt/resize](https://github.com/nfnt/resize) | 图像缩放 |
| [mongo-driver](https://go.mongodb.org/mongo-driver) | MongoDB 官方驱动 |
| [gopkg.in/mgo.v2](https://gopkg.in/mgo.v2) | 旧版 MongoDB 驱动封装 |
| [golang.org/x/crypto](https://golang.org/x/crypto) | 加密原语 |
| [golang.org/x/image](https://golang.org/x/image) | 图像处理 |
| [golang.org/x/net](https://golang.org/x/net) | 网络工具 |
| [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) | YAML 配置 |

## 🧪 测试

```bash
# 运行所有测试
go test ./...

# 运行测试并显示覆盖率
go test -cover ./...

# 运行特定包
go test ./netx/tcpx
```

## 📖 文档

- [pkg.go.dev](https://pkg.go.dev/github.com/xuzhuoxi/infra-go) — API 参考
- 各包 `doc.go` — 设计说明与能力概述
- 仓库内 `*_test.go` — 可运行的用法示例

## 🤝 贡献

欢迎贡献！请按以下步骤操作：

1. **Fork** 仓库
2. **创建** 功能分支（`git checkout -b feature/my-feature`）
3. **提交** 更改并附上清晰的提交说明
4. **推送** 分支并 **打开** Pull Request

### 开发指南

- 遵循 Go 规范；提交前运行 `gofmt`
- 为新功能添加测试
- 修改公开 API 时同步更新文档
- 提交前确保 `go test ./...` 通过

## 📄 许可证

本项目采用 MIT 许可证 — 详见 [LICENSE](LICENSE)。

## 👥 作者

- **xuzhuoxi** — [GitHub](https://github.com/xuzhuoxi)

## 📞 联系方式

- **邮箱**：xuzhuoxi@gmail.com 或 mailxuzhuoxi@163.com
- **GitHub**：[@xuzhuoxi](https://github.com/xuzhuoxi)

---

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>
