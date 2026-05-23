package astar

import (
	"errors"

	"github.com/xuzhuoxi/infra-go/mathx"
)

const (
	DEFAULT_MASK_ADD = 1048576 // Mask increment per search pass / 每次寻路的掩码增量
	//DEFAULT_MASK_ADD = 10
	MAX_MASK = 2000000000 // Upper bound for path masks / 路径掩码上限
)

// FuncHn estimates the cost from a cell to the goal (heuristic Hn).
// FuncHn 估值函数 Hn，表示下一点到终点的估算代价。
type FuncHn func(cx, cy, cz, ex, ey, ez int) int

// DefaultFuncHn returns the default Manhattan-distance heuristic.
// DefaultFuncHn 返回默认的曼哈顿距离估值函数。
func DefaultFuncHn(cx, cy, cz, ex, ey, ez int) int {
	return mathx.AbsInt(cx-ex) + mathx.AbsInt(cy-ey) + mathx.AbsInt(cz-ez)
}

// IAStarAlg defines the A* pathfinding algorithm API.
// IAStarAlg 定义 A* 寻路算法接口。
type IAStarAlg interface {
	// SetAllowedDirections sets which directions may be explored (for special maps).
	// SetAllowedDirections 设置允许检索的方向（可用于特殊地图）。
	SetAllowedDirections(direction []Direction)
	// SetCustomFuncHn sets a custom heuristic function.
	// SetCustomFuncHn 设置自定义的估值函数。
	SetCustomFuncHn(hn FuncHn)

	// InitMapSize initializes the map size for 3D grids.
	// InitMapSize 初始化三维地图尺寸。
	InitMapSize(width, height, depth int)
	// InitMapSize2D initializes the map size for 2D grids.
	// InitMapSize2D 初始化二维地图尺寸。
	InitMapSize2D(width, height int)
	// SetData loads flat map cell data.
	// SetData 设置一维展开的地图数据。
	SetData(data []int) (sourceData [][][]int, err error)
	// SetData2D loads 2D map cell data.
	// SetData2D 设置二维地图数据。
	SetData2D(data [][]int) (sourceData [][][]int, err error)
	// SetData3D loads 3D map cell data.
	// SetData3D 设置三维地图数据。
	SetData3D(data [][][]int) (sourceData [][][]int, err error)
	// Search2D searches a path on a 2D grid.
	// Search2D 在二维网格上寻路。
	// sx is StartX; sy is StartY; ex is EndX; ey is EndY.
	// sx 为起点 X；sy 为起点 Y；ex 为终点 X；ey 为终点 Y。
	Search2D(sx, sy, ex, ey int) (path []Position, ok bool)
	// Search searches a path on a 3D grid.
	// Search 在三维网格上寻路。
	// sx is StartX; sy is StartY; sz is StartZ; ex is EndX; ey is EndY; ez is EndZ.
	// sx 为起点 X；sy 为起点 Y；sz 为起点 Z；ex 为终点 X；ey 为终点 Y；ez 为终点 Z。
	Search(sx, sy, sz, ex, ey, ez int) (path []Position, ok bool)
	// SearchPosition searches a path between two positions.
	// SearchPosition 在两点之间寻路。
	SearchPosition(startPos, endPos Position) (path []Position, ok bool)
}

// NewIAStartAlg returns a new IAStarAlg implementation.
// NewIAStartAlg 创建 IAStarAlg 实现实例。
func NewIAStartAlg() IAStarAlg {
	return NewAStartAlg()
}

// NewAStartAlg constructs a new AStarAlg.
// NewAStartAlg 创建 AStarAlg 实例。
func NewAStartAlg() *AStarAlg {
	return &AStarAlg{openQueue: newSortedQueue()}
}

type history struct {
	startPos Position
	endPos   Position
	// 包含startPos,endPos
	path   []Position
	pathOk bool
}

func (h history) getSubPath(sIndex, eIndex int) []Position {
	if nil == h.path {
		return nil
	}
	if 0 == len(h.path) {
		return h.path
	}
	ln := eIndex - sIndex + 1
	rs := make([]Position, ln)
	copy(rs, h.path[sIndex:eIndex+1])
	return rs
}

func (h history) checkSubPath(startPos Position, endPos Position) (sub bool, sIndex int, eIndex int) {
	if !h.pathOk {
		return
	}
	sIndex = h.getFirstPositionIndex(startPos)
	eIndex = h.getLastPositionIndex(endPos)
	if -1 == sIndex || -1 == eIndex {
		return
	}
	if sIndex > eIndex {
		return
	}
	return true, sIndex, eIndex
}

func (h history) getFirstPositionIndex(pos Position) int {
	for index := range h.path {
		if h.path[index].EqualTo(pos) {
			return index
		}
	}
	return -1
}

func (h history) getLastPositionIndex(pos Position) int {
	for index := len(h.path) - 1; index >= 0; index-- {
		if h.path[index].EqualTo(pos) {
			return index
		}
	}
	return -1
}

// AStarAlg implements the A* pathfinding algorithm on grid maps.
// AStarAlg 在网格地图上实现 A* 寻路算法。
type AStarAlg struct {
	// 地图的大小
	size *Size
	// 可行方向
	nextDirs []Direction
	// 自定义估值函数
	customFuncHn FuncHn
	// 初始地图数据
	sourceMap [][][]int
	// 检索地图数据
	searchMap [][][]int

	baseMask  int
	maxMask   int
	startPos  Position
	endPos    Position
	openQueue *sortedQueue
	history   history
}

// SetAllowedDirections sets which directions may be explored.
// SetAllowedDirections 设置允许检索的方向。
func (alg *AStarAlg) SetAllowedDirections(direction []Direction) {
	cp := make([]Direction, len(direction))
	copy(cp, direction)
	alg.nextDirs = cp
}

// SetCustomFuncHn sets a custom heuristic function.
// SetCustomFuncHn 设置自定义估值函数。
func (alg *AStarAlg) SetCustomFuncHn(hn FuncHn) {
	alg.customFuncHn = hn
}

// InitMapSize initializes the map size for 3D grids.
// InitMapSize 初始化三维地图尺寸。
func (alg *AStarAlg) InitMapSize(width, height, depth int) {
	alg.size = &Size{Width: width, Height: height, Depth: depth}
}

// InitMapSize2D initializes the map size for 2D grids.
// InitMapSize2D 初始化二维地图尺寸。
func (alg *AStarAlg) InitMapSize2D(width, height int) {
	alg.size = &Size{Width: width, Height: height, Depth: 1}
}

// SetData loads flat map cell data.
// SetData 设置一维展开的地图数据。
func (alg *AStarAlg) SetData(data []int) (sourceData [][][]int, err error) {
	if len(data) != alg.size.Area() {
		return nil, errors.New("Data Error. ")
	}
	alg.sourceMap, _ = alg.copyData(data)
	alg.updateSearchMapFromSource()
	alg.clearHistory()
	return alg.sourceMap, nil
}

// SetData2D loads 2D map cell data.
// SetData2D 设置二维地图数据。
func (alg *AStarAlg) SetData2D(data [][]int) (sourceData [][][]int, err error) {
	if nil == data || len(data) == 0 {
		return nil, errors.New("Data Error. ")
	}
	var data3d [][][]int
	var ln int
	if data3d, ln = alg.copyData2D(data); ln != alg.size.Area() {
		return nil, errors.New("Data Error. ")
	}
	alg.sourceMap = data3d
	alg.updateSearchMapFromSource()
	alg.clearHistory()
	return data3d, nil
}

// SetData3D loads 3D map cell data.
// SetData3D 设置三维地图数据。
func (alg *AStarAlg) SetData3D(data [][][]int) (sourceData [][][]int, err error) {
	if nil == data || len(data) == 0 {
		return nil, errors.New("Data Error. ")
	}
	var data3d [][][]int
	var ln int
	if data3d, ln = alg.copyData3D(data); ln != alg.size.Area() {
		return nil, errors.New("Data Error. ")
	}
	alg.sourceMap = data3d
	alg.updateSearchMapFromSource()
	alg.clearHistory()
	return data3d, nil
}

// Search2D searches a path on a 2D grid.
// Search2D 在二维网格上寻路。
// sx is StartX; sy is StartY; ex is EndX; ey is EndY.
// sx 为起点 X；sy 为起点 Y；ex 为终点 X；ey 为终点 Y。
func (alg *AStarAlg) Search2D(sx, sy, ex, ey int) (path []Position, ok bool) {
	startPos, endPos := Position{sx, sy, 0}, Position{ex, ey, 0}
	return alg.SearchPosition(startPos, endPos)
}

// Search searches a path on a 3D grid.
// Search 在三维网格上寻路。
// sx is StartX; sy is StartY; sz is StartZ; ex is EndX; ey is EndY; ez is EndZ.
// sx 为起点 X；sy 为起点 Y；sz 为起点 Z；ex 为终点 X；ey 为终点 Y；ez 为终点 Z。
func (alg *AStarAlg) Search(sx, sy, sz, ex, ey, ez int) (path []Position, ok bool) {
	startPos, endPos := Position{sx, sy, sz}, Position{ex, ey, ez}
	return alg.SearchPosition(startPos, endPos)
}

// SearchPosition searches a path between two positions.
// SearchPosition 在两点之间寻路。
func (alg *AStarAlg) SearchPosition(startPos, endPos Position) (path []Position, ok bool) {
	if ok, s, e := alg.history.checkSubPath(startPos, endPos); ok {
		return alg.history.getSubPath(s, e), ok
	}
	alg.startPos, alg.endPos = startPos, endPos
	alg.initPath()
	if alg.searchPath() {
		path := alg.genPath()
		alg.history.startPos, alg.history.endPos = startPos, endPos
		alg.history.path, alg.history.pathOk = path, true
		return path, true
	} else {
		return nil, false
	}
}

//-------------------------------------------------------

func (alg *AStarAlg) initPath() {
	// 通过幅度提高Mask值，达到不用重置检索地图的效果
	alg.baseMask += DEFAULT_MASK_ADD
	// 当前可使用的最高Mask值
	alg.maxMask = alg.baseMask + DEFAULT_MASK_ADD
	// 清空open队列
	alg.openQueue.Clear()
	if alg.baseMask > MAX_MASK {
		alg.baseMask = DEFAULT_MASK_ADD
		alg.maxMask = alg.baseMask + DEFAULT_MASK_ADD
		alg.updateSearchMapFromSource()
	}
}

// 寻路
// 执行过程修改检索地图数据
// 通过数据值来标记地图路径
// 从开始点向目标点检索
func (alg *AStarAlg) searchPath() bool {
	if alg.isOutOfMap(alg.startPos) || alg.isOutOfMap(alg.endPos) {
		return false
	}
	// 以开始点为初始点搜索
	start, end, openQueue := alg.startPos, alg.endPos, alg.openQueue

	alg.searchMap[start.Z][start.Y][start.X] = alg.baseMask
	openQueue.Push3D(start.X, start.Y, start.Z, alg.baseMask)

	var gn int //当前点到下一点的实际代价
	var hn int //下一点到终点的估算代价
	var fn int //总的启发函数：fn = gn + hn + nextValue
	for openQueue.Len() > 0 {
		info, _ := openQueue.Shift()
		pos := info.Position
		for _, n := range alg.nextDirs {
			nextVector := GetDirVector(n)
			nextPos := pos.AddVector(nextVector)
			// 如果下一个超出地图范围
			if alg.isOutOfMap(nextPos) {
				continue
			}
			adjX, adjY, adjZ := nextPos.X, nextPos.Y, nextPos.Z
			// 如果下一个是 不可走区域 或 内部地图边界
			if alg.searchMap[adjZ][adjY][adjX] == GridObstacle || alg.searchMap[adjZ][adjY][adjX] == GridOut {
				continue
			}
			gn = alg.searchMap[pos.Z][pos.Y][pos.X] + nextVector.V() //当前代价 + 方向代价 + 地形代表
			//fmt.Println("AStarAlg.searchPath:下一个路径点：", adjX, adjY, gn, nextDir)
			// 已走过的区域
			if alg.searchMap[adjZ][adjY][adjX] >= alg.baseMask {
				if gn < alg.searchMap[adjZ][adjY][adjX] { //实际代价更小，路径更优
					alg.searchMap[adjZ][adjY][adjX] = gn
				}
				continue
			}
			if gn >= alg.maxMask {
				panic("Mask Error: gn >= maxMask")
			}
			alg.searchMap[adjZ][adjY][adjX] = gn
			if adjX == end.X && adjY == end.Y { //找到目标点
				//fmt.Println("AStarAlg.searchPath:检索结束！成功", adjX, adjY, gn)
				return true
			}
			if nil != alg.customFuncHn {
				hn = alg.customFuncHn(adjX, adjY, adjZ, end.X, end.Y, end.Z)
			} else {
				hn = DefaultFuncHn(adjX, adjY, adjZ, end.X, end.Y, end.Z)
			}
			fn = gn + hn + alg.sourceMap[adjZ][adjY][adjX] //fn只参与排序
			openQueue.Push3D(adjX, adjY, adjZ, fn)
			//fmt.Println("AStarAlg.searchPath:\t加入路径点：", adjX, adjY, fn)
		}
		//fmt.Println("AStarAlg.searchPath:余下长度：", openQueue.Len())
	}
	//fmt.Println("AStarAlg.searchPath:检索结束！失败")
	return false
}

// 生成搜索成功的路径
// 从目标点向开始点以取小值方式取出路径
func (alg *AStarAlg) genPath() []Position {
	walk := []Position{alg.endPos}
	nextPos := alg.endPos
	minGnVal := alg.searchMap[nextPos.Z][nextPos.Y][nextPos.X]
	for !nextPos.EqualTo(alg.startPos) {
		//fmt.Println("for 0", pos, minGnVal)
		var nextDirPos Position
		for _, n := range alg.nextDirs {
			nextDirVector := GetDirVector(n)
			adjPos := nextPos.AddVector(nextDirVector)
			// 如果下一个超出地图范围
			if alg.isOutOfMap(adjPos) {
				continue
			}
			//到达目标点
			if alg.startPos.EqualTo(adjPos) {
				nextDirPos = adjPos
				break
			}
			adjX, adjY, adjZ := adjPos.X, adjPos.Y, adjPos.Z
			//fmt.Println("for 1", nextPos, nextDir, alg.searchMap[adjY][adjX], alg.baseMask)
			// 如果不是检索过的地方
			if alg.searchMap[adjZ][adjY][adjX] < alg.baseMask {
				continue
			}
			//fmt.Println("比较：", alg.searchMap[adjY][adjX], minGnVal)
			if alg.searchMap[adjZ][adjY][adjX] < minGnVal {
				minGnVal = alg.searchMap[adjZ][adjY][adjX]
				nextDirPos = adjPos
			}
		}
		nextPos = nextDirPos
		walk = append(walk, nextDirPos)
	}
	for sIndex, eIndex := 0, len(walk)-1; sIndex < eIndex; sIndex, eIndex = sIndex+1, eIndex-1 { //反序
		walk[sIndex], walk[eIndex] = walk[eIndex], walk[sIndex]
	}
	//fmt.Println("AStarAlg.genPath:路径生成！", walk)
	return walk
}

func (alg *AStarAlg) isOutOfMap(pos Position) bool {
	return pos.X < 0 || pos.Y < 0 || pos.Z < 0 || pos.X >= alg.size.Width || pos.Y >= alg.size.Height || pos.Z >= alg.size.Depth
}

func (alg *AStarAlg) isOutOfMap2(x, y, z int) bool {
	return x < 0 || y < 0 || z < 0 || x >= alg.size.Width || y >= alg.size.Height || z >= alg.size.Depth
}

// 重置检索地图
func (alg *AStarAlg) updateSearchMapFromSource() {
	alg.searchMap, _ = alg.copyData3D(alg.sourceMap)
}

// 清除历史记录
func (alg *AStarAlg) clearHistory() {
	alg.history = history{}
}

//----------------------------------------

func (alg *AStarAlg) containsDirection(directions []Direction, dir Direction) bool {
	for _, d := range directions {
		if d == dir {
			return true
		}
	}
	return false
}

func (alg *AStarAlg) copyData(data []int) (data3D [][][]int, len int) {
	size := *alg.size
	cp := make([]int, size.Area())
	ln := copy(cp, data)
	var rsz [][][]int
	for z := 0; z < size.Depth; z++ {
		var rsy [][]int
		startIndexZ := z * size.Depth
		for y := 0; y < alg.size.Height; y++ {
			startIndexY := startIndexZ + y*size.Width
			slice := cp[startIndexY : startIndexY+alg.size.Width]
			rsy = append(rsy, slice)
		}
		rsz = append(rsz, rsy)
	}
	return rsz, ln
}

func (alg *AStarAlg) copyData2D(data [][]int) (data3D [][][]int, len int) {
	return alg.copyData3D([][][]int{data})
}

func (alg *AStarAlg) copyData3D(data [][][]int) (data3D [][][]int, len int) {
	size := *alg.size
	cp := make([]int, size.Area())
	var rsz [][][]int
	var ln int
	for z := 0; z < size.Depth; z++ {
		var rsy [][]int
		startIndexZ := z * size.Depth
		for y := 0; y < size.Height; y++ {
			startIndexY := startIndexZ + y*size.Width
			slice := cp[startIndexY : startIndexY+size.Width]
			l := copy(slice, data[z][y])
			rsy = append(rsy, slice)
			ln = ln + l
		}
		rsz = append(rsz, rsy)
	}
	return rsz, ln
}
