package main

import (
	"fmt"
	"log"
)

/* SUMMARY
 * Recover
 * - Used to recover from panics
 * - Only useful in deferred functions
 * - Current function will not attempt to continue,
 *   but functions higher up in the callstack will
 */
func main() {
	/* "recover" */
	// fmt.Println("start")
	// // anonymous function
	// defer func() {
	// 	// recover() returns nil if the app is NOT "panic"-ing
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()
	// panic("something bad happened")
	// fmt.Println("end") // unreachable
	// // NOTE: "defer" takes in function CALLS, not functions themselves. That's what the parentheses are for at the closing curly
	fmt.Println("start")
	panicker()
	// if panic is not rethrown, funcs higher up on the callstack can still continue
	fmt.Println("end")
}

func panicker() {
	// This function will recover, but it will not continue
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// rethrowing panic up the callstack if it can't be handled after recover()
			// panic(err)
		}
	}()
	panic("something bad happened")
	// fmt.Println("done panicking") // unreachable
}
