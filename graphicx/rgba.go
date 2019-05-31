//
//Created by xuzhuoxi
//on 2019-05-29.
//@author xuzhuoxi
//
package graphicx

import (
	"image/color"
)

// 亮度 [0,1]
// Y = 0.299 * r + 0.587 * g + 0.114 * b
func Luminosity(c color.RGBA) float64 {
	return (0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B)) / 255
}

// 饱和度
// V = 0.5 * r - 0.4187*g - 0.0813*b + 128
func Saturation(c color.RGBA) float64 {
	return 0.5*float64(c.R) + 0.4187*float64(c.G) + 0.0813*float64(c.B) + 128
}

// 色度
// U = 0.1687* r - 0.3313* g + 0.5 * b + 128
func Chroma(c color.RGBA) float64 {
	return 0.1687*float64(c.R) + 0.3313*float64(c.G) + 0.5*float64(c.B) + 128
}

// 灰度
// 1.浮点算法：Gray=R*0.3+G*0.59+B*0.11
// 2.整数方法：Gray=(R*30+G*59+B*11)/100
// 3.移位方法：Gray =(R*76+G*151+B*28)>>8
// 4.平均值法：Gray=（R+G+B）/3
// 5.仅取绿色：Gray=G
func Gray(c color.RGBA) color.Gray {
	gray := uint8((uint16(c.R)*76 + uint16(c.G)*151 + uint16(c.B)*28) >> 8)
	return color.Gray{Y:gray}
}

// 反相
func Inverse(c color.RGBA) color.RGBA {
	c.R = 255 - c.R
	c.G = 255 - c.G
	c.B = 255 - c.B
	c.A = 255 - c.A
	return c
}
