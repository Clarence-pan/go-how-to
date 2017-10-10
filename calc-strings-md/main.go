package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	text := "This text will be MD5 hashed."

	h := md5.New()
	h.Write([]byte(text))

	md5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("Text MD5: %s\n", md5)
}

/** output: *****************************************************
Text MD5: 787343302e1052aba4ef9cb1f7071654
**************************************************************/
