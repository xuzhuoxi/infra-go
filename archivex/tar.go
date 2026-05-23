package archivex

import (
	"archive/tar"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/xuzhuoxi/infra-go/filex"
)

// ArchiveMultiToTarRoot archives multiple paths into a tar using the archive root as the base path.
// ArchiveMultiToTarRoot 把多个路径同时归档，归档目录为根目录。
func ArchiveMultiToTarRoot(filePaths []string, archPath string) error {
	return ArchiveMultiToTar(filePaths, archPath, "")
}

// ArchiveMultiToTar archives multiple paths into a tar under dirHeaderName.
// ArchiveMultiToTar 把多个路径同时归档，并指定归档目录。
func ArchiveMultiToTar(filePaths []string, archPath string, dirHeaderName string) error {
	if len(filePaths) == 0 {
		return errors.New(fmt.Sprintf("ArchiveMultiToTar: filepaths is empty! "))
	}
	f, err := os.Create(archPath)
	if err != nil {
		return errors.New(fmt.Sprintf("ArchiveMultiToTar: create arch file[%s] fail! ", archPath))
	}
	defer f.Close()
	tarWriter := tar.NewWriter(f)
	defer tarWriter.Close()
	for _, filePath := range filePaths {
		if !filex.IsExist(filePath) {
			return errors.New(fmt.Sprintf("ArchiveMultiToTar: filePath[%s] is not exits.", filePath))
		}

		_, fileName := filex.Split(filePath)
		newHeaderName := filex.Combine(dirHeaderName, fileName)
		if filex.IsDir(filePath) {
			err = appendDirToTar(filePath, newHeaderName, tarWriter)
		} else {
			err = appendFileToTar(filePath, newHeaderName, tarWriter)
		}
		if nil != err {
			return err
		}
	}
	return nil
}

// ArchiveChildrenToTar archives all files under dirPath without the parent directory entry.
// ArchiveChildrenToTar 把目录内容进行归档。
func ArchiveChildrenToTar(dirPath string, archPath string) error {
	return ArchiveToTar(dirPath, archPath, "")
}

// ArchiveToTarDefault archives a file or directory using its original name as the tar entry header.
// ArchiveToTarDefault 把文件或目录归档，使用原文件名或目录名作为 HeaderName。
func ArchiveToTarDefault(filePath string, archPath string) error {
	_, headerName := filex.Split(filePath)
	return ArchiveToTar(filePath, archPath, headerName)
}

// ArchiveToTar archives a file or directory with a custom entry header name.
// ArchiveToTar 把文件或目录归档并指定新的 HeaderName，可用于定制目录结构。
func ArchiveToTar(filePath string, archPath string, newHeaderName string) error {
	if !filex.IsExist(filePath) {
		return errors.New(fmt.Sprintf("ArchiveToTar: filepath[%s] is not exist! ", filePath))
	}
	f, err := os.Create(archPath)
	if err != nil {
		return errors.New(fmt.Sprintf("ArchiveToTar: create arch file[%s] fail! ", archPath))
	}
	defer f.Close()
	tarWriter := tar.NewWriter(f)
	defer tarWriter.Close()
	if filex.IsFile(filePath) {
		return appendFileToTar(filePath, newHeaderName, tarWriter)
	} else {
		return appendDirToTar(filePath, newHeaderName, tarWriter)
	}
}

// AppendDirToTarRoot appends a directory to tar using the directory name as the entry header.
// AppendDirToTarRoot 追加目录到 tar 文件，使用目录名作为 HeaderName，保持目录内容结构不变。
func AppendDirToTarRoot(filePath string, tarWriter *tar.Writer) error {
	_, newHeaderName := filex.Split(filePath)
	return appendDirToTar(filePath, newHeaderName, tarWriter)
}

// AppendDirToTar appends a directory to tar with a custom entry header name.
// AppendDirToTar 追加目录到 tar 文件并为目录指定新的 HeaderName，保持目录内容结构不变。
func AppendDirToTar(filePath string, newHeaderName string, tarWriter *tar.Writer) error {
	return appendDirToTar(filePath, newHeaderName, tarWriter)
}

// AppendChildrenToTarRoot appends directory contents to tar without a parent directory header.
// AppendChildrenToTarRoot 追加目录下内容到 tar，不使用目录 HeaderName，保持目录内容结构不变。
func AppendChildrenToTarRoot(filePath string, tarWriter *tar.Writer) error {
	return appendChildrenToTar(filePath, "", tarWriter)
}

// AppendChildrenToTar appends directory contents to tar under dirHeaderName.
// AppendChildrenToTar 追加目录下内容到 tar 并为目录指定新的 HeaderName，保持目录内容结构不变。
func AppendChildrenToTar(dirPath string, dirHeaderName string, tarWriter *tar.Writer) error {
	return appendChildrenToTar(dirPath, dirHeaderName, tarWriter)
}

// AppendFileToTarRoot appends a file to tar using the file name as the entry header.
// AppendFileToTarRoot 追加文件到 tar，使用文件名作为 HeaderName。
func AppendFileToTarRoot(filePath string, tarWriter *tar.Writer) error {
	_, newHeaderName := filex.Split(filePath)
	return appendFileToTar(filePath, newHeaderName, tarWriter)
}

// AppendFileToTar appends a file to tar with a custom entry header name.
// AppendFileToTar 追加文件到 tar 并为文件指定新的 HeaderName。
func AppendFileToTar(filePath string, newHeaderName string, tarWriter *tar.Writer) error {
	return appendFileToTar(filePath, newHeaderName, tarWriter)
}

// 把目录写入到tar, 并指定目录的headerName
func appendDirToTar(dirPath string, newHeaderName string, tarWriter *tar.Writer) error {
	return appendChildrenToTar(dirPath, newHeaderName, tarWriter)
}

// 把目录的内容写入到tar, 并指定目标目录的headerName为dirHeaderName
func appendChildrenToTar(dirPath string, dirHeaderName string, tarWriter *tar.Writer) error {
	files, err := filex.GetPathsAll(dirPath, func(path string, info os.FileInfo) bool {
		return !info.IsDir()
	})
	if nil != err {
		return errors.New(fmt.Sprintf("appendChildrenToTar[%s] Error[%s]. ", dirPath, err))
	}
	filePathLen := len(dirPath)
	for index := range files {
		relativeFilePath := files[index][filePathLen+1:]
		headerName := filex.Combine(dirHeaderName, relativeFilePath)
		err := appendFileToTar(files[index], headerName, tarWriter)
		if nil != err {
			return err
		}
	}
	return nil
}

// 把文件写入到tar中, 并指定headerName
func appendFileToTar(filePath string, newHeaderName string, tarWriter *tar.Writer) error {
	fileInfo, err := os.Stat(filePath)
	if nil != err {
		return errors.New(fmt.Sprintf("appendFileToTar[file is not exist][%s]", err))
	}

	srcFile, err := os.Open(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("appendFileToTar[Open][%s]", err))
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	hdr, _ := tar.FileInfoHeader(fileInfo, "")
	hdr.Name = newHeaderName
	hdr.Format = tar.FormatGNU

	tarWriter.WriteHeader(hdr)
	_, err = io.Copy(tarWriter, reader)
	if err != nil {
		return errors.New(fmt.Sprintf("appendFileToTar[Copy][%s]", err))
	}
	return nil
}
