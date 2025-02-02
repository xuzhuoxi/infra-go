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

English | [简体中文](./README_zh-CN.md)

## Features
+ **Core dependencies**: mainly depend on go standard library and golang.org/x extension library.  
+ **Business dependency**: Functions with business nature may depend on third-party libraries, for example: reading yaml configuration, using quic communication, etc.  
+ **Coupling degree**: The coupling degree of core dependencies related packages is average, and the coupling degree of business-related packages is low.  
+ **Reusable**  
+ **Extensible**  

## Installation
1. It is required to use go1.16 to have the above version.  
2. `go get github.com/xuzhuoxi/infra-go`  

## Usage
+ infra-go is organized in a package structure, and it can be used immediately after importing.  
+ For example: using the TCP server function  
```go
import "github.com/xuzhuoxi/infra-go/netx/tcpx"
```

## Example
```go

```

## Documentation
<details>
<summary>Expand view</summary>
<pre><code>.
├── alg: Common algorithm
│   ├── astar: AStar algorithm supported 2D 3D static path finding.
├── binaryx: Binary data serialization and deserialization.
├── bytex: Byte slice and byte buff serialization and deserialization.
├── cmdx: Command line input listening, interpretation and processing.
├── cryptox: Encrypt.
├── encodingx: Encode and decode.
│   ├── gobx: Gob encode and decode.
│   ├── jsonx: Json encode and decode.
├── graphicx: Graphic and color processing library, image color correlation function
│   ├── blendx: blend mode support
├── imagex: Image processing library, including loading and saving processing of various image formats
│   ├── formatx: Image format support
│   │   ├── jpegx: jpg,jpeg,jps support
│   │   ├── pngx: png support
│   ├── resizex: resize support
├── errorsx: error
├── eventx: A simple event module.
├── extendx: Common extension.
│   ├── protox: Proto Extension.
├── lang: Some commonly used functions for go language.
│   ├── listx: go list
├── logx:  A log module
├── mathx: A set of math methods.
├── netx:  Net module, include server and client module.
├── osxu:  A set of function for OS.
├── regexpx: A set of commonly used regular expressions
├── slicex: A set of slice functions for basic structure.
├── stringx: A set of functions for string.
├── timex: A set of functions for timer
</code></pre>
</details>  

## Contact
xuzhuoxi   
<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>  

## License
"infra-go" source code is available under the MIT [License](/LICENSE).