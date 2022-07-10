package main

import (
	"fmt"
)

/* SUMMARY
* Immutable
* Can be shadowed
* Replaced by compiler at compile time (must be calculable at compile)
* Named like variables (PascalCase for exports, camelCase for internal)
* Typed constants work like immutable variables (can interoperate only with same type)
* Infer typed constants work like literals (can interoperate with similar types)

* Enumerated constants (iota)
	* allows easy-creation of related constants,
	* starts from 0
	* be careful of constant values that match zero values for variables
* Enumerated expressions
	* Ops that can be calculated at compile are allowed
		* Arithmetic
		* Bitwise ops
		* Bitshifting
*/
// iota/enumerated constants
const (
	_  = iota // ignore first value (0) by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

// setting bool flags in a single byte
const (
	// setting pattern
	isAdmin          = 1 << iota // 0000 0001
	isHeadquarters               // 0000 0010
	canSeeFinancials             // 0000 0100

	canSeeAfrica       // 0000 1000
	canSeeAsia         // 0001 0000
	canSeeEurope       // 0010 0000
	canSeeNorthAmerica // 0100 0000
	canSeeSouthAmerica // 1000 0000
)

// pkg level, typed constant
const myConst int16 = 27

func main() {
	// block scoped, typed constant
	const myConst int = 42
	// constants can be shadowed
	fmt.Printf("%v, %T\n", myConst, myConst)
	// NOTE: constants can't be something that has to be determined at runtime, so it must be calculable at compile
	// e.g. NO const myConst float64 = math.Sin(1.57)

	// constant with inferred typing
	const a = 42
	var b int16 = 27
	fmt.Printf("%v, %T\n", a+b, a+b)
	// NOTE: this works because the compiler is replacing every instance of the constant with a "literal" and it's interpreted as int16 as it's being added (implicit conversion)

	// filesize example
	fileSize := 4e9
	fmt.Printf("%.2fGB\n", fileSize/GB)

	// user roles/access rights example
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)            // bitmask
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters) // bitmask
}
