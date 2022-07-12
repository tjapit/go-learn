package main

import "fmt"

/* SUMMARY
 * Basic syntax
 * - func foo() {
 * 	...
 * }
 *
 * Parameters
 * - Comma delimited list of variables and types
 *   - func foo(bar string, baz int)
 * - Parameters of same type list type once
 *   - func foo(bar, baz int)
 * - Function can change value in the caller when pointers are passed
 *   - This is always true for maps & slices
 * - Use variadic parameters to send list of same types in
 *   - Must be last parameter
 *   - Received as a slice
 *   - func foo(bar string, baz ...int)
 *
 * Return values
 * - Single return values just list type
 *   - func foo() int
 * - Multiple return value list types surrounded by parentheses
 *   - func foo() (int, error)
 *   - The (result type, error) paradigm is a common idiom in Go s.t.
 *     the callee doesn't have to determine what the application has to do
 *     if it does receive an error, it can defer to the judgement of the calling func
 * - Can use named return values
 *   - Initializes returned variable
 *   - Return using return keyword on its own
 * - Can return addresses of local variables
 *   - Automatically from local stack to shared heap memory
 *
 * Anonymous funcs
 * - Functions don't have names if they are:
 *   - Immediately invoked
 *   - Assigned to a variable or passed as an argument to a func
 *
 * Functions as types
 * - Can assign functions to variables or use as arguments and
 *   return values in functions
 * - Type signature is like function signature, with no parameter names
 *   - var f func(string, string, int) (int, error)
 *
 * Methods
 * - Function that xcutes in context of a type (doesn't have to be a struct)
 *   - Could be a type of integer and then add methods to that custom type
 * - Format
 *   - func (localName typeContext) funcName() {
 *	...
 *	}
 * - Receiver can be value or pointer
 *   - Value receiver gets copy of type
 *   - Pointer receiver gets pointer of type
 */
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

	// variadic
	sum("The sum is", 1, 2, 3, 4, 5)
	fmt.Println()
	// returning pointer from local call stack
	s := sumReturn(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)
	fmt.Println()
	// named return value on func signature
	s1 := sumNamedReturn(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s1)
	fmt.Println()

	// multiple return values
	d, err := divide(5., 3.)
	if err != nil {
		fmt.Println(err)
		// exit from calling func to avoid chains of if-elses
		return
	}
	fmt.Println(d)
	fmt.Println()

	// functions can be treated as types
	var div func(float64, float64) (float64, error)
	div = func(f1, f2 float64) (float64, error) {
		if f2 == 0. {
			return 0., fmt.Errorf("Cannot divide by zero")
		}
		return f1 / f2, nil
	}
	t, err := div(5., 3.)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
	fmt.Println()

	// methods
	g := greeter{
		greeting: "Oh hi",
		name:     "Tim",
	}
	g.greet1()
	g.greet2()
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

/* Methods are just functions + a known context
 * This one is a value receiver, the method is getting a copy
 * of the struct. Not the struct itself.
 */
func (g greeter) greet1() {
	fmt.Println(g.greeting, g.name)
}

// This one is a pointer receiver, we can manipulate the underlying data
func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = "Emily"
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

// multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
