package filex

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type absEle struct {
	Path  string
	IsAbs bool
}

var absEles = []absEle{
	absEle{Path: `c:/website/index.htm`, IsAbs: true},
	absEle{Path: `c:/website/img/photo.jpg`, IsAbs: true},
	absEle{Path: `C:\windows\system32\cmd.exe`, IsAbs: true},
	absEle{Path: `system32\cmd.exe`, IsAbs: false},
	absEle{Path: `..\windows\system32\cmd.exe`, IsAbs: false},
	absEle{Path: `..\..\windows\system32\cmd.exe`, IsAbs: false},
	absEle{Path: `/home/user1/abc.txt`, IsAbs: true},
	absEle{Path: `home/user1/abc.txt`, IsAbs: false},
	absEle{Path: `./home/user1/abc.txt`, IsAbs: false},
	absEle{Path: `http://localhost/index.html`, IsAbs: true},
	absEle{Path: `\\192.168.2.20`, IsAbs: true},
	absEle{Path: `ftp://test:test@192.168.0.1:21/profile`, IsAbs: true},
}

func TestAbs(t *testing.T) {
	for _, e := range absEles {
		fmt.Println(e.Path, e.IsAbs, IsAbsFormat(e.Path))
	}
}

func TestFormatPath(t *testing.T) {
	paths := []string{"/", "/a", "/a\\", "/a/b", "/a\\b/"}
	for index, path := range paths {
		fmt.Println(index, ":", fmt.Sprint(FormatPath(path)))
	}
}

func TestSplit(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "/a/b", "/a/b/", "o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	for index, path := range paths {
		dir, name := Split(path)
		fmt.Println(index, ":", dir, ",", name)
	}
}

func TestGetUpDir(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "\\a\\b", "/a/b/"}
	for index, path := range paths {
		fmt.Println(index, ":", fmt.Sprintln(GetUpDir(path)))
	}
}

func TestSplitFileName(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "/a/b", "/a/b/", "o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	for index, path := range paths {
		shortName, _, ext := SplitFileName(path)
		fmt.Println(index, ":", shortName+" , "+ext)
	}
}

func TestCheckExtName(t *testing.T) {
	fileNames := []string{"o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	extNames := []string{".exe", ".exe", ".", "abc", "abc"}
	for index, fn := range fileNames {
		fmt.Println(CheckExt(fn, extNames[index]))
	}
}

func TestWalkCurrent(t *testing.T) {
	var current = FormatPath(filepath.Dir(os.Args[0]))
	var dir = current + "/source"
	WalkInDir(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

func TestCombine(t *testing.T) {
	dirs := []string{"/", "/a", "/a/", "/a/b", "/a/b/", "o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	adds := []string{"/", "/a", "/a/", "/a/b", "/a/b/", "o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	for index, dir := range dirs {
		fmt.Println(index, ":", fmt.Sprint(Combine(dir, adds[index])))
	}
}

func TestWalk(t *testing.T) {
	var current = FormatPath(filepath.Dir(os.Args[0]))
	var dir = current + "/source"
	WalkAll(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
