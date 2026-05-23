package bytex

import "encoding/binary"

// DefaultOrder is the default byte order used by package-level helpers.
// DefaultOrder 是包级 helper 使用的默认字节序。
var DefaultOrder binary.ByteOrder = binary.LittleEndian

// DefaultDataBlockHandler is the default length-prefix block encoder and decoder.
// DefaultDataBlockHandler 是默认的长度前缀块编解码处理器。
var DefaultDataBlockHandler = NewDefaultDataBlockHandler()
