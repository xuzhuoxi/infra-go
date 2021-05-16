package filex

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	ExtSeparator      = '.' // OS-specific path separator
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator

	ExtSeparatorStr      = "." // OS-specific path separator
	PathSeparatorStr     = "/" // OS-specific path separator
	PathListSeparatorStr = ":" // OS-specific path list separator
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

// 遍历指定目录
// 只对dir下一级文件或目录执行walkFn
// 返回的path已使用FormatPath处理
// 注意：WalkInDir 可以 对dir目录中的文件进行增删改
func WalkInDir(dir string, walkFn filepath.WalkFunc) error {
	dir = FormatPath(dir)
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, file := range list {
		path := Combine(dir, file.Name())
		err = walkFn(path, file, nil)
		if nil != err {
			return nil
		}
	}
	return nil
}

// 遍历,包含自身与子目录的全部
// 返回的path已使用FormatPath处理
// 注意：在Walk过程中 不可以 对dir目录(包括子目录)中的文件进行增删
func WalkAll(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return filepath.Walk(path, func(eachPath string, info os.FileInfo, err error) error {
		return walkFn(FormatPath(eachPath), info, err)
	})
}

// 遍历文件,包含自身与子目录文件
// 返回的path已使用FormatPath处理
// 注意：在Walk过程中 不可以 对dir目录(包括子目录)中的文件进行增删
func WaldAllFiles(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return WalkAll(path, func(eachPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return walkFn(eachPath, info, err)
	})
}

// 遍历文件夹,包含自身与子目录
// 返回的path已使用FormatPath处理
// 注意：在Walk过程中 不可以 对dir目录(包括子目录)中的文件进行增删
func WaldAllDirs(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return WalkAll(path, func(eachPath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		return walkFn(eachPath, info, err)
	})
}

// 遍历并根据筛选器获取全部路径
// 当filter=nil时，默认为命中
// 路径已进行FormatPath处理
func GetPathsAll(dir string, filter PathFilter) (paths []string, err error) {
	err = WalkAll(dir, func(path string, info os.FileInfo, err1 error) error {
		if nil != err1 {
			return err1
		}
		if nil == filter || filter(path, info) {
			paths = append(paths, path)
		}
		return nil
	})
	return
}

// 遍历指定目录，并根据筛选器获取全部路径
// 当filter=nil时，默认为命中
// 不对子目录内容与当前目录进行筛选
func GetPathsInDir(dir string, filter PathFilter) (paths []string, err error) {
	dir = FormatPath(dir)
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range list {
		path := Combine(dir, file.Name())
		if nil == filter {
			paths = append(paths, path)
			continue
		}
		stat, err := os.Stat(path)
		if nil != err {
			return paths, err
		}
		if filter(path, stat) {
			paths = append(paths, path)
		}
	}
	return paths, nil
}

//File Func ------------------------------------

// 合并路径
// 不检测有效性
func Combine(dir string, add string) string {
	dir = FormatPath(dir)
	if "" == add {
		return dir
	}
	add = FormatPath(add)
	if IsPathSeparatorStr(dir) && IsPathSeparatorStr(add) {
		return PathSeparatorStr
	}
	if IsPathSeparator(add[0]) {
		return dir + add
	}
	return dir + PathSeparatorStr + add
}

// 取不包含扩展名的部分的文件名
// 不检查存在性
// 如果目录名字包含".",同样只截取"."前部分
func GetShortName(path string) string {
	shortName, _, _ := SplitFileName(path)
	return shortName
}

// 检查文件名的扩展名, 支持带"."
func CheckExt(path string, extName string) bool {
	if "" == extName || "." == extName {
		return false
	}
	_, dotExt, ext := SplitFileName(path)
	if IsExtSeparator(extName[0]) {
		return strings.ToLower(dotExt) == strings.ToLower(extName)
	} else {
		return strings.ToLower(ext) == strings.ToLower(extName)
	}
}

//取文件扩展名，不包含"."
func GetExtWithoutDot(path string) string {
	_, _, ext := SplitFileName(path)
	return ext
}

//取文件扩展名，包含"."
func GetExtWithDot(path string) string {
	_, ext, _ := SplitFileName(path)
	return ext
}

// 拆分文件名[shortName + dotExt + ext]
// shortName: 不带扩展名的文件名
// dotExt: 带“.”的扩展名
// ext: 不带“.”的扩展名
func SplitFileName(path string) (shortName string, dotExt string, ext string) {
	_, fileName := Split(path)
	if "" == fileName {
		return
	}
	dot := strings.LastIndexByte(fileName, ExtSeparator)
	if -1 == dot {
		return fileName, "", ""
	}
	if len(fileName)-1 == dot {
		return fileName, ExtSeparatorStr, ""
	}
	shortName, dotExt, ext = fileName[:dot], fileName[dot:], fileName[dot+1:]
	return
}

// 拆分路为目录+文件， 或父级目录+当前目录
// 返回的目录格式经过FormatPath处理
func Split(path string) (formattedDir string, fileName string) {
	path = FormatPath(path)
	if IsPathSeparatorStr(path) {
		return PathSeparatorStr, ""
	}
	dot := strings.LastIndexByte(path, PathSeparator)
	if -1 == dot { //只有文件名
		return "", path
	}
	if 0 == dot { //根目录
		return PathSeparatorStr, path[1:]
	}
	formattedDir, fileName = path[:dot], path[dot+1:]
	return
}

// 取上级目录，如果没有目录分隔符，返回失败
// 根目录的上级目录为空，并返回失败
// dir要求是经过FormatPath处理后的路径格式
// 注意：如果文件名或目录名中使用了"/"字符，会造成结果错误
func GetUpDir(dir string) (upDir string, err error) {
	up, cur := Split(dir)
	if "" == up {
		return up, &os.PathError{Op: "GetUpDir", Path: dir, Err: ErrNoUpperDir}
	}
	if IsPathSeparatorStr(up) && "" == cur {
		return upDir, &os.PathError{Op: "GetUpDir", Path: dir, Err: ErrRootPath}
	}
	return up, nil
}

//Format ------------------------------------

// 标准化路径
// 转换为"/"形式路径
// 如果结果路径为目录，并以"/"结尾，清除"/"
// 不检测有效性
func FormatPath(path string) string {
	path = strings.Replace(path, `\`, PathSeparatorStr, -1)
	var ln = len(path)
	if IsPathSeparator(path[ln-1]) && ln > 1 {
		return path[:ln-1]
	}
	return path
}

//Basic ------------------------------------

func IsPathSeparator(c uint8) bool {
	return PathSeparator == c
}

func IsPathSeparatorStr(str string) bool {
	return PathSeparatorStr == str
}

func IsExtSeparator(c uint8) bool {
	return ExtSeparator == c
}
