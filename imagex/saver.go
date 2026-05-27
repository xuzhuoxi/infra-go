package imagex

import (
	"image"
	"os"

	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

// SaveImage encodes img to fullPath using format (Auto selects format from file extension).
// SaveImage 将 img 编码保存到 fullPath；format 为 Auto 时按扩展名选择格式。
func SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	os.Open(fullPath)
	file, err := os.Create(fullPath)
	if nil != err {
		return err
	}
	defer file.Close()
	if formatx.Auto == format {
		extF := formatx.ImageFormat(filex.GetExtWithoutDot(fullPath))
		return extF.Encode(file, img, options)
	} else {
		return format.Encode(file, img, options)
	}
}
