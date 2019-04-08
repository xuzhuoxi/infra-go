//
//Created by xuzhuoxi
//on 2019-04-03.
//@author xuzhuoxi
//
package astar

type Size struct {
	Width  int
	Height int
}

func (s Size) Area() int {
	return s.Width * s.Height
}

func (s Size) Empty() bool {
	return 0 == s.Width || 0 == s.Height
}

const (
	//地图外
	GridOut = -1
	//通路
	GridPath = 0
	//障碍
	GridObstacle = 1
)
