package main

import (
	"fmt"
)

/* SUMMARY
 * Arrays
 * - Collection of items with same type
 * - Fixed size
 * - Declaration styles
 * 	 - a := [3]int{1, 2, 3}
 * 	 - a := [...]int{1, 2, 3}
 * 	 - var a [3]int
 * - Access is zero-based index
 * - len() returns size of array, cap() returns capacity
 * - Arrays are value-based, copies refer to different data
 *
 * Slices
 * - Backed by array
 * - Creation styles
 *   - Slice existing array or slice
 *   - Literal
 *   - make()
 *     - a := make([]int, 10, 100) // length 10, cap 100, all zeros
 * - len() returns length of slice
 * - cap() returns length of underlying array
 * - append() adds elements to slice
 *   - BUT, may cause expensive copy ops if underlying array is too small
 * - Slices are reference-based, copies refer to the same array
 */
func main() {
	// initializing array literal with just enough space (...)
	grades := [...]int{97, 85, 83}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Println()

	// declaring an empty string array
	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Tim"
	students[2] = "Jerome"
	students[1] = "Ahmed"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("#of Students: %v\n", len(students))
	fmt.Println()

	// 2D matrix
	var eye [3][3]int
	eye[0] = [3]int{1, 0, 0}
	eye[1] = [3]int{0, 1, 0}
	eye[2] = [3]int{0, 0, 1}
	fmt.Printf("Identity: %v\n", eye)
	fmt.Println()

	/* Arrays are considered values,
	 * copying does not pass pointers around,
	 * it actually generates a new array with the same elements
	 */
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Printf("a: %v", a)
	fmt.Printf("b: %v\n", b)
	// BUT, we can pass pointers around if we assign the pointer instead (&)
	c := &a
	c[1] = 4
	// we see that changing 2nd value in c changed the 2nd value in a
	fmt.Printf("a: %v", a)
	fmt.Printf("c: %v\n", *c) // dereference with *
	fmt.Println()

	/* Slices are naturally reference types, unlike arrays in Go
	 * copying passes the pointer around, therefore points to the same data
	 */
	s := []int{1, 2, 3}
	fmt.Printf("s: %v\n", s)
	fmt.Printf("Length(s): %v\n", len(s))
	fmt.Printf("Capacity(s): %v\n", cap(s))
	t := s
	t[1] = 5
	fmt.Printf("s: %v\n", s)
	fmt.Printf("t: %v\n", t)
	fmt.Println()

	// Actual slicing
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := s1[:]   // slice of all elements
	s3 := s1[3:]  // slice from 4th element to end
	s4 := s1[:6]  // slice first 6 elements
	s5 := s1[3:6] // slice from index 3-5, [inclusiveStart:xclusiveEnd]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println()

	// bulit-in make() to make slices
	sMake := make([]int, 3, 100) // (type, length, capacity)
	fmt.Printf("sMake: %v\n", sMake)
	fmt.Printf("Length(sMake): %v\n", len(sMake))
	fmt.Printf("Capacity(sMake): %v\n", cap(sMake))
	fmt.Println()

	// Arrays have fixed size. Slices do not
	slice := []int{}
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("Length(slice): %v\n", len(slice))
	fmt.Printf("Capacity(slice): %v\n", cap(slice))
	slice = append(slice, 1)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("Length(slice): %v\n", len(slice))
	fmt.Printf("Capacity(slice): %v\n", cap(slice))
	slice = append(slice, 2, 3, 4, 5)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("Length(slice): %v\n", len(slice))
	fmt.Printf("Capacity(slice): %v\n", cap(slice))
	fmt.Println()

	// concatenating slices
	// slice = append(slice, []int{6, 7, 8})
	// NOTE: the syntax above doesn't wokr, but we can use JS analogue of a spread operator to make it work
	slice = append(slice, []int{6, 7, 8}...)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("Length(slice): %v\n", len(slice))
	fmt.Printf("Capacity(slice): %v\n", cap(slice))
	fmt.Println()

	// removing elements
	// beginning
	sliceBegin := slice[1:]
	fmt.Printf("sliceBegin: %v\n", sliceBegin)
	fmt.Printf("original slice: %v\n", slice)
	fmt.Println()
	// end
	sliceEnd := slice[:len(slice)-1]
	fmt.Printf("sliceEnd: %v\n", sliceEnd)
	fmt.Printf("original slice: %v\n", slice)
	fmt.Println()
	// middle (e.g. removing element from index 3)
	sliceMid := append(slice[:3], slice[4:]...)
	fmt.Printf("sliceMid: %v\n", sliceMid)
	fmt.Printf("original slice: %v\n", slice)
	// NOTE: careful with append ops because it's using the same slice pointer, therefore modifying the original slice
	fmt.Println()

}
