//
//Created by xuzhuoxi
//on 2019-02-18.
//@author xuzhuoxi
//
package osxu

import "os"

//判断一个进程是否启动
func ProcessExist(pid int) bool {
	_, err := os.FindProcess(pid)
	return nil == err
}
