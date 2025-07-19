# infra-go

<div align="center">

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.16-9cf)
[![Tags](https://img.shields.io/badge/tags-1.0.4-green.svg)](https://github.com/xuzhuoxi/infra-go/tags)
[![GoDoc](https://godoc.org/github.com/xuzhuoxi/infra-go?status.svg)](https://pkg.go.dev/github.com/xuzhuoxi/infra-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuzhuoxi/infra-go)](https://goreportcard.com/report/github.com/xuzhuoxi/infra-go)
[![test](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xuzhuoxi/infra-go/blob/master/LICENSE)

**A comprehensive Go infrastructure library providing essential utilities and tools for modern Go applications**

[English](./README.md) | [简体中文](./README_zh-CN.md)

</div>

## 📖 Overview

`infra-go` is a comprehensive Go infrastructure library that provides a wide range of utilities, tools, and components for building robust Go applications. It offers modular, reusable, and extensible solutions for common development tasks including networking, cryptography, image processing, data serialization, and more.

### ✨ Key Features

- **🔄 Modular Design**: Organized in package structure for easy import and use
- **🔗 Low Coupling**: Business-related packages have minimal dependencies
- **📦 Reusable**: Components can be used independently
- **🚀 Extensible**: Easy to extend and customize
- **⚡ Performance**: Optimized for high-performance applications
- **🛡️ Reliable**: Comprehensive test coverage and error handling

## 🚀 Quick Start

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

```bash
# Install the library
go get github.com/xuzhuoxi/infra-go

# Or using go mod
go mod init your-project
go get github.com/xuzhuoxi/infra-go
```

### Basic Usage

```go
package main

import (
    "github.com/xuzhuoxi/infra-go/netx/tcpx"
    "github.com/xuzhuoxi/infra-go/logx"
)

func main() {
    // Initialize logger
    logger := logx.NewLogger()
    
    // Create TCP server
    server := tcpx.NewServer(":8080")
    
    // Start server
    if err := server.Start(); err != nil {
        logger.Error("Failed to start server:", err)
    }
}
```

## 🛠️ Technology Stack

### Core Dependencies
- **Go Standard Library**: Primary foundation
- **golang.org/x**: Extended functionality

### Key Third-party Libraries
- **json-iterator/go**: High-performance JSON processing
- **lucas-clemente/quic-go**: QUIC protocol implementation
- **nfnt/resize**: Image resizing capabilities
- **mongo-driver**: MongoDB database driver
- **golang.org/x/crypto**: Cryptographic functions
- **golang.org/x/image**: Image processing
- **golang.org/x/net**: Network utilities
- **gopkg.in/yaml.v2**: YAML configuration support

## 📁 Project Structure

```
infra-go/
├── alg/                    # Common algorithms
│   └── astar/             # A* pathfinding algorithm (2D/3D)
├── archivex/              # Archive utilities (tar, zip)
├── binaryx/               # Binary data serialization
├── bytex/                 # Byte slice and buffer operations
├── cmdx/                  # Command line interface utilities
├── cryptox/               # Cryptography and encryption
│   ├── asymmetric/        # Asymmetric encryption (RSA, PKCS)
│   └── symmetric/         # Symmetric encryption (AES, DES, XOR)
├── encodingx/             # Data encoding/decoding
│   ├── gobx/             # Gob encoding
│   ├── jsonx/            # JSON encoding
│   └── yamlx/            # YAML encoding
├── errorsx/               # Error handling utilities
├── eventx/                # Event system
├── filex/                 # File operations
├── graphicx/              # Graphics and color processing
│   └── blendx/           # Image blending modes
├── imagex/                # Image processing library
│   ├── featurex/         # Image feature extraction
│   ├── filterx/          # Image filters and effects
│   ├── formatx/          # Image format support
│   └── resizex/          # Image resizing
├── lang/                  # Language utilities
│   ├── collectionx/      # Collection utilities
│   └── listx/            # List implementations
├── logx/                  # Logging system
├── mathx/                 # Mathematical utilities
│   └── gaussx/           # Gaussian functions
├── netx/                  # Network utilities
│   ├── httpx/            # HTTP client/server
│   ├── quicx/            # QUIC protocol
│   ├── rpcx/             # RPC utilities
│   ├── tcpx/             # TCP utilities
│   ├── udpx/             # UDP utilities
│   └── wsx/              # WebSocket utilities
├── osxu/                  # Operating system utilities
├── poolx/                 # Object pooling
├── regexpx/               # Regular expression utilities
├── serialx/               # Serialization utilities
├── servicex/              # Service management
├── slicex/                # Slice utilities
├── storagex/              # Storage utilities
│   ├── mongox/           # MongoDB utilities
│   └── redisx/           # Redis utilities
├── stringx/               # String utilities
├── syncx/                 # Synchronization utilities
└── timex/                 # Time utilities
```

## 📚 Usage Examples

### Network Operations

```go
// TCP Server
import "github.com/xuzhuoxi/infra-go/netx/tcpx"

server := tcpx.NewServer(":8080")
server.Start()

// HTTP Client
import "github.com/xuzhuoxi/infra-go/netx/httpx"

client := httpx.NewClient()
resp, err := client.Get("https://api.example.com/data")
```

### Image Processing

```go
// Load and process image
import "github.com/xuzhuoxi/infra-go/imagex"

img, err := imagex.LoadImage("input.jpg")
if err != nil {
    log.Fatal(err)
}

// Apply blur filter
blurred := imagex.ApplyBlur(img, 5.0)
imagex.SaveImage(blurred, "output.jpg")
```

### Cryptography

```go
// AES encryption
import "github.com/xuzhuoxi/infra-go/cryptox/symmetric"

key := []byte("your-secret-key-32-bytes-long!")
data := []byte("Hello, World!")

encrypted, err := symmetric.AESEncrypt(data, key)
if err != nil {
    log.Fatal(err)
}
```

### Data Serialization

```go
// JSON encoding
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

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./netx/tcpx
```

## 📖 Documentation

- [GoDoc](https://pkg.go.dev/github.com/xuzhuoxi/infra-go) - API documentation
- [Examples](./examples/) - Code examples
- [Tests](./*/_test.go) - Test files with usage examples

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Development Guidelines

- Follow Go coding standards and conventions
- Add tests for new functionality
- Update documentation for API changes
- Ensure all tests pass before submitting PR
- Use meaningful commit messages

### Code Style

- Use `gofmt` for code formatting
- Follow Go naming conventions
- Add comments for exported functions
- Keep functions focused and concise

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **xuzhuoxi** - *Initial work* - [GitHub](https://github.com/xuzhuoxi)

## 🙏 Acknowledgments

- Go community for excellent tooling and libraries
- Contributors and users of this project
- Open source projects that inspired this library

## 📞 Contact

- **Email**: xuzhuoxi@gmail.com or mailxuzhuoxi@163.com
- **GitHub**: [@xuzhuoxi](https://github.com/xuzhuoxi)

## 🔗 Related Projects

- [Go Standard Library](https://golang.org/pkg/)
- [golang.org/x](https://golang.org/x/)

---

<div align="center">

**Made with ❤️ by the Go community**

[![Go](https://img.shields.io/badge/Go-1.16+-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>