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

	err = os.MkdirAll(dst, os.ModeDir|0755)
	CheckError(err)

	for _, file := range files {
		if file.IsDir() == true {
			CopyDir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			continue
		}
		CopyFile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
	}
}

func ClearDir(path string, except []string) {
	files, err := ioutil.ReadDir(path)
	CheckError(err)

	for _, file := range files {
		isExcept := false
		for _, name := range except {
			if file.Name() == name {
				isExcept = true
				break
			}
		}
		if isExcept {
			continue
		}
		joinedPath := filepath.Join(path, file.Name())
		if file.IsDir() == true {
			os.RemoveAll(joinedPath)
		}
		os.Remove(joinedPath)
	}
}
