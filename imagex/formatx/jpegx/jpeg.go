//
//Created by xuzhuoxi
//on 2019-04-27.
//@author xuzhuoxi
//
package jpegx

import (
	"io"
	"image"
	"image/jpeg"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
)

func init() {
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jps", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	formatx.RegisterFormat(string(formatx.JPEG), EncodeJPEG, DecodeJPEG)
	formatx.RegisterFormat(string(formatx.JPG), EncodeJPEG, DecodeJPEG)
	formatx.RegisterFormat(string(formatx.JPS), EncodeJPEG, DecodeJPEG)
}

func EncodeJPEG(w io.Writer, m image.Image, options interface{}) error {
	if nil == options {
		return jpeg.Encode(w, m, nil)
	} else {
		return jpeg.Encode(w, m, options.(*jpeg.Options))
	}
}

func DecodeJPEG(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}
