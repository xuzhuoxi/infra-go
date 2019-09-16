package imagex

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
)

type PixelDirection int

// 包含方向
func (d PixelDirection) IncludeDirection(direction PixelDirection) bool {
	return d&direction > 0
}

//方向偏移量
type PixelDirectionAdd struct{ X, Y int }

const (
	Left PixelDirection = 1 << iota
	LeftUp
	Up
	RightUp
	Right
	RightDown
	Down
	LeftDown
)

const (
	//水平方向，包含Left,Right
	Horizontal = Left | Right
	//垂直方向，包含Up,down
	Vertical = Up | Down
	//斜方向，包含LeftDown,LeftUp,RightDown,RightUp
	Oblique = LeftDown | LeftUp | RightDown | RightUp
	//全部八个方向
	AllDirection = Horizontal | Vertical | Oblique
)

var (
	LeftDirAdd      = PixelDirectionAdd{-1, 0}
	LeftUpDirAdd    = PixelDirectionAdd{-1, -1}
	UpDirAdd        = PixelDirectionAdd{0, -1}
	RightUpDirAdd   = PixelDirectionAdd{1, -1}
	RightDirAdd     = PixelDirectionAdd{1, 0}
	RightDownDirAdd = PixelDirectionAdd{1, 1}
	DownDirAdd      = PixelDirectionAdd{0, 1}
	LeftDownDirAdd  = PixelDirectionAdd{-1, 1}
)

var (
	dirs    = []PixelDirection{Left, LeftUp, Up, RightUp, Right, RightDown, Down, LeftDown}
	dirAdds = []PixelDirectionAdd{LeftDirAdd, LeftUpDirAdd, UpDirAdd, RightUpDirAdd, RightDirAdd, RightDownDirAdd, DownDirAdd, LeftDownDirAdd}
	dirMap  map[PixelDirection]PixelDirectionAdd
)

func init() {
	dirMap = make(map[PixelDirection]PixelDirectionAdd)
	dirMap[Left] = LeftDirAdd
	dirMap[LeftUp] = LeftUpDirAdd
	dirMap[Up] = UpDirAdd
	dirMap[RightUp] = RightUpDirAdd
	dirMap[Right] = RightDirAdd
	dirMap[RightDown] = RightDownDirAdd
	dirMap[Down] = DownDirAdd
	dirMap[LeftDown] = LeftDownDirAdd
}

//图像字符串化
func SprintImage(img image.Image) string {
	bs := bytes.NewBufferString("")
	rect := img.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; y < rect.Max.X; x++ {
			bs.WriteString(fmt.Sprint(img.At(x, y)))
		}
		bs.WriteString("\n")
	}
	return bs.String()
}

// 取向方向
func ReverseDirection(direction PixelDirection) PixelDirection {
	return AllDirection ^ direction
}

//取方向坐标增加值
func GetPixelDirectionAdd(direction PixelDirection) (PixelDirectionAdd, error) {
	rs, ok := dirMap[direction]
	if ok {
		return rs, nil
	} else {
		return PixelDirectionAdd{0, 0}, errors.New("Direction Error! ")
	}
}

//取多方向坐标增加值
func GetPixelDirectionAdds(directions PixelDirection) []PixelDirectionAdd {
	if directions <= 0 {
		return nil
	}
	var rs []PixelDirectionAdd
	for index, dir := range dirs {
		if dir&directions > 0 {
			rs = append(rs, dirAdds[index])
		}
	}
	return rs
}

// 新建灰度图像
func NewGray(rect image.Rectangle, grayColor uint8) *image.Gray {
	rs := image.NewGray(rect)
	FillImage(rs, &color.Gray{Y: grayColor})
	return rs
}

// 新建灰度图像
func NewGray16(rect image.Rectangle, grayColor uint16) *image.Gray16 {
	rs := image.NewGray16(rect)
	FillImage(rs, &color.Gray16{Y: grayColor})
	return rs
}

// 新建RGBA图像
func NewRGBA(rect image.Rectangle, rgbaColor uint32) *image.RGBA {
	rs := image.NewRGBA(rect)
	r := (rgbaColor & 0xff000000) >> 24
	g := (rgbaColor & 0x00ff0000) >> 16
	b := (rgbaColor & 0x0000ff00) >> 8
	a := rgbaColor & 0x000000ff
	FillImage(rs, &color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
	return rs
}

// 新建RGBA64图像
func NewRGBA64(rect image.Rectangle, rgbaColor uint64) *image.RGBA64 {
	rs := image.NewRGBA64(rect)
	r := (rgbaColor & 0xffff000000000000) >> 24
	g := (rgbaColor & 0x0000ffff00000000) >> 16
	b := (rgbaColor & 0x00000000ffff0000) >> 8
	a := rgbaColor & 0x000000000000ffff
	FillImage(rs, &color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
	return rs
}

//使用颜色填充图像
func FillImage(img draw.Image, color color.Color) {
	rect := img.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.Set(x, y, color)
		}
	}
}

//使用颜色填充图像部分区域
func FillImagetAt(img draw.Image, color color.Color, rect image.Rectangle) {
	rect2 := img.Bounds()
	minX := mathx.MaxInt(rect.Min.X, rect2.Min.X)
	minY := mathx.MaxInt(rect.Min.Y, rect2.Min.Y)
	maxX := mathx.MinInt(rect.Max.X, rect2.Max.X)
	maxY := mathx.MinInt(rect.Max.Y, rect2.Max.Y)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			img.Set(x, y, color)
		}
	}
}
