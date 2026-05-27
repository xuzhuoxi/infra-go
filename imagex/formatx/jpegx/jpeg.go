package jpegx

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

func init() {
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jps", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	formatx.RegisterFormat(string(formatx.JPEG), EncodeJPEG, DecodeJPEG)
	formatx.RegisterFormat(string(formatx.JPG), EncodeJPEG, DecodeJPEG)
	formatx.RegisterFormat(string(formatx.JPS), EncodeJPEG, DecodeJPEG)
	formatx.RegisterFormatExt(string(formatx.JPEG), "jpg")
	formatx.RegisterFormatExt(string(formatx.JPG), "jpg")
	formatx.RegisterFormatExt(string(formatx.JPS), "jps")
}

// EncodeJPEG encodes m as JPEG to w; options may be *jpeg.Options or nil.
// EncodeJPEG 将 m 以 JPEG 格式编码写入 w；options 可为 *jpeg.Options 或 nil。
func EncodeJPEG(w io.Writer, m image.Image, options interface{}) error {
	if nil == options {
		return jpeg.Encode(w, m, nil)
	} else {
		return jpeg.Encode(w, m, options.(*jpeg.Options))
	}
}

// DecodeJPEG decodes a JPEG image from r.
// DecodeJPEG 从 r 解码 JPEG 图像。
func DecodeJPEG(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}
