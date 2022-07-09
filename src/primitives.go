package main

import (
	"fmt"
)

func main() {
	// zero value for a boolean is false
	var n bool
	fmt.Printf("%v, %T\n", n, n)

	// signed ints
	// int8 (2^8 divide by 2 for + and -) -128 - +127
	// int16 (2^16) -32,768 - +32,767
	// int32
	// int64
	var a int32 = 42
	fmt.Printf("%v, %T\n", a, a)
	// unsigned int
	// uint8 0-255
	// uint16 0-65,535
	// uint32 0-A lot
	b := uint32(a)
	fmt.Printf("%v, %T\n", b, b)

	// primitive bit ops
	c := 10             // 1010
	d := 3              // 0011
	fmt.Println(c & d)  // 0010 = 2
	fmt.Println(c | d)  // 1011 = 11
	fmt.Println(c ^ d)  // 1001 = 9 (Bitwise XOR)
	fmt.Println(c &^ d) // 1000 = 8	(Bitclear AND NOT)
	// bit shifts
	e := 8              // 2^3
	fmt.Println(e << 3) // 2^3 * 2^3 = 2^6 (Left shift == * 2^nShifts)
	fmt.Println(e >> 3) // 2^3 / 2^3 = 2^6 (Right shift == / 2^nShifts)

	// floats
	// float32 +/-1.18E-38 - +/-3.4E38
	// float64 +/-2.23E-308 - +/-1.8E308
	f := 3.14
	f = 13.7e72
	// f = 2.1E14 // saving file lowercases the E
	fmt.Printf("%v, %T\n", f, f)

	// complex
	var g complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", real(g), real(g))
	fmt.Printf("%v, %T\n", imag(g), imag(g))
	var h complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", h, h)

	// string (utf8)
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	// strings in GO are an alias for a byte (uint8)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	// string concats
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
	// convert to a "slice" of bytes
	t := []byte(s)
	fmt.Printf("%v, %T\n", t, t)

	// rune (utf32)
	r := 'a'
	// runes are int32s
	fmt.Printf("%v, %T\n", r, r)

}
