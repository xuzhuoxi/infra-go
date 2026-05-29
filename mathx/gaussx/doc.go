// Package gaussx provides two-dimensional Gaussian functions and convolution kernel builders.
// Package gaussx 提供二维高斯函数及卷积核构造能力。
//
// Capabilities:
//   - GaussFunc2: evaluate the 2D Gaussian PDF at integer offsets with a given sigma
//   - CreateGaussKernel / CreateGaussKernel2: normalized 2D matrix or flattened 1D kernels
//   - CreateGaussKernelInt / CreateGaussKernelInt2: integer-scaled kernel variants for fixed-point use
//
// 主要职责：
//   - GaussFunc2：按给定 sigma 在整数偏移处计算二维高斯概率密度
//   - CreateGaussKernel / CreateGaussKernel2：构造归一化的二维矩阵或一维展平卷积核
//   - CreateGaussKernelInt / CreateGaussKernelInt2：整数放大版卷积核，便于定点运算
package gaussx
