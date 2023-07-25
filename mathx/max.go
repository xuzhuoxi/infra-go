// Package mathx
// Create on 2023/7/23
// @author xuzhuoxi
package mathx

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Max3Int(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func MaxUInt(a, b uint) uint {
	if a > b {
		return a
	} else {
		return b
	}
}

func Max3UInt(a, b, c uint) uint {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func Max3Float64(a, b, c float64) float64 {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	} else {
		return b
	}
}
