package main

import "C"
import "fmt"

//export PrintVersion
func PrintVersion() int32 {
	fmt.Println("From DLL: v1.0!")
	return 12
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
