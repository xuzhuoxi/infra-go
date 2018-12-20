package osxu

import (
	"github.com/xuzhuoxi/go-util/stringsx"
	"io/ioutil"
	"os"
	"strings"
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

//取文件或文件夹大小
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

//取文件大小
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

//取文件夹大小，递归全部文件的大小之和
func GetFolderSize(dirPath string) (uint64, error) {
	list, err := GetFolderFileList(dirPath)
	if err != nil {
		return 0, err
	}
	var size uint64 = 0
	for _, file := range list {
		size += uint64(file.Size())
	}
	return size, nil
}

//取文件夹下全部文件
func GetFolderFileList(dirPath string) ([]os.FileInfo, error) {
	dirPath = GetUnitePath(dirPath)
	_, err := os.Stat(dirPath)
	if nil != err {
		return nil, err
	}
	dirLen := stringsx.GetCharCount(dirPath)
	if stringsx.LastIndexOfChar(dirPath, "/") != dirLen-1 { //最后一个不是"/"
		dirPath = dirPath + "/"
	}
	var rs []os.FileInfo
	var recursiveFunc func(folderPath string) = nil
	recursiveFunc = func(folderPath string) { //folderPath必须为"/"结尾
		list, e := ioutil.ReadDir(folderPath)
		if nil != e {
			return
		}
		for _, file := range list {
			if file.IsDir() {
				recursiveFunc(folderPath + file.Name() + "/")
			} else {
				rs = append(rs, file)
			}
		}
	}
	recursiveFunc(dirPath)
	return rs, nil
}

//取扩展名
func GetExtensionName(fileName string) string {
	_, eName := SplitFileName(fileName)
	return eName
}

//取文件名
func GetFilePrefixName(fileName string) string {
	bName, _ := SplitFileName(fileName)
	return bName
}

//取父文件夹(父目录)
func GetParentDir(dirPath string) (string, bool) {
	newPath := GetUnitePath(dirPath)
	pathLen := stringsx.GetCharCount(newPath)
	dot := stringsx.LastIndexOfChar(newPath, "/")
	if dot == -1 { //无效路径 或 windows顶级路径
		return "", false
	}
	if dot == 0 { //在头部
		if pathLen == 1 {
			return "", false
		} else {
			return "/", true
		}
	}
	var f = func(str string) (string, bool) { //保证"/"不在最后一个字符
		d := stringsx.LastIndexOfChar(str, "/")
		if -1 == d {
			return "", false
		}
		if 0 == d {
			return "/", true
		}
		return stringsx.SubPrefix(str, d+1), true
	}
	if dot < pathLen-1 { //在中间
		return f(newPath)
	}
	if dot == pathLen-1 { //在尾部
		return f(stringsx.SubPrefix(newPath, dot))
	}
	return "", false
}

//转换为"/"形式路径
func GetUnitePath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

//把文件名拆分
func SplitFileName(fileName string) (string, string) {
	if "" == fileName || 0 == stringsx.GetCharCount(fileName) {
		return "", ""
	}
	dot := stringsx.LastIndexOfChar(fileName, ".")
	if -1 == dot {
		return fileName, ""
	}
	return stringsx.CutString(fileName, dot, false)
}

//private ------------------------------------
