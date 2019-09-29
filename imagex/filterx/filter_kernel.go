package filterx

import "sort"

// 滤波器向量单元
type KernelVector struct {
	//向量X
	X int
	//向量Y
	Y int
	//向量值
	Value int
}

type FilterKernel []KernelVector

func (fk FilterKernel) Len() int {
	return len(fk)
}

func (fk FilterKernel) Less(i, j int) bool {
	return (fk[i].X < fk[j].X) || (fk[i].Y < fk[j].Y)
}

func (fk FilterKernel) Swap(i, j int) {
	fk[i], fk[j] = fk[j], fk[i]
}

//顺时针旋转90度
func (fk FilterKernel) RotateClockwise90() FilterKernel {
	ln := len(fk)
	kernel := make([]KernelVector, ln, ln)
	for i, kv := range fk {
		kernel[i] = kv
		if 0 == kv.X && 0 == kv.Y {
			continue
		}
		if (kv.X < 0 && kv.Y < 0) || (kv.X > 0 && kv.Y > 0) { //左上 或 右下
			kernel[i].X = -kernel[i].X
		} else { //左下 或 右下
			kernel[i].Y = -kernel[i].Y
		}
	}
	return kernel
}

//顺时针旋转180度
func (fk FilterKernel) RotateClockwise180() FilterKernel {
	ln := len(fk)
	kernel := make([]KernelVector, ln, ln)
	for i, kv := range fk {
		kernel[i] = kv
		if 0 == kv.X && 0 == kv.Y {
			continue
		}
		kernel[i].X, kernel[i].Y = -kernel[i].X, -kernel[i].Y
	}
	return kernel
}

//顺时针旋转270度
func (fk FilterKernel) RotateClockwise270() FilterKernel {
	ln := len(fk)
	kernel := make([]KernelVector, ln, ln)
	for i, kv := range fk {
		kernel[i] = kv
		if 0 == kv.X && 0 == kv.Y {
			continue
		}
		if (kv.X < 0 && kv.Y < 0) || (kv.X > 0 && kv.Y > 0) { //左上 或 右下
			kernel[i].Y = -kernel[i].Y
		} else { //左下 或 右下
			kernel[i].X = -kernel[i].X
		}
	}
	return kernel
}

//核排序
func (fk FilterKernel) Sorted() {
	sort.Sort(fk)
}
