package filex

import (
	"bufio"
	"io"
	"os"
)

const (
	LittleFile = 32 * 1024  // LittleFile is the small-file copy threshold (32 KiB). 小文件复制阈值（32 KiB）.
	MiddleFile = 256 * 1024 // MiddleFile is the medium-file copy threshold (256 KiB). 中等文件复制阈值（256 KiB）.
)

// Copy copies srcFile to dstFile using a size-based strategy; the destination parent directory must exist.
// Copy 按文件大小选择策略复制 srcFile 到 dstFile；要求目标路径的父目录已存在。
func Copy(srcFile string, dstFile string) (written int64, err error) {
	srcFile = FormatPath(srcFile)
	stat, err := os.Stat(srcFile)
	if nil != err {
		return 0, err
	}
	dstFile = FormatPath(dstFile)
	return copy(srcFile, dstFile, stat, stat.Mode())
}

// CopyAuto copies srcFile to dstFile and creates missing parent directories with autoPerm.
// CopyAuto 复制 srcFile 到 dstFile；若父目录不存在则按 autoPerm 自动补全。
func CopyAuto(srcFile string, dstFile string, autoPerm os.FileMode) (written int64, err error) {
	srcFile = FormatPath(srcFile)
	stat, err := os.Stat(srcFile)
	if nil != err {
		return 0, err
	}
	dstFile = FormatPath(dstFile)
	err = CompleteParentPath(dstFile, autoPerm)
	if nil != err {
		return 0, err
	}
	return copy(srcFile, dstFile, stat, stat.Mode())
}

// CopyTo copies srcFile into targetDir using the base file name.
// CopyTo 将 srcFile 复制到 targetDir，保留源文件名。
func CopyTo(srcFile string, targetDir string) (written int64, err error) {
	var _, name = Split(srcFile)
	var newPath = Combine(targetDir, name)
	return Copy(srcFile, newPath)
}

// CopyToAuto copies srcFile into targetDir and creates missing parent directories with autoPerm.
// CopyToAuto 将 srcFile 复制到 targetDir；若父目录不存在则按 autoPerm 自动补全。
func CopyToAuto(srcFile string, targetDir string, autoPerm os.FileMode) (written int64, err error) {
	var _, name = Split(srcFile)
	var newPath = Combine(targetDir, name)
	return CopyAuto(srcFile, newPath, autoPerm)
}

// CopyMod copies srcFile to dstFile and sets the destination permission to filePerm.
// CopyMod 复制 srcFile 到 dstFile，并将目标权限设为 filePerm。
func CopyMod(srcFile string, dstFile string, filePerm os.FileMode) (written int64, err error) {
	srcFile = FormatPath(srcFile)
	stat, err := os.Stat(srcFile)
	if nil != err {
		return 0, err
	}
	return copy(srcFile, dstFile, stat, filePerm)
}

// CopyModAuto copies srcFile to dstFile with filePerm and creates missing parents with autoPerm.
// CopyModAuto 复制并设置 filePerm；若父目录不存在则按 autoPerm 自动补全。
func CopyModAuto(srcFile string, dstFile string, filePerm os.FileMode, autoPerm os.FileMode) (written int64, err error) {
	srcFile = FormatPath(srcFile)
	stat, err := os.Stat(srcFile)
	if nil != err {
		return 0, err
	}
	dstFile = FormatPath(dstFile)
	err = CompleteParentPath(dstFile, autoPerm)
	if nil != err {
		return 0, err
	}
	return copy(srcFile, dstFile, stat, filePerm)
}

// CopyModTo copies srcFile into targetDir with filePerm.
// CopyModTo 将 srcFile 复制到 targetDir，并将目标权限设为 filePerm。
func CopyModTo(srcFile string, targetDir string, filePerm os.FileMode) (written int64, err error) {
	var _, name = Split(srcFile)
	var newPath = Combine(targetDir, name)
	return CopyMod(srcFile, newPath, filePerm)
}

// CopyModToAuto copies srcFile into targetDir with filePerm and creates missing parents with autoPerm.
// CopyModToAuto 将 srcFile 复制到 targetDir 并设置 filePerm；若父目录不存在则按 autoPerm 自动补全。
func CopyModToAuto(srcFile string, targetDir string, filePerm os.FileMode, autoPerm os.FileMode) (written int64, err error) {
	var _, name = Split(srcFile)
	var newPath = Combine(targetDir, name)
	return CopyModAuto(srcFile, newPath, filePerm, autoPerm)
}

// 以下为内部函数-------------------------

// 复制文件
// 根据文件大小选择不同的复制算法
func copy(srcFile string, dstFile string, srcStat os.FileInfo, perm os.FileMode) (written int64, err error) {
	if nil == srcStat {
		srcStat, err = os.Stat(srcFile)
		if nil != err {
			return 0, err
		}
	}
	size := srcStat.Size()
	if size <= LittleFile {
		return copy1(srcFile, dstFile, perm)
	}
	if size <= MiddleFile {
		return copy2(srcFile, dstFile, perm)
	}
	return copy3(srcFile, dstFile, perm)
}

// 使用ioutil包中的API进行文件复制
// 不建议使用于大文件
func copy1(srcFile string, dstFile string, perm os.FileMode) (written int64, err error) {
	srcBytes, err := os.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}

	//打开dstFileName
	err = os.WriteFile(dstFile, srcBytes, perm)
	if err != nil {
		return 0, err
	}

	return int64(len(srcBytes)), nil
}

// 使用io包中API进行文件复制
// 复制数据时使用缓冲区，可用于大文件复制
func copy2(src string, dst string, perm os.FileMode) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//打开dstFileName
	dstFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

// 使用io包中API进行文件复制
// 使用了三层缓冲：1.读文件使用了缓冲；2.写文件使用了缓冲；3.复制数据时使用了缓冲
// 可用于大文件复制
func copy3(src string, dst string, perm os.FileMode) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//打开dstFileName
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	return io.Copy(writer, reader)
}
