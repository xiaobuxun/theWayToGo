package main

import (
	"fmt"
	"io"
	"os"
)

func filecopy(srcFile, dstFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open file %s file.\n", srcFile)
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFile, os.O_WRONLY | os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file %s file.\n", dstFile)
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	filecopy("test.txt", "copy.txt")
	fmt.Println("copy done.")
}