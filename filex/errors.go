package filex

import "errors"

var (
	ErrRootPath   = errors.New("Path is root. ")   // ErrRootPath indicates the path is already root. 路径已是根目录.
	ErrNoUpperDir = errors.New("No upper dir. ")   // ErrNoUpperDir indicates no parent directory exists. 不存在上级目录.
)

var (
	ErrPathNotExist = errors.New("Path is not exist. ") // ErrPathNotExist indicates the path does not exist. 路径不存在.
)
