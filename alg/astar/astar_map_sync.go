package astar

import "sync"

func NewIGridMapSync() IGridMap {
	return NewGridMapSync()
}

func NewGridMapSync() *GridMapSync {
	return &GridMapSync{}
}

type GridMapSync struct {
	gMap GridMap
	rwMu sync.RWMutex
}

func (m *GridMapSync) InitGridMap(dataSize, gridSize Size) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.gMap.InitGridMap(dataSize, gridSize)
}

func (m *GridMapSync) GetGridSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetGridSize()
}

func (m *GridMapSync) GetDataSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetDataSize()
}

func (m *GridMapSync) GetPixelSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetPixelSize()
}

func (m *GridMapSync) GetAStartAlg() IAStarAlg {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetAStartAlg()
}

func (m *GridMapSync) GetDataValue(pos Position) int {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetDataValue(pos)
}

func (m *GridMapSync) CheckPath(path []Position) bool {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.CheckPath(path)
}

func (m *GridMapSync) CanLineTo(startPos, endPos Position) bool {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.CanLineTo(startPos, endPos)
}

func (m *GridMapSync) SetAllowedDirections(direction ...Direction) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.gMap.SetAllowedDirections(direction...)
}

func (m *GridMapSync) SetMapData(data interface{}) error {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	return m.gMap.SetMapData(data)
}

func (m *GridMapSync) SetSearchHn(hn FuncHn) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.gMap.SetSearchHn(hn)
}

func (m *GridMapSync) SearchPath(startPos, endPos Position) (path []Position, ok bool) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	return m.gMap.SearchPath(startPos, endPos)
}

func (m *GridMapSync) SearchPathWithInflection(startPos, endPos Position) (path []Position, ok bool) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	return m.gMap.SearchPathWithInflection(startPos, endPos)
}
