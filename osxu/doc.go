// Package osxu provides lightweight operating-system helpers for process checks and runtime executable paths.
// Package osxu 提供进程检测与运行时可执行文件路径等轻量级操作系统辅助能力。
//
// Created by xuzhuoxi on 2019-02-18.
// @author xuzhuoxi
//
// Capabilities:
//   - ProcessExist: check whether a process with the given PID is present
//   - GetRunningDir: return the directory of the running executable (os.Args[0])
//   - GetRunningExecFile: return the base name of the running executable
//
// 主要职责：
//   - ProcessExist：判断指定 PID 的进程是否存在
//   - GetRunningDir：返回当前运行可执行文件所在目录（基于 os.Args[0]）
//   - GetRunningExecFile：返回当前运行可执行文件的文件名
package osxu
