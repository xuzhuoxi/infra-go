package filex

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

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
