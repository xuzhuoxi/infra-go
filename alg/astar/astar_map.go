//
//Created by xuzhuoxi
//on 2019-04-05.
//@author xuzhuoxi
//
package astar

import (
	"sync"
)

type IGridMap interface {
	//初始化地图大小
	InitGridMap(dataSize, gridSize Size)
	//格式大小
	GetGridSize() Size
	//地图数据大小
	GetDataSize() Size
	//地图像素大小
	GetPixelSize() Size
	//AStart算法
	GetAStartAlg() IAStarAlg

	//设置允许寻路的方向
	SetAllowedDirections(direction ...Direction)
	//设置地图数据
	SetMapData(data interface{}) error
	//格式数据值
	GetDataValue(pos Position) int
	//判断路径是否通路
	CheckPath(path []Position) bool
	//判断是否两点直通
	CanLineTo(startPos, endPos Position) bool
	//设置自定义估值函数
	SetSearchHn(hn FuncHn)
	//检索路径
	SearchPath(startPos, endPos Position) (path []Position, ok bool)
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
	oblique  bool
	aStarAlg IAStarAlg

	rwMu sync.RWMutex
}

func (m *GridMap) GetGridSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gridSize
}

func (m *GridMap) GetDataSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.dataSize
}

func (m *GridMap) GetAStartAlg() IAStarAlg {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.aStarAlg
}

func (m *GridMap) GetPixelSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return Size{Width: m.gridSize.Width * m.dataSize.Width, Height: m.gridSize.Height * m.dataSize.Height}
}

func (m *GridMap) SetAllowedDirections(direction ...Direction) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.aStarAlg.SetAllowedDirections(direction)
}

func (m *GridMap) InitGridMap(dataSize, gridSize Size) {
	if dataSize.Empty() || gridSize.Empty() {
		panic("GridMap Empty! ")
	}
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.dataSize = dataSize
	m.gridSize = gridSize
	m.aStarAlg = NewAStartAlg()
	m.aStarAlg.InitMapSize(dataSize.Width, dataSize.Height, dataSize.Depth)
}

func (m *GridMap) SetMapData(data interface{}) error {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
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
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.aStarAlg.SetCustomFuncHn(hn)
}

func (m *GridMap) GetDataValue(pos Position) int {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.getDataValue(pos)
}

func (m *GridMap) CheckPath(path []Position) bool {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
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

//判断是否两点直通
func (m *GridMap) CanLineTo(startPos, endPos Position) bool {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.canLineTo(startPos, endPos)
}

func (m *GridMap) SearchPath(startPos, endPos Position) (path []Position, ok bool) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
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
			path = ClearInflection(path)
			return path, ok
		}
		return nil, false
	}
}

//--------------------------------------

//是否可以直线行走
func (m *GridMap) canLineTo(startPos, endPos Position) bool {
	if !IsInStandardLine(startPos, endPos, m.oblique) {
		return false
	}
	funcGetAdd := func(start, end int) int {
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
	addX := funcGetAdd(startPos.X, endPos.X)
	addY := funcGetAdd(startPos.Y, endPos.Y)
	temp := startPos
	for !temp.EqualTo(endPos) {
		if !m.isPathGrid(temp) {
			return false
		}
		(&temp).X += addX
		(&temp).Y += addY
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
