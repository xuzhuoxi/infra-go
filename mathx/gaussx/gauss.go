package gaussx

import (
	"math"
)

// 二维高斯函数
// sigma: 标准差
func GaussFunc2(x, y int, sigma float64) float64 {
	return (1.0 / (2.0 * math.Pi * sigma * sigma)) * math.Exp(-(float64(x*x+y*y) / (2.0 * sigma * sigma)))
}

// 计算高斯卷积核(二维数据)
// radius： 半径
// sigma： 	标准差
func CreateGaussKernel(radius int, sigma float64) [][]float64 {
	kSize := radius*2 + 1
	center := kSize / 2
	sum := 0.0
	rs := make([][]float64, kSize, kSize)
	for i := 0; i < kSize; i++ {
		arr := make([]float64, kSize, kSize)
		for j := 0; j < kSize; j++ {
			gr := GaussFunc2(i-center, j-center, sigma)
			arr[j] = gr
			sum += gr
		}
		rs[i] = arr
	}
	for i := 0; i < kSize; i++ {
		for j := 0; j < kSize; j++ {
			rs[i][j] = rs[i][j] / sum
		}
	}
	return rs
}

// 计算高斯卷积核(一维数据)
// radius： 半径
// sigma： 	标准差
func CreateGaussKernel2(radius int, sigma float64) []float64 {
	kSize := radius*2 + 1
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
	for index, _ := range rs {
		rs[index] /= sum
	}
	return rs
}

// 计算高斯卷积核(二维数据)
// 整数值卷积核
// radius： 	半径
// sigma： 		标准差
// integer： 	整数乘数
func CreateGaussKernelInteger(radius int, sigma float64, scale float64) [][]int {
	kernel := CreateGaussKernel(radius, sigma)
	kSize := radius*2 + 1
	rs := make([][]int, kSize, kSize)
	for i := 0; i < kSize; i++ {
		arr := make([]int, kSize, kSize)
		for j := 0; j < kSize; j++ {
			arr[j] = int(kernel[i][j] * scale)
		}
		rs[i] = arr
	}
	return rs
}

// 计算高斯卷积核(一维数据)
// 整数值卷积核
// radius： 	半径
// sigma： 		标准差
// integer： 	整数乘数
func CreateGaussKernelInteger2(radius int, sigma float64, scale float64) []int {
	kernel := CreateGaussKernel2(radius, sigma)
	kSize := radius*2 + 1
	rs := make([]int, kSize*kSize, kSize*kSize)
	for index, _ := range kernel {
		rs[index] = int(kernel[index] * float64(scale))
	}
	return rs
}
