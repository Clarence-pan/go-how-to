package main

import (
	"log"
	"sync"
)

func main() {
	once := new(sync.Once)

	for i := 0; i < 10; i++ {
		once.Do(func() {
			log.Println("This will only print once.")
		})
	}
}
