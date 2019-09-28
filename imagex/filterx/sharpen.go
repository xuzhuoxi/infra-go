// 锐化
package filterx

// 锐化滤波器
var (
	//3x3 简单锐化滤波器
	Sharpen3AllSimple = FilterTemplate{Radius: 1, Size: 3, Scale: 1,
		Offsets: []FilterOffset{
			{0, -1, -1}, {-1, +0, -1}, {0, +0, +5}, {1, +0, -1}, {0, +1, -1}}}
	//3x3 全方向锐化滤波器
	Sharpen3All = FilterTemplate{Radius: 1, Size: 3, Scale: 1,
		Offsets: []FilterOffset{
			{-1, -1, -1}, {0, -1, -1}, {1, -1, -1},
			{-1, +0, -1}, {0, +0, +9}, {1, +0, -1},
			{-1, +1, -1}, {0, +1, -1}, {1, +1, -1}}}
	//3x3 全方向锐化滤波器(强调边缘)
	SharpenStrengthen3All = FilterTemplate{Radius: 1, Size: 3, Scale: 1,
		Offsets: []FilterOffset{
			{-1, -1, +1}, {0, -1, +1}, {1, -1, +1},
			{-1, +0, +1}, {0, +0, -7}, {1, +0, +1},
			{-1, +1, +1}, {0, +1, +1}, {1, +1, +1}}}
	//5x5 全方向锐化滤波器
	Sharpen5All = FilterTemplate{Radius: 2, Size: 5, Scale: 1,
		Offsets: []FilterOffset{
			{-2, -2, -1}, {-1, -2, -1}, {0, -2, -1}, {1, -2, -1}, {2, -2, -1},
			{-2, -1, -1}, {-1, -1, +2}, {0, -1, +2}, {1, -1, +2}, {2, -1, -1},
			{-2, +0, -1}, {-1, +0, +2}, {0, +0, +8}, {1, +0, +2}, {2, +0, -1},
			{-2, +1, -1}, {-1, +1, +2}, {0, +1, +2}, {1, +1, +2}, {2, +1, -1},
			{-2, +2, -1}, {-1, +2, -1}, {1, +2, -1}, {1, +2, -1}, {2, +2, -1}}}
)