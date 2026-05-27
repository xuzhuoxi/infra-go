package filterx

import "sort"

// KernelVector is one sparse offset-weight pair in a convolution kernel.
// KernelVector 是卷积核中的一个稀疏偏移-权重对。
type KernelVector struct {
	//向量X
	X int
	//向量Y
	Y int
	//向量值
	Value int
}

// FilterKernel is a sparse list of KernelVector entries.
// FilterKernel 是 KernelVector 稀疏列表。
type FilterKernel []KernelVector

// Len returns the number of kernel vectors (sort.Interface).
// Len 返回核向量数量（sort.Interface）。
func (fk FilterKernel) Len() int {
	return len(fk)
}

// Less compares two kernel vectors for sorting (sort.Interface).
// Less 比较两个核向量以排序（sort.Interface）。
func (fk FilterKernel) Less(i, j int) bool {
	if fk[i].Y == fk[j].Y {
		return fk[i].X < fk[j].X
	} else {
		return fk[i].Y < fk[j].Y
	}
}

// Swap exchanges two kernel vectors (sort.Interface).
// Swap 交换两个核向量（sort.Interface）。
func (fk FilterKernel) Swap(i, j int) {
	fk[i], fk[j] = fk[j], fk[i]
}

// Clone returns a shallow copy of the kernel slice.
// Clone 返回核向量的浅拷贝。
func (fk FilterKernel) Clone() FilterKernel {
	ln := len(fk)
	kernel := make([]KernelVector, ln, ln)
	for i, kv := range fk {
		kernel[i] = kv
	}
	return kernel
}

// FlipUDSelf flips Y offsets in place (up-down).
// FlipUDSelf 原地上下翻转 Y 偏移。
func (fk FilterKernel) FlipUDSelf() {
	for index := range fk {
		if 0 == fk[index].Y {
			continue
		}
		fk[index].Y = -fk[index].Y
	}
}

// FlipUD returns a vertically flipped copy.
// FlipUD 返回上下翻转后的副本。
func (fk FilterKernel) FlipUD() FilterKernel {
	rs := fk.Clone()
	rs.FlipUDSelf()
	return rs
}

// FlipLRSelf flips X offsets in place (left-right).
// FlipLRSelf 原地左右翻转 X 偏移。
func (fk FilterKernel) FlipLRSelf() {
	for index := range fk {
		if 0 == fk[index].X {
			continue
		}
		fk[index].X = -fk[index].X
	}
}

// FlipLR returns a horizontally flipped copy.
// FlipLR 返回左右翻转后的副本。
func (fk FilterKernel) FlipLR() FilterKernel {
	rs := fk.Clone()
	rs.FlipLRSelf()
	return rs
}

// Rotate90Self rotates the kernel 90° in place.
// Rotate90Self 原地旋转卷积核 90°。
// clockwise: true for clockwise rotation.
// clockwise：true 表示顺时针旋转。
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

// Rotate90 returns a copy rotated 90°.
// Rotate90 返回旋转 90° 后的副本。
// clockwise: true for clockwise rotation.
// clockwise：true 表示顺时针旋转。
func (fk FilterKernel) Rotate90(clockwise bool) FilterKernel {
	rs := fk.Clone()
	rs.Rotate90Self(clockwise)
	return rs
}

// RotateSelf rotates the kernel count90 quarter-turns in place.
// RotateSelf 原地旋转 count90 个 90°。
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

// Rotate returns a copy rotated count90 quarter-turns.
// Rotate 返回旋转 count90 个 90° 后的副本。
func (fk FilterKernel) Rotate(clockwise bool, count90 int) FilterKernel {
	rs := fk.Clone()
	rs.RotateSelf(clockwise, count90)
	return rs
}

// Sorted sorts kernel vectors by (Y, X).
// Sorted 按 (Y, X) 对核向量排序。
func (fk FilterKernel) Sorted() {
	sort.Sort(fk)
}
