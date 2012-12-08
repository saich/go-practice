// channels
package main

import (
	"fmt"
)

// Go encourages a different approach in which shared values are passed around on channels and, in fact, never 
// actively shared by separate threads of execution. Only one goroutine has access to the value at any given time.

// "Do not communicate by sharing memory; instead, share memory by communicating."

// Channels are a type that can be used send or receive data even between goroutines (without mutex / semaphores)
// By default, sends and receives are blocking until other side is ready.  This allows goroutines to synchronize
// without explicit locks or condition variables.

// ch <- v // send v to channel ch
// v <- ch // receive into v from channel ch

// The data flows in the direction of the arrow

func sum(a []int, ch chan int) {
	sum := 0
	for _, val := range a {
		sum += val
	}
	ch <- sum // send the sum through the channel

}

func main() {

	ch := make(chan int) // Create a channel of integers
	a := []int{7, 2, 8, -9, 4, 0}
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)
	x, y := <-ch, <-ch // receive the values from the channel
	fmt.Println(x, y)

	// Channels can be buffered. Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.

	c := make(chan int, 2)
	c <- 1
	c <- 2
	// c <- 3 // adding this will casue a deadlock - "throw: all goroutines are asleep - deadlock!" since the buffer is already full and read happens only after this write..
	fmt.Println(<-c)
	fmt.Println(<-c)

	// Like we say earlier, a single channel can be used to send multiple messages.
	// A sender can close the channel to indicate that no more values will be sent.
	// Receiver can test whether a channel has been closed by assigning a second parameter to the receive expressioon:
	//		v, ok := <-ch
	// ok is false if there are no more values to receive and the channel is closed.

	// The loop i := range ch receives values from the channel repeatedly until the channel is closed

	// NOTE: Only the sender must close the channel, never a receiver. Sending on a closed channel will cause panic.

	// NOTE: Channels are not like files - closing channels is not always required. Closing is required only when the reciever must be told
	// there are no more values coming, such as like terminating the range loop

	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	ch2 := make(chan int, 10)
	go fibonacci(cap(ch2), ch2)
	fmt.Println("Fibonacci Numbers:")
	for val := range ch2 {
		fmt.Println(val)
	}

	// SELECT: The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. 
	// It chooses one at random if multiple are ready.

	// As you can see from this function, the select statement works both for sending as receivng data.
	fibonacci2 := func(c, quit chan int) {
		x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("Quitting...")
				return
			}
		}
	}
	fibs, quit := make(chan int), make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-fibs)
		}
		quit <- 0
	}()
	fmt.Println("fibonacci numbers set 2:")
	fibonacci2(fibs, quit)

	// TODO: default case in select... Not clear... Read more on this..
}
