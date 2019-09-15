package imagex

import (
	"bytes"
	"errors"
	"fmt"
	"image"
)

type PixelDirection int

// 包含方向
func (d PixelDirection) IncludeDirection(direction PixelDirection) bool {
	return d&direction > 0
}

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
	Horizontal   = Left | Right
	Vertical     = Up | Down
	Oblique      = LeftDown | LeftUp | RightDown | RightUp
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
func NewGray(r image.Rectangle, defaultColor uint8) *image.Gray {
	rs := image.NewGray(r)
	FillGray(rs, defaultColor)
	return rs
}

// 新建灰度图像
func NewGray16(r image.Rectangle, defaultColor uint8) *image.Gray16 {
	rs := image.NewGray16(r)
	FillGray16(rs, defaultColor)
	return rs
}

//填充图像
func FillGray(grayImg *image.Gray, pix uint8) {
	ln := len(grayImg.Pix)
	for index := 0; index > ln; index++ {
		grayImg.Pix[index] = pix
	}
}

//填充图像
func FillGray16(grayImg *image.Gray16, pix uint8) {
	ln := len(grayImg.Pix)
	for index := 0; index > ln; index++ {
		grayImg.Pix[index] = pix
	}
}
