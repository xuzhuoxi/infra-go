package astar

import "fmt"

// NewPosition2D creates a 2D grid position.
// NewPosition2D 新建二维坐标点。
func NewPosition2D(x, y int) Position {
	return Position{X: x, Y: y}
}

// NewPosition3D creates a 3D grid position.
// NewPosition3D 新建三维坐标点。
func NewPosition3D(x, y, z int) Position {
	return Position{X: x, Y: y, Z: z}
}

// Position is a grid cell coordinate.
// Position 表示网格坐标点。
type Position struct {
	X int
	Y int
	Z int
}

// String formats the position as "[x,y,z]".
// String 将坐标格式化为 "[x,y,z]"。
func (pos Position) String() string {
	return fmt.Sprintf("[%d,%d,%d]", pos.X, pos.Y, pos.Z)
}

// EqualTo reports whether two positions are the same.
// EqualTo 判断坐标点是否相同。
func (pos Position) EqualTo(pos2 Position) bool {
	return pos.X == pos2.X && pos.Y == pos2.Y && pos.Z == pos2.Z
}

// AddVector returns pos offset by vector.
// AddVector 按方向矢量位移坐标点。
func (pos Position) AddVector(vector DirectionVector) Position {
	return Position{X: pos.X + vector.X(), Y: pos.Y + vector.Y(), Z: pos.Z + vector.Z()}
}

//-----------------------

// PriorityPosition is a position with search priority.
// PriorityPosition 表示带权值的坐标点。
type PriorityPosition struct {
	Position
	Priority int
}

// String formats the position as "[x,y,z,priority]".
// String 将坐标格式化为 "[x,y,z,priority]"。
func (pos PriorityPosition) String() string {
	return fmt.Sprintf("[%d,%d,%d,%d]", pos.X, pos.Y, pos.Z, pos.Priority)
}

// NewPriorityPosition2D creates a 2D priority position.
// NewPriorityPosition2D 新建二维带权值的坐标点。
func NewPriorityPosition2D(x, y, p int) *PriorityPosition {
	return &PriorityPosition{Position: Position{x, y, 0}, Priority: p}
}

// NewPriorityPosition3D creates a 3D priority position.
// NewPriorityPosition3D 新建三维带权值的坐标点。
func NewPriorityPosition3D(x, y, z, p int) *PriorityPosition {
	return &PriorityPosition{Position: Position{x, y, z}, Priority: p}
}
