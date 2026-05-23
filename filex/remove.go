package filex

import (
	"io/ioutil"
	"os"
)

// RemoveAll removes the file or directory tree at path.
// RemoveAll 删除 path 处的文件或整个目录树。
func RemoveAll(path string) error {
	path = FormatPath(path)
	return removeAll(path)
}

// Remove deletes a file or an empty directory.
// Remove 删除文件或空目录。
func Remove(path string) error {
	path = FormatPath(path)
	return remove(path)
}

// RemoveChildren clears all entries inside dir but keeps dir itself.
// RemoveChildren 清空 dir 内的全部内容，但保留 dir 目录本身。
func RemoveChildren(dir string) error {
	dir = FormatPath(dir)
	return clearDir(dir)
}

// RemoveEmptyDir removes empty subdirectories under dir; also removes dir if it becomes empty.
// RemoveEmptyDir 清除 dir 下的空子目录；若 dir 自身也为空则一并删除。
//
// It returns the number of removed directories and any error encountered.
// 返回已删除的空目录数量及遇到的错误。
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

// RemoveFiles deletes all files under dir while preserving the directory structure.
// RemoveFiles 删除 dir 下的全部文件，但保留目录结构。
func RemoveFiles(dir string) (count int, err error) {
	return RemoveByFilter(dir, func(path string, info os.FileInfo) bool {
		return !info.IsDir()
	})
}

// RemoveByFilter deletes paths under dir that match filter.
// RemoveByFilter 按 filter 删除 dir 下匹配的路径。
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
