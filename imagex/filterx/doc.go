// Package filterx provides convolution filters, fast blur, grayscale conversion, binarization, and morphological operations.
// Package filterx 提供卷积滤波、快速模糊、灰度转换、二值化及形态学（腐蚀/膨胀）等图像滤波处理能力。
//
// Capabilities:
//   - FilterMatrix / FilterKernel: sparse convolution kernels with flip, rotate, and validity checks
//   - Blur: box, Gaussian, motion presets; Create* builders; FastBlur stack-blur variants
//   - Edge and Sharpen: Laplace and directional kernels; preset and custom radius helpers
//   - Emboss: symmetry and asymmetry emboss filter builders
//   - GrayFromNRGBA / GrayFromRGBA; NRGBAAt* alpha compositing onto backgrounds
//   - CVTGray, OTSU / Kittler thresholds; ErodeGray / DilateGray with PixelDirection
//
// 主要职责：
//   - FilterMatrix / FilterKernel：稀疏卷积核，支持翻转、旋转与有效性校验
//   - 模糊：均值、高斯、运动模糊预设与 Create* 构造；FastBlur 栈模糊实现
//   - 边缘与锐化：拉普拉斯及方向性卷积核；预设与自定义半径辅助函数
//   - 浮雕：对称与非对称浮雕滤波器构造
//   - GrayFromNRGBA / GrayFromRGBA；NRGBAAt* 将 RGBA 按背景色去 Alpha
//   - CVTGray、OTSU / Kittler 阈值；ErodeGray / DilateGray 配合 PixelDirection
package filterx
