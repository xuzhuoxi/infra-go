package filex

import (
	"io/ioutil"
	"os"
)

// 删除全部
// 路径为文件时，删除文件
// 路径为目录时，删除目录
func RemoveAll(path string) error {
	path = FormatDirPath(path)
	return os.RemoveAll(path)
}

// 删除文件
// 或
// 删除空目录
func Remove(path string) error {
	path = FormatDirPath(path)
	return remove(path)
}

// 清空目录
func ClearDir(dir string) error {
	dir = FormatDirPath(dir)
	return clearDir(dir)
}

func remove(dir string) error {
	return os.Remove(dir)
}

func clearDir(dir string) error {
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, file := range list {
		path := Combine(dir, file.Name())
		err := os.RemoveAll(path)
		if nil != err {
			return err
		}
	}
	return nil
}
