package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	/* No inheritance in Go, instead composition.
	 * We "embed" Animal inside the bird struct.
	 */
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// declaring and then filling in details, no need to worry about embedding
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Animal)
	fmt.Println(b.Name)
	fmt.Println()

	// declaring with literal structure, need to be aware of embedding
	c := Bird{
		Animal:   Animal{Name: "Penguin", Origin: "Antarctica"},
		SpeedKPH: 10,
		CanFly:   false,
	}
	fmt.Println(c)
	fmt.Println(c.Animal)
	fmt.Println(c.Name)

}
