package jpegx

import "image/jpeg"

// DefaultJPEGOptions is the default JPEG encoder options (quality 75).
// DefaultJPEGOptions 为默认 JPEG 编码选项（品质 75）。
var DefaultJPEGOptions = &jpeg.Options{Quality: 75}

// NewJpegOptions returns JPEG options with the given quality.
// NewJpegOptions 返回指定品质的 JPEG 编码选项。
func NewJpegOptions(quality int) *jpeg.Options {
	return &jpeg.Options{Quality: quality}
}
