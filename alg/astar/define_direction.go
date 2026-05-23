package astar

// Direction identifies a neighbor step on the grid.
// Direction 表示网格上的邻域移动方向。
type Direction uint8

// 3D directions (Z0 plane, Z+1 plane, Z-1 plane).
// 三维方向（Z0 平面、Z 增加平面、Z 减小平面）。
const (
	X0_Y0_Z0 Direction = iota // Horizontal center / 水平方向：中心 ----------------
	X0_Y1_Z0                  // Horizontal north / 水平方向：↑
	X1_Y1_Z0                  // Horizontal northeast / 水平方向：↗
	X1_Y0_Z0                  // Horizontal east / 水平方向：→
	X1_Y__Z0                  // Horizontal southeast / 水平方向：↘
	X0_Y__Z0                  // Horizontal south / 水平方向：↓
	X__Y__Z0                  // Horizontal southwest / 水平方向：↙
	X__Y0_Z0                  // Horizontal west / 水平方向：←
	X__Y1_Z0                  // Horizontal northwest / 水平方向：↖
	X0_Y0_Z1                  // Z+ center / Z 增加方向：中心 ----------------
	X0_Y1_Z1                  // Z+ north / Z 增加方向：↑
	X1_Y1_Z1                  // Z+ northeast / Z 增加方向：↗
	X1_Y0_Z1                  // Z+ east / Z 增加方向：→
	X1_Y__Z1                  // Z+ southeast / Z 增加方向：↘
	X0_Y__Z1                  // Z+ south / Z 增加方向：↓
	X__Y__Z1                  // Z+ southwest / Z 增加方向：↙
	X__Y0_Z1                  // Z+ west / Z 增加方向：←
	X__Y1_Z1                  // Z+ northwest / Z 增加方向：↖
	X0_Y0_Z_                  // Z- center / Z 减小方向：中心 ----------------
	X0_Y1_Z_                  // Z- north / Z 减小方向：↑
	X1_Y1_Z_                  // Z- northeast / Z 减小方向：↗
	X1_Y0_Z_                  // Z- east / Z 减小方向：→
	X1_Y__Z_                  // Z- southeast / Z 减小方向：↘
	X0_Y__Z_                  // Z- south / Z 减小方向：↓
	X__Y__Z_                  // Z- southwest / Z 减小方向：↙
	X__Y0_Z_                  // Z- west / Z 减小方向：←
	X__Y1_Z_                  // Z- northwest / Z 减小方向：↖
)

// 2D direction aliases on the Z0 plane.
// 二维方向别名（Z0 平面）。
const (
	Center    = X0_Y0_Z0 // Center / 中心
	North     = X0_Y1_Z0 // North / 北
	EastNorth = X1_Y1_Z0 // Northeast / 东北
	East      = X1_Y0_Z0 // East / 东
	EastSouth = X1_Y__Z0 // Southeast / 东南
	South     = X0_Y__Z0 // South / 南
	WestSouth = X__Y__Z0 // Southwest / 西南
	West      = X__Y0_Z0 // West / 西
	WestNorth = X__Y1_Z0 // Northwest / 西北
)

var (
	// DefaultDirections2D lists default 2D pathfinding directions.
	// DefaultDirections2D 默认二维寻路方向。
	DefaultDirections2D = []Direction{North, EastNorth, East, EastSouth, South, WestSouth, West, WestNorth}
	// DefaultDirections3D lists default 3D pathfinding directions.
	// DefaultDirections3D 默认三维寻路方向。
	DefaultDirections3D = []Direction{
		X0_Y1_Z0, X1_Y1_Z0, X1_Y0_Z0, X1_Y__Z0, X0_Y__Z0, X__Y__Z0, X__Y0_Z0, X__Y1_Z0,
		X0_Y0_Z1, X0_Y1_Z1, X1_Y1_Z1, X1_Y0_Z1, X1_Y__Z1, X0_Y__Z1, X__Y__Z1, X__Y0_Z1, X__Y1_Z1,
		X0_Y0_Z_, X0_Y1_Z_, X1_Y1_Z_, X1_Y0_Z_, X1_Y__Z_, X0_Y__Z_, X__Y__Z_, X__Y0_Z_, X__Y1_Z_}
	// ObliqueDirections2D lists 2D diagonal directions only.
	// ObliqueDirections2D 二维斜向方向。
	ObliqueDirections2D = []Direction{EastNorth, EastSouth, WestSouth, WestNorth}
	// ObliqueDirections3D lists 3D diagonal directions only.
	// ObliqueDirections3D 三维斜向方向。
	ObliqueDirections3D = []Direction{
		X0_Y1_Z0, X1_Y0_Z0, X0_Y__Z0, X__Y0_Z0,
		X0_Y0_Z1, X0_Y1_Z1, X1_Y0_Z1, X0_Y__Z1, X__Y0_Z1,
		X0_Y0_Z_, X0_Y1_Z_, X1_Y0_Z_, X0_Y__Z_, X__Y0_Z_}
)

//-------------------------------------------------------

// DirectionVector is a step delta [dx, dy, dz, cost].
// DirectionVector 表示方向矢量 [dx, dy, dz, 代价]。
type DirectionVector [4]int

// X returns the X delta.
// X 返回 X 方向增量。
func (v DirectionVector) X() int {
	return v[0]
}

// Y returns the Y delta.
// Y 返回 Y 方向增量。
func (v DirectionVector) Y() int {
	return v[1]
}

// Z returns the Z delta.
// Z 返回 Z 方向增量。
func (v DirectionVector) Z() int {
	return v[2]
}

// V returns the movement cost for this step.
// V 返回该方向的移动代价值。
func (v DirectionVector) V() int {
	return v[3]
}

// Default 3D direction vectors.
// 默认三维方向矢量定义。
var (
	Vector_X0_Y0_Z0 = DirectionVector{0, 0, 0, 5}    // Horizontal center / 水平方向：中心 ----------------
	Vector_X0_Y1_Z0 = DirectionVector{0, 1, 0, 5}    // Horizontal north / 水平方向：↑
	Vector_X1_Y1_Z0 = DirectionVector{1, 1, 0, 5}    // Horizontal northeast / 水平方向：↗
	Vector_X1_Y0_Z0 = DirectionVector{1, 0, 0, 5}    // Horizontal east / 水平方向：→
	Vector_X1_Y__Z0 = DirectionVector{1, -1, 0, 5}   // Horizontal southeast / 水平方向：↘
	Vector_X0_Y__Z0 = DirectionVector{0, -1, 0, 5}   // Horizontal south / 水平方向：↓
	Vector_X__Y__Z0 = DirectionVector{-1, -1, 0, 5}  // Horizontal southwest / 水平方向：↙
	Vector_X__Y0_Z0 = DirectionVector{-1, 0, 0, 5}   // Horizontal west / 水平方向：←
	Vector_X__Y1_Z0 = DirectionVector{-1, 1, 0, 5}   // Horizontal northwest / 水平方向：↖
	Vector_X0_Y0_Z1 = DirectionVector{0, 0, 1, 5}    // Z+ center / Z 增加方向：中心 ----------------
	Vector_X0_Y1_Z1 = DirectionVector{0, 1, 1, 5}    // Z+ north / Z 增加方向：↑
	Vector_X1_Y1_Z1 = DirectionVector{1, 1, 1, 5}    // Z+ northeast / Z 增加方向：↗
	Vector_X1_Y0_Z1 = DirectionVector{1, 0, 1, 5}    // Z+ east / Z 增加方向：→
	Vector_X1_Y__Z1 = DirectionVector{1, -1, 1, 5}   // Z+ southeast / Z 增加方向：↘
	Vector_X0_Y__Z1 = DirectionVector{0, -1, 1, 5}   // Z+ south / Z 增加方向：↓
	Vector_X__Y__Z1 = DirectionVector{-1, -1, 1, 5}  // Z+ southwest / Z 增加方向：↙
	Vector_X__Y0_Z1 = DirectionVector{-1, 0, 1, 5}   // Z+ west / Z 增加方向：←
	Vector_X__Y1_Z1 = DirectionVector{-1, 1, 1, 5}   // Z+ northwest / Z 增加方向：↖
	Vector_X0_Y0_Z_ = DirectionVector{0, 0, -1, 5}   // Z- center / Z 减小方向：中心 ----------------
	Vector_X0_Y1_Z_ = DirectionVector{0, 1, -1, 5}   // Z- north / Z 减小方向：↑
	Vector_X1_Y1_Z_ = DirectionVector{1, 1, -1, 5}   // Z- northeast / Z 减小方向：↗
	Vector_X1_Y0_Z_ = DirectionVector{1, 0, -1, 5}   // Z- east / Z 减小方向：→
	Vector_X1_Y__Z_ = DirectionVector{1, -1, -1, 5}  // Z- southeast / Z 减小方向：↘
	Vector_X0_Y__Z_ = DirectionVector{0, -1, -1, 5}  // Z- south / Z 减小方向：↓
	Vector_X__Y__Z_ = DirectionVector{-1, -1, -1, 5} // Z- southwest / Z 减小方向：↙
	Vector_X__Y0_Z_ = DirectionVector{-1, 0, -1, 5}  // Z- west / Z 减小方向：←
	Vector_X__Y1_Z_ = DirectionVector{-1, 1, -1, 5}  // Z- northwest / Z 减小方向：↖
)

// Default 2D direction vector aliases.
// 默认二维方向矢量别名。
var (
	VectorCenter    = Vector_X0_Y0_Z0 // Center / 中心
	VectorNorth     = Vector_X0_Y1_Z0 // North / 北
	VectorEastNorth = Vector_X1_Y1_Z0 // Northeast / 东北
	VectorEast      = Vector_X1_Y0_Z0 // East / 东
	VectorEastSouth = Vector_X1_Y__Z0 // Southeast / 东南
	VectorSouth     = Vector_X0_Y__Z0 // South / 南
	VectorWestSouth = Vector_X__Y__Z0 // Southwest / 西南
	VectorWest      = Vector_X__Y0_Z0 // West / 西
	VectorWestNorth = Vector_X__Y1_Z0 // Northwest / 西北
)

var vectors = []DirectionVector{
	Vector_X0_Y0_Z0, Vector_X0_Y1_Z0, Vector_X1_Y1_Z0, Vector_X1_Y0_Z0, Vector_X1_Y__Z0, Vector_X0_Y__Z0, Vector_X__Y__Z0, Vector_X__Y0_Z0, Vector_X__Y1_Z0,
	Vector_X0_Y0_Z1, Vector_X0_Y1_Z1, Vector_X1_Y1_Z1, Vector_X1_Y0_Z1, Vector_X1_Y__Z1, Vector_X0_Y__Z1, Vector_X__Y__Z1, Vector_X__Y0_Z1, Vector_X__Y1_Z1,
	Vector_X0_Y0_Z_, Vector_X0_Y1_Z_, Vector_X1_Y1_Z_, Vector_X1_Y0_Z_, Vector_X1_Y__Z_, Vector_X0_Y__Z_, Vector_X__Y__Z_, Vector_X__Y0_Z_, Vector_X__Y1_Z_}

// GetDirVector returns the direction vector for dir.
// GetDirVector 返回方向对应的矢量定义。
func GetDirVector(dir Direction) DirectionVector {
	return vectors[dir]
}
