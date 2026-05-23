package filex

import (
	"os"
	"path/filepath"
)

// GetSize returns the size of a file or the total size of all files under a directory.
// GetSize 返回文件大小，或目录下全部文件大小之和。
func GetSize(path string) (uint64, error) {
	fi, err := os.Stat(path)
	if nil != err {
		return 0, err
	}
	if fi.IsDir() {
		return GetFolderSize(path)
	} else {
		return GetFileSize(path)
	}
}

// GetFileSize returns the size of a file; returns 0 when path is a directory.
// GetFileSize 返回文件大小；若路径为目录则返回 0。
func GetFileSize(filePath string) (uint64, error) {
	fi, err := os.Stat(filePath)
	if nil != err {
		return 0, err
	}
	if fi.IsDir() {
		return 0, nil
	}
	return uint64(fi.Size()), nil
}

// GetFolderSize returns the sum of all file sizes under dirPath recursively.
// GetFolderSize 递归累加 dirPath 下全部文件的大小。
func GetFolderSize(dirPath string) (size uint64, err error) {
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err1 error) error {
		if nil != err1 {
			err = err1
			return err
		}
		if info.IsDir() {
			return nil
		}
		size += uint64(info.Size())
		return nil
	})
	return
}
