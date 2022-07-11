package main

import "fmt"

/* SUMMARY
 * Creating pointers
 * - Pointer types use an asterisk (*) as a prefix to type poined to
 *   - *int - a pointer to an integer
 * - Use the addressof op (&) to get address of variable
 *
 * Dereferencing pointers
 * - Dereference a pointer by preceding with an asterisk (*)
 * - Complex types (e.g. structs) are automatically dereferenced
 *
 * Create pointers to objects
 * - Can use the addressof op (&) if value type already exists
 *   - ms := myStruct{foo: 42}
 *   - p := &ms
 * - Use addressof op before initializer
 *   - &myStruct{foo:42}
 * - Use the "new" keyword
 *   - Can't initialize fields at the same time
 *
 * Types with internal pointers
 * - All assignment ops in Go are copy operations
 * - Slices and maps contain internal pointers,
 *   so copies point to same underlying data
 */
func main() {
	// address of (&), pointer declaration, dereferencing (*)
	var a int = 42
	var b *int = &a
	fmt.Printf("a: %v, b: %v\n", a, *b)
	fmt.Printf("addr_a: %p, addr_b: %p\n", &a, b)
	fmt.Printf("%p, %T\n", b, b)
	a = 27
	fmt.Printf("a: %v, b: %v\n", a, *b)
	*b = 69
	fmt.Printf("a: %v, b: %v\n", a, *b)
	fmt.Println()

	// NO pointer arithmetic, unless you import "unsafe"
	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	g := &d[2]
	fmt.Printf("%v %p %p %p\n", d, e, f, g)
	fmt.Println()

	// no need to declare the underlying data type first in order to create a pointer
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms) // ms is holding the address of an object that has a field with value 42 in it

	// the "new" keyword initializes a variable to a pointe to an object
	ms2 := new(myStruct)
	fmt.Println(ms2)
	(*ms2).foo = 42
	// NOTE: dereference op (*) has a lower precedence than dot op (.)
	fmt.Println((*ms2).foo)
	fmt.Println(ms2.foo)
	// NOTE: Go interprets ms2.foo as dereferencing the field within the data pointed by the pointer. The pointer doesn't have any fields, it just points to the underlying object that actually has the field
	fmt.Println()

	// zero value for a pointer when it is declared: <nil>
	var ms3 *myStruct
	fmt.Println(ms3)
	ms3 = new(myStruct)
	fmt.Println(ms3)

}

type myStruct struct {
	foo int
}
