package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // 0 memory allocation to send signals (empty struct)

var wg = sync.WaitGroup{}

/* SUMMARY
 * Channel basics
 * - Create channel with make
 *   - make(chan int)
 * - Send message into channel
 *   - ch <- val
 * - Receive message from channel
 *   - val := <-ch
 * - Can have multiple senders and receivers
 *
 * Restricting data flow
 * - Channel can be cast into send/receive-only versions
 *   - Send-only: chan <- int
 *   - Receive-only: <-chan int
 *
 * Buffered channels
 * - Channels block sender side `til receiver is available
 * - Block receiver side 'til message is available
 * - Can decouple sender/receiver with buffered channels
 *   - make(chan int, 50)
 * - Use buffered channels when sender/receiver have asymmetric loading
 *
 * For...range loops with channels
 * - First parameter is the value itself, instead of the index in arrays, slices, maps
 * - Use to monitor channel and process messages as they arrive
 * - Loop exits when channel is closed
 *
 * Select statements
 * - Allows Goroutine to monitor several channels at once
 *   - Blocks if all channels block
 *   - If multiple channels receiver value simultaneously, behavior is undefined
 * - default causes the select statement to be non-blocking
 */
func main() {
	// Buffered channel
	ch := make(chan int, 50)
	wg.Add(2)

	// Receive only channel
	go func(ch <-chan int) {
		// receiving from channel
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		// OR use for range
		// for i := range ch {
		// 	fmt.Println(i)
		// }

		wg.Done()
	}(ch)

	// Send only channel
	go func(ch chan<- int) {
		// sending into channel
		ch <- 42
		ch <- 27
		// close channel to signify no more data being sent
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
	fmt.Println()

	/* Select statements
	 * Always know the exit strategy for a Goroutine before creating one,
	 * In this case, we can use "defer" to close the channel or
	 * use an empty struct to signify process complete.
	 *
	 * select with default is a non-blocking statement, without a default
	 * it becomes a blocking statement.
	 */
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	// for entry := range logCh {
	// 	fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	// }
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
