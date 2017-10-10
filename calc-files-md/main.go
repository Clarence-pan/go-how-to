package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := os.Args[0]
	fmt.Printf("Processing file: %s\n", filepath)

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	h := md5.New()
	written, err := io.Copy(h, f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File Length: %d\n", written)

	md5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("File MD5: %s\n", md5)
}

/** output: *****************************************************
Processing file: ***\go-build118277518\command-line-arguments\_obj\exe\main.exe
File Length: 1335808
File MD5: b5ed52f871ceb36da4442ad71564d594
**************************************************************/
