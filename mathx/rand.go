package mathx

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandBool returns a pseudo-random boolean value.
// RandBool 返回伪随机布尔值。
func RandBool() bool {
	return rand.Int63() > math.MaxInt32
}

// RandFloat64 returns a pseudo-random float64 in [min, max).
// RandFloat64 返回 [min, max) 范围内的伪随机 float64。
func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandFloat32 returns a pseudo-random float32 in [min, max).
// RandFloat32 返回 [min, max) 范围内的伪随机 float32。
func RandFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// RandInt64 returns a pseudo-random int64 in [min, max).
// RandInt64 返回 [min, max) 范围内的伪随机 int64。
func RandInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// RandInt32 returns a pseudo-random int32 in [min, max).
// RandInt32 返回 [min, max) 范围内的伪随机 int32。
func RandInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min)
}

// RandUint64 returns a pseudo-random uint64 in [min, max).
// RandUint64 返回 [min, max) 范围内的伪随机 uint64。
func RandUint64(min, max uint64) uint64 {
	return min + uint64(rand.Int63n(int64(max-min)))
}

// RandUint32 returns a pseudo-random uint32 in [min, max).
// RandUint32 返回 [min, max) 范围内的伪随机 uint32。
func RandUint32(min, max uint32) uint32 {
	return min + uint32(rand.Int31n(int32(max-min)))
}
