package main

import (
	"fmt"
	"sync"
	"time"
)

type TMOCond struct {
	L  sync.Locker
	ch chan bool
}

func NewTMOCond(l sync.Locker) *TMOCond {
	return &TMOCond{ch: make(chan bool), L: l}
}

func (t *TMOCond) Wait() {
	t.L.Unlock()
	<-t.ch
	t.L.Lock()
}

func (t *TMOCond) WaitOrTimeout(d time.Duration) bool {
	tmo := time.NewTimer(d)
	t.L.Unlock()
	var r bool
	select {
	case <-tmo.C:
		r = false
	case <-t.ch:
		r = true
	}
	if !tmo.Stop() {
		select {
		case <-tmo.C:
		default:
		}
	}
	t.L.Lock()
	return r
}

func (t *TMOCond) Signal() {
	t.signal()
}

func (t *TMOCond) Broadcast() {
	for {
		// Stop when we run out of waiters
		//
		if !t.signal() {
			return
		}
	}
}

func (t *TMOCond) signal() bool {
	select {
	case t.ch <- true:
		return true
	default:
		return false
	}
}

// **** TEST CASES ****
func lockAndSignal(t *TMOCond) {
	t.L.Lock()
	t.Signal()
	t.L.Unlock()
}

func waitAndPrint(t *TMOCond, i int) {
	t.L.Lock()
	fmt.Println("Goroutine", i, "waiting...")
	ok := t.WaitOrTimeout(10 * time.Second)
	t.L.Unlock()
	fmt.Println("This is goroutine", i, "ok:", ok)
}

func main() {
	var m sync.Mutex
	t := NewTMOCond(&m)

	// // Simple wait
	// //
	// t.L.Lock()
	// go lockAndSignal(t)
	// t.Wait()
	// t.L.Unlock()
	// fmt.Println("Simple wait finished.")

	// Wait that times out
	//
	t.L.Lock()
	ok := t.WaitOrTimeout(100 * time.Millisecond)
	t.L.Unlock()
	fmt.Println("Timeout wait finished. Timeout:", !ok)

	// // Broadcast. All threads should finish.
	// //
	// for i := 0; i < 10; i++ {
	// 	go waitAndPrint(t, i)
	// }
	// time.Sleep(1 * time.Second)
	// t.L.Lock()
	// fmt.Println("About to signal")
	// t.Broadcast()
	// t.L.Unlock()
	// time.Sleep(10 * time.Second)
}
