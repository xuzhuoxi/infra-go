// Package astar
// Created by xuzhuoxi
// on 2019-04-05.
// @author xuzhuoxi
//
package astar

type IGridMap interface {
	// InitGridMap 初始化地图大小
	InitGridMap(dataSize, gridSize Size)
	// GetGridSize 格式大小
	GetGridSize() Size
	// GetDataSize 地图数据大小
	GetDataSize() Size
	// GetPixelSize 地图像素大小
	GetPixelSize() Size
	// GetAStartAlg AStart算法
	GetAStartAlg() IAStarAlg

	// GetDataValue 格式数据值
	GetDataValue(pos Position) int
	// CheckPath 判断路径是否通路
	CheckPath(path []Position) bool
	// CanLineTo 判断是否两点直通
	CanLineTo(startPos, endPos Position) bool

	// SetAllowedDirections 设置允许寻路的方向
	SetAllowedDirections(direction ...Direction)
	// SetMapData 设置地图数据
	SetMapData(data interface{}) error
	// SetSearchHn 设置自定义估值函数
	SetSearchHn(hn FuncHn)
	// SearchPath
	// 检索路径
	// 默认清除拐点
	SearchPath(startPos, endPos Position) (path []Position, ok bool)
	// SearchPathWithInflection
	// 检索路径
	// 保留拐点
	SearchPathWithInflection(startPos, endPos Position) (path []Position, ok bool)
}

func NewIGridMap() IGridMap {
	return NewGridMap()
}

func NewGridMap() *GridMap {
	return &GridMap{}
}

type GridMap struct {
	gridSize Size
	dataSize Size

	mapData  [][][]int
	aStarAlg IAStarAlg
}

//---------------

func (m *GridMap) InitGridMap(dataSize, gridSize Size) {
	if dataSize.Empty() || gridSize.Empty() {
		panic("GridMap Empty! ")
	}
	m.dataSize = dataSize
	m.gridSize = gridSize
	m.aStarAlg = NewAStartAlg()
	m.aStarAlg.InitMapSize(dataSize.Width, dataSize.Height, dataSize.Depth)
}

func (m *GridMap) GetGridSize() Size {
	return m.gridSize
}

func (m *GridMap) GetDataSize() Size {
	return m.dataSize
}

func (m *GridMap) GetPixelSize() Size {
	return Size{Width: m.gridSize.Width * m.dataSize.Width, Height: m.gridSize.Height * m.dataSize.Height}
}

func (m *GridMap) GetAStartAlg() IAStarAlg {
	return m.aStarAlg
}

//---------------

func (m *GridMap) GetDataValue(pos Position) int {
	return m.getDataValue(pos)
}

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

// CanLineTo 判断是否两点直通
func (m *GridMap) CanLineTo(startPos, endPos Position) bool {
	return m.canLineTo(startPos, endPos)
}

//---------------

func (m *GridMap) SetAllowedDirections(direction ...Direction) {
	m.aStarAlg.SetAllowedDirections(direction)
}

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

func (m *GridMap) SetSearchHn(hn FuncHn) {
	m.aStarAlg.SetCustomFuncHn(hn)
}

func (m *GridMap) SearchPath(startPos, endPos Position) (path []Position, ok bool) {
	return m.searchPath(startPos, endPos, false)
}

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

//是否可以直线行走
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
