package main

import (
	"fmt"
	"math"
)

/* SUMMARY
 * If statements
 * - Initializer syntax
 * - Comparison ops
 * - Logical ops
 * - Short circuiting
 * - If - else statements
 * - If - else if statements
 * - Equality and floats, error function
 *
 * Switch statements
 * - Switching on a tag
 * - A single case can have multiple tests unlike other languages
 * - Initializers
 * - Tagless syntax
 * - Implicit breaks
 * - "fallthrough"
 * - Type switches
 * - "break" out early (e.g. invalid data, shouldn't be saved to a DB)
 */
func main() {
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	// the "initializer" syntax
	if pop, ok := statePopulations["Florida"]; ok {
		// pop is only defined within the if block
		fmt.Printf("Florida: %v\n", pop)
	}

	// number guess example
	number := 50
	guess := 105
	// no single line ifs
	if guess < 1 {
		fmt.Println("The guess must be greater than 0!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than or equal to 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}
	fmt.Println()

	// equalities with floats
	myNum := .123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are same")
	} else {
		fmt.Println("These are different")
	}
	// better to use an error threshold
	if math.Abs(myNum-math.Pow(math.Sqrt(myNum), 2)) < 1e-9 {
		fmt.Println("These are same")
	} else {
		fmt.Println("These are different")
	}
	fmt.Println()

	/* switch/case
	 * can also use initializer syntax
	 * in Go, cases have implicit breaks
	 * need to use "fallthrough" to continue to the next case
	 */
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five, or ten")
	case 2, 4, 6:
		fmt.Println("two, four, or six")
	default:
		fmt.Println("another number")
	}
	// tagless syntax, cases can overlap. If they do, first one is picked
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
	fmt.Println()

	// type switching
	var j interface{} = [3]int{}
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float64")
	case string:
		fmt.Println("j is a string")
	case [3]int:
		fmt.Println("j is [3]int")
	default:
		fmt.Println("j is another type")
	}
	// NOTE: Arrays have to have the same data type AND the same length to be considered the same type
}
