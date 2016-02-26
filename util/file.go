package util

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CopyFile(src string, dst string) {
	srcFile, err := os.Open(src)
	defer srcFile.Close()
	CheckError(err)
	dstFile, err := os.Create(dst)
	defer dstFile.Close()
	CheckError(err)

	_, err = io.Copy(dstFile, srcFile)
	CheckError(err)
}

func CopyDir(src string, dst string) {
	files, err := ioutil.ReadDir(src)
	CheckError(err)

	for _, file := range files {
		if file.IsDir() == true {
			continue
		}
		CopyFile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
	}
}
