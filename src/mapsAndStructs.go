package main

import (
	"fmt"
)

/* SUMMARY
 * Maps
 * - Collection of key:value pairs
 * - Created by literals or make()
 * - Access by keys: map["key"] = "value"
 * - Check for presence by "value, ok" = map["key"]
 * - Map is reference typed
 *
 * Structs
 * - Collection of different data types, kinda like Objects
 * - Keyed by named fields
 * - Normally created as types, but can do anonymous structs (usually short-lived)
 * - Struct is value typed
 * - No inheritance, but can do composition with embedding
 * - Tags can be added to struct fields to describe field
 */

// struct
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	// declaring empty map with make()
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	fmt.Println(statePopulations)

	// add
	statePopulations["Georgia"] = 10310371
	fmt.Printf("Georgia: %v\n", statePopulations["Georgia"])
	fmt.Println()

	// delete
	delete(statePopulations, "Georgia")
	// getting value by key returns 0
	fmt.Printf("Georgia: %v\n", statePopulations["Georgia"])

	// use comma-ok to find out if key is actually in the map
	_, ok := statePopulations["Georgia"]
	fmt.Printf("keyExists: %v\n", ok)
	fmt.Println()

	// maps are passed by reference
	sp := statePopulations
	delete(sp, "New York")
	fmt.Println(sp)
	fmt.Println(statePopulations)
	fmt.Println()

	// struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[2])
	fmt.Println()

	// anonymous struct
	aPokemon := struct{ name string }{name: "Bulbasaur"}
	bPokemon := aPokemon
	bPokemon.name = "Charmander"
	fmt.Println(aPokemon)
	fmt.Println(bPokemon)
}
