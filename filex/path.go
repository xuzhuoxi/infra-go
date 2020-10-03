package filex

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	PathSeparator     = "/" // OS-specific path separator
	PathListSeparator = ":" // OS-specific path list separator
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

//Folder Func ------------------------------------

// 遍历当前目录
func WalkCurrent(currentDir string, walkFn filepath.WalkFunc) error {
	currentDir = FormatDirPath(currentDir)
	list, err := ioutil.ReadDir(currentDir)
	if err != nil {
		return walkFn(currentDir, nil, err)
	}
	for _, file := range list {
		path := currentDir + PathSeparator + file.Name()
		err := walkFn(path, file, nil)
		if nil != err {
			return nil
		}
	}
	return nil
}

// 遍历,包含子目录的全部
func Walk(dir string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(dir, walkFn)
}

//File Func ------------------------------------

// 取不包含扩展名的部分的文件名
// 不检查存在性
// 如果目录名字包含".",同样只截取"."前部分
func GetShortName(path string) string {
	_, fileName := Split(path)
	if "" == fileName {
		return ""
	}
	var dot = strings.LastIndex(fileName, ".")
	if -1 == dot {
		return fileName
	}
	return fileName[:dot]
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

// 拆分路为目录+文件， 或父级目录+当前目录
// 返回的目录格式经过FormatDirPath处理
func Split(path string) (formatDir string, fileName string) {
	formatDir, fileName = filepath.Split(FormatDirPath(path))
	return FormatDirPath(formatDir), fileName
}

// 取上级目录，如果没有目录分隔符，返回失败
// 根目录的上级目录为空，并返回失败
// dir要求是经过FormatPath处理后的路径格式
func GetUpDir(dir string) (upDir string, ok bool) {
	dir = FormatDirPath(dir)
	if PathSeparator == dir {
		return "", false
	}
	upDir, _ = filepath.Split(dir)
	if upDir == "" {
		return upDir, false
	}
	return FormatDirPath(upDir), true
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
