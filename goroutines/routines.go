// Package goroutines
// Created by RTT.
// Author: teocci@yandex.com on 2022-Apr-28
package main

// #cgo LDFLAGS: -L. -lroutines
// #include "routines.hpp"
import "C"
import "fmt"
import "time"
import "sync"

func cpuIntensive(n int) int {
	fmt.Println("[go] dispatching cpuIntensive(", n, ")")
	return int(C.cpu_intensive(C.int(n)))
}

func ioIntensive() int {
	fmt.Println("[go] dispatching ioIntensive")
	return int(C.io_intensive())
}

func dispatchCPUIntensive(wg *sync.WaitGroup, n int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		cpuIntensive(n)
	}()
}

func dispatchIOIntensive(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		ioIntensive()
	}()
}

func main() {
	// A dirty example launching a bunch of C++ calls
	tbeg := time.Now()
	var wg sync.WaitGroup
	dispatchCPUIntensive(&wg, 30)
	dispatchCPUIntensive(&wg, 35)
	dispatchCPUIntensive(&wg, 40)
	dispatchCPUIntensive(&wg, 45)
	dispatchCPUIntensive(&wg, 45)
	dispatchCPUIntensive(&wg, 50)
	dispatchCPUIntensive(&wg, 30)
	dispatchCPUIntensive(&wg, 35)
	dispatchCPUIntensive(&wg, 40)
	dispatchCPUIntensive(&wg, 45)
	dispatchCPUIntensive(&wg, 45)
	dispatchCPUIntensive(&wg, 50)
	dispatchCPUIntensive(&wg, 30)
	dispatchCPUIntensive(&wg, 35)
	dispatchCPUIntensive(&wg, 40)
	dispatchCPUIntensive(&wg, 45)
	dispatchCPUIntensive(&wg, 45)
	dispatchCPUIntensive(&wg, 50)
	dispatchIOIntensive(&wg)
	dispatchIOIntensive(&wg)
	dispatchIOIntensive(&wg)
	fmt.Println("[go] dispatched all tasks")
	wg.Wait()
	fmt.Println("[go] done (elapsed:", time.Since(tbeg).Seconds(), "seconds)")
}
