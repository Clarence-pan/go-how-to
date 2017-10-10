package main

import (
	"log"
	"time"
)

func main() {
	timeout := time.After(10 * time.Second)
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			log.Printf("tick!")
		case <-timeout:
			log.Printf("timeout!")
			return
		}
	}

}
