package filex

import (
	"os"
	"path/filepath"
)

// 取文件或文件夹大小
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

// 取文件大小
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

// 取文件夹大小，递归全部文件的大小之和
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
