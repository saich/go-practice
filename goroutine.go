// goroutine
package main

import (
	"fmt"
	"runtime"
)

// A goroutine is a lightweight thread, but managed by the go runtime (i.e; they may not correspond to native threads)
// go f(x,y,z) starts a goroutine that runs f(x,y,z)

// The parameters will be executed in the current routine, at the "go" statement itself.

// Goroutines run in the same address space, so access to shared memeory must be synchronized.

// By default, all goroutines run in the same thread and Go's scheduler takes care of scheduling
// between go-routines where there is waiting on I/O or the code forces a switch.

// In future, the Go's scheduler shall use all available CPUs by default. But for now,
// We can ask the scheduler to use multiple CPUs by setting GOMAXPROCS

// Also, the go routine DOESN'T wait for all goroutines to finish. It finishes ant the pending ones may not
// execute at all..

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // To understand why this is here, read http://stackoverflow.com/questions/13107958/what-exactly-does-runtime-gosched-do
		fmt.Println(s)
	}
}

func main() {
	// runtime.GOMAXPROCS(4) // We can use multiple CPUs (basically, native threads) in which case the output will not be predictable (even without the Gosched())
	go say("world")
	say("hello")
}
