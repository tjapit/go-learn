package main

import (
	"fmt"
)

/* SUMMARY
 * Simple loops
 * - for initializer; test; incrementer {}
 * - for test {} --> while
 * - for {} --> run 'til "break"
 *
 * Exiting early
 * - break
 * - continue
 * - Label
 *
 * Looping over collections
 * - for k, v := range collection {}
 * **NOTE: Go is unforgiving with unused vars so...
 * - for k := ... 		--> only keys
 * - for _, v := ... 	--> only values
 */
func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
	fmt.Println()

	// NOTE: the increment operator (++) is a STATEMENT in Go, it's an EXPRESSION in others
	// for i, j := 0, 0; i < 5; i, j = <<i++, j++>> NOT ALLOWED
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// NOT RECOMMENDED manipulating counter
		// if i%2 == 0 {
		// 	i /= 2
		// } else {
		// 	i = 2*i + 1
		// }
	}
	fmt.Println()

	// without initializer, scoped to the main() block
	i := 2
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Println(i)
	fmt.Println()

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// exiting loop with complex logic
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}
	fmt.Println()

	// "continue"
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// breaking out of nested loop: Labels (kinda like ASM)
	// example, want to leave loop as soon as we get i*j >= 3
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
	fmt.Println()

	// looping thru collection
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	fmt.Println()

	s := "Oh hi"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
