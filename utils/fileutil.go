package utils

import (
	"os"
	"path"
)

//判断文件|文件夹是否存在
func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

//创建目录
func CreateDir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

//创建目录下及目录下的文件
func CreateFile(filepath string, perm os.FileMode) error {
	dir, _ := path.Split(filepath)
	err := os.MkdirAll(dir, perm)
	if err != nil {
		return err
	}
	_, err = os.Create(filepath)
	if err != nil {
		return err
	}
	return nil
}
