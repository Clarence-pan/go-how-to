package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	log.Printf("[main] enter")

	for i := 0; i < 3; i++ {
		i := i
		wg.Add(1)
		go func() {
			log.Printf("[go%d] enter", i)

			time.Sleep(time.Duration(rand.Int()%100) * time.Millisecond)

			log.Printf("[go%d] leave", i)
			wg.Done()
		}()
	}

	log.Printf("[main] waiting")
	wg.Wait()

	log.Printf("[main] done.")
}

/************** OUTPUT **********************
2017/09/28 14:39:52 [main] enter
2017/09/28 14:39:52 [main] waiting
2017/09/28 14:39:52 [go2] enter
2017/09/28 14:39:52 [go0] enter
2017/09/28 14:39:52 [go1] enter
2017/09/28 14:39:52 [go2] leave
2017/09/28 14:39:52 [go1] leave
2017/09/28 14:39:52 [go0] leave
2017/09/28 14:39:52 [main] done.
*************************************************/
