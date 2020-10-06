package filex

import "errors"

var (
	ErrRootPath   = errors.New("Path is root. ")
	ErrNoUpperDir = errors.New("No upper dir. ")
)

var (
	ErrPathNotExist = errors.New("Path is not exist. ")
)
