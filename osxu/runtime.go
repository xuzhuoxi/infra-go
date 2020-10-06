package osxu

import (
	"os"
	"path/filepath"
)

//// 运行时的当前目录
//// 以"/"结尾
//func RunningBaseDir() string {
//	return FormatPath(filepath.Dir(os.Args[0]))
//}
//
//// 运行时的可执行文件名
//func RunningExecFile() string {
//	_, file := SplitFilePath(os.Args[0])
//	return file
//}

func GetRunningDir() string {
	return filepath.Dir(os.Args[0])
}

func GetRunningExecFile() string {
	_, name := filepath.Split(os.Args[0])
	return name
}
