package astar

// Size represents width, height, and depth of a grid region.
// Size 表示网格区域的宽、高、深。
type Size struct {
	Width  int
	Height int
	Depth  int
}

// Area returns width * height * depth.
// Area 求体积或面积（宽 × 高 × 深）。
func (s Size) Area() int {
	return s.Width * s.Height * s.Depth
}

// Empty reports whether any dimension is zero.
// Empty 判断尺寸是否为空（任一边为 0）。
func (s Size) Empty() bool {
	return 0 == s.Width || 0 == s.Height || 0 == s.Depth
}
