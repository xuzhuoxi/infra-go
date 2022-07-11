package filex

import (
	"os"
)

// 移动文件或目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 新路径无效：忽略
// 新路径已存在：覆盖
func Move(srcPath string, dstPath string) error {
	return move(FormatPath(srcPath), FormatPath(dstPath))
}

// 移动文件或目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 新路径无效：补全
// 新路径已存在：覆盖
func MoveAuto(srcPath string, dstPath string, autoPerm os.FileMode) error {
	return moveFileAuto(FormatPath(srcPath), FormatPath(dstPath), autoPerm)
}

// 移动文件或目录到指定目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 目标路径无效：忽略
// 目标已存在：覆盖
func MoveTo(srcPath string, dstDir string) error {
	return moveTo(srcPath, dstDir)
}

// 移动文件或目录
// 文件移动：保留原有属性
// 目录移动：保留原有属性，同时移动目录包含的文件及子目录
// 旧路径无效：忽略
// 目标路径为文件：忽略
// 目标路径无效：自动补全
func MoveToAuto(srcPath string, dstDir string, autoPerm os.FileMode) error {
	return moveToDirAuto(srcPath, dstDir, autoPerm)
}

//
//func MoveFilesByFilter(srcDir string, newDir string, filter PathFilter) (count int, err error) {
//	srcDir = FormatPath(srcDir)
//	if !IsFolder(srcDir) {
//		return 0, errors.New(fmt.Sprintf("Path(%s) is not exist. ", srcDir))
//	}
//
//}

// -------------

// 移动文件或目录
func move(srcPath string, dstPath string) error {
	return os.Rename(srcPath, dstPath)
}

// 移动文件或目录到目标目录
func moveTo(srcPath string, dstDir string) error {
	_, fileName := Split(srcPath)
	return move(srcPath, Combine(dstDir, fileName))
}

// 移动文件或目录
// 自动补全目标路径
func moveFileAuto(srcPath string, dstPath string, autoPerm os.FileMode) error {
	var dir, _ = Split(dstPath)
	if !IsExist(dir) {
		err := os.MkdirAll(dir, autoPerm)
		if nil != err {
			return err
		}
	}
	return move(srcPath, dstPath)
}

// 移动文件或目录
// 自动补全目标路径
func moveToDirAuto(srcPath string, dstDir string, autoPerm os.FileMode) error {
	dstDir = FormatPath(dstDir)
	if !IsExist(dstDir) {
		err := os.MkdirAll(dstDir, autoPerm)
		if nil != err {
			return err
		}
	}
	return moveTo(srcPath, dstDir)
}
