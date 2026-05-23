package filex

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuzhuoxi/infra-go/osxu"
)

const (
	ExtSeparator         = '.' // Extension separator. 扩展名分隔符.
	ExtSeparatorStr      = "." // Extension separator string. 扩展名分隔符字符串.
	PathListSeparator    = ';' // Path list separator. 路径列表分隔符.
	PathListSeparatorStr = ";" // Path list separator string. 路径列表分隔符字符串.
)

const (
	UnixSeparator       = '/'  // Unix path separator. Linux 路径分隔符.
	UnixSeparatorStr    = "/"  // Unix path separator string. Linux 路径分隔符字符串.
	WindowsSeparator    = '\\' // Windows path separator. Windows 路径分隔符.
	WindowsSeparatorStr = "\\" // Windows path separator string. Windows 路径分隔符字符串.
)

// RunningDir holds the process running directory from osxu.GetRunningDir.
// RunningDir 为进程运行目录（由 osxu.GetRunningDir 取得）。
var (
	RunningDir = osxu.GetRunningDir()
)

// IsRelativeFormat reports whether path is in relative path format.
// IsRelativeFormat 检查路径是否为相对路径格式。
func IsRelativeFormat(path string) bool {
	return !IsAbsFormat(path)
}

// IsAbsFormat reports whether path is in absolute path format.
// IsAbsFormat 检查路径是否为绝对路径格式。
func IsAbsFormat(path string) bool {
	if strings.Contains(path, ":") {
		return true
	}
	if strings.HasPrefix(path, "/") || strings.HasPrefix(path, `\\`) {
		return true
	}
	return IsUnixAbs(path) || IsWinAbs(path)
}

// IsUnixAbs reports whether path is a Unix absolute path.
// IsUnixAbs 检查路径是否为 Unix 绝对路径格式。
func IsUnixAbs(path string) bool {
	return isUnixAbs(path)
}

// IsWinAbs reports whether path is a Windows absolute path.
// IsWinAbs 检查路径是否为 Windows 绝对路径格式。
func IsWinAbs(path string) bool {
	return isWinAbs(path)
}

// IsExist reports whether path exists.
// IsExist 检查路径是否存在。
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	//return err == nil || errors.Is(err, os.ErrExist)
}

// IsFile reports whether path exists and is a file.
// IsFile 检查路径是否存在且为文件。
func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if nil != err && !os.IsExist(err) {
		return false
	}
	return !fi.IsDir()
}

// IsDir reports whether path exists and is a directory.
// IsDir 检查路径是否存在且为目录。
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if nil != err && !os.IsExist(err) {
		return false
	}
	return fi.IsDir()
}

// IsFolder reports whether path exists and is a directory.
// IsFolder 检查路径是否存在且为文件夹。
func IsFolder(path string) bool {
	return IsDir(path)
}

// IsEmptyDir reports whether path is an existing empty directory.
// IsEmptyDir 检查路径是否为已存在的空目录。
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

// IsEmptyFolder reports whether path is an existing empty directory.
// IsEmptyFolder 检查路径是否为已存在的空文件夹。
func IsEmptyFolder(path string) bool {
	return IsEmptyDir(path)
}

// CompleteParentPath creates missing parent directories for absPath.
// CompleteParentPath 根据文件路径补全缺失的父目录。
//
// absPath must be an absolute path.
// absPath 必须为绝对路径。
func CompleteParentPath(absPath string, perm os.FileMode) error {
	dstUpDir, err := GetUpDir(absPath)
	if nil != err {
		return err
	}
	filepath.Abs(dstUpDir)
	return CompleteDir(dstUpDir, perm)
}

// CompleteDir creates dirPath and any missing parents with perm.
// CompleteDir 补全缺失的目录路径。
func CompleteDir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

//Folder Func ------------------------------------

// WalkInDir walks immediate children of dir and calls walkFn for each entry.
// WalkInDir 遍历指定目录的直接子项并对每一项调用 walkFn。
//
// Returned paths are formatted with FormatPath.
// 返回的 path 已使用 FormatPath 处理。
//
// Files under dir may be added, removed, or modified during the walk.
// 注意：WalkInDir 过程中可对 dir 目录中的文件进行增删改。
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

// WalkAll walks path recursively, including path itself and all descendants.
// WalkAll 递归遍历 path，包含自身与子目录的全部内容。
//
// Returned paths are formatted with FormatPath.
// 返回的 path 已使用 FormatPath 处理。
//
// Do not add, remove, or modify files under dir (including subdirectories) during the walk.
// 注意：Walk 过程中不可对 dir 目录（含子目录）中的文件进行增删。
func WalkAll(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return filepath.Walk(path, func(eachPath string, info os.FileInfo, err error) error {
		return walkFn(FormatPath(eachPath), info, err)
	})
}

// WalkAllFiles walks all files under path recursively, excluding directories.
// WalkAllFiles 递归遍历 path 下的全部文件，不包含目录本身作为文件项。
//
// Returned paths are formatted with FormatPath.
// 返回的 path 已使用 FormatPath 处理。
//
// Do not add, remove, or modify files under dir (including subdirectories) during the walk.
// 注意：Walk 过程中不可对 dir 目录（含子目录）中的文件进行增删。
func WalkAllFiles(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return WalkAll(path, func(eachPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return walkFn(eachPath, info, err)
	})
}

// WalkAllDirs walks all directories under path recursively, including path when it is a directory.
// WalkAllDirs 递归遍历 path 下的全部目录，包含自身。
//
// Returned paths are formatted with FormatPath.
// 返回的 path 已使用 FormatPath 处理。
//
// Do not add, remove, or modify files under dir (including subdirectories) during the walk.
// 注意：Walk 过程中不可对 dir 目录（含子目录）中的文件进行增删。
func WalkAllDirs(path string, walkFn filepath.WalkFunc) error {
	path = FormatPath(path)
	return WalkAll(path, func(eachPath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		return walkFn(eachPath, info, err)
	})
}

// GetPathsAll recursively collects paths under dir that match filter.
// GetPathsAll 递归遍历并根据 filter 获取全部路径。
//
// When filter is nil, every path matches.
// 当 filter 为 nil 时，默认为全部命中。
//
// Paths are formatted with FormatPath.
// 路径已进行 FormatPath 处理。
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

// GetPathsInDir collects paths of immediate children of dir that match filter.
// GetPathsInDir 遍历指定目录并根据 filter 获取一级子路径。
//
// When filter is nil, every child matches; neither the directory itself nor deeper levels are included.
// 当 filter 为 nil 时默认为全部命中；不对子目录内容与当前目录本身进行筛选。
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

// FixFilePath returns filePath when the file exists, otherwise RunningDir joined with filePath.
// FixFilePath 修复文件路径。
//
// When the file exists, the input path is returned; otherwise RunningDir + input path is returned.
// 文件存在时返回传入路径；不存在时返回运行时目录与传入路径的合并结果。
func FixFilePath(filePath string) string {
	if IsFile(filePath) {
		return filePath
	}
	return Combine(RunningDir, filePath)
}

// FixDirPath returns dirPath when the directory exists, otherwise RunningDir joined with dirPath.
// FixDirPath 修复目录路径。
//
// When the directory exists, the input path is returned; otherwise RunningDir + input path is returned.
// 目录存在时返回传入路径；不存在时返回运行时目录与传入路径的合并结果。
func FixDirPath(dirPath string) string {
	if IsDir(dirPath) {
		return dirPath
	}
	return Combine(RunningDir, dirPath)
}

// FixPath returns path when it exists, otherwise RunningDir joined with path.
// FixPath 修复路径。
//
// When the path exists, the input path is returned; otherwise RunningDir + input path is returned.
// 路径存在时返回传入路径；不存在时返回运行时目录与传入路径的合并结果。
func FixPath(path string) string {
	if IsExist(path) {
		return path
	}
	return Combine(RunningDir, path)
}

// Combine joins dir, add, and optional subs without validating existence.
// Combine 合并路径，不检测有效性。
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

// GetShortName returns the base name without extension; existence is not checked.
// GetShortName 取不包含扩展名的文件名，不检查存在性。
//
// If a directory name contains ".", only the part before the last "." is returned.
// 若目录名包含 "."，同样只截取 "." 前部分。
func GetShortName(path string) string {
	shortName, _, _ := SplitFileName(path)
	return shortName
}

// CheckExt reports whether path has extension extName; extName may include a leading ".".
// CheckExt 检查文件名的扩展名，支持带 "." 的 extName。
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

// GetExtWithoutDot returns the file extension without the leading ".".
// GetExtWithoutDot 取文件扩展名，不包含 "."。
func GetExtWithoutDot(path string) string {
	_, _, ext := SplitFileName(path)
	return ext
}

// GetExtWithDot returns the file extension including the leading ".".
// GetExtWithDot 取文件扩展名，包含 "."。
func GetExtWithDot(path string) string {
	_, ext, _ := SplitFileName(path)
	return ext
}

// SplitFileName splits the base name into shortName, dotExt, and ext.
// SplitFileName 拆分文件名为 shortName、dotExt、ext 三部分。
//
// shortName is the name without extension; dotExt includes "."; ext excludes ".".
// shortName 为不带扩展名的文件名；dotExt 为带 "." 的扩展名；ext 为不带 "." 的扩展名。
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

// Split separates path into formatted directory and final name component.
// Split 拆分路径为目录与文件名，或父级目录与当前目录名。
//
// The returned directory is formatted with FormatPath.
// 返回的目录格式经过 FormatPath 处理。
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

// GetUpDir returns the parent directory of dir.
// GetUpDir 取上级目录。
//
// dir should already be formatted with FormatPath.
// dir 要求是经过 FormatPath 处理后的路径格式。
//
// Returns an error when no separator is present or when dir is root.
// 若没有目录分隔符或 dir 为根目录，则返回错误。
//
// A "/" character inside a file or directory name may produce incorrect results.
// 注意：若文件名或目录名中包含 "/"，可能造成结果错误。
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

// FormatPath normalizes path to Unix-style separators.
// FormatPath 标准化路径（转为 Linux 路径）。
//
// Converts to "/" form; trailing "/" on non-root paths is removed; existence is not checked.
// 转换为 "/" 形式；若结果为目录且以 "/" 结尾则清除末尾 "/"；不检测有效性。
func FormatPath(path string) string {
	return ToUnixPath(path)
}

// ToUnixPath formats path as a Unix path.
// ToUnixPath 格式化为 Linux 路径。
//
// Trailing "/" on non-root paths is removed; existence is not checked.
// 若结果为目录且以 "/" 结尾则清除末尾 "/"；不检测有效性。
func ToUnixPath(p string) string {
	p = strings.Replace(p, WindowsSeparatorStr, UnixSeparatorStr, -1)
	var ln = len(p)
	if IsUnixSeparator(p[ln-1]) && ln > 1 {
		return p[:ln-1]
	}
	return p
}

// ToWindowsPath formats path as a Windows path.
// ToWindowsPath 格式化为 Windows 路径。
//
// Trailing separator on non-root paths is removed; existence is not checked.
// 若结果为目录且以 "/" 结尾则清除末尾 "/"；不检测有效性。
func ToWindowsPath(p string) string {
	p = strings.Replace(p, UnixSeparatorStr, WindowsSeparatorStr, -1)
	var ln = len(p)
	if IsWindowsSeparator(p[ln-1]) && ln > 1 {
		return p[:ln-1]
	}
	return p
}

//Basic ------------------------------------

// IsUnixSeparator reports whether c is the Unix path separator.
// IsUnixSeparator 判断 c 是否为 Unix 路径分隔符。
func IsUnixSeparator(c uint8) bool {
	return UnixSeparator == c
}

// IsUnixSeparatorStr reports whether str is the Unix path separator string.
// IsUnixSeparatorStr 判断 str 是否为 Unix 路径分隔符字符串。
func IsUnixSeparatorStr(str string) bool {
	return UnixSeparatorStr == str
}

// IsWindowsSeparator reports whether c is the Windows path separator.
// IsWindowsSeparator 判断 c 是否为 Windows 路径分隔符。
func IsWindowsSeparator(c uint8) bool {
	return WindowsSeparator == c
}

// IsWindowsSeparatorStr reports whether str is the Windows path separator string.
// IsWindowsSeparatorStr 判断 str 是否为 Windows 路径分隔符字符串。
func IsWindowsSeparatorStr(str string) bool {
	return WindowsSeparatorStr == str
}

// IsExtSeparator reports whether c is the extension separator.
// IsExtSeparator 判断 c 是否为扩展名分隔符。
func IsExtSeparator(c uint8) bool {
	return ExtSeparator == c
}

// IsExtSeparatorStr reports whether str is the extension separator string.
// IsExtSeparatorStr 判断 str 是否为扩展名分隔符字符串。
func IsExtSeparatorStr(str string) bool {
	return ExtSeparatorStr == str
}

// IsListSeparator reports whether c is the path list separator.
// IsListSeparator 判断 c 是否为路径列表分隔符。
func IsListSeparator(c uint8) bool {
	return PathListSeparator == c
}

// IsListSeparatorStr reports whether str is the path list separator string.
// IsListSeparatorStr 判断 str 是否为路径列表分隔符字符串。
func IsListSeparatorStr(str string) bool {
	return PathListSeparatorStr == str
}
