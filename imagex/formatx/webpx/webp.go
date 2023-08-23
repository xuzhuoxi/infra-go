// Package webpx
// Create on 2023/5/3
// @author xuzhuoxi
package webpx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"image"
	"io"
)

func init() {
	formatx.RegisterFormat(string(formatx.WEBP), EncodeWEBP, DecodeWEBP)
	formatx.RegisterFormatExt(string(formatx.WEBP), "webp")
}

func EncodeWEBP(w io.Writer, m image.Image, options interface{}) error {
	return errors.New("EncodeWEBP is not implemented")
}

func DecodeWEBP(r io.Reader) (image.Image, error) {
	return nil, errors.New("DecodeWEBP is not implemented")
}

//
//func EncodeWEBP(w io.Writer, m image.Image, options interface{}) error {
//	if nil == options {
//		return encodeWEBPDefault(w, m)
//	} else {
//		switch webpOptions := options.(type) {
//		case encoder.Options:
//			return encodeWEBP(w, m, &webpOptions)
//		case *encoder.Options:
//			return encodeWEBP(w, m, webpOptions)
//		default:
//			return encodeWEBPDefault(w, m)
//		}
//	}
//}
//
//func DecodeWEBP(r io.Reader) (image.Image, error) {
//	return webp.Decode(r, &decoder.Options{})
//}
//
//func encodeWEBPDefault(w io.Writer, m image.Image) error {
//	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
//	if err != nil {
//		return err
//	}
//	return encodeWEBP(w, m, options)
//}
//
//func encodeWEBP(w io.Writer, m image.Image, options *encoder.Options) error {
//	return webp.Encode(w, m, options)
//}
