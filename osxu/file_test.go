package osxu

import (
	"fmt"
	"testing"
)

func TestGetParentDir(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "/a/b", "/a/b/"}
	for index, path := range paths {
		fmt.Println(index, ":", fmt.Sprintln(GetParentDir(path)))
	}
}

func TestGetFolderFileList(t *testing.T) {
	path := "D://VMOS//Windows 7 x64//"
	list, _ := GetFolderFileList(path, true, nil)
	for _, e := range list {
		fmt.Println(e.Name())
	}
}

func TestGetFolderSize(t *testing.T) {
	path := "D://VMOS//Windows 7 x64//"
	fmt.Println(GetFolderSize(path))
}
