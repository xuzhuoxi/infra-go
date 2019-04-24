//
//Created by xuzhuoxi
//on 2019-04-05.
//@author xuzhuoxi
//
package astar

// 新建2D坐标点
func NewPosition2D(x, y int) Position {
	return Position{X: x, Y: y}
}

// 新建3D坐标点
func NewPosition3D(x, y, z int) Position {
	return Position{X: x, Y: y, Z: z}
}

// 坐标点
type Position struct {
	X int
	Y int
	Z int
}

// 判断坐标点是否相同
func (pos Position) EqualTo(pos2 Position) bool {
	return pos.X == pos2.X && pos.Y == pos2.Y && pos.Z == pos2.Z
}

// 坐标点位移
func (pos Position) AddVector(vector DirectionVector) Position {
	return Position{X: pos.X + vector.X(), Y: pos.Y + vector.Y(), Z: pos.Z + vector.Z()}
}

//-----------------------

// 带权值的坐标点
type PriorityPosition struct {
	Position
	Priority int
}

// 新建2D带权值的坐标点
func NewPriorityPosition2D(x, y, p int) *PriorityPosition {
	return &PriorityPosition{Position: Position{x, y, 0}, Priority: p}
}

// 新建3D带权值的坐标点
func NewPriorityPosition3D(x, y, z, p int) *PriorityPosition {
	return &PriorityPosition{Position: Position{x, y, 0}, Priority: p}
}
