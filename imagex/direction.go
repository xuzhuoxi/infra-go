package imagex

import "errors"

type PixelDirection int

// IncludeDirection 包含方向
func (d PixelDirection) IncludeDirection(direction PixelDirection) bool {
	return d&direction > 0
}

func (d PixelDirection) String() string {
	return d.Name()
}

func (d PixelDirection) Name() string {
	if d == 0 {
		return "None"
	}
	if s, ok := dirNameMap[d]; ok {
		return s
	}
	return "Multi"
}

// IsNone 无方向
func (d PixelDirection) IsNone() bool {
	return 0 == d
}

// IsMulti 多方向
func (d PixelDirection) IsMulti() bool {
	if d.IsNone() {
		return false
	}
	if len(DecomposeDirection(d)) > 1 {
		return true
	}
	return false
}

// IsSingle 单方向
func (d PixelDirection) IsSingle() bool {
	if d.IsNone() {
		return false
	}
	return len(DecomposeDirection(d)) == 1
}

// ReverseDirection 取反方向
func (d PixelDirection) ReverseDirection() PixelDirection {
	var rs PixelDirection
	halfLn := len(dirs) / 2
	var index = 0
	var index2 = halfLn
	for index < halfLn {
		if d&dirs[index] > 0 {
			rs = rs | dirs[index2]
		}
		if d&dirs[index2] > 0 {
			rs = rs | dirs[index]
		}
		index++
		index2++
	}
	return rs
}

// PixelDirectionAdd 方向偏移量
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
	// Horizontal 水平方向，包含Left,Right
	Horizontal = Left | Right
	// Vertical 垂直方向，包含Up,down
	Vertical = Up | Down
	// Oblique 斜方向，包含LeftDown,LeftUp,RightDown,RightUp
	Oblique = LeftDown | LeftUp | RightDown | RightUp
	// AllDirection 全部八个方向
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
	dirs       = []PixelDirection{Left, LeftUp, Up, RightUp, Right, RightDown, Down, LeftDown} //顺序有意义
	dirAdds    = []PixelDirectionAdd{LeftDirAdd, LeftUpDirAdd, UpDirAdd, RightUpDirAdd, RightDirAdd, RightDownDirAdd, DownDirAdd, LeftDownDirAdd}
	dirMap     map[PixelDirection]PixelDirectionAdd
	dirNameMap map[PixelDirection]string
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
	dirNameMap = make(map[PixelDirection]string)
	dirNameMap[Left] = "Left"
	dirNameMap[LeftUp] = "LeftUp"
	dirNameMap[Up] = "Up"
	dirNameMap[RightUp] = "RightUp"
	dirNameMap[Right] = "Right"
	dirNameMap[RightDown] = "RightDown"
	dirNameMap[Down] = "Down"
	dirNameMap[LeftDown] = "LeftDown"
}

// GetPixelDirectionAdd
// 取方向坐标增加值
func GetPixelDirectionAdd(direction PixelDirection) (PixelDirectionAdd, error) {
	rs, ok := dirMap[direction]
	if ok {
		return rs, nil
	} else {
		return PixelDirectionAdd{}, errors.New("Direction Error! ")
	}
}

// GetPixelDirectionAdds
// 取多方向坐标增加值
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

// DecomposeDirection
// 分解方向
func DecomposeDirection(directions PixelDirection) []PixelDirection {
	if directions <= 0 {
		return nil
	}
	var rs []PixelDirection
	for index, dir := range dirs {
		if dir&directions > 0 {
			rs = append(rs, dirs[index])
		}
	}
	return rs
}
