// 边缘化
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
)

// 边缘滤波器
var (
	//5x5 水平边缘滤波器
	Edge5Horizontal = FilterTemplate{Radius: 1, Size: 3, Scale: 0,
		Offsets: []FilterOffset{
			{-2, 0, -1}, {-1, 0, -1}, {0, 0, 4}, {1, 0, 1}, {2, 0, 1}}}
	//5x5 垂直边缘滤波器
	Edge5Vertical = FilterTemplate{Radius: 1, Size: 3, Scale: 0,
		Offsets: []FilterOffset{
			{0, -2, -1}, {0, -1, -1}, {0, 0, 4}, {0, 1, 1}, {0, 2, 1}}}
	//5x5 45度边缘滤波器(左上右下)
	Edge5Oblique45 = FilterTemplate{Radius: 1, Size: 3, Scale: 0,
		Offsets: []FilterOffset{
			{-2, -2, -1}, {-1, -1, -1}, {0, 0, 4}, {1, 1, 1}, {1, 2, 1}}}
	//5x5 135度边缘滤波器(左下右上)
	Edge5Oblique135 = FilterTemplate{Radius: 1, Size: 3, Scale: 0,
		Offsets: []FilterOffset{
			{-2, 2, -1}, {-1, 1, -1}, {0, 0, 4}, {1, -1, 1}, {2, -2, 1}}}
	//5x5 全方向边缘滤波器
	Edge3All = FilterTemplate{Radius: 1, Size: 3, Scale: 0,
		Offsets: []FilterOffset{
			{-1, -1, -1}, {0, -1, -1}, {1, -1, -1},
			{-1, +0, -1}, {0, +0, +8}, {1, +0, -1},
			{-1, +1, -1}, {0, +1, -1}, {1, +1, -1}}}
)

// 创建边缘检测滤波器
// radius：		卷积核半径
// direction: 	运动方向
// diff:		梯度差
func CreateEdgeFilter(radius int, direction imagex.PixelDirection, diff uint) (filter *FilterTemplate, err error) {
	if radius < 1 {
		return nil, errors.New("Radius < 1. ")
	}
	dirAdds := imagex.GetPixelDirectionAdds(direction)
	if nil == dirAdds {
		return nil, errors.New("Direction Error. ")
	}
	kSize := radius + radius + 1
	ln := len(dirAdds)*radius + 1
	var offsets = make([]FilterOffset, 0, ln)
	var value int
	var sumValue int
	for _, add := range dirAdds {
		for i := radius; i >= 1; i-- {
			value = (i-radius)*int(diff) - 1
			offsets = append(offsets, FilterOffset{X: add.X * i, Y: add.Y * i, Value: value})
			sumValue += value
		}
	}
	offsets = append(offsets, FilterOffset{X: 0, Y: 0, Value: -sumValue})
	return &FilterTemplate{Radius: radius, Size: kSize, Scale: 0, Offsets: offsets}, nil
}
