package filex

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFolderSize(t *testing.T) {
	base := filepath.Dir(os.Args[0])
	path := base + "/source"
	fmt.Println(GetFolderSize(path))
}
