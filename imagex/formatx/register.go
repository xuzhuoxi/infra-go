//
//Created by xuzhuoxi
//on 2019-05-31.
//@author xuzhuoxi
//
package formatx

import (
	"io"
	"image"
	"strings"
)

type ImageEncodeFunc = func(w io.Writer, m image.Image, options interface{}) error
type ImageDecodeFunc = func(w io.Reader) (image.Image, error)

type formatDefined struct {
	name   string
	encode ImageEncodeFunc
	decode ImageDecodeFunc
}

var formats = make([]formatDefined, 128)[0:0]

func CheckFormatRegistered(format string) bool {
	format = strings.ToLower(format)
	_, ok := getFormat(format)
	return ok
}

func RegisterFormat(name string, encodeFunc ImageEncodeFunc, decodeFunc ImageDecodeFunc) {
	if _, ok := getFormat(name); ok {
		panic("Repeat registration! ")
	}
	formats = append(formats, formatDefined{name, encodeFunc, decodeFunc})
}

func getFormat(name string) (fm formatDefined, ok bool) {
	for _, fm := range formats {
		if name == fm.name {
			return fm, true
		}
	}
	return formatDefined{}, false
}
