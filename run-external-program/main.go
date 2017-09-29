package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-al")

	// 直接run的话是获取不到输出的
	// err := cmd.Run()
	// if err != nil {
	// 	panic(err)
	// }

	// CombinedOutput() 会运行并返回stderr/stdout混合后的输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))

}
