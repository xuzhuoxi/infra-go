// Package astar
// Created by xuzhuoxi
// on 2019-04-03.
// @author xuzhuoxi
//
package astar

import (
	"github.com/xuzhuoxi/infra-go/mathx"
)

// ClearInflection 清除拐点
func ClearInflection(path []Position) []Position {
	ln := len(path)
	if ln <= 2 {
		return path
	}
	for index := ln - 2; index >= 1; index-- {
		if IsInLine(path[index-1], path[index], path[index+1]) {
			path = append(path[:index], path[index+1:]...)
		}
	}
	return path
}

// IsInLine 判断三点是否一线
func IsInLine(first, second, third Position) bool {
	return (second.Y-first.Y)*(third.X-first.X) == (third.Y-first.Y)*(second.X-first.X)
}

// IsSamePosition 判断是为同一点
func IsSamePosition(pos1, pos2 Position) bool {
	return pos1.EqualTo(pos2)
}

// IsInStandardLine
// 是否标准线向
// includeOblique:包含斜向
func IsInStandardLine(pos1, pos2 Position, includeOblique bool) bool {
	if IsSamePosition(pos1, pos2) {
		return false
	}
	var xe = pos1.X == pos2.X
	var ye = pos1.Y == pos2.Y
	var ze = pos1.Z == pos2.Z

	if (xe != ye) != ze {
		return true
	}

	if !includeOblique {
		return false
	}

	var lx = mathx.AbsInt(pos1.X - pos2.X)
	var ly = mathx.AbsInt(pos1.Y - pos2.Y)
	var lz = mathx.AbsInt(pos1.Z - pos2.Z)

	return (lx == ly && lz == 0) || (lx == lz && ly == 0) || (ly == lz && lx == 0) || (lx == ly && ly == lz)
}

// GetDirection
// 判断方向,前提是两点为线向
// 采用笛卡尔坐标系
func GetDirection(sourcePos, targetPos Position) Direction {
	if sourcePos.EqualTo(targetPos) {
		return Center
	}
	if sourcePos.X == targetPos.X || sourcePos.Y == targetPos.Y {
		if sourcePos.X == sourcePos.X { //垂直
			if sourcePos.Y < targetPos.Y {
				return North
			} else {
				return South
			}
		} else {
			if sourcePos.X < targetPos.X {
				return East
			} else {
				return West
			}
		}
	} else {
		if sourcePos.X-targetPos.X == sourcePos.Y-targetPos.Y { //左下角 或 右上角
			if sourcePos.X < targetPos.X {
				return EastNorth
			} else {
				return WestSouth
			}
		} else {
			if sourcePos.Y < targetPos.Y {
				return WestNorth
			} else {
				return EastSouth
			}
		}
	}
}
