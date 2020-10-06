package filex

import "os"

type PathFilter func(path string, info os.FileInfo) bool
