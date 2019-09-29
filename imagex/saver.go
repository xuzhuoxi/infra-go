//
//Created by xuzhuoxi
//on 2019-05-31.
//@author xuzhuoxi
//
package imagex

import (
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"image"
	"os"
)

func SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	os.Open(fullPath)
	file, _ := os.Create(fullPath)
	defer file.Close()
	if formatx.Auto == format {
		extF := formatx.ImageFormat(osxu.GetExtensionName(fullPath))
		return extF.Encode(file, img, options)
	} else {
		return format.Encode(file, img, options)
	}
}
