package imagex

import "errors"

// PixelDirection is a bitmask of eight neighbor directions for morphological and convolution ops.
// PixelDirection 表示八邻域方向的位掩码，用于形态学与卷积运算。
type PixelDirection int

// IncludeDirection reports whether d includes the given single direction bit.
// IncludeDirection 判断 d 是否包含指定单一方向位。
func (d PixelDirection) IncludeDirection(direction PixelDirection) bool {
	return d&direction > 0
}

// String returns the direction name for godoc and debugging.
// String 返回方向名称，供 godoc 与调试使用。
func (d PixelDirection) String() string {
	return d.Name()
}

// Name returns "None", a single direction name, or "Multi".
// Name 返回 "None"、单一方向名或 "Multi"。
func (d PixelDirection) Name() string {
	if d == 0 {
		return "None"
	}
	if s, ok := dirNameMap[d]; ok {
		return s
	}
	return "Multi"
}

// IsNone reports whether d has no direction bits set.
// IsNone 判断是否为无方向。
func (d PixelDirection) IsNone() bool {
	return 0 == d
}

// IsMulti reports whether d combines more than one direction bit.
// IsMulti 判断是否为多方向组合。
func (d PixelDirection) IsMulti() bool {
	if d.IsNone() {
		return false
	}
	if len(DecomposeDirection(d)) > 1 {
		return true
	}
	return false
}

// IsSingle reports whether d is exactly one direction bit.
// IsSingle 判断是否为单一方向。
func (d PixelDirection) IsSingle() bool {
	if d.IsNone() {
		return false
	}
	return len(DecomposeDirection(d)) == 1
}

// ReverseDirection returns the opposite direction bits in d.
// ReverseDirection 返回 d 中各方向的反向组合。
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

// PixelDirectionAdd holds integer (X, Y) offsets for a single direction step.
// PixelDirectionAdd 表示单一方向的整数 (X, Y) 偏移量。
type PixelDirectionAdd struct{ X, Y int }

const (
	Left PixelDirection = 1 << iota // Left neighbor direction. // 左邻方向。
	LeftUp                          // Upper-left neighbor direction. // 左上邻方向。
	Up                              // Up neighbor direction. // 上邻方向。
	RightUp                         // Upper-right neighbor direction. // 右上邻方向。
	Right                           // Right neighbor direction. // 右邻方向。
	RightDown                       // Lower-right neighbor direction. // 右下邻方向。
	Down                            // Down neighbor direction. // 下邻方向。
	LeftDown                        // Lower-left neighbor direction. // 左下邻方向。
)

const (
	// Horizontal is Left | Right.
	// Horizontal 水平方向，包含 Left、Right。
	Horizontal = Left | Right
	// Vertical is Up | Down.
	// Vertical 垂直方向，包含 Up、Down。
	Vertical = Up | Down
	// Oblique is the four diagonal directions.
	// Oblique 斜方向，包含 LeftDown、LeftUp、RightDown、RightUp。
	Oblique = LeftDown | LeftUp | RightDown | RightUp
	// AllDirection is all eight neighbor directions.
	// AllDirection 全部八个方向。
	AllDirection = Horizontal | Vertical | Oblique
)

var (
	LeftDirAdd      = PixelDirectionAdd{-1, 0}  // Offset for Left. // Left 方向偏移。
	LeftUpDirAdd    = PixelDirectionAdd{-1, -1} // Offset for LeftUp. // LeftUp 方向偏移。
	UpDirAdd        = PixelDirectionAdd{0, -1}  // Offset for Up. // Up 方向偏移。
	RightUpDirAdd   = PixelDirectionAdd{1, -1}  // Offset for RightUp. // RightUp 方向偏移。
	RightDirAdd     = PixelDirectionAdd{1, 0}   // Offset for Right. // Right 方向偏移。
	RightDownDirAdd = PixelDirectionAdd{1, 1}   // Offset for RightDown. // RightDown 方向偏移。
	DownDirAdd      = PixelDirectionAdd{0, 1}   // Offset for Down. // Down 方向偏移。
	LeftDownDirAdd  = PixelDirectionAdd{-1, 1}  // Offset for LeftDown. // LeftDown 方向偏移。
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

// GetPixelDirectionAdd returns the offset for a single direction bit.
// GetPixelDirectionAdd 返回单一方向位对应的坐标偏移量。
func GetPixelDirectionAdd(direction PixelDirection) (PixelDirectionAdd, error) {
	rs, ok := dirMap[direction]
	if ok {
		return rs, nil
	} else {
		return PixelDirectionAdd{}, errors.New("Direction Error! ")
	}
}

// GetPixelDirectionAdds returns offsets for all bits set in directions.
// GetPixelDirectionAdds 返回 directions 中各方向位对应的坐标偏移量列表。
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

// DecomposeDirection splits directions into individual direction bits.
// DecomposeDirection 将 directions 分解为单一方向位列表。
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
