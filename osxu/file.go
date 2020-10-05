package osxu

//
//type FileDetailInfo interface {
//	os.FileInfo
//	FullPath() string
//}
//
//type fileDetailInfo struct {
//	os.FileInfo
//	fullPath string
//}
//
//func (f *fileDetailInfo) FullPath() string {
//	return f.fullPath
//}

//
////检查路径是否存在
//func IsExist(path string) bool {
//	_, err := os.Stat(path)
//	return err == nil || os.IsExist(err)
//}
//
////是否为文件夹
//func IsFolder(path string) bool {
//	if !IsExist(path) {
//		return false
//	}
//	fi, err := os.Stat(path)
//	if nil != err {
//		return false
//	}
//	return fi.IsDir()
//}
//
////取文件或文件夹大小
//func GetSize(path string) (uint64, error) {
//	fi, err := os.Stat(path)
//	if nil != err {
//		return 0, err
//	}
//	if fi.IsDir() {
//		return GetFolderSize(path)
//	} else {
//		return GetFileSize(path)
//	}
//}
//
////取文件大小
//func GetFileSize(filePath string) (uint64, error) {
//	fi, err := os.Stat(filePath)
//	if nil != err {
//		return 0, err
//	}
//	if fi.IsDir() {
//		return 0, nil
//	}
//	return uint64(fi.Size()), nil
//}
//
////取文件夹大小，递归全部文件的大小之和
//func GetFolderSize(dirPath string) (uint64, error) {
//	list, err := GetFolderFileList(dirPath, true, nil)
//	if err != nil {
//		return 0, err
//	}
//	var size uint64 = 0
//	for _, file := range list {
//		size += uint64(file.Size())
//	}
//	return size, nil
//}
//
////取扩展名
//func GetExtensionName(fileName string) string {
//	_, eName := SplitFileName(fileName)
//	return eName
//}
//
////取文件名
//func GetFilePrefixName(fileName string) string {
//	bName, _ := SplitFileName(fileName)
//	return bName
//}

////取文件夹下全部文件
////recursive 是否递归子文件夹
////filter 过滤器，=nil时为不增加过滤,返回true时的FileInfo将包含到返回结果中
//func GetFolderFileList(dirPath string, recursive bool, filter func(fileInfo os.FileInfo) bool) ([]FileDetailInfo, error) {
//	dirPath = FormatDirPath(dirPath)
//	_, err := os.Stat(dirPath)
//	if nil != err {
//		return nil, err
//	}
//	var rs []FileDetailInfo
//	var recursiveFunc func(folderPath string) = nil
//	recursiveFunc = func(folderPath string) { //folderPath必须为"/"结尾
//		list, e := ioutil.ReadDir(folderPath)
//		if nil != e {
//			return
//		}
//		for _, file := range list {
//			if file.IsDir() {
//				if recursive {
//					recursiveFunc(folderPath + file.Name() + "/")
//				}
//			} else {
//				if nil == filter || filter(file) {
//					rs = append(rs, &fileDetailInfo{fullPath: folderPath + file.Name(), FileInfo: file})
//				}
//			}
//		}
//	}
//	recursiveFunc(dirPath)
//	return rs, nil
//}
//
//// 取路径的父目录
//// 如果当前路径为文件夹，返回上一级文件夹路径
//// 如果当前路径为文件，返回所在目录
//func GetParentDir(dirPath string) (dir string, ok bool) {
//	fileDir, _ := SplitFilePath(dirPath)
//	if "" == fileDir || "/" == fileDir {
//		return "", false
//	}
//	fileDir = stringx.SubString(fileDir, 0, len(fileDir)-1) //取掉最后一个"/"
//	dot := stringx.LastIndexOfChar(fileDir, "/")
//	// 已经是顶级目录
//	if -1 == dot {
//		return "", false
//	}
//	dir, _ = stringx.CutString(fileDir, dot+1, true)
//	return dir, true
//}
//
//// 取路径的当前目录
//// 如果当前路径是文件夹，返回当前路径
//// 如果当前路径是文件，返回所在目录
//func GetCurrentDir(fullPath string) (string, bool) {
//	fileDir, _ := SplitFilePath(fullPath)
//	if "" == fileDir {
//		return "", false
//	}
//	return fileDir, true
//}
//
//// 把文件名拆分
//// fileName不能包含目录
//// 路径已经转化为"/"格式
//func SplitFileName(fileName string) (fileBaseName string, fileExtName string) {
//	if "" == fileName || 0 == stringx.GetCharCount(fileName) {
//		return "", ""
//	}
//	dot := stringx.LastIndexOfChar(fileName, ".")
//	if -1 == dot {
//		return fileName, ""
//	}
//	return stringx.CutString(fileName, dot, false)
//}
//
//// 拆分目录路径与文件,目录路径以"/"结尾
//// 取绝对路径下对应的文件全名
//// 路径已经转化为"/"格式
//func SplitFilePath(fileFullPath string) (fileDir string, fileName string) {
//	fileFullPath = FormatPath(fileFullPath)
//	if IsFolder(fileFullPath) {
//		return FormatDirPath(fileFullPath), ""
//	}
//	dot := stringx.LastIndexOfChar(fileFullPath, "/")
//	if -1 == dot {
//		return "", fileFullPath
//	}
//	fileDir, fileName = stringx.CutString(fileFullPath, dot+1, true)
//	return
//}
//
//// 检查文件名的扩展名, 要求带"."
//func CheckExtensionName(fileName string, extensionName string) bool {
//	if len(fileName) < len(extensionName) {
//		return false
//	}
//	fs := []rune(fileName)
//	es := []rune(extensionName)
//	fIndex := len(fs) - 1
//	eIndex := len(es) - 1
//	for eIndex >= 0 {
//		if fs[fIndex] != es[eIndex] {
//			return false
//		}
//		eIndex -= 1
//		fIndex -= 1
//	}
//	return true
//}
//
////Format ------------------------------------
//
//// 标准化目录路径，以"/"结尾
//// 非"/"结尾补全
//// 已经转换为"/"形式路径
//// 不检测有效性
//func FormatDirPath(dirPath string) string {
//	fDirPath := FormatPath(dirPath)
//	dot := stringx.LastIndexOfChar(fDirPath, "/")
//	// 非"/"结尾
//	if dot != stringx.GetCharCount(fDirPath)-1 {
//		return fDirPath + "/"
//	}
//	return fDirPath
//}
//
//// 标准化路径
//// 转换为"/"形式路径
//// 不检测有效性
//func FormatPath(path string) string {
//	return strings.Replace(path, "\\", "/", -1)
//}
