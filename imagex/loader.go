//
//Created by xuzhuoxi
//on 2019-05-31.
//@author xuzhuoxi
//
package imagex

import (
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"image"
	"os"
)

func LoadImage(fullPath string, format formatx.ImageFormat) (img image.Image, err error) {
	file, e := os.Open(fullPath)
	defer file.Close()
	if nil != e {
		return nil, e
	}
	if formatx.Auto == format {
		img, _, err = image.Decode(file)
	} else {
		img, err = format.Decode(file)
	}
	return
}
