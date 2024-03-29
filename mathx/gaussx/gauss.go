package gaussx

import (
	"math"
)

// GaussFunc2
// 二维高斯函数
// sigma: 标准差
func GaussFunc2(x, y int, sigma float64) float64 {
	return (1.0 / (2.0 * math.Pi * sigma * sigma)) * math.Exp(-(float64(x*x+y*y) / (2.0 * sigma * sigma)))
}

// CreateGaussKernel
// 计算高斯卷积核(二维数据)
// radius： 半径
// sigma： 	标准差
func CreateGaussKernel(radius int, sigma float64) [][]float64 {
	kSize := radius + radius + 1
	center := kSize / 2
	sum := 0.0
	rs := make([][]float64, kSize, kSize)
	for i := 0; i < kSize; i++ {
		rs[i] = make([]float64, kSize, kSize)
		for j := 0; j < kSize; j++ {
			gr := GaussFunc2(i-center, j-center, sigma)
			rs[i][j] = gr
			sum += gr
		}
	}
	for i := 0; i < kSize; i++ {
		for j := 0; j < kSize; j++ {
			rs[i][j] = rs[i][j] / sum
		}
	}
	return rs
}

//
//func GetAvgArr(radius int, sigma float64) [][]float64 {
//	kSize := radius + radius + 1
//	sum := 0.0
//	arr := make([][]float64, kSize, kSize)
//	for i := 0; i < kSize; i++ {
//		arr[i] = make([]float64, kSize, kSize)
//	}
//	for i := 0; i < radius; i++ {
//		weight := GaussFunc2(i-radius, 0, sigma)
//		arr[i][radius] = weight
//		sum += 4 * weight
//		for j := 0; j < radius; j++ {
//			thisGaussResult := GaussFunc2(i-radius, j-radius, sigma)
//			arr[i][j] = thisGaussResult
//			sum += 4 * thisGaussResult
//		}
//	}
//	weight := GaussFunc2(0, 0, sigma)
//	arr[radius][radius] = weight
//	sum += weight
//
//	for i := 0; i < radius; i++ {
//		arr[i][radius] /= sum
//		arr[2*radius-i][radius], arr[radius][i], arr[radius][2*radius-i] = arr[i][radius], arr[i][radius], arr[i][radius]
//
//		for j := 0; j < radius; j++ {
//			arr[i][j] /= sum
//			arr[i][2*radius-j], arr[2*radius-i][j], arr[2*radius-i][2*radius-j] = arr[i][j], arr[i][j], arr[i][j]
//
//		}
//	}
//	arr[radius][radius] /= sum
//	return arr
//}

// CreateGaussKernel2
// 计算高斯卷积核(一维数据)
// radius： 半径
// sigma： 	标准差
func CreateGaussKernel2(radius int, sigma float64) []float64 {
	kSize := radius + radius + 1
	center := kSize / 2
	sum := 0.0
	rs := make([]float64, kSize*kSize, kSize*kSize)
	for i := 0; i < kSize; i++ {
		for j := 0; j < kSize; j++ {
			index := i*kSize + j
			gr := GaussFunc2(i-center, j-center, sigma)
			rs[index] = gr
			sum += gr
		}
	}
	for index := range rs {
		rs[index] /= sum
	}
	return rs
}

// CreateGaussKernelInt
// 计算高斯卷积核(二维数据)
// 整数值卷积核
// radius： 	半径
// sigma： 		标准差
// scale： 		整数乘数, =0时使用默认放大模式
func CreateGaussKernelInt(radius int, sigma float64, scale float64) [][]int {
	kernel := CreateGaussKernel(radius, sigma)
	kSize := radius + radius + 1
	rs := make([][]int, kSize, kSize)
	if scale <= 0 {
		scale = math.Ceil(1 / kernel[0][0])
	}
	for i := 0; i < kSize; i++ {
		rs[i] = make([]int, kSize, kSize)
		for j := 0; j < kSize; j++ {
			rs[i][j] = int(kernel[i][j] * scale)
		}
	}
	return rs
}

// CreateGaussKernelInt2
// 计算高斯卷积核(一维数据)
// 整数值卷积核
// radius： 	半径
// sigma： 		标准差
// scale： 		整数乘数, =0时使用默认放大模式
func CreateGaussKernelInt2(radius int, sigma float64, scale float64) []int {
	kernel := CreateGaussKernel2(radius, sigma)
	kSize := radius + radius + 1
	rs := make([]int, kSize*kSize, kSize*kSize)
	if scale <= 0 {
		scale = 1 / kernel[0]
	}
	for index := range kernel {
		rs[index] = int(math.Round(kernel[index] * float64(scale)))
	}
	return rs
}
