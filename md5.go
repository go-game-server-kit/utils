package utils

import (
	"crypto/md5"
	"io"
	"os"
)

func Md5File(file string) (string, error) {
	fp, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fp); err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}
