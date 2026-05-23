// Package astar provides A* pathfinding on 2D and 3D grid maps with optional path smoothing.
// Package astar 在二维与三维网格地图上提供 A* 寻路，并支持路径平滑与并发安全封装。
//
// Created by xuzhuoxi on 2019-04-03.
// @author xuzhuoxi
//
// Core pieces:
//   - IAStarAlg / AStarAlg: configurable heuristic (FuncHn), allowed directions, and mask-based search state reuse
//   - IGridMap / GridMap: grid data loading, line-of-sight checks, path search with optional inflection removal
//   - GridMapSync: RWMutex-wrapped GridMap for concurrent use
//   - Position, Size, Direction, DirectionVector: grid coordinates, extents, and neighbor step definitions
//   - GridOut, GridPath, GridObstacle: cell walkability markers
//   - ClearInflection, IsInStandardLine, GetDirection: geometry helpers for path post-processing
//
// 主要职责：
//   - IAStarAlg / AStarAlg：可配置启发函数、可行方向及基于掩码的检索状态复用
//   - IGridMap / GridMap：地图数据装载、直线可达判断、可选拐点清除的寻路
//   - GridMapSync：带读写锁的并发安全 GridMap 封装
//   - Position、Size、Direction、DirectionVector：坐标、范围与邻域方向矢量
//   - GridOut、GridPath、GridObstacle：格子可走性标记
//   - ClearInflection、IsInStandardLine、GetDirection：路径后处理相关的几何工具
package astar
