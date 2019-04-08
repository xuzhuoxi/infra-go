//
//Created by xuzhuoxi
//on 2019-04-05.
//@author xuzhuoxi
//
package astar

type Position struct {
	X int
	Y int
}

func (pos Position) EqualTo(pos2 Position) bool {
	return pos.X == pos2.X && pos.Y == pos2.Y
}

//-----------------------

type PriorityPosition struct {
	Position
	Priority int
}

func NewPriorityPosition(x, y, p int) *PriorityPosition {
	return &PriorityPosition{Position: Position{x, y}, Priority: p}
}
