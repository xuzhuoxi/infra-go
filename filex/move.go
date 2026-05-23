package filex

import (
	"os"
)

// Move renames or moves srcPath to dstPath, overwriting dstPath if it exists.
// Move 将 srcPath 移动/重命名为 dstPath；目标已存在时覆盖。
//
// Invalid srcPath is ignored; invalid dstPath is ignored; directory moves include all contents.
// 源路径无效时忽略；目标路径无效时忽略；移动目录时包含其全部内容。
func Move(srcPath string, dstPath string) error {
	return move(FormatPath(srcPath), FormatPath(dstPath))
}

// MoveAuto moves srcPath to dstPath and creates missing parent directories with autoPerm.
// MoveAuto 移动 srcPath 到 dstPath；目标父目录不存在时按 autoPerm 补全。
//
// Invalid srcPath is ignored; invalid dstPath is completed; directory moves include all contents.
// 源路径无效时忽略；目标路径无效时补全；移动目录时包含其全部内容。
func MoveAuto(srcPath string, dstPath string, autoPerm os.FileMode) error {
	return moveFileAuto(FormatPath(srcPath), FormatPath(dstPath), autoPerm)
}

// MoveTo moves srcPath into dstDir using the base name of srcPath.
// MoveTo 将 srcPath 移动到 dstDir，保留源路径的文件名。
//
// Invalid srcPath is ignored; invalid dstDir is ignored; existing destination is overwritten.
// 源路径无效时忽略；目标目录无效时忽略；目标已存在时覆盖。
func MoveTo(srcPath string, dstDir string) error {
	return moveTo(srcPath, dstDir)
}

// MoveToAuto moves srcPath into dstDir and creates dstDir with autoPerm when missing.
// MoveToAuto 将 srcPath 移动到 dstDir；目标目录不存在时按 autoPerm 创建。
//
// Invalid srcPath is ignored; if dstDir is a file, the move is ignored.
// 源路径无效时忽略；若 dstDir 为文件则忽略。
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
