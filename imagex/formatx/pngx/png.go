//
//Created by xuzhuoxi
//on 2019-04-27.
//@author xuzhuoxi
//
package pngx

import (
	"image/png"
	"io"
	"image"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

func init() {
	formatx.RegisterFormat(string(formatx.PNG), EncodePNG, DecodePNG)
}

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

func DecodePNG(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func encodePNG(w io.Writer, m image.Image, level png.CompressionLevel) error {
	var e png.Encoder
	e.CompressionLevel = level
	return e.Encode(w, m)
}
