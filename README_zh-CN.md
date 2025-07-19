# infra-go

<div align="center">

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.16-9cf)
[![Tags](https://img.shields.io/badge/tags-1.0.4-green.svg)](https://github.com/xuzhuoxi/infra-go/tags)
[![GoDoc](https://godoc.org/github.com/xuzhuoxi/infra-go?status.svg)](https://pkg.go.dev/github.com/xuzhuoxi/infra-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuzhuoxi/infra-go)](https://goreportcard.com/report/github.com/xuzhuoxi/infra-go)
[![test](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xuzhuoxi/infra-go/blob/master/LICENSE)

**一个全面的 Go 基础设施库，为现代 Go 应用程序提供必要的工具和实用程序**

[English](./README.md) | [简体中文](./README_zh-CN.md)

</div>

## 📖 项目概述

`infra-go` 是一个全面的 Go 基础设施库，为构建健壮的 Go 应用程序提供广泛的工具、实用程序和组件。它为常见的开发任务提供模块化、可重用和可扩展的解决方案，包括网络通信、密码学、图像处理、数据序列化等。

### ✨ 主要特性

- **🔄 模块化设计**：以包结构组织，便于导入和使用
- **🔗 低耦合**：业务相关包具有最少的依赖关系
- **📦 可重用**：组件可以独立使用
- **🚀 可扩展**：易于扩展和自定义
- **⚡ 高性能**：针对高性能应用程序优化
- **🛡️ 可靠**：全面的测试覆盖和错误处理

## 🚀 快速开始

### 前置条件

- Go 1.16 或更高版本
- Git

### 安装

```bash
# 安装库
go get github.com/xuzhuoxi/infra-go

# 或使用 go mod
go mod init your-project
go get github.com/xuzhuoxi/infra-go
```

### 基本使用

```go
package main

import (
    "github.com/xuzhuoxi/infra-go/netx/tcpx"
    "github.com/xuzhuoxi/infra-go/logx"
)

func main() {
    // 初始化日志记录器
    logger := logx.NewLogger()
    
    // 创建 TCP 服务器
    server := tcpx.NewServer(":8080")
    
    // 启动服务器
    if err := server.Start(); err != nil {
        logger.Error("启动服务器失败:", err)
    }
}
```

## 🛠️ 技术栈

### 核心依赖
- **Go 标准库**：主要基础
- **golang.org/x**：扩展功能

### 关键第三方库
- **json-iterator/go**：高性能 JSON 处理
- **lucas-clemente/quic-go**：QUIC 协议实现
- **nfnt/resize**：图像缩放功能
- **mongo-driver**：MongoDB 数据库驱动
- **golang.org/x/crypto**：加密函数
- **golang.org/x/image**：图像处理
- **golang.org/x/net**：网络工具
- **gopkg.in/yaml.v2**：YAML 配置支持

## 📁 项目结构

```
infra-go/
├── alg/                    # 通用算法
│   └── astar/             # A* 寻路算法（2D/3D）
├── archivex/              # 归档工具（tar, zip）
├── binaryx/               # 二进制数据序列化
├── bytex/                 # 字节切片和缓冲区操作
├── cmdx/                  # 命令行界面工具
├── cryptox/               # 密码学和加密
│   ├── asymmetric/        # 非对称加密（RSA, PKCS）
│   └── symmetric/         # 对称加密（AES, DES, XOR）
├── encodingx/             # 数据编码/解码
│   ├── gobx/             # Gob 编码
│   ├── jsonx/            # JSON 编码
│   └── yamlx/            # YAML 编码
├── errorsx/               # 错误处理工具
├── eventx/                # 事件系统
├── filex/                 # 文件操作
├── graphicx/              # 图形和颜色处理
│   └── blendx/           # 图像混合模式
├── imagex/                # 图像处理库
│   ├── featurex/         # 图像特征提取
│   ├── filterx/          # 图像滤镜和效果
│   ├── formatx/          # 图像格式支持
│   └── resizex/          # 图像缩放
├── lang/                  # 语言工具
│   ├── collectionx/      # 集合工具
│   └── listx/            # 列表实现
├── logx/                  # 日志系统
├── mathx/                 # 数学工具
│   └── gaussx/           # 高斯函数
├── netx/                  # 网络工具
│   ├── httpx/            # HTTP 客户端/服务器
│   ├── quicx/            # QUIC 协议
│   ├── rpcx/             # RPC 工具
│   ├── tcpx/             # TCP 工具
│   ├── udpx/             # UDP 工具
│   └── wsx/              # WebSocket 工具
├── osxu/                  # 操作系统工具
├── poolx/                 # 对象池
├── regexpx/               # 正则表达式工具
├── serialx/               # 序列化工具
├── servicex/              # 服务管理
├── slicex/                # 切片工具
├── storagex/              # 存储工具
│   ├── mongox/           # MongoDB 工具
│   └── redisx/           # Redis 工具
├── stringx/               # 字符串工具
├── syncx/                 # 同步工具
└── timex/                 # 时间工具
```

## 📚 使用示例

### 网络操作

```go
// TCP 服务器
import "github.com/xuzhuoxi/infra-go/netx/tcpx"

server := tcpx.NewServer(":8080")
server.Start()

// HTTP 客户端
import "github.com/xuzhuoxi/infra-go/netx/httpx"

client := httpx.NewClient()
resp, err := client.Get("https://api.example.com/data")
```

### 图像处理

```go
// 加载和处理图像
import "github.com/xuzhuoxi/infra-go/imagex"

img, err := imagex.LoadImage("input.jpg")
if err != nil {
    log.Fatal(err)
}

// 应用模糊滤镜
blurred := imagex.ApplyBlur(img, 5.0)
imagex.SaveImage(blurred, "output.jpg")
```

### 密码学

```go
// AES 加密
import "github.com/xuzhuoxi/infra-go/cryptox/symmetric"

key := []byte("your-secret-key-32-bytes-long!")
data := []byte("Hello, World!")

encrypted, err := symmetric.AESEncrypt(data, key)
if err != nil {
    log.Fatal(err)
}
```

### 数据序列化

```go
// JSON 编码
import "github.com/xuzhuoxi/infra-go/encodingx/jsonx"

data := map[string]interface{}{
    "name": "John",
    "age":  30,
}

jsonBytes, err := jsonx.Encode(data)
if err != nil {
    log.Fatal(err)
}
```

## 🧪 测试

运行测试套件：

```bash
# 运行所有测试
go test ./...

# 运行测试并显示覆盖率
go test -cover ./...

# 运行特定包的测试
go test ./netx/tcpx
```

## 📖 文档

- [GoDoc](https://pkg.go.dev/github.com/xuzhuoxi/infra-go) - API 文档
- [示例](./examples/) - 代码示例
- [测试](./*/_test.go) - 包含使用示例的测试文件

## 🤝 贡献

我们欢迎贡献！请按照以下步骤：

1. **Fork** 仓库
2. **创建** 功能分支 (`git checkout -b feature/amazing-feature`)
3. **提交** 您的更改 (`git commit -m '添加令人惊叹的功能'`)
4. **推送** 到分支 (`git push origin feature/amazing-feature`)
5. **打开** 拉取请求

### 开发指南

- 遵循 Go 编码标准和约定
- 为新功能添加测试
- 更新 API 更改的文档
- 在提交 PR 前确保所有测试通过
- 使用有意义的提交消息

### 代码风格

- 使用 `gofmt` 进行代码格式化
- 遵循 Go 命名约定
- 为导出的函数添加注释
- 保持函数专注和简洁

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 👥 作者

- **xuzhuoxi** - *初始工作* - [GitHub](https://github.com/xuzhuoxi)

## 🙏 致谢

- Go 社区提供优秀的工具和库
- 本项目的贡献者和用户
- 启发本库的开源项目

## 📞 联系方式

- **邮箱**：xuzhuoxi@gmail.com 或 mailxuzhuoxi@163.com
- **GitHub**：[@xuzhuoxi](https://github.com/xuzhuoxi)

## 🔗 相关项目

- [Go 标准库](https://golang.org/pkg/)
- [golang.org/x](https://golang.org/x/)

---

<div align="center">

**由 Go 社区用 ❤️ 制作**

[![Go](https://img.shields.io/badge/Go-1.16+-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>
