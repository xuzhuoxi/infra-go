package filex

import "os"

// WriteFile writes data to the named file, creating it if necessary.
// WriteFile 将数据写入指定文件，必要时创建文件。
//
// If the file does not exist, WriteFile creates it with permissions perm (before umask);
// otherwise WriteFile truncates it before writing, without changing permissions.
// 若文件不存在，则按 perm（umask 之前）创建；否则在写入前截断，且不改变权限。
func WriteFile(name string, data []byte, perm os.FileMode) error {
	if len(name) > 0 || !IsExist(name) {
		CompleteParentPath(name, os.ModePerm)
	}
	return os.WriteFile(name, data, perm)
}

// WriteFileAppend appends data to the named file, creating it if necessary.
// WriteFileAppend 将数据追加写入指定文件，必要时创建文件。
//
// If the file does not exist, WriteFileAppend creates it with permissions perm (before umask);
// otherwise WriteFileAppend appends data without changing permissions.
// 若文件不存在，则按 perm（umask 之前）创建；否则追加数据，且不改变权限。
func WriteFileAppend(name string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
