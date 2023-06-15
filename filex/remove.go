package filex

import (
	"io/ioutil"
	"os"
)

// RemoveAll
// 删除全部
// 路径为文件时，删除文件
// 路径为目录时，删除目录
func RemoveAll(path string) error {
	path = FormatPath(path)
	return removeAll(path)
}

// Remove
// 删除文件 或 删除空目录
func Remove(path string) error {
	path = FormatPath(path)
	return remove(path)
}

// RemoveChildren
// 清空目录内容
// 保留当前目录
func RemoveChildren(dir string) error {
	dir = FormatPath(dir)
	return clearDir(dir)
}

// RemoveEmptyDir
// 清除空目录
// 如果当前路径也为空目录，清除之
// 返回清除数量与错误
func RemoveEmptyDir(dir string) (count int, err error) {
	dir = FormatPath(dir)
	var paths []string
	WalkAll(dir, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	index := len(paths) - 1
	for index > 0 {
		list, err := ioutil.ReadDir(paths[index])
		index -= 1
		if nil != err {
			return count, err
		}
		if len(list) == 0 {
			err = remove(paths[index])
			if nil != err {
				return count, err
			}
			count += 1
		}
	}
	return count, nil
}

// RemoveFiles
// 清空目录中全部文件
// 保留目录结构
func RemoveFiles(dir string) (count int, err error) {
	return RemoveByFilter(dir, func(path string, info os.FileInfo) bool {
		return !info.IsDir()
	})
}

// RemoveByFilter
// 使用PathFilter进行可选删除
func RemoveByFilter(dir string, filter PathFilter) (count int, err error) {
	paths, err := GetPathsAll(dir, filter)
	if nil != err {
		return 0, err
	}
	var index = 0
	for index = range paths {
		err = removeAll(paths[index])
		if nil != err {
			return index, err
		}
	}
	return index + 1, nil
}

func remove(path string) error {
	return os.Remove(path)
}

func removeAll(path string) error {
	return os.RemoveAll(path)
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
