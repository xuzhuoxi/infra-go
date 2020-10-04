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

func TestFormatDirPath(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "/a\\b", "/a\\b/"}
	for index, path := range paths {
		fmt.Println(index, ":", fmt.Sprint(FormatDirPath(FormatPath(path))))
	}
}

func TestGetUpDir(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "\\a\\b", "/a/b/"}
	for index, path := range paths {
		fmt.Println(index, ":", fmt.Sprintln(GetUpDir(path)))
	}
}

func TestSplit(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "/a/b", "/a/b/", "o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	for index, path := range paths {
		dir, name := Split(path)
		fmt.Println(index, ":", dir, ",", name)
	}
}

func TestCheckExtName(t *testing.T) {
	fileNames := []string{"o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	extNames := []string{".exe", ".exe", ".", "abc", "abc"}
	for index, fn := range fileNames {
		fmt.Println(CheckExt(fn, extNames[index]))
	}
}

func TestGetShortName(t *testing.T) {
	paths := []string{"/", "/a", "/a/", "/a/b", "/a/b/", "o.exe2", "a.exe", "b", "c.abc", "d.abc"}
	for index, path := range paths {
		fmt.Println(index, ":", fmt.Sprint(GetShortName(path)))
	}
}

func TestWalkCurrent(t *testing.T) {
	var current = FormatDirPath(filepath.Dir(os.Args[0]))
	var dir = current + "/source"
	WalkCurrent(dir, func(path string, info os.FileInfo, err error) error {
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
	var current = FormatDirPath(filepath.Dir(os.Args[0]))
	var dir = current + "/source"
	Walk(dir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
