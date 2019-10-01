//浮雕
package filterx

// 浮雕滤波器
var (
	//3x3 45度浮雕滤波器 右下
	Emboss3Oblique45 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, -1}, {-0, -1, -1},
			{-1, -0, -1}, {+1, +0, +1},
			{+0, +1, +1}, {+1, +1, +1}}}
	//5x5 浮雕滤波器
	Emboss5Oblique45 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -2, -1}, {+0, -2, -1}, {+1, -2, -1},
			{-2, -1, -1}, {-1, -1, -1}, {+0, -1, -1}, {+2, -1, +1},
			{-2, -0, -1}, {-1, -0, -1}, {+1, -0, +1}, {+2, -0, +1},
			{-2, +1, -1}, {+0, +1, +1}, {+1, +1, +1}, {+2, +1, +1},
			{-1, +2, +1}, {+0, +2, +1}, {+1, +2, +1}, {+2, +2, +1}}}
	//3x3 非对称浮雕滤波器
	Emboss3Asymmetrical = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, 2},
			{+0, +0, -1},
			{+1, +1, -1}}}
)

// 创建对称浮雕滤波器
func CreateSymmetryEmbossFilter(radius int, value int, offset int, angle int) (filter FilterMatrix, err error) {
	return
}

// 创建对称浮雕滤波器
func CreateAsymmetricalEmbossFilter(radius int, value int, offset int, angle int) (filter FilterMatrix, err error) {
	return
}
