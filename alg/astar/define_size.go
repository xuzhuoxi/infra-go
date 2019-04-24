//
//Created by xuzhuoxi
//on 2019-04-03.
//@author xuzhuoxi
//
package astar

// 范围定义
type Size struct {
	Width  int
	Height int
	Depth  int
}

// 求体积/面积
func (s Size) Area() int {
	return s.Width * s.Height * s.Depth
}

// 是否为空
func (s Size) Empty() bool {
	return 0 == s.Width || 0 == s.Height || 0 == s.Depth
}
