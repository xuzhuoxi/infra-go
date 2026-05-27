package pngx

import (
	"image"
	"image/png"
	"io"

	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

func init() {
	formatx.RegisterFormat(string(formatx.PNG), EncodePNG, DecodePNG)
	formatx.RegisterFormatExt(string(formatx.PNG), "png")
}

// EncodePNG encodes m as PNG to w; options may be png.CompressionLevel or *png.CompressionLevel.
// EncodePNG 将 m 以 PNG 格式编码写入 w；options 可为 png.CompressionLevel 或 *png.CompressionLevel。
func EncodePNG(w io.Writer, m image.Image, options interface{}) error {
	if nil == options {
		return encodePNG(w, m, png.DefaultCompression)
	} else {
		switch level := options.(type) {
		case png.CompressionLevel:
			return encodePNG(w, m, level)
		case *png.CompressionLevel:
			return encodePNG(w, m, *level)
		default:
			return encodePNG(w, m, png.DefaultCompression)
		}
	}
}

// DecodePNG decodes a PNG image from r.
// DecodePNG 从 r 解码 PNG 图像。
func DecodePNG(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func encodePNG(w io.Writer, m image.Image, level png.CompressionLevel) error {
	var e png.Encoder
	e.CompressionLevel = level
	return e.Encode(w, m)
}
