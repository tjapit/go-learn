package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/* BEST PRACTICES
 * - Don't create Goroutines in libraries
 *   - Let consumer control concurrency
 * - When creating a Goroutine, know how it will end
 *   - Avoids subtle memory leaks
 * - Check for race conditions at compile time with the race flag
 *   - go run -race <app_name>
 */
var wg = sync.WaitGroup{}
var counter = 0
var mutex = sync.RWMutex{}

/* SUMMARY
 * - Creating Goroutines
 *   - Use go keyword in front of function call
 *   - When using anonymous functions, pass data as local variables
 *     to avoid race conditions
 * - Synchronization
 *   - Use sync.WaitGroup to wait for groups of Goroutines to complete
 *     - WaitGroup.Add(int)
 *     - WaitGroup.Done()
 *     - WaitGroup.Wait()
 *   - Use sync.Mutex and sync.RWMutex to protect data access
 * - Parallelism
 *   - By default, Go will use CPU threads equal to available cores (1 thread/core)
 *   - Change with runtime.GOMAXPROCS
 *   - More threads can increase performance, but too many can slow it down
 */
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

	/* Beware of race conditions */
	var msg = "Oh hi"
	go func() {
		// closure
		fmt.Println(msg)
	}()
	// race condition
	msg = "Bye"
	time.Sleep(100 * time.Millisecond)

	/* Solution: pass data into the Goroutine */
	msg = "Oh hi"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Bye"
	time.Sleep(100 * time.Millisecond)
	// NOTE: sleep is bad practice because we're binding the application's performance and its clock cycle to the IRL clock.

	/* Solution: use Wait Groups */
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

	/* Multiple Goroutines working on the same data */
	// without syncing
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
	fmt.Println()

	// syncing with Mutex (within the context of the Goroutines, still causes problems)
	counter = 0 // restarting counter
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHelloMutexInContext()
		go incrementMutexInContext()
	}
	wg.Wait()
	fmt.Println()

	// syncing w/Mutex (outside of the contexts of the Goroutines)
	counter = 0
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go sayHelloMutexOutContext()
		mutex.Lock()
		go incrementMutexOutContext()
	}
	wg.Wait()
	// NOTE: there is no concurrency here because the mutexes are forcing the data to be sync-ed and run in a single-threaded way. The potential benefits using Goroutine are not applicable in this example.
	fmt.Println()

	/* GOMAXPROCS
	 * Tuning variable, generally 1 thread/core at a minimum, but app performance
	 * can actually increase by increasing GOMAXPROCS beyond that value. Up to a
	 * point of diminishing returns.
	 */
	runtime.GOMAXPROCS(12)
	fmt.Printf("Threads : %v\n", runtime.GOMAXPROCS(-1))
}

func sayHelloMutexOutContext() {
	fmt.Printf("Oh hi #%v\n", counter)
	mutex.RUnlock()
	wg.Done()
}

func incrementMutexOutContext() {
	counter++
	mutex.Unlock()
	wg.Done()
}

func sayHelloMutexInContext() {
	mutex.RLock()
	fmt.Printf("Oh hi #%v\n", counter)
	wg.Done()
	mutex.RUnlock()
}

func incrementMutexInContext() {
	mutex.Lock()
	counter++
	wg.Done()
	mutex.Unlock()
}

func sayHello() {
	fmt.Printf("Oh hi #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
