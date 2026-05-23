// Package filex provides cross-platform path utilities, directory walking, and file copy/move/remove helpers.
// Package filex 提供跨平台路径工具、目录遍历，以及文件复制、移动、删除等辅助能力。
//
// Created by xuzhuoxi on 2022-07-08.
// @author xuzhuoxi
//
// Capabilities:
//   - Normalize paths (Unix/Windows), detect absolute/relative format, combine/split names and extensions
//   - Existence and type checks (file, dir, empty dir); complete missing parent directories
//   - Walk directories (single level or recursive), collect paths with PathFilter
//   - Copy/move/remove files and trees; size helpers; WriteFile with optional parent-dir creation
//
// 主要职责：
//   - 路径标准化（Unix/Windows）、绝对/相对判定、合并/拆分与扩展名处理
//   - 存在性与类型检查（文件、目录、空目录）；补全缺失的父级目录
//   - 目录遍历（单层或递归），配合 PathFilter 收集路径
//   - 文件/目录树的复制、移动、删除与大小统计；WriteFile 及可选的父目录自动创建
package filex
