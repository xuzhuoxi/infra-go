package formatx

import (
	"errors"
	"image"
	"io"
)

// ImageFormat names a registered image encoding format.
// ImageFormat 表示已注册的图像编码格式名称。
type ImageFormat string

const (
	Auto ImageFormat = "" // Auto detects format from content or file extension. // Auto 自动识别格式。
	PNG  ImageFormat = "png"  // PNG image format. // PNG 图像格式。
	JPEG ImageFormat = "jpeg" // JPEG image format. // JPEG 图像格式。
	JPG  ImageFormat = "jpg"  // JPG image format (alias of JPEG). // JPG 图像格式（JPEG 别名）。
	JPS  ImageFormat = "jps"  // JPS stereo JPEG format. // JPS 立体 JPEG 格式。
	BMP  ImageFormat = "bmp"  // BMP image format. // BMP 图像格式。
	WEBP ImageFormat = "webp" // WebP image format. // WebP 图像格式。
)

// Encode writes image m to w using the registered encoder for f.
// Encode 使用 f 对应的已注册编码器将图像 m 写入 w。
func (f ImageFormat) Encode(w io.Writer, m image.Image, options interface{}) error {
	if fm, ok := getFormat(string(f)); ok {
		return fm.encode(w, m, options)
	}
	return errors.New("No RegisterEncode:" + string(f))
}

// Decode reads an image from r using the registered decoder for f.
// Decode 使用 f 对应的已注册解码器从 r 读取图像。
func (f ImageFormat) Decode(r io.Reader) (image.Image, error) {
	if fm, ok := getFormat(string(f)); ok {
		return fm.decode(r)
	}
	return nil, errors.New("No RegisterDecode:" + string(f))
}
