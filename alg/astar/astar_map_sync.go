package astar

import "sync"

// NewIGridMapSync returns a new thread-safe IGridMap implementation.
// NewIGridMapSync 创建线程安全的 IGridMap 实现实例。
func NewIGridMapSync() IGridMap {
	return NewGridMapSync()
}

// NewGridMapSync constructs a new GridMapSync.
// NewGridMapSync 创建 GridMapSync 实例。
func NewGridMapSync() *GridMapSync {
	return &GridMapSync{}
}

// GridMapSync is a mutex-protected GridMap safe for concurrent use.
// GridMapSync 为带读写锁的 GridMap，适用于并发访问。
type GridMapSync struct {
	gMap GridMap
	rwMu sync.RWMutex
}

// InitGridMap initializes logical data size and display grid size.
// InitGridMap 初始化地图数据尺寸与显示格尺寸。
func (m *GridMapSync) InitGridMap(dataSize, gridSize Size) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.gMap.InitGridMap(dataSize, gridSize)
}

// GetGridSize returns the display grid cell size.
// GetGridSize 返回显示格尺寸。
func (m *GridMapSync) GetGridSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetGridSize()
}

// GetDataSize returns the logical map data size.
// GetDataSize 返回地图数据尺寸。
func (m *GridMapSync) GetDataSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetDataSize()
}

// GetPixelSize returns the map size in pixels.
// GetPixelSize 返回地图像素尺寸。
func (m *GridMapSync) GetPixelSize() Size {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetPixelSize()
}

// GetAStartAlg returns the underlying A* algorithm instance.
// GetAStartAlg 返回底层 A* 算法实例。
func (m *GridMapSync) GetAStartAlg() IAStarAlg {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetAStartAlg()
}

// GetDataValue returns the cell value at pos.
// GetDataValue 返回指定格子的数据值。
func (m *GridMapSync) GetDataValue(pos Position) int {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.GetDataValue(pos)
}

// CheckPath reports whether every cell on path is walkable.
// CheckPath 判断路径上各点是否均可通行。
func (m *GridMapSync) CheckPath(path []Position) bool {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.CheckPath(path)
}

// CanLineTo reports whether startPos and endPos are connected by a straight walkable line.
// CanLineTo 判断两点是否可直线直达。
func (m *GridMapSync) CanLineTo(startPos, endPos Position) bool {
	m.rwMu.RLock()
	defer m.rwMu.RUnlock()
	return m.gMap.CanLineTo(startPos, endPos)
}

// SetAllowedDirections sets allowed pathfinding directions.
// SetAllowedDirections 设置允许寻路的方向。
func (m *GridMapSync) SetAllowedDirections(direction ...Direction) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.gMap.SetAllowedDirections(direction...)
}

// SetMapData loads map data ([]int, [][]int, or [][][]int).
// SetMapData 设置地图数据。
func (m *GridMapSync) SetMapData(data interface{}) error {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	return m.gMap.SetMapData(data)
}

// SetSearchHn sets a custom heuristic for path search.
// SetSearchHn 设置自定义估值函数。
func (m *GridMapSync) SetSearchHn(hn FuncHn) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	m.gMap.SetSearchHn(hn)
}

// SearchPath finds a path and removes collinear inflection points by default.
// SearchPath 检索路径，默认清除拐点。
func (m *GridMapSync) SearchPath(startPos, endPos Position) (path []Position, ok bool) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	return m.gMap.SearchPath(startPos, endPos)
}

// SearchPathWithInflection finds a path and keeps all inflection points.
// SearchPathWithInflection 检索路径并保留拐点。
func (m *GridMapSync) SearchPathWithInflection(startPos, endPos Position) (path []Position, ok bool) {
	m.rwMu.Lock()
	defer m.rwMu.Unlock()
	return m.gMap.SearchPathWithInflection(startPos, endPos)
}
