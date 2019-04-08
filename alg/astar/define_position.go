//
//Created by xuzhuoxi
//on 2019-04-05.
//@author xuzhuoxi
//
package astar

func NewPosition2D(x, y int) Position {
	return Position{X: x, Y: y}
}

func NewPosition3D(x, y, z int) Position {
	return Position{X: x, Y: y, Z: z}
}

type Position struct {
	X int
	Y int
	Z int
}

func (pos Position) EqualTo(pos2 Position) bool {
	return pos.X == pos2.X && pos.Y == pos2.Y && pos.Z == pos2.Z
}

func (pos Position) AddVector(vector DirectionVector) Position {
	return Position{X: pos.X + vector.X(), Y: pos.Y + vector.Y(), Z: pos.Z + vector.Z()}
}

//-----------------------

type PriorityPosition struct {
	Position
	Priority int
}

func NewPriorityPosition2D(x, y, p int) *PriorityPosition {
	return &PriorityPosition{Position: Position{x, y, 0}, Priority: p}
}

func NewPriorityPosition3D(x, y, z, p int) *PriorityPosition {
	return &PriorityPosition{Position: Position{x, y, 0}, Priority: p}
}
