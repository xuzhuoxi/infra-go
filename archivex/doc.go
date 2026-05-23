// Package archivex provides helpers to build and append zip and tar archives from local paths.
// Package archivex 提供从本地路径创建与追加 zip、tar 归档的工具函数。
//
// Created by xuzhuoxi on 2022-07-12.
// @author xuzhuoxi
//
// Capabilities:
//   - Create archives from a single path or multiple paths (ArchiveToZip/Tar, ArchiveMultiToZip/Tar)
//   - Archive directory contents without wrapping the parent folder name (ArchiveChildrenToZip/Tar)
//   - Append files or directories to an existing zip.Writer or tar.Writer with customizable entry names
//
// 主要职责：
//   - 从单一路径或多个路径创建归档（ArchiveToZip/Tar、ArchiveMultiToZip/Tar）
//   - 仅打包目录内容、不包含父目录名（ArchiveChildrenToZip/Tar）
//   - 向已有 zip.Writer / tar.Writer 追加文件或目录，并支持自定义条目名称
package archivex
