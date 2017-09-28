package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	log.Println("no_atomic_version:")
	no_atomic_version()

	log.Println("with_atomic_version:")
	with_atomic_version()
}

func with_atomic_version() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	n := int32(0)

	for i := 0; i < 1000; i++ {
		// i := i
		wg.Add(1)
		go func() {
			atomic.AddInt32(&n, 1)

			time.Sleep(time.Duration(rand.Int()%100) * time.Millisecond)

			atomic.AddInt32(&n, -1)

			wg.Done()
		}()
	}

	// 如果不用atomic.xxx，最后的n可能非0
	// 使用atomic.xxx后，最后的n总是0
	wg.Wait()
	log.Printf("final n = %d", n)
}

func no_atomic_version() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	n := 0

	for i := 0; i < 1000; i++ {
		// i := i
		wg.Add(1)
		go func() {
			n++

			time.Sleep(time.Duration(rand.Int()%100) * time.Millisecond)

			n--

			wg.Done()
		}()
	}

	// 如果不用atomic.xxx，最后的n可能非0
	// 使用atomic.xxx后，最后的n总是0
	wg.Wait()
	log.Printf("final n = %d", n)
}
