package jpegx

import "image/jpeg"

var DefaultJPEGOptions = &jpeg.Options{Quality: 75}

// NewJpegOptions
// 品质设置
func NewJpegOptions(quality int) *jpeg.Options {
	return &jpeg.Options{Quality: quality}
}
