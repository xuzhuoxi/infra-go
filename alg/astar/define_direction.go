//
//Created by xuzhuoxi
//on 2019-04-05.
//@author xuzhuoxi
//
package astar

// 方向定义
type Direction uint8

// 3D方向
const (
	X0_Y0_Z0 Direction = iota // 水平方向：中心 ----------------
	X0_Y1_Z0                  // 水平方向：↑
	X1_Y1_Z0                  // 水平方向：↗
	X1_Y0_Z0                  // 水平方向：→
	X1_Y__Z0                  // 水平方向：↘
	X0_Y__Z0                  // 水平方向：↓
	X__Y__Z0                  // 水平方向：↙
	X__Y0_Z0                  // 水平方向：←
	X__Y1_Z0                  // 水平方向：↖
	X0_Y0_Z1                  // Z增加方向：中心 ----------------
	X0_Y1_Z1                  // Z增加方向：↑
	X1_Y1_Z1                  // Z增加方向：↗
	X1_Y0_Z1                  // Z增加方向：→
	X1_Y__Z1                  // Z增加方向：↘
	X0_Y__Z1                  // Z增加方向：↓
	X__Y__Z1                  // Z增加方向：↙
	X__Y0_Z1                  // Z增加方向：←
	X__Y1_Z1                  // Z增加方向：↖
	X0_Y0_Z_                  // Z减小方向：中心 ----------------
	X0_Y1_Z_                  // Z减小方向：↑
	X1_Y1_Z_                  // Z减小方向：↗
	X1_Y0_Z_                  // Z减小方向：→
	X1_Y__Z_                  // Z减小方向：↘
	X0_Y__Z_                  // Z减小方向：↓
	X__Y__Z_                  // Z减小方向：↙
	X__Y0_Z_                  // Z减小方向：←
	X__Y1_Z_                  // Z减小方向：↖
)

// 2D方向
const (
	Center    = X0_Y0_Z0
	North     = X0_Y1_Z0
	EastNorth = X1_Y1_Z0
	East      = X1_Y0_Z0
	EastSouth = X1_Y__Z0
	South     = X0_Y__Z0
	WestSouth = X__Y__Z0
	West      = X__Y0_Z0
	WestNorth = X__Y1_Z0
)

// 默认寻路方向
var (
	DefaultDirections2D = []Direction{North, EastNorth, East, EastSouth, South, WestSouth, West, WestNorth}
	DefaultDirections3D = []Direction{
		X0_Y1_Z0, X1_Y1_Z0, X1_Y0_Z0, X1_Y__Z0, X0_Y__Z0, X__Y__Z0, X__Y0_Z0, X__Y1_Z0,
		X0_Y0_Z1, X0_Y1_Z1, X1_Y1_Z1, X1_Y0_Z1, X1_Y__Z1, X0_Y__Z1, X__Y__Z1, X__Y0_Z1, X__Y1_Z1,
		X0_Y0_Z_, X0_Y1_Z_, X1_Y1_Z_, X1_Y0_Z_, X1_Y__Z_, X0_Y__Z_, X__Y__Z_, X__Y0_Z_, X__Y1_Z_}
	ObliqueDirections2D = []Direction{EastNorth, EastSouth, WestSouth, WestNorth}
	ObliqueDirections3D = []Direction{
		X0_Y1_Z0, X1_Y0_Z0, X0_Y__Z0, X__Y0_Z0,
		X0_Y0_Z1, X0_Y1_Z1, X1_Y0_Z1, X0_Y__Z1, X__Y0_Z1,
		X0_Y0_Z_, X0_Y1_Z_, X1_Y0_Z_, X0_Y__Z_, X__Y0_Z_}
)

//-------------------------------------------------------

// 方向矢量
type DirectionVector [4]int

// 坐标X方向增量
func (v DirectionVector) X() int {
	return v[0]
}

// 坐标Y方向增量
func (v DirectionVector) Y() int {
	return v[1]
}

// 坐标Z方向增量
func (v DirectionVector) Z() int {
	return v[2]
}

// 代价值
func (v DirectionVector) V() int {
	return v[3]
}

// 默认3D方向矢量定义
var (
	Vector_X0_Y0_Z0 = DirectionVector{0, 0, 0, 5}    // 水平方向：中心 ----------------
	Vector_X0_Y1_Z0 = DirectionVector{0, 1, 0, 5}    // 水平方向：↑
	Vector_X1_Y1_Z0 = DirectionVector{1, 1, 0, 5}    // 水平方向：↗
	Vector_X1_Y0_Z0 = DirectionVector{1, 0, 0, 5}    // 水平方向：→
	Vector_X1_Y__Z0 = DirectionVector{1, -1, 0, 5}   // 水平方向：↘
	Vector_X0_Y__Z0 = DirectionVector{0, -1, 0, 5}   // 水平方向：↓
	Vector_X__Y__Z0 = DirectionVector{-1, -1, 0, 5}  // 水平方向：↙
	Vector_X__Y0_Z0 = DirectionVector{-1, 0, 0, 5}   // 水平方向：←
	Vector_X__Y1_Z0 = DirectionVector{-1, 1, 0, 5}   // 水平方向：↖
	Vector_X0_Y0_Z1 = DirectionVector{0, 0, 1, 5}    // Z增加方向：中心 ----------------
	Vector_X0_Y1_Z1 = DirectionVector{0, 1, 1, 5}    // Z增加方向：↑
	Vector_X1_Y1_Z1 = DirectionVector{1, 1, 1, 5}    // Z增加方向：↗
	Vector_X1_Y0_Z1 = DirectionVector{1, 0, 1, 5}    // Z增加方向：→
	Vector_X1_Y__Z1 = DirectionVector{1, -1, 1, 5}   // Z增加方向：↘
	Vector_X0_Y__Z1 = DirectionVector{0, -1, 1, 5}   // Z增加方向：↓
	Vector_X__Y__Z1 = DirectionVector{-1, -1, 1, 5}  // Z增加方向：↙
	Vector_X__Y0_Z1 = DirectionVector{-1, 0, 1, 5}   // Z增加方向：←
	Vector_X__Y1_Z1 = DirectionVector{-1, 1, 1, 5}   // Z增加方向：↖
	Vector_X0_Y0_Z_ = DirectionVector{0, 0, -1, 5}   // Z减小方向：中心 ----------------
	Vector_X0_Y1_Z_ = DirectionVector{0, 1, -1, 5}   // Z减小方向：↑
	Vector_X1_Y1_Z_ = DirectionVector{1, 1, -1, 5}   // Z减小方向：↗
	Vector_X1_Y0_Z_ = DirectionVector{1, 0, -1, 5}   // Z减小方向：→
	Vector_X1_Y__Z_ = DirectionVector{1, -1, -1, 5}  // Z减小方向：↘
	Vector_X0_Y__Z_ = DirectionVector{0, -1, -1, 5}  // Z减小方向：↓
	Vector_X__Y__Z_ = DirectionVector{-1, -1, -1, 5} // Z减小方向：↙
	Vector_X__Y0_Z_ = DirectionVector{-1, 0, -1, 5}  // Z减小方向：←
	Vector_X__Y1_Z_ = DirectionVector{-1, 1, -1, 5}  // Z减小方向：↖
)

// 默认2D方向矢量定义
var (
	VectorCenter    = Vector_X0_Y0_Z0
	VectorNorth     = Vector_X0_Y1_Z0
	VectorEastNorth = Vector_X1_Y1_Z0
	VectorEast      = Vector_X1_Y0_Z0
	VectorEastSouth = Vector_X1_Y__Z0
	VectorSouth     = Vector_X0_Y__Z0
	VectorWestSouth = Vector_X__Y__Z0
	VectorWest      = Vector_X__Y0_Z0
	VectorWestNorth = Vector_X__Y1_Z0
)

var vectors = []DirectionVector{
	Vector_X0_Y0_Z0, Vector_X0_Y1_Z0, Vector_X1_Y1_Z0, Vector_X1_Y0_Z0, Vector_X1_Y__Z0, Vector_X0_Y__Z0, Vector_X__Y__Z0, Vector_X__Y0_Z0, Vector_X__Y1_Z0,
	Vector_X0_Y0_Z1, Vector_X0_Y1_Z1, Vector_X1_Y1_Z1, Vector_X1_Y0_Z1, Vector_X1_Y__Z1, Vector_X0_Y__Z1, Vector_X__Y__Z1, Vector_X__Y0_Z1, Vector_X__Y1_Z1,
	Vector_X0_Y0_Z_, Vector_X0_Y1_Z_, Vector_X1_Y1_Z_, Vector_X1_Y0_Z_, Vector_X1_Y__Z_, Vector_X0_Y__Z_, Vector_X__Y__Z_, Vector_X__Y0_Z_, Vector_X__Y1_Z_}

// 取方向对应的矢量定义
func GetDirVector(dir Direction) DirectionVector {
	return vectors[dir]
}
