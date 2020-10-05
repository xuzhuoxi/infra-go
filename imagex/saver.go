//
//Created by xuzhuoxi
//on 2019-05-31.
//@author xuzhuoxi
//
package imagex

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"image"
	"os"
)

func SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	os.Open(fullPath)
	file, _ := os.Create(fullPath)
	defer file.Close()
	if formatx.Auto == format {
		extF := formatx.ImageFormat(filex.GetExtWithoutDot(fullPath))
		return extF.Encode(file, img, options)
	} else {
		return format.Encode(file, img, options)
	}
}
