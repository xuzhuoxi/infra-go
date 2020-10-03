package filex

import (
	"os"
	"path/filepath"
)

// 移动文件或目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 新路径无效：忽略，
// 新路径已存在：覆盖
func Move(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// 移动文件或目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 新路径无效：忽略，
// 新路径已存在：覆盖
func MoveAuto(oldPath string, newPath string, autoPerm os.FileMode) error {
	var dir, _ = filepath.Split(newPath)
	if !IsExist(dir) {
		os.MkdirAll(dir, autoPerm)
	}
	return Move(oldPath, newPath)
}

// 移动文件或目录到指定目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 目标路径无效：忽略
// 目标已存在：覆盖
func MoveTo(oldPath string, targetDir string) error {
	var _, name = filepath.Split(oldPath)
	var newPath = targetDir + PathSeparator + name
	return os.Rename(oldPath, newPath)
}

// 移动文件或目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 目标路径为文件：忽略
// 目标路径无效：自动补全
func MoveToAuto(oldPath string, targetDir string, autoPerm os.FileMode) error {
	if !IsExist(targetDir) {
		os.MkdirAll(targetDir, autoPerm)
	}
	return MoveTo(oldPath, targetDir)
}
