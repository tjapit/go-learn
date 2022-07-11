package main

import "fmt"

func main() {
	sayMessage("Oh hi")
	fmt.Println()
	greeting := "Hello"
	name := "Tim"
	// passing by pointers, really the only way to manipulate the original args passed
	sayGreeting(&greeting, &name)
	fmt.Println(name)
	// NOTE: passing in pointers is more efficient when working with large data structures, because usually args are passed in by value (i.e. copied whole as a separate entity, EXCEPT for maps and slices because they are naturally reference-typed). Pointers avoid the expensive copying cost.
	fmt.Println()

	sum("The sum is", 1, 2, 3, 4, 5)
	fmt.Println()
	s := sumReturn(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)
	fmt.Println()
	s1 := sumNamedReturn(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s1)

}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// syntax sugars, Go infers the type of the first param
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// variadic
// can only have 1, and has to be at the end
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	// values will act like a slice
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

// return values
func sumReturn(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	/* In other languages, as soon as the function finishes execution,
	 * its call stack is freed. (i.e. freeing the memory the function took up)
	 *
	 * Go instead will promote our return value to the shared memory (heap)
	 * s.t. we can do things like return a pointer from the local stack
	 */
	return &result
}

// named return value on the signature
func sumNamedReturn(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	// no need to name return value because it's already done on signature
	return
}
