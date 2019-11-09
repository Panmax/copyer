package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var BufferSize = 1024 * 1014
var BlockSize = 1024 * 1014 * 100

func RecordSum(srcFile, dstFile *os.File) (err error) {
	sumFilePath := getSumPath(dstFile)
	err = os.Mkdir(filepath.Dir(sumFilePath), os.ModePerm)
	sumFile, err := os.Create(sumFilePath)
	if err != nil {
		return
	}
	defer sumFile.Close()

	_, err = sumFile.WriteString(GetFileSum(srcFile))
	return
}

func getSumPath(file *os.File) string {
	return fmt.Sprintf("%s/.copyer/%s.md5", filepath.Dir(file.Name()), filepath.Base(file.Name()))
}

// 每 100M 的前 1M 作 md5 然后整体再 md5 一次
func GetFileSum(file *os.File) string {
	_, err := file.Seek(0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	md5Builder := strings.Builder{}
	offset := 0
	for {
		buffer := make([]byte, BufferSize)
		if n, err := file.ReadAt(buffer, int64(offset*BlockSize)); n <= 0 || (err != nil && err != io.EOF) {
			if err != nil && err != io.EOF {
				log.Fatalln(err)
			}
			break
		}

		md5Builder.WriteString(fmt.Sprintf("%x", md5.Sum(buffer)))
		offset += 1
	}

	return fmt.Sprintf("%x", md5.Sum([]byte(md5Builder.String())))
}
