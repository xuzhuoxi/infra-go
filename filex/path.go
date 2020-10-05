package filex

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	PathSeparator     = "/" // OS-specific path separator
	PathListSeparator = ":" // OS-specific path list separator
)

var (
	errPathEqualSeparator = errors.New("Path equal to separator. ")
	errRootPath           = errors.New("Path is root. ")
)

//检查路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//是否为文件夹
func IsFolder(path string) bool {
	if !IsExist(path) {
		return false
	}
	fi, err := os.Stat(path)
	if nil != err {
		return false
	}
	return fi.IsDir()
}

// 根据文件路径补全缺失目录路径
// filePath必须为文件路径
func CompletePath(filePath string, perm os.FileMode) error {
	dstUpDir, err := GetUpDir(filePath)
	if nil != err {
		return err
	}
	return CompleteDir(dstUpDir, perm)
}

// 补全缺失目录路径
func CompleteDir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

//Folder Func ------------------------------------

// 遍历当前目录
// 返回的path已使用FormatDirPath处理
// 注意：WalkCurrent过程中 可以 对dir目录中的文件进行增改删
func WalkCurrent(currentDir string, walkFn filepath.WalkFunc) error {
	currentDir = FormatDirPath(currentDir)
	stat, err := os.Stat(currentDir)
	if nil != err {
		return err
	}
	err = walkFn(currentDir, stat, nil)
	if nil != err {
		return err
	}
	list, err := ioutil.ReadDir(currentDir)
	if err != nil {
		return nil
	}
	for _, file := range list {
		path := Combine(currentDir, file.Name())
		err = walkFn(path, file, nil)
		if nil != err {
			return nil
		}
	}
	return nil
}

// 遍历,包含子目录的全部
// 返回的path已使用FormatDirPath处理
// 注意：在Walk过程中 不可以 对dir目录(包括子目录)中的文件进行增删
func Walk(dir string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		return walkFn(FormatDirPath(path), info, err)
	})
}

//File Func ------------------------------------

// 合并路径
// 不检测有效性
func Combine(dir string, add string) string {
	dir = FormatDirPath(dir)
	if "" == add {
		return dir
	}
	add = FormatDirPath(add)
	if dir == PathSeparator && add == PathSeparator {
		return PathSeparator
	}
	if add[0] == '/' {
		return dir + add
	}
	return dir + "/" + add
}

// 取不包含扩展名的部分的文件名
// 不检查存在性
// 如果目录名字包含".",同样只截取"."前部分
func GetShortName(path string) string {
	shortName, _ := SplitFileName(path, true)
	return shortName
}

// 检查文件名的扩展名, 支持带"."
func CheckExt(path string, extName string) bool {
	if "" == extName || "." == extName {
		return false
	}
	if extName[0] == '.' {
		return strings.ToLower(GetExtWithDot(path)) == strings.ToLower(extName)
	} else {
		return strings.ToLower(GetExtWithoutDot(path)) == strings.ToLower(extName)
	}
}

//取文件扩展名，不包含"."
func GetExtWithoutDot(path string) string {
	var ext = GetExtWithDot(path)
	if "" == ext {
		return ""
	}
	return ext[1:]
}

//取文件扩展名，包含"."
func GetExtWithDot(path string) string {
	path = FormatPath(path)
	return filepath.Ext(path)
}

// 拆分文件名[shortName+ext]
func SplitFileName(path string, keepDot bool) (shortName string, ext string) {
	_, path = Split(path)
	for i := len(path) - 1; i >= 0 && '/' != path[i]; i-- {
		if path[i] == '.' {
			if keepDot {
				return path[:i], path[i:]
			} else {
				return path[:i], path[i+1:]
			}
		}
	}
	return path, ""
}

// 拆分路为目录+文件， 或父级目录+当前目录
// 返回的目录格式经过FormatDirPath处理
func Split(path string) (formatDir string, fileName string) {
	formatDir, fileName = filepath.Split(FormatDirPath(path))
	return FormatDirPath(formatDir), fileName
}

// 取上级目录，如果没有目录分隔符，返回失败
// 根目录的上级目录为空，并返回失败
// dir要求是经过FormatPath处理后的路径格式
// 注意：如果文件名或目录名中使用了"/"字符，会造成结果错误
func GetUpDir(dir string) (upDir string, err error) {
	dir = FormatDirPath(dir)
	if PathSeparator == dir {
		return "", &os.PathError{Op: "GetUpDir", Path: dir, Err: errPathEqualSeparator}
	}
	upDir, _ = filepath.Split(dir)
	if upDir == "" {
		return upDir, &os.PathError{Op: "GetUpDir", Path: dir, Err: errRootPath}
	}
	return FormatDirPath(upDir), nil
}

//Format ------------------------------------

// 如果当前路径以目录分隔符结尾，则去除
// 路径要求是经过FormatPath处理后的路径格式
func FormatDirPath(path string) string {
	path = FormatPath(path)
	var ln = len(path)
	if path[ln-1] == '/' && ln > 1 {
		return path[:ln-1]
	}
	return path
}

// 标准化路径
// 转换为"/"形式路径
// 如果结果路径为目录，并以"/"结尾，清除"/"
// 不检测有效性
func FormatPath(path string) string {
	return strings.Replace(path, `\`, PathSeparator, -1)
}
