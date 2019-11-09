package utils

import (
	"io"
	"os"
)

func Copy(srcFile, dstFile *os.File) error {
	buf := make([]byte, 4096)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := dstFile.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

func ChModifyTime(srcFile, dstFile *os.File) error {
	srcStat, _ := srcFile.Stat()
	return os.Chtimes(dstFile.Name(), srcStat.ModTime(), srcStat.ModTime())
}

func Chmod(srcFile, dstFile *os.File) error {
	srcStat, _ := srcFile.Stat()
	return dstFile.Chmod(srcStat.Mode())
}
