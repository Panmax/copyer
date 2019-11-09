package main

import (
	"copyer/utils"
	"log"
	"os"
)

func main() {
	src := "/tmp/copyer/src.md"
	dst := "/tmp/copyer/dst.md"

	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	utils.Chmod(srcFile, dstFile)
	utils.Copy(srcFile, dstFile)
	utils.ChModifyTime(srcFile, dstFile)
	utils.RecordSum(srcFile, dstFile)

}
