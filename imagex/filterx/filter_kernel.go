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
	if fk[i].Y == fk[j].Y {
		return fk[i].X < fk[j].X
	} else {
		return fk[i].Y < fk[j].Y
	}
}

func (fk FilterKernel) Swap(i, j int) {
	fk[i], fk[j] = fk[j], fk[i]
}

func (fk FilterKernel) Clone() FilterKernel {
	ln := len(fk)
	kernel := make([]KernelVector, ln, ln)
	for i, kv := range fk {
		kernel[i] = kv
	}
	return kernel
}

//上下翻转自身
func (fk FilterKernel) FlipUDSelf() {
	for index := range fk {
		if 0 == fk[index].Y {
			continue
		}
		fk[index].Y = -fk[index].Y
	}
}

//上下翻转
func (fk FilterKernel) FlipUD() FilterKernel {
	rs := fk.Clone()
	rs.FlipUDSelf()
	return rs
}

//左右翻转自身
func (fk FilterKernel) FlipLRSelf() {
	for index := range fk {
		if 0 == fk[index].X {
			continue
		}
		fk[index].X = -fk[index].X
	}
}

//左右翻转
func (fk FilterKernel) FlipLR() FilterKernel {
	rs := fk.Clone()
	rs.FlipLRSelf()
	return rs
}

// 旋转90度
// clockwise ;是否为顺时针
func (fk FilterKernel) Rotate90Self(clockwise bool) {
	if clockwise {
		for index := range fk {
			fk[index].X, fk[index].Y = -fk[index].Y, fk[index].X
		}
	} else {
		for index := range fk {
			fk[index].X, fk[index].Y = fk[index].Y, -fk[index].X
		}
	}
}

//旋转90度
// clockwise ;是否为顺时针
func (fk FilterKernel) Rotate90(clockwise bool) FilterKernel {
	rs := fk.Clone()
	rs.Rotate90Self(clockwise)
	return rs
}

//旋转
func (fk FilterKernel) RotateSelf(clockwise bool, count90 int) {
	c := count90
	for c < 0 {
		c += 4
	}
	c = c % 4
	for c > 0 {
		fk.Rotate90Self(clockwise)
		c--
	}
}

//旋转
func (fk FilterKernel) Rotate(clockwise bool, count90 int) FilterKernel {
	rs := fk.Clone()
	rs.RotateSelf(clockwise, count90)
	return rs
}

//核排序
func (fk FilterKernel) Sorted() {
	sort.Sort(fk)
}
