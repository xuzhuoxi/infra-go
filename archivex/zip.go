// Package archivex
// Create on 2022/7/12
// @author xuzhuoxi
package archivex

import (
	"archive/zip"
	"bufio"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"io"
	"os"
)

// ArchiveMultiToZipRoot
// 把多个路径同时归档, 归档目录为根目录
func ArchiveMultiToZipRoot(filePaths []string, archPath string) error {
	return ArchiveMultiToZip(filePaths, archPath, "")
}

// ArchiveMultiToZip
// 把多个路径同时归档, 并指定归档目录
func ArchiveMultiToZip(filePaths []string, archPath string, dirHeaderName string) error {
	if len(filePaths) == 0 {
		return errors.New(fmt.Sprintf("ArchiveMultiToZip: filepaths is empty! "))
	}
	f, err := os.Create(archPath)
	if err != nil {
		return errors.New(fmt.Sprintf("ArchiveMultiToZip: create arch file[%s] fail! ", archPath))
	}
	defer f.Close()
	zipWriter := zip.NewWriter(f)
	defer zipWriter.Close()
	for _, filePath := range filePaths {
		if !filex.IsExist(filePath) {
			return errors.New(fmt.Sprintf("ArchiveMultiToZip: filePath[%s] is not exits.", filePath))
		}

		_, fileName := filex.Split(filePath)
		newHeaderName := filex.Combine(dirHeaderName, fileName)
		if filex.IsDir(filePath) {
			err = appendDirToZip(filePath, newHeaderName, zipWriter)
		} else {
			err = appendFileToZip(filePath, newHeaderName, zipWriter)
		}
		if nil != err {
			return err
		}
	}
	return nil
}

// ArchiveChildrenToZip
// 把目录内容进行归档
func ArchiveChildrenToZip(dirPath string, archPath string) error {
	return ArchiveToZip(dirPath, archPath, "")
}

// ArchiveToZipDefault
// 把文件或目录归档，使用原文件名或目录名代为HeaderName
func ArchiveToZipDefault(filePath string, archPath string) error {
	_, headerName := filex.Split(filePath)
	return ArchiveToZip(filePath, archPath, headerName)
}

// ArchiveToZip
// 把文件或目录归档，并指定新的HeaderName
// 可用于定制目录结构
func ArchiveToZip(filePath string, archPath string, newHeaderName string) error {
	if !filex.IsExist(filePath) {
		return errors.New(fmt.Sprintf("ArchiveToZip: filepath[%s] is not exist! ", filePath))
	}
	f, err := os.Create(archPath)
	if err != nil {
		return errors.New(fmt.Sprintf("ArchiveToZip: create arch file[%s] fail! ", archPath))
	}
	defer f.Close()
	zipWriter := zip.NewWriter(f)
	defer zipWriter.Close()
	if filex.IsFile(filePath) {
		return appendFileToZip(filePath, newHeaderName, zipWriter)
	} else {
		return appendDirToZip(filePath, newHeaderName, zipWriter)
	}
}

// AppendDirToZipRoot
// 追加目录到zip文件中
// 使用目录名作为HeaderName
// 保持目录内容结构不变
func AppendDirToZipRoot(filePath string, zipWriter *zip.Writer) error {
	_, newHeaderName := filex.Split(filePath)
	return appendDirToZip(filePath, newHeaderName, zipWriter)
}

// AppendDirToZip
// 追加目录到zip文件中
// 为目录指定新的HeaderName
// 保持目录内容结构不变
func AppendDirToZip(filePath string, newHeaderName string, zipWriter *zip.Writer) error {
	return appendDirToZip(filePath, newHeaderName, zipWriter)
}

// AppendChildrenToZipRoot
// 追加目录下列表zip文件中
// 不使用目录HeaderName
// 保持目录内容结构不变
func AppendChildrenToZipRoot(filePath string, zipWriter *zip.Writer) error {
	return appendChildrenToZip(filePath, "", zipWriter)
}

// AppendChildrenToZip
// 追加目录下列表zip文件中
// 为目录指定新的HeaderName
// 保持目录内容结构不变
func AppendChildrenToZip(dirPath string, dirHeaderName string, zipWriter *zip.Writer) error {
	return appendChildrenToZip(dirPath, dirHeaderName, zipWriter)
}

// AppendFileToZipRoot
// 追加文件到zip文件中
// 使用文件名作为HeaderName
func AppendFileToZipRoot(filePath string, zipWriter *zip.Writer) error {
	_, newHeaderName := filex.Split(filePath)
	return appendFileToZip(filePath, newHeaderName, zipWriter)
}

// AppendFileToZip
// 追加文件到zip文件中
// 为文件指定新的HeaderName
func AppendFileToZip(filePath string, newHeaderName string, zipWriter *zip.Writer) error {
	return appendFileToZip(filePath, newHeaderName, zipWriter)
}

// 把目录写入到zip, 并指定目录的headerName
func appendDirToZip(dirPath string, newHeaderName string, zipWriter *zip.Writer) error {
	return appendChildrenToZip(dirPath, newHeaderName, zipWriter)
}

// 把目录的内容写入到zip, 并指定目标目录的headerName为dirHeaderName
func appendChildrenToZip(dirPath string, dirHeaderName string, zipWriter *zip.Writer) error {
	files, err := filex.GetPathsAll(dirPath, func(path string, info os.FileInfo) bool {
		return !info.IsDir()
	})
	if nil != err {
		return errors.New(fmt.Sprintf("appendChildrenToZip[%s] Error[%s]. ", dirPath, err))
	}
	filePathLen := len(dirPath)
	for index := range files {
		relativeFilePath := files[index][filePathLen+1:]
		headerName := filex.Combine(dirHeaderName, relativeFilePath)
		err := appendFileToZip(files[index], headerName, zipWriter)
		if nil != err {
			return err
		}
	}
	return nil
}

// 把文件写入到zip中, 并指定headerName
func appendFileToZip(filePath string, newHeaderName string, zipWriter *zip.Writer) error {
	fileInfo, err := os.Stat(filePath)
	if nil != err {
		return errors.New(fmt.Sprintf("appendFileToZip[file is not exist][%s]", err))
	}

	srcFile, err := os.Open(filePath)
	if err != nil {
		return errors.New(fmt.Sprintf("appendFileToZip[Open][%s]", err))
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	hdr, _ := zip.FileInfoHeader(fileInfo)
	hdr.Name = newHeaderName
	subWriter, err := zipWriter.CreateHeader(hdr)
	if err != nil {
		return errors.New(fmt.Sprintf("appendFileToZip[CreateHeader][%s]", err))
	}

	_, err = io.Copy(subWriter, reader)
	if err != nil {
		return errors.New(fmt.Sprintf("appendFileToZip[Copy][%s]", err))
	}
	return nil
}
