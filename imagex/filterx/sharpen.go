// 锐化
package filterx

// 锐化滤波器
var (
	//3x3 简单锐化滤波器
	Sharpen3AllSimple = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 1,
		Kernel: []KernelVector{
			{0, -1, -1}, {-1, +0, -1}, {0, +0, +5}, {1, +0, -1}, {0, +1, -1}}}
	//3x3 全方向锐化滤波器
	Sharpen3All = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 1,
		Kernel: []KernelVector{
			{-1, -1, -1}, {0, -1, -1}, {1, -1, -1},
			{-1, +0, -1}, {0, +0, +9}, {1, +0, -1},
			{-1, +1, -1}, {0, +1, -1}, {1, +1, -1}}}
	//3x3 全方向锐化滤波器(强调边缘)
	SharpenStrengthen3All = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 1,
		Kernel: []KernelVector{
			{-1, -1, +1}, {0, -1, +1}, {1, -1, +1},
			{-1, +0, +1}, {0, +0, -7}, {1, +0, +1},
			{-1, +1, +1}, {0, +1, +1}, {1, +1, +1}}}
	//5x5 全方向锐化滤波器
	Sharpen5All = FilterMatrix{KernelRadius: 2, KernelSize: 5, KernelScale: 1,
		Kernel: []KernelVector{ // 中间改为1，原来是8过不了检验
			{-2, -2, -1}, {-1, -2, -1}, {0, -2, -1}, {1, -2, -1}, {2, -2, -1},
			{-2, -1, -1}, {-1, -1, +2}, {0, -1, +2}, {1, -1, +2}, {2, -1, -1},
			{-2, +0, -1}, {-1, +0, +2}, {0, +0, +1}, {1, +0, +2}, {2, +0, -1},
			{-2, +1, -1}, {-1, +1, +2}, {0, +1, +2}, {1, +1, +2}, {2, +1, -1},
			{-2, +2, -1}, {-1, +2, -1}, {1, +2, -1}, {1, +2, -1}, {2, +2, -1}}}
)
