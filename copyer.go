package main

import (
	"crypto/md5"
	"log"
	"os"
)

var BufferSize = 1024 * 10 * 10

func main() {
	filepath := "/Users/jiapan/Downloads/test.js"

	info, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	size := info.Size()
	log.Println(size)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)
	file.Read(buffer)

	md5N := md5.New()
	log.Printf("%x\n",md5N.Sum(buffer))



	offset, _:=file.Seek(int64(-size),  2)
	log.Println(offset)
	buffer = make([]byte, BufferSize)
	file.Read(buffer)

	log.Printf("%x\n",md5N.Sum(buffer))
}
