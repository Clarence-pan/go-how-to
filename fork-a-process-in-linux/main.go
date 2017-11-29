package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	ws := &sync.WaitGroup{}
	name := ""

	pid, err := fork()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if pid < 0 {
		fmt.Printf("Error: failed to fork!\n")
		os.Exit(1)
	} else if pid > 0 {
		name = "parent"
		fmt.Printf("[parent]: forked process pid = %d\n", pid)
	} else {
		name = "child"
		fmt.Printf("[child]: this is child!\n")
	}

	for i := 0; i < 10; i++ {
		i := i
		ws.Add(1)
		go func() {
			defer ws.Done()
			fmt.Printf("[%s][%d]: in goroutine\n", name, i)
		}()
	}

	ws.Wait()
	fmt.Printf("[%s] done.\n", name)

	if pid > 0 {
		waitPid(pid)
	}
}
