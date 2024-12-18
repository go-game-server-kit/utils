package utils

import (
	"crypto/md5"
	"encoding/hex"
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
	return hex.EncodeToString(hash.Sum(nil)), nil
}
