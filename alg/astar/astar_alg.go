//
//Created by xuzhuoxi
//on 2019-04-03.
//@author xuzhuoxi
//
package astar

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/mathx"
)

const (
	DEFAULT_MASK_ADD = 1048576
	//DEFAULT_MASK_ADD = 10
	MAX_MASK = 2000000000
)

// 估值函数Hn
// 下一点到终点的估算代价
type FuncHn func(cx, cy, ex, ey int) int

//默认估值函数
func DefaultFuncHn(cx, cy, ex, ey int) int {
	return mathx.AbsInt(cx-ex) + mathx.AbsInt(cy-ey)
}

type IAStarAlg interface {
	//设置禁止检索的方向(可用于特殊地图)
	SetAllowedDirections(direction []Direction)
	//设置自定义的估值函数
	SetCustomFuncHn(hn FuncHn)

	//初始化地图Size
	InitMapSize(width, height int)
	//设置地图数据
	InitData(data []int) error
	//设置地图数据
	InitData2(data [][]int) error
	// 寻路
	// sx:StartX; sy:StartY
	// ex:EndX; ey:EndY
	Search(sx, sy, ex, ey int) (path []Position, ok bool)
}

func NewIAStartAlg() IAStarAlg {
	return NewAStartAlg()
}

func NewAStartAlg() *AStarAlg {
	return &AStarAlg{openQueue: newSortedQueue()}
}

type history struct {
	startPos Position
	endPos   Position
	path     []Position
	pathOk   bool
}

func (h history) getPath() []Position {
	if nil == h.path {
		return nil
	}
	if 0 == len(h.path) {
		return []Position{}
	}
	rs := make([]Position, len(h.path))
	copy(rs, h.path)
	return rs
}

type AStarAlg struct {
	//地图的大小
	size *Size
	//可行方向
	nextDirs []Direction
	//自定义估值函数
	customFuncHn FuncHn
	//初始地图数据
	sourceMap [][]int
	//检索地图数据
	searchMap [][]int

	baseMask  int
	startPos  *Position
	endPos    *Position
	openQueue *sortedQueue
	history   history
}

func (alg *AStarAlg) SetAllowedDirections(direction []Direction) {
	cp := make([]Direction, len(direction))
	copy(cp, direction)
	alg.nextDirs = cp
}

func (alg *AStarAlg) SetCustomFuncHn(hn FuncHn) {
	alg.customFuncHn = hn
}

func (alg *AStarAlg) InitMapSize(width, height int) {
	alg.size = &Size{Width: width, Height: height}
}

func (alg *AStarAlg) InitData(data []int) error {
	if len(data) != alg.size.Area() {
		return errors.New("Data Error. ")
	}
	alg.sourceMap = alg.copyData(data)
	alg.resetMapSearch()
	return nil
}

func (alg *AStarAlg) InitData2(data [][]int) error {
	if nil == data || len(data) == 0 {
		return errors.New("Data Error. ")
	}
	ln := 0
	for _, ds := range data {
		ln += len(ds)
	}
	if ln != alg.size.Area() {
		return errors.New("Data Error. ")
	}
	alg.sourceMap = alg.copyData2(data)
	alg.resetMapSearch()
	return nil
}

func (alg *AStarAlg) Search(sx, sy, ex, ey int) (path []Position, ok bool) {
	startPos := Position{sx, sy}
	endPos := Position{ex, ey}
	if alg.history.startPos.EqualTo(startPos) && alg.history.endPos.EqualTo(endPos) { //相同的检索
		return alg.history.getPath(), alg.history.pathOk
	}
	alg.startPos = &startPos
	alg.endPos = &endPos
	history := alg.history
	history.startPos = startPos
	history.endPos = endPos
	alg.initPath()
	if alg.searchPath() {
		path := alg.genPath()
		history.path, history.pathOk = path, true
		return path, true
	} else {
		history.path, history.pathOk = nil, false
		return nil, false
	}
}

//-------------------------------------------------------

func (alg *AStarAlg) initPath() {
	// 通过幅度提高Mask值，达到不用重置检索地图的效果
	alg.baseMask += DEFAULT_MASK_ADD
	// 清空open队列
	alg.openQueue.Clear()
	if alg.baseMask > MAX_MASK {
		alg.baseMask = DEFAULT_MASK_ADD
		alg.resetMapSearch()
	}
}

// 寻路
// 执行过程修改检索地图数据
// 通过数据值来标记地图路径
// 从开始点向目标点检索
func (alg *AStarAlg) searchPath() bool {
	if alg.outMap(alg.startPos.X, alg.startPos.Y) || alg.outMap(alg.endPos.X, alg.endPos.Y) {
		return false
	}
	// 以开始点为初始点搜索
	start := alg.startPos
	end := alg.endPos
	openQueue := alg.openQueue

	alg.searchMap[start.Y][start.X] = alg.baseMask
	openQueue.Push(start.X, start.Y, alg.baseMask)

	var gn int //当前点到下一点的实际代价
	var hn int //下一点到终点的估算代价
	var fn int //总的启发函数：fn = gn + hn + nextValue
	for openQueue.Len() > 0 {
		info, _ := openQueue.Shift()
		pos := info.Position
		for _, n := range alg.nextDirs {
			nextDir := GetNextDir(n)
			adjX := pos.X + int(nextDir.X())
			adjY := pos.Y + int(nextDir.Y())
			// 如果下一个超出地图范围
			if alg.outMap(adjX, adjY) {
				continue
			}
			// 如果下一个是 不可走区域 或 内部地图边界
			if alg.searchMap[adjY][adjX] == GridObstacle || alg.searchMap[adjY][adjX] == GridOut {
				continue
			}
			gn = alg.searchMap[pos.Y][pos.X] + nextDir.V() //当前代价 + 方向代价 + 地形代表
			//fmt.Println("AStarAlg.searchPath:下一个路径点：", adjX, adjY, gn, nextDir)
			// 已走过的区域
			if alg.searchMap[adjY][adjX] >= alg.baseMask {
				if gn < alg.searchMap[adjY][adjX] { //实际代价更小，路径更优
					alg.searchMap[adjY][adjX] = gn
				}
				continue
			}
			alg.searchMap[adjY][adjX] = gn
			if adjX == end.X && adjY == end.Y { //找到目标点
				//fmt.Println("AStarAlg.searchPath:检索结束！成功", adjX, adjY, gn)
				return true
			}
			if nil != alg.customFuncHn {
				hn = alg.customFuncHn(adjX, adjY, end.X, end.Y)
			} else {
				hn = DefaultFuncHn(adjX, adjY, end.X, end.Y)
			}
			fn = gn + hn + alg.sourceMap[adjY][adjX] //fn只参与排序
			openQueue.Push(adjX, adjY, fn)
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
	walk := []Position{*alg.endPos}
	px := alg.endPos.X
	py := alg.endPos.Y
	minGnVal := alg.searchMap[py][px]
	for px != alg.startPos.X || py != alg.startPos.Y {
		//fmt.Println("for 0", px, py, minGnVal)
		nx := -1
		ny := -1
		for _, n := range alg.nextDirs {
			nextDir := GetNextDir(n)
			adjX := px + nextDir.X()
			adjY := py + nextDir.Y()
			//fmt.Println("for 1", adjX, adjY, nextDir, alg.searchMap[adjY][adjX], alg.baseMask)
			// 如果下一个超出地图范围
			if alg.outMap(adjX, adjY) {
				continue
			}
			//到达目标点
			if adjX == alg.startPos.X && adjY == alg.startPos.Y {
				nx = adjX
				ny = adjY
				break
			}
			// 如果不是检索过的地方
			if alg.searchMap[adjY][adjX] < alg.baseMask {
				continue
			}
			//fmt.Println("比较：", alg.searchMap[adjY][adjX], minGnVal)
			if alg.searchMap[adjY][adjX] < minGnVal {
				minGnVal = alg.searchMap[adjY][adjX]
				nx = adjX
				ny = adjY
			}
		}
		if nx < 0 {
			break
		}
		px = nx
		py = ny
		walk = append(walk, Position{nx, ny})
	}
	for sIndex, eIndex := 0, len(walk)-1; sIndex < eIndex; sIndex, eIndex = sIndex+1, eIndex-1 { //反序
		walk[sIndex], walk[eIndex] = walk[eIndex], walk[sIndex]
	}
	//fmt.Println("AStarAlg.genPath:路径生成！", walk)
	return walk
}

func (alg *AStarAlg) outMap(x, y int) bool {
	return x < 0 || y < 0 || x >= alg.size.Width || y >= alg.size.Height
}

//重置检索地图
func (alg *AStarAlg) resetMapSearch() {
	alg.searchMap = alg.copyData2(alg.sourceMap)
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

func (alg *AStarAlg) copyData(data []int) [][]int {
	cp := make([]int, alg.size.Area())
	copy(cp, data)
	var rs [][]int
	for y := 0; y < alg.size.Height; y++ {
		startIndex := y * alg.size.Width
		slice := cp[startIndex : startIndex+alg.size.Width]
		rs = append(rs, slice)
	}
	return rs
}

func (alg *AStarAlg) copyData2(data [][]int) [][]int {
	cp := make([]int, alg.size.Area())
	var rs [][]int
	for y := 0; y < alg.size.Height; y++ {
		startIndex := y * alg.size.Width
		slice := cp[startIndex : startIndex+alg.size.Width]
		copy(slice, data[y])
		rs = append(rs, slice)
	}
	return rs
}
