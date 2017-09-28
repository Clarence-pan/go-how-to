package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	log.Println("no_mutex_version:")
	no_mutex_version()

	log.Println("with_mutex_version:")
	with_mutex_version()
}

func with_mutex_version() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)

	n := 0

	for i := 0; i < 1000; i++ {
		// i := i
		wg.Add(1)
		go func() {
			mutex.Lock()
			n++
			// log.Printf("[go%d] n++ -> %d", i, n)
			mutex.Unlock()

			time.Sleep(time.Duration(rand.Int()%100) * time.Millisecond)

			mutex.Lock()
			n--
			// log.Printf("[go%d] n-- -> %d", i, n)
			mutex.Unlock()

			wg.Done()
		}()
	}

	// 如果不用mutex，最后的n可能非0
	// 使用mutex后，最后的n总是0
	wg.Wait()
	log.Printf("[main] n = %d", n)
}

func no_mutex_version() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	n := 0

	for i := 0; i < 1000; i++ {
		// i := i
		wg.Add(1)
		go func() {
			n++
			// log.Printf("[go%d] n++ -> %d", i, n)

			time.Sleep(time.Duration(rand.Int()%100) * time.Millisecond)

			n--
			// log.Printf("[go%d] n-- -> %d", i, n)

			wg.Done()
		}()
	}

	// 如果不用mutex，最后的n可能非0
	// 使用mutex后，最后的n总是0
	wg.Wait()
	log.Printf("[main] n = %d", n)
}
