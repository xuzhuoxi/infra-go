package imagex

import (
	"image"
	"os"

	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

// LoadImage loads an image from fullPath using the given format (Auto detects from content or extension).
// LoadImage 从 fullPath 加载图像；format 为 Auto 时自动识别格式。
func LoadImage(fullPath string, format formatx.ImageFormat) (img image.Image, imgFormat string, err error) {
	file, e := os.Open(fullPath)
	defer file.Close()
	if nil != e {
		return nil, "", e
	}
	if formatx.Auto == format {
		img, imgFormat, err = image.Decode(file)
	} else {
		img, err = format.Decode(file)
		imgFormat = string(format)
	}
	return
}
