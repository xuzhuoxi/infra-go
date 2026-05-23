package astar

// IGridMap defines grid-map pathfinding with data and pixel sizing.
// IGridMap 定义带数据格与像素尺寸换算的网格地图寻路接口。
type IGridMap interface {
	// InitGridMap initializes logical data size and display grid size.
	// InitGridMap 初始化地图数据尺寸与显示格尺寸。
	InitGridMap(dataSize, gridSize Size)
	// GetGridSize returns the display grid cell size.
	// GetGridSize 返回显示格尺寸。
	GetGridSize() Size
	// GetDataSize returns the logical map data size.
	// GetDataSize 返回地图数据尺寸。
	GetDataSize() Size
	// GetPixelSize returns the map size in pixels.
	// GetPixelSize 返回地图像素尺寸。
	GetPixelSize() Size
	// GetAStartAlg returns the underlying A* algorithm instance.
	// GetAStartAlg 返回底层 A* 算法实例。
	GetAStartAlg() IAStarAlg

	// GetDataValue returns the cell value at pos.
	// GetDataValue 返回指定格子的数据值。
	GetDataValue(pos Position) int
	// CheckPath reports whether every cell on path is walkable.
	// CheckPath 判断路径上各点是否均可通行。
	CheckPath(path []Position) bool
	// CanLineTo reports whether startPos and endPos are connected by a straight walkable line.
	// CanLineTo 判断两点是否可直线直达。
	CanLineTo(startPos, endPos Position) bool

	// SetAllowedDirections sets allowed pathfinding directions.
	// SetAllowedDirections 设置允许寻路的方向。
	SetAllowedDirections(direction ...Direction)
	// SetMapData loads map data ([]int, [][]int, or [][][]int).
	// SetMapData 设置地图数据。
	SetMapData(data interface{}) error
	// SetSearchHn sets a custom heuristic for path search.
	// SetSearchHn 设置自定义估值函数。
	SetSearchHn(hn FuncHn)
	// SearchPath finds a path and removes collinear inflection points by default.
	// SearchPath 检索路径，默认清除拐点。
	SearchPath(startPos, endPos Position) (path []Position, ok bool)
	// SearchPathWithInflection finds a path and keeps all inflection points.
	// SearchPathWithInflection 检索路径并保留拐点。
	SearchPathWithInflection(startPos, endPos Position) (path []Position, ok bool)
}

// NewIGridMap returns a new IGridMap implementation.
// NewIGridMap 创建 IGridMap 实现实例。
func NewIGridMap() IGridMap {
	return NewGridMap()
}

// NewGridMap constructs a new GridMap.
// NewGridMap 创建 GridMap 实例。
func NewGridMap() *GridMap {
	return &GridMap{}
}

// GridMap wraps grid data and A* search with line-of-sight optimization.
// GridMap 封装网格数据与带直线可达优化的 A* 寻路。
type GridMap struct {
	gridSize Size
	dataSize Size

	mapData  [][][]int
	aStarAlg IAStarAlg
}

//---------------

// InitGridMap initializes logical data size and display grid size.
// InitGridMap 初始化地图数据尺寸与显示格尺寸。
func (m *GridMap) InitGridMap(dataSize, gridSize Size) {
	if dataSize.Empty() || gridSize.Empty() {
		panic("GridMap Empty! ")
	}
	m.dataSize = dataSize
	m.gridSize = gridSize
	m.aStarAlg = NewAStartAlg()
	m.aStarAlg.InitMapSize(dataSize.Width, dataSize.Height, dataSize.Depth)
}

// GetGridSize returns the display grid cell size.
// GetGridSize 返回显示格尺寸。
func (m *GridMap) GetGridSize() Size {
	return m.gridSize
}

// GetDataSize returns the logical map data size.
// GetDataSize 返回地图数据尺寸。
func (m *GridMap) GetDataSize() Size {
	return m.dataSize
}

// GetPixelSize returns the map size in pixels.
// GetPixelSize 返回地图像素尺寸。
func (m *GridMap) GetPixelSize() Size {
	return Size{Width: m.gridSize.Width * m.dataSize.Width, Height: m.gridSize.Height * m.dataSize.Height}
}

// GetAStartAlg returns the underlying A* algorithm instance.
// GetAStartAlg 返回底层 A* 算法实例。
func (m *GridMap) GetAStartAlg() IAStarAlg {
	return m.aStarAlg
}

//---------------

// GetDataValue returns the cell value at pos.
// GetDataValue 返回指定格子的数据值。
func (m *GridMap) GetDataValue(pos Position) int {
	return m.getDataValue(pos)
}

// CheckPath reports whether every cell on path is walkable.
// CheckPath 判断路径上各点是否均可通行。
func (m *GridMap) CheckPath(path []Position) bool {
	if len(path) == 0 {
		return false
	}
	for _, pos := range path {
		if !m.isPathGrid(pos) {
			return false
		}
	}
	return true
}

// CanLineTo reports whether startPos and endPos are connected by a straight walkable line.
// CanLineTo 判断两点是否可直线直达。
func (m *GridMap) CanLineTo(startPos, endPos Position) bool {
	return m.canLineTo(startPos, endPos)
}

//---------------

// SetAllowedDirections sets allowed pathfinding directions.
// SetAllowedDirections 设置允许寻路的方向。
func (m *GridMap) SetAllowedDirections(direction ...Direction) {
	m.aStarAlg.SetAllowedDirections(direction)
}

// SetMapData loads map data ([]int, [][]int, or [][][]int).
// SetMapData 设置地图数据。
func (m *GridMap) SetMapData(data interface{}) error {
	var source [][][]int
	var err error
	switch d := data.(type) {
	case []int:
		if source, err = m.aStarAlg.SetData(d); nil != err {
			return err
		}
		m.mapData = source
	case [][]int:
		if source, err = m.aStarAlg.SetData2D(d); nil != err {
			return err
		}
		m.mapData = source
	case [][][]int:
		if source, err = m.aStarAlg.SetData3D(d); nil != err {
			return err
		}
		m.mapData = source
	}
	return nil
}

// SetSearchHn sets a custom heuristic for path search.
// SetSearchHn 设置自定义估值函数。
func (m *GridMap) SetSearchHn(hn FuncHn) {
	m.aStarAlg.SetCustomFuncHn(hn)
}

// SearchPath finds a path and removes collinear inflection points by default.
// SearchPath 检索路径，默认清除拐点。
func (m *GridMap) SearchPath(startPos, endPos Position) (path []Position, ok bool) {
	return m.searchPath(startPos, endPos, false)
}

// SearchPathWithInflection finds a path and keeps all inflection points.
// SearchPathWithInflection 检索路径并保留拐点。
func (m *GridMap) SearchPathWithInflection(startPos, endPos Position) (path []Position, ok bool) {
	return m.searchPath(startPos, endPos, true)
}

//--------------------------------------

func (m *GridMap) searchPath(startPos, endPos Position, keepInflection bool) (path []Position, ok bool) {
	if !m.isPathGrid(startPos) || !m.isPathGrid(endPos) { //位置不可行走
		return nil, false
	}
	if startPos.EqualTo(endPos) { //两点相同
		return nil, false
	}
	if m.canLineTo(startPos, endPos) {
		return []Position{startPos, endPos}, true
	} else {
		if path, ok := m.aStarAlg.SearchPosition(startPos, endPos); ok {
			if keepInflection {
				return path, ok
			} else {
				return ClearInflection(path), ok
			}
		}
		return nil, false
	}
}

// 是否可以直线行走
func (m *GridMap) canLineTo(startPos, endPos Position) bool {
	if !IsInStandardLine(startPos, endPos, true) {
		return false
	}
	funcGetAddDiff := func(start, end int) int {
		diff := end - start
		switch {
		case diff == 0:
			return 0
		case diff > 0:
			return 1
		default:
			return -1
		}
	}
	addX := funcGetAddDiff(startPos.X, endPos.X)
	addY := funcGetAddDiff(startPos.Y, endPos.Y)
	addZ := funcGetAddDiff(startPos.Z, endPos.Z)
	temp := startPos
	for !temp.EqualTo(endPos) {
		if !m.isPathGrid(temp) {
			return false
		}
		(&temp).X += addX
		(&temp).Y += addY
		(&temp).Z += addZ
	}
	return true
}

// 格子是否为可走
func (m *GridMap) isPathGrid(pos Position) bool {
	val := m.getDataValue(pos)
	return val != GridOut && val != GridObstacle
}

// 取格式数据值
func (m *GridMap) getDataValue(pos Position) int {
	x, y, z := pos.X, pos.Y, pos.Z
	if x < 0 || x >= m.dataSize.Width || y < 0 || y >= m.dataSize.Height || z < 0 || z >= m.dataSize.Height {
		return GridOut
	}
	return m.mapData[z][y][x]
}
