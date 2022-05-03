package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var increment int64

func main() {
	fmt.Println("Num of Goroutine:", runtime.NumGoroutine())

	gr := 50
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&increment, 1)
			fmt.Println("Counter", atomic.LoadInt64(&increment))
			fmt.Println("The Increment Count (Goroutine) ", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Go routine :", runtime.NumGoroutine())
	fmt.Println("Count", increment)
}
