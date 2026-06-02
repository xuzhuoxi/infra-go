package osxu

import "os"

// ProcessExist reports whether a process with the given PID exists.
// ProcessExist 判断指定 PID 的进程是否已启动。
func ProcessExist(pid int) bool {
	_, err := os.FindProcess(pid)
	return nil == err
}
