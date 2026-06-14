# infra-go

<div align="center">

![Go version](https://img.shields.io/badge/go-1.24-9cf)
[![Tags](https://img.shields.io/badge/tags-v2.0.0-green.svg)](https://github.com/xuzhuoxi/infra-go/tags)
[![GoDoc](https://godoc.org/github.com/xuzhuoxi/infra-go?status.svg)](https://pkg.go.dev/github.com/xuzhuoxi/infra-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuzhuoxi/infra-go)](https://goreportcard.com/report/github.com/xuzhuoxi/infra-go)
[![test](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml/badge.svg?branch=main&event=push)](https://github.com/xuzhuoxi/infra-go/actions/workflows/codecov.yml)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/xuzhuoxi/infra-go/blob/master/LICENSE)

**A comprehensive Go infrastructure library providing essential utilities and tools for modern Go applications**

[English](./README.md) | [简体中文](./README_zh-CN.md)

</div>

## 📋 Overview

`infra-go` is a modular Go infrastructure library for building robust applications. It covers networking (TCP/UDP/QUIC/WebSocket/HTTP/RPC), cryptography, image processing, data serialization, file I/O, logging, collections, and more. Packages are designed to be imported independently with minimal coupling.

### ✨ Key Features

- **Modular Design**: One package per concern; import only what you need
- **Low Coupling**: Business packages depend mainly on the standard library and `golang.org/x`
- **Reusable**: Each component works standalone
- **Extensible**: Interface-driven design for custom handlers and codecs
- **Performance-Oriented**: Async JSON codecs, object pools, and efficient binary framing
- **Reliable**: Broad `_test.go` coverage across packages

## 🚀 Quick Start

### Prerequisites

- Go 1.24 or higher
- Git

### Installation

In Go modules mode, add imports in your code; dependencies are resolved automatically. The current release tag is **v2.0.0**.

```bash
go mod init your-project   # skip if go.mod already exists
go mod tidy                # after adding imports
```

### Basic Usage

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
    logger.Infoln("starting server")

    server := tcpx.NewTCPServer()
    if err := server.StartServer(netx.SockParams{LocalAddress: ":8080"}); err != nil {
        logger.Errorln("failed to start server:", err)
    }
}
```

## 📦 Package Reference

| Package | Description |
|---------|-------------|
| `alg/astar` | 2D/3D A* pathfinding and path smoothing |
| `archivex` | ZIP and TAR archive creation and extraction |
| `binaryx` | Binary read/write over `io.Reader` / `io.Writer` |
| `bytex` | Length-prefixed blocks, thread-safe buffers, byte pipelines |
| `cmdx` | Interactive CLI and extended flag/argument parsing |
| `cryptox` | Cipher interfaces, hashing, padding, block modes |
| `cryptox/asymmetric` | RSA key management, PKCS/PEM, hybrid encryption |
| `cryptox/key` | SHA-256/PBKDF2 key derivation and DH key exchange |
| `cryptox/symmetric` | AES (CBC/CTR/GCM), DES/3DES, XOR |
| `encodingx` | Serialization abstractions, buffered codecs, key-value maps |
| `encodingx/gobx` | Gob encoding handler |
| `encodingx/jsonx` | JSON encoding handler (json-iterator) |
| `encodingx/yamlx` | YAML encoding handler |
| `errorsx` | Domain-specific error types |
| `eventx` | Lightweight publish-subscribe event dispatcher |
| `filex` | Cross-platform paths, traversal, copy/move/delete |
| `graphicx` | Pixel packing, color space conversion, grayscale |
| `graphicx/blendx` | Photoshop-style image blend modes |
| `imagex` | Image I/O, pixel buffers, copy/fill, neighborhood ops |
| `imagex/featurex` | Feature descriptors and histograms |
| `imagex/filterx` | Convolution filters, blur, grayscale, morphology |
| `imagex/formatx` | Pluggable image format registry |
| `imagex/projectionx` | Horizontal/vertical pixel projection statistics |
| `imagex/resizex` | Lanczos3 resampling resize |
| `lang` | Value equality, object pools, channel limits, callbacks |
| `lang/collectionx` | ID-indexed ordered collections |
| `lang/listx` | `IList` with slice and linked-list implementations |
| `logx` | Multi-sink logging with levels, rotation, named managers |
| `mathx` | abs/min/max/clamp, bit ops, random, units, radix |
| `mathx/gaussx` | 2D Gaussian functions and convolution kernels |
| `netx` | Unified socket abstractions and pack framing |
| `netx/httpx` | HTTP client helpers and `ServeMux` server wrapper |
| `netx/quicx` | QUIC client and server |
| `netx/rpcx` | Lightweight `net/rpc` wrapper |
| `netx/tcpx` | TCP client and server |
| `netx/udpx` | UDP client and server |
| `netx/wsx` | WebSocket client and server |
| `osxu` | Process detection and executable path utilities |
| `poolx` | Reusable object pools (including byte-slice pools) |
| `regexpx` | Predefined regular expression patterns |
| `serialx` | Sequential module lifecycle manager |
| `servicex` | Service lifecycle orchestration (init/data/save pipeline) |
| `slicex` | Multi-type slice merge/insert/remove/search utilities |
| `storagex/mongox` | Shared MongoDB error types |
| `storagex/mongox/mgox` | Official `mongo-driver` multi-client lifecycle |
| `storagex/mongox/mgo.v2x` | `gopkg.in/mgo.v2` multi-session lifecycle |
| `storagex/redisx` | Redis storage helpers |
| `stringx` | Unicode string utilities and password validation |
| `timex` | Time, timestamp, duration parsing and formatting |

## 📁 Project Structure

```
infra-go/
├── alg/
│   └── astar/                  # A* pathfinding (2D/3D)
├── archivex/                   # Archive utilities (tar, zip)
├── binaryx/                    # Binary serialization
├── bytex/                      # Byte slice and buffer operations
├── cmdx/                       # Command-line utilities
├── cryptox/                    # Cryptography core
│   ├── asymmetric/             # RSA, PKCS
│   ├── key/                    # Key derivation, DH
│   └── symmetric/              # AES, DES, XOR
├── encodingx/                  # Serialization abstractions
│   ├── gobx/                   # Gob encoding
│   ├── jsonx/                  # JSON encoding
│   └── yamlx/                  # YAML encoding
├── errorsx/                    # Error handling utilities
├── eventx/                     # Event system
├── filex/                      # File operations
├── graphicx/                   # Graphics and color processing
│   └── blendx/                 # Image blend modes
├── imagex/                     # Image processing
│   ├── featurex/               # Feature extraction
│   ├── filterx/                # Filters and effects
│   ├── formatx/                # Format registry
│   │   ├── jpegx/
│   │   ├── pngx/
│   │   └── webpx/
│   ├── projectionx/            # Projection statistics
│   └── resizex/                # Image resizing
├── lang/                       # Language utilities
│   ├── collectionx/            # Ordered collections
│   └── listx/                  # List implementations
├── logx/                       # Logging system
├── mathx/                      # Mathematical utilities
│   └── gaussx/                 # Gaussian functions
├── netx/                       # Network core (socket abstractions)
│   ├── httpx/                  # HTTP client/server
│   ├── quicx/                  # QUIC protocol
│   ├── rpcx/                   # RPC utilities
│   ├── tcpx/                   # TCP utilities
│   ├── udpx/                   # UDP utilities
│   └── wsx/                    # WebSocket utilities
├── osxu/                       # OS utilities
├── poolx/                      # Object pooling
├── regexpx/                    # Regular expression utilities
├── serialx/                    # Serialization lifecycle
├── servicex/                   # Service management
├── slicex/                     # Slice utilities
├── storagex/
│   ├── mongox/                 # MongoDB shared types
│   │   ├── mgox/               # mongo-driver wrapper
│   │   └── mgo.v2x/            # mgo.v2 wrapper
│   └── redisx/                 # Redis utilities
├── stringx/                    # String utilities
└── timex/                      # Time utilities
```

## 📚 Usage Examples

### Logging

```go
import "github.com/xuzhuoxi/infra-go/logx"

logger := logx.NewLogger()
logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
logger.Traceln("hello", "world")

// Package-level helpers on the default logger
logx.Infoln("using default logger")
```

### Network — TCP

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

### Network — HTTP

```go
import (
    "fmt"
    "net/http"

    "github.com/xuzhuoxi/infra-go/netx/httpx"
)

// Package-level helpers
httpx.HttpGet("https://api.example.com/data", func(res *http.Response, body *[]byte) {
    fmt.Println(res.StatusCode, string(*body))
})

// Or use a client instance with custom timeout
client := httpx.NewHttpClient()
client.HttpGet("https://api.example.com/data", func(res *http.Response, body *[]byte) {
    fmt.Println(string(*body))
})
```

### Image Processing

```go
import (
    "image"
    "image/color"

    "github.com/xuzhuoxi/infra-go/imagex"
    "github.com/xuzhuoxi/infra-go/imagex/filterx"
    "github.com/xuzhuoxi/infra-go/imagex/formatx"
    "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"

    _ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx" // register JPEG codec
)

// Create and save
rect := image.Rect(0, 0, 640, 480)
img := imagex.NewRGBA(rect, color.White)
imagex.SaveImage(img, "output.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)

// Load and apply blur
src, _, err := imagex.LoadImage("input.png", formatx.PNG)
dst := imagex.NewRGBA(src.Bounds(), color.Transparent)
filterx.FastBlur2(src, dst, 5)
```

### Cryptography — AES

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

### Data Serialization — JSON

```go
import "github.com/xuzhuoxi/infra-go/encodingx/jsonx"

handler := jsonx.NewJsonCodingHandlerAsync()

data := map[string]interface{}{"name": "John", "age": 30}
jsonBytes, err := handler.HandleEncode(data)

var decoded map[string]interface{}
err = handler.HandleDecode(jsonBytes, &decoded)
```

## 🛠️ Technology Stack

### Core

- **Go Standard Library** — primary foundation
- **golang.org/x** — crypto, image, net extensions

### Third-party Dependencies

| Library | Purpose |
|---------|---------|
| [json-iterator/go](https://github.com/json-iterator/go) | High-performance JSON |
| [quic-go/quic-go](https://github.com/quic-go/quic-go) | QUIC protocol |
| [nfnt/resize](https://github.com/nfnt/resize) | Image resizing |
| [mongo-driver](https://go.mongodb.org/mongo-driver) | MongoDB official driver |
| [gopkg.in/mgo.v2](https://gopkg.in/mgo.v2) | Legacy MongoDB driver wrapper |
| [golang.org/x/crypto](https://golang.org/x/crypto) | Cryptographic primitives |
| [golang.org/x/image](https://golang.org/x/image) | Image processing |
| [golang.org/x/net](https://golang.org/x/net) | Network utilities |
| [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) | YAML configuration |

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run a specific package
go test ./netx/tcpx
```

## 📖 Documentation

- [pkg.go.dev](https://pkg.go.dev/github.com/xuzhuoxi/infra-go) — API reference
- Each package's `doc.go` — design notes and capability overview
- `*_test.go` files throughout the repo — runnable usage examples

## 🤝 Contributing

Contributions are welcome. Please:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/my-feature`)
3. **Commit** your changes with a clear message
4. **Push** to the branch and **open** a Pull Request

### Guidelines

- Follow Go conventions; run `gofmt` before committing
- Add tests for new functionality
- Update documentation when changing public APIs
- Ensure `go test ./...` passes before submitting

## 📄 License

This project is licensed under the MIT License — see [LICENSE](LICENSE) for details.

## 👥 Authors

- **xuzhuoxi** — [GitHub](https://github.com/xuzhuoxi)

## 📞 Contact

- **Email**: xuzhuoxi@gmail.com or mailxuzhuoxi@163.com
- **GitHub**: [@xuzhuoxi](https://github.com/xuzhuoxi)

---

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

</div>
