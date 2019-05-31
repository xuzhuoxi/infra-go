//
//Created by xuzhuoxi
//on 2019-04-27.
//@author xuzhuoxi
//
package pngx

import (
	"image/png"
	"io"
	"image"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

func init() {
	formatx.RegisterFormat(string(formatx.PNG), EncodePNG, DecodePNG)
}

func EncodePNG(w io.Writer, m image.Image, _ interface{}) error {
	return png.Encode(w, m)
}

func DecodePNG(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}
