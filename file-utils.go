package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func getMD5Hash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func CompareFiles(file1, file2 string) (bool, error) {
	size1, err := getFileSize(file1)
	if err != nil {
		return false, err
	}
	size2, err := getFileSize(file2)
	if err != nil {
		return false, err
	}

	if size1 != size2 {
		return false, nil // 文件大小不同，内容必然不同
	}

	hash1, err := getMD5Hash(file1)
	if err != nil {
		return false, err
	}
	hash2, err := getMD5Hash(file2)
	if err != nil {
		return false, err
	}

	return hash1 == hash2, nil // 比较MD5哈希值
}
