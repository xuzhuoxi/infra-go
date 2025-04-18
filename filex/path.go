package filex

import (
	"github.com/xuzhuoxi/infra-go/osxu"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	ExtSeparator         = '.' // 扩展名分隔符
	ExtSeparatorStr      = "." // 扩展名分隔符
	PathListSeparator    = ';' // 路径列表分隔符
	PathListSeparatorStr = ";" // 路径列表分隔符
)

const (
	UnixSeparator       = '/'  // Linux路径级分隔符
	UnixSeparatorStr    = "/"  // Linux路径级分隔符
	WindowsSeparator    = '\\' // Windows路径级分隔符
	WindowsSeparatorStr = "\\" // Windows路径级分隔符
)

var (
	RunningDir = osxu.GetRunningDir()
)

// IsRelativeFormat 检查路径是否为相对路径格式
func IsRelativeFormat(path string) bool {
	return !IsAbsFormat(path)
}

// IsAbsFormat 检查路径是否为绝对路径格式
func IsAbsFormat(path string) bool {
	if strings.Contains(path, ":") {
		return true
	}
	if strings.HasPrefix(path, "/") || strings.HasPrefix(path, `\\`) {
		return true
	}
	return IsUnixAbs(path) || IsWinAbs(path)
}

// IsUnixAbs 检查路径是否为Unix绝对路径格式
func IsUnixAbs(path string) bool {
	return isUnixAbs(path)
}

// IsWinAbs 检查路径是否为Windows绝对路径格式
func IsWinAbs(path string) bool {
	return isWinAbs(path)
}

// IsExist 检查路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	//return err == nil || errors.Is(err, os.ErrExist)
}

// IsFile 是否为文件
func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if nil != err && !os.IsExist(err) {
		return false
	}
	return !fi.IsDir()
}

// IsDir 是否为文件夹
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if nil != err && !os.IsExist(err) {
		return false
	}
	return fi.IsDir()
}

// IsFolder 是否为文件夹
func IsFolder(path string) bool {
	return IsDir(path)
}

// IsEmptyDir 是否为空文件夹
func IsEmptyDir(path string) bool {
	fi, err := os.Stat(path)
	if nil != err && !os.IsExist(err) {
		return false
	}
	if !fi.IsDir() {
		return false
	}
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}

	return len(list) == 0
}

// IsEmptyFolder 是否为空文件夹
func IsEmptyFolder(path string) bool {
	return IsEmptyDir(path)
}

// CompleteParentPath 根据文件路径补全缺失父目录路径
// filePath必须绝对路径
func CompleteParentPath(absPath string, perm os.FileMode) error {
	dstUpDir, err := GetUpDir(absPath)
	if nil != err {
		return err
	}
	filepath.Abs(dstUpDir)
	return CompleteDir(dstUpDir, perm)
}

// CompleteDir 补全缺失目录路径
func CompleteDir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

//Folder Func ------------------------------------

// WalkInDir
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

// WalkAll
// 遍历,包含自身与子目录的全部
// 返回的path已使用FormatPath处理
// 注意：在Walk过程中 不可以 对dir目录(包括子目录)中的文件进行增删
func WalkAll(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return filepath.Walk(path, func(eachPath string, info os.FileInfo, err error) error {
		return walkFn(FormatPath(eachPath), info, err)
	})
}

// WalkAllFiles
// 遍历文件,包含自身与子目录文件
// 返回的path已使用FormatPath处理
// 注意：在Walk过程中 不可以 对dir目录(包括子目录)中的文件进行增删
func WalkAllFiles(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return WalkAll(path, func(eachPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return walkFn(eachPath, info, err)
	})
}

// WalkAllDirs
// 遍历文件夹,包含自身与子目录
// 返回的path已使用FormatPath处理
// 注意：在Walk过程中 不可以 对dir目录(包括子目录)中的文件进行增删
func WalkAllDirs(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return WalkAll(path, func(eachPath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		return walkFn(eachPath, info, err)
	})
}

// GetPathsAll
// 遍历并根据筛选器获取全部路径(递归)
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

// GetPathsInDir
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

//FilePath Func ------------------------------------

// FixFilePath
// 修复文件路径
// 文件存在：则返回 传入路径
// 文件不存在：则返回 运行时目录+传入路径
func FixFilePath(filePath string) string {
	if IsFile(filePath) {
		return filePath
	}
	return Combine(RunningDir, filePath)
}

// FixDirPath
// 修复目录路径
// 目录存在：则返回 传入路径
// 目录不存在：则返回 运行时目录+传入路径
func FixDirPath(dirPath string) string {
	if IsDir(dirPath) {
		return dirPath
	}
	return Combine(RunningDir, dirPath)
}

// FixPath
// 修复路径
// 路径存在：则返回 传入路径
// 路径不存在：则返回 运行时目录+传入路径
func FixPath(path string) string {
	if IsExist(path) {
		return path
	}
	return Combine(RunningDir, path)
}

// Combine
// 合并路径
// 不检测有效性
func Combine(dir string, add string, subs ...string) string {
	path := combine(dir, add)
	for _, sub := range subs {
		path = combine(path, sub)
	}
	return path
}

func combine(dir string, add string) string {
	if "" == dir && "" == add {
		return ""
	}
	if "" != dir {
		dir = FormatPath(dir)
	}
	if "" != add {
		add = FormatPath(add)
	}
	if "" == add {
		return dir
	}
	if "" == dir {
		return add
	}
	if IsUnixSeparatorStr(dir) && IsUnixSeparatorStr(add) {
		return UnixSeparatorStr
	}
	if IsUnixSeparator(add[0]) {
		return dir + add
	}
	return dir + UnixSeparatorStr + add
}

// GetShortName
// 取不包含扩展名的部分的文件名
// 不检查存在性
// 如果目录名字包含".",同样只截取"."前部分
func GetShortName(path string) string {
	shortName, _, _ := SplitFileName(path)
	return shortName
}

// CheckExt
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

// GetExtWithoutDot
// 取文件扩展名，不包含"."
func GetExtWithoutDot(path string) string {
	_, _, ext := SplitFileName(path)
	return ext
}

// GetExtWithDot
// 取文件扩展名，包含"."
func GetExtWithDot(path string) string {
	_, ext, _ := SplitFileName(path)
	return ext
}

// SplitFileName
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

// Split
// 拆分路为目录+文件， 或父级目录+当前目录
// 返回的目录格式经过FormatPath处理
func Split(path string) (formattedDir string, fileName string) {
	path = FormatPath(path)
	if IsUnixSeparatorStr(path) {
		return UnixSeparatorStr, ""
	}
	dot := strings.LastIndexByte(path, UnixSeparator)
	if -1 == dot { //只有文件名
		return "", path
	}
	if 0 == dot { //根目录
		return UnixSeparatorStr, path[1:]
	}
	formattedDir, fileName = path[:dot], path[dot+1:]
	return
}

// GetUpDir
// 取上级目录，如果没有目录分隔符，返回失败
// 根目录的上级目录为空，并返回失败
// dir要求是经过FormatPath处理后的路径格式
// 注意：如果文件名或目录名中使用了"/"字符，会造成结果错误
func GetUpDir(dir string) (upDir string, err error) {
	up, cur := Split(dir)
	if "" == up {
		return up, &os.PathError{Op: "GetUpDir", Path: dir, Err: ErrNoUpperDir}
	}
	if IsUnixSeparatorStr(up) && "" == cur {
		return upDir, &os.PathError{Op: "GetUpDir", Path: dir, Err: ErrRootPath}
	}
	return up, nil
}

//Format ------------------------------------

// FormatPath
// 标准化路径(转为Linux路径)
// 转换为"/"形式路径
// 如果结果路径为目录，并以"/"结尾，清除"/"
// 不检测有效性
func FormatPath(path string) string {
	return ToUnixPath(path)
}

// ToUnixPath
// 格式化为Linux路径
// 如果结果路径为目录，并以"/"结尾，清除"/"
// 不检测有效性
func ToUnixPath(p string) string {
	p = strings.Replace(p, WindowsSeparatorStr, UnixSeparatorStr, -1)
	var ln = len(p)
	if IsUnixSeparator(p[ln-1]) && ln > 1 {
		return p[:ln-1]
	}
	return p
}

// ToWindowsPath
// 格式化为Windows路径
// 如果结果路径为目录，并以"/"结尾，清除"/"
// 不检测有效性
func ToWindowsPath(p string) string {
	p = strings.Replace(p, UnixSeparatorStr, WindowsSeparatorStr, -1)
	var ln = len(p)
	if IsWindowsSeparator(p[ln-1]) && ln > 1 {
		return p[:ln-1]
	}
	return p
}

//Basic ------------------------------------

func IsUnixSeparator(c uint8) bool {
	return UnixSeparator == c
}

func IsUnixSeparatorStr(str string) bool {
	return UnixSeparatorStr == str
}

func IsWindowsSeparator(c uint8) bool {
	return WindowsSeparator == c
}

func IsWindowsSeparatorStr(str string) bool {
	return WindowsSeparatorStr == str
}

func IsExtSeparator(c uint8) bool {
	return ExtSeparator == c
}

func IsExtSeparatorStr(str string) bool {
	return ExtSeparatorStr == str
}

func IsListSeparator(c uint8) bool {
	return PathListSeparator == c
}

func IsListSeparatorStr(str string) bool {
	return PathListSeparatorStr == str
}
