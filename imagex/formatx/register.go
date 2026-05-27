package formatx

import (
	"image"
	"io"
	"strings"
)

// ImageEncodeFunc encodes an image to w with optional format-specific options.
// ImageEncodeFunc 将图像编码写入 w，options 为格式相关选项。
type ImageEncodeFunc = func(w io.Writer, m image.Image, options interface{}) error

// ImageDecodeFunc decodes an image from r.
// ImageDecodeFunc 从 r 解码图像。
type ImageDecodeFunc = func(w io.Reader) (image.Image, error)

type formatDefined struct {
	name   string
	encode ImageEncodeFunc
	decode ImageDecodeFunc
}

var (
	formats    = make([]formatDefined, 0, 128)
	format2ext = make(map[string]string)
)

// CheckFormatRegistered reports whether format is registered.
// CheckFormatRegistered 检查 format 是否已注册。
func CheckFormatRegistered(format string) bool {
	format = strings.ToLower(format)
	_, ok := getFormat(format)
	return ok
}

// RegisterFormat registers encode and decode handlers for name.
// RegisterFormat 为 name 注册编码与解码处理函数。
func RegisterFormat(name string, encodeFunc ImageEncodeFunc, decodeFunc ImageDecodeFunc) {
	if _, ok := getFormat(name); ok {
		panic("Repeat registration! ")
	}
	formats = append(formats, formatDefined{name, encodeFunc, decodeFunc})
}

func getFormat(name string) (fm formatDefined, ok bool) {
	for _, fm := range formats {
		if name == fm.name {
			return fm, true
		}
	}
	return formatDefined{}, false
}

// RegisterFormatExt maps formatName to a preferred file extension.
// RegisterFormatExt 将 formatName 映射到首选文件扩展名。
func RegisterFormatExt(formatName string, extName string) {
	format2ext[formatName] = extName
}

// GetExtName returns the registered extension for formatName, or formatName itself.
// GetExtName 返回 formatName 的已注册扩展名，未注册则返回 formatName 本身。
func GetExtName(formatName string) string {
	formatName = strings.ToLower(formatName)
	if rs, ok := format2ext[formatName]; ok {
		return rs
	}
	return formatName
}
