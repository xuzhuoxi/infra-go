//浮雕
package filterx

// 浮雕滤波器
var (
	//3x3 浮雕滤波器
	Emboss3Oblique45 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-1, -1, -1}, {-0, -1, -1},
			{-1, -0, -1}, {+1, +0, +1},
			{+0, +1, +1}, {+1, +1, +1}}}
	//5x5 浮雕滤波器
	Emboss5Oblique45 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -2, -1}, {+0, -2, -1}, {+1, -2, -1},
			{-2, -1, -1}, {-1, -1, -1}, {+0, -1, -1}, {+2, -1, +1},
			{-2, -0, -1}, {-1, -0, -1}, {+1, -0, +1}, {+2, -0, +1},
			{-2, +1, -1}, {+0, +1, +1}, {+1, +1, +1}, {+2, +1, +1},
			{-1, +2, +1}, {+0, +2, +1}, {+1, +2, +1}, {+2, +2, +1}}}
	//3x3 非对称浮雕滤波器
	Emboss3Asymmetrical = FilterMatrix{Radius: 1, Size: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-1, -1, 2},
			{+0, +0, -1},
			{+1, +1, -1}}}
)
