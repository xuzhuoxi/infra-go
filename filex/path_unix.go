// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package filex

func isPathSeparatorStr(str string) bool {
	return UnixSeparatorStr == str
}
