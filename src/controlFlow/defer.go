package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* SUMMARY
 * Defer
 * - Delay execution of a statement until function exits
 * - Useful to group "open" and "close" functions together
 *   - Be careful in loops, might wanna explicitly handle
 * 	   without defer if a bunch of resources is open during execution
 * - Runs in LIFO (last-in, first-out)
 * - Arguments evaluated at time "defer" is EXECUTED,
 *   not at time of CALLED FUNCTION execution
 */
func main() {
	/* "defer" is executed in LIFO
	 * it executes after main() in this case
	 */
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	/* This pattern is often used
	 * getting a resource, error checking, and then closing.
	 *
	 * Why error check first instead of closing?
	 * because if we got an error from the request,
	 * closing an error would be a bad time
	 */
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	fmt.Println()

	/* "defer" takes the argument at the time the function
	 * is called, NOT at the time the called function
	 * is executed
	 */
	defer fmt.Println()
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
