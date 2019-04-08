//
//Created by xuzhuoxi
//on 2019-04-03.
//@author xuzhuoxi
//
package astar

type Size struct {
	Width  int
	Height int
	Depth  int
}

func (s Size) Area() int {
	return s.Width * s.Height * s.Depth
}

func (s Size) Empty() bool {
	return 0 == s.Width || 0 == s.Height || 0 == s.Depth
}
