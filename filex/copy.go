package filex

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

const (
	LittleFile = 32 * 1024
	MiddleFile = 256 * 1024
)

// 复制文件
// 根据文件大小选择不同的复制算法
// 要求目录路径的父目录存在
func Copy(srcFile string, dstFile string) (written int64, err error) {
	srcFile = FormatPath(srcFile)
	stat, err := os.Stat(srcFile)
	if nil != err {
		return 0, err
	}
	dstFile = FormatPath(dstFile)
	return copy(srcFile, dstFile, stat, stat.Mode())
}

// 复制文件
// 根据文件大小选择不同的复制算法
// 如果目录路径的父目录不存在，自动补全
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

// 复制文件到指定目录
// 根据文件大小选择不同的复制算法
func CopyTo(srcFile string, targetDir string) (written int64, err error) {
	var _, name = Split(srcFile)
	var newPath = Combine(targetDir, name)
	return Copy(srcFile, newPath)
}

// 复制文件到指定目录
// 根据文件大小选择不同的复制算法
// 如果目录路径的父目录不存在，自动补全
func CopyToAuto(srcFile string, targetDir string, autoPerm os.FileMode) (written int64, err error) {
	var _, name = Split(srcFile)
	var newPath = Combine(targetDir, name)
	return CopyAuto(srcFile, newPath, autoPerm)
}

// 复制文件
// 根据文件大小选择不同的复制算法
// 同时设置新的FileMode
func CopyMod(srcFile string, dstFile string, filePerm os.FileMode) (written int64, err error) {
	srcFile = FormatPath(srcFile)
	stat, err := os.Stat(srcFile)
	if nil != err {
		return 0, err
	}
	return copy(srcFile, dstFile, stat, filePerm)
}

// 复制文件
// 根据文件大小选择不同的复制算法
// 同时设置新的FileMode
// 如果目录路径的父目录不存在，自动补全
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

// 复制文件到指定目录
// 根据文件大小选择不同的复制算法
// 同时设置新的FileMode
func CopyModTo(srcFile string, targetDir string, filePerm os.FileMode) (written int64, err error) {
	var _, name = Split(srcFile)
	var newPath = Combine(targetDir, name)
	return CopyMod(srcFile, newPath, filePerm)
}

// 复制文件到指定目录
// 根据文件大小选择不同的复制算法
// 同时设置新的FileMode
// 如果目录路径的父目录不存在，自动补全
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
	srcBytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}

	//打开dstFileName
	err = ioutil.WriteFile(dstFile, srcBytes, perm)
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
