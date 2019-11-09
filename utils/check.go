package utils

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var BufferSize = 100000
var SubSize = 32

func RecordSum(srcFile, dstFile *os.File) (err error) {
	sumFilePath := getSumPath(dstFile)
	err = os.Mkdir(filepath.Dir(sumFilePath), os.ModePerm)
	sumFile, err := os.Create(sumFilePath)
	if err != nil {
		return
	}
	defer sumFile.Close()

	n, err := sumFile.WriteString(getSum(srcFile))
	log.Println(n)
	return
}

func getSumPath(file *os.File) string {
	return fmt.Sprintf("%s/.copyer/%s.md5", filepath.Dir(file.Name()), filepath.Base(file.Name()))
}

func getSum(file *os.File) string {
	stat, _ := file.Stat()
	size := stat.Size()
	log.Println(size)

	_, _ = file.Seek(0, 0)
	md5N := md5.New()

	bufferStart := make([]byte, BufferSize)
	n, err := file.Read(bufferStart)
	log.Println(n)
	log.Println(err)

	_, _ = file.Seek(size/2, 0)
	bufferMiddle := make([]byte, BufferSize)
	n, err = file.Read(bufferMiddle)
	log.Println(n)
	log.Println(err)

	_, _ = file.Seek(int64(-BufferSize), 2)
	bufferEnd := make([]byte, BufferSize)
	n, err = file.Read(bufferEnd)
	log.Println(n)
	log.Println(err)

	return fmt.Sprintf("%s:%s:%s",
		fmt.Sprintf("%x", md5N.Sum(bufferStart))[:SubSize],
		fmt.Sprintf("%x", md5N.Sum(bufferMiddle))[:SubSize],
		fmt.Sprintf("%x", md5N.Sum(bufferEnd))[:SubSize])
}
