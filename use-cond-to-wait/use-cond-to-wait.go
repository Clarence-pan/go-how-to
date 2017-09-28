package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	done := make(chan bool)

	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	test := time.Now().Second()%2 == 0

	go func() {
		log.Printf("[go] enter and do something")

		for i := 3; i > 0; i-- {
			log.Printf("[go] %d", i)
			time.Sleep(time.Second)
		}

		log.Printf("[go] leave and send signal")
		cond.Signal()

		done <- true
	}()

	log.Printf("[main] lock cond")
	cond.L.Lock()

	log.Printf("[main] use the cond...")

	log.Printf("[main] test=%v", test)
	if test {
		log.Printf("[main] wait cond")
		cond.Wait()
		log.Printf("[main] wait end")
	}

	log.Printf("[main] use the cond...")

	cond.L.Unlock()
	log.Printf("[main] unlock cond")

	<-done
	log.Printf("[main] done")

}
