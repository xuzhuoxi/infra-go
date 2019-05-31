//
//Created by xuzhuoxi
//on 2019-05-31.
//@author xuzhuoxi
//
package resizex

import (
	"image"
	"github.com/nfnt/resize"
)

func ResizeImage(source image.Image, width, height uint) (img image.Image, err error) {
	return resize.Resize(width, height, source, resize.Lanczos3), nil
}
