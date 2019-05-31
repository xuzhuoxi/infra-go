//
//Created by xuzhuoxi
//on 2019-05-31.
//@author xuzhuoxi
//
package imagex

import (
	"image"
	"os"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

func SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	os.Open(fullPath)
	file, _ := os.Create(fullPath)
	defer file.Close()
	return format.Encode(file, img, options)
}
