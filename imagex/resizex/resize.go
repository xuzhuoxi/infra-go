package resizex

import (
	"image"

	"github.com/nfnt/resize"
)

// ResizeImage scales source to width×height using Lanczos3 resampling.
// ResizeImage 使用 Lanczos3 重采样将 source 缩放到 width×height。
func ResizeImage(source image.Image, width, height uint) (img image.Image, err error) {
	return resize.Resize(width, height, source, resize.Lanczos3), nil
}
