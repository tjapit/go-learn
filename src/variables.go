package main

import (
	"fmt"
	"strconv"
)

/* Variable name convention
- short and to the point, within block scope it's fine to have short names (i, j, k) which might be used as flags or counters, etc.
- camelCase and more verbose/longer names can be used in pkg level or longer lifespans
- acronyms should be all caps (URL, HTTP, TCP, etc.) for readability
*/

// variable on the pkg level shadowed by the one in main function
var i int = 30

/* only 3 levels of visibility for variables */
// lowercase at pkg level: scoped to the pkg (protected), any file in the same pkg can access it
var j int = 39

// uppercase at pkg level: globally visible (public), exported from pkg
var J int = 39

// lowercase in a block: only visible within the block

/* variable block */
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// declaring variables w/out initialization
	var i int
	i = 42
	// explicit initialization of variables
	var j int = 27
	// letting Go infer variable type
	k := 69
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// explicit type conversion/casting variables
	var l int = 42
	fmt.Printf("%v, %T\n", l, l)

	var m float32
	m = float32(l)
	fmt.Printf("%v, %T\n", m, m)
	// NOTE: Go won't let you cast the other way round (i.e. float32 -> int) because it's going to lose precision

	// converting int -> string
	var n string
	n = string(l) // NOTE: returns ASCII 42 (*) instead of the integer 42
	fmt.Printf("%v, %T\n", n, n)
	n = strconv.Itoa(l)
	fmt.Printf("%v, %T\n", n, n)

}
