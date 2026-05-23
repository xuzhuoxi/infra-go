package filex

import "os"

// PathFilter decides whether a path matches during walk or filter operations.
// PathFilter 用于在遍历或筛选时判断路径是否命中。
type PathFilter func(path string, info os.FileInfo) bool
