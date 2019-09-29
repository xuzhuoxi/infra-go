//
//Created by xuzhuoxi
//on 2019-04-26.
//@author xuzhuoxi
//
package formatx

import (
	"errors"
	"image"
	"io"
)

type ImageFormat string

const (
	Auto ImageFormat = ""
	PNG              = "png"
	JPEG             = "jpeg"
	JPG              = "jpg"
	JPS              = "jps"
)

func (f ImageFormat) Encode(w io.Writer, m image.Image, options interface{}) error {
	if fm, ok := getFormat(string(f)); ok {
		return fm.encode(w, m, options)
	}
	return errors.New("No RegisterEncode:" + string(f))
}

func (f ImageFormat) Decode(r io.Reader) (image.Image, error) {
	if fm, ok := getFormat(string(f)); ok {
		return fm.decode(r)
	}
	return nil, errors.New("No RegisterDecode:" + string(f))
}
