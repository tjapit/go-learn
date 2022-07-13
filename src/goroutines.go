package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	/* Green Thread
	 * Most programming languages use OS threads. They are expensive
	 * to create and destroy that's why there are concepts like Thread
	 * Pools to avoid thread cycle and resource thrashing.
	 *
	 * Go uses Green Threads which creates an abstraction of a thread
	 * called a Goroutine. In the Go runtime there is a scheduler that
	 * maps the Goroutines onto the OS threads for peiods of time.
	 *
	 */

	// Beware of race conditions
	var msg = "Oh hi"
	go func() {
		// closure
		fmt.Println(msg)
	}()
	// race condition
	msg = "Bye"
	time.Sleep(100 * time.Millisecond)

	// Solution: pass data into the Goroutine
	msg = "Oh hi"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Bye"
	time.Sleep(100 * time.Millisecond)
	// NOTE: sleep is bad practice because we're binding the application's performance and its clock cycle to the IRL clock.

	// Solution: use Wait Groups
	msg = "Oh hi"
	// add #of Goroutine to synchronize to WaitGroup
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		// let the wait group know that it has completed xcution, it decrements the #of groups that the WaitGroup is waiting on
		wg.Done()
	}(msg)
	msg = "Bye"
	// wait for the Goroutine to be done xcuting
	wg.Wait()
	fmt.Println()

	// Multiple Goroutines working on the same data

}
