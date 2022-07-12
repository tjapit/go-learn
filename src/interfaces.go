package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

/* BEST PRACTICES
 * - Use many, small interfaces
 *   - Single method interfaces are some of the most powerful and flexible
 *     - io.Writer, io.Reader, interface{}
 * - Don't export interfaces for types that will be consumed
 * - Do export interfaces for types that will be used by package
 * - Design functions and methods to receive interfaces whenever possible
 *
 * NOTE: GO INTERFACE DIFFER FROM OTHER LANGUAGES
 * source: https://youtu.be/YS4e4q9oBaU?t=19495
 */

/* SUMMARY
 * - Basics
 *   - Defining an interface, define behavior
 *   - Implicit implementation, having a method that match the signature defined
 * - Interface composition (e.g. WriterCloser)
 * - Type conversion
 *   - _, ok syntax to check if the conversion succeeds
 * - Empty interfaces and type switches
 *   - every type in Go implements the empty interface
 *   - empty interfaces usually combined with type switches
 * - Implementing with values vs. pointers
 *   - method set of VALUE is all methods with VALUE receivers
 *   - method set of POINTER is all methods, REGARDLESS of receiver type
 */
func main() {
	// polymorphism
	var w Writer = ConsoleWriter{}
	/* We know how to call the Write() func because it's
	 * defined in the Writer interface. But the concrete
	 * type (i.e. ConsoleWriter) determines the
	 * implementation
	 */
	w.Write([]byte("Oh hi"))
	fmt.Println()

	/* Any type can have methods associated with it,
	 * doesn't have to be a struct
	 */
	myInt := IntCounter(0) // casting an integer to an IntCounter
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		w.Write([]byte(strconv.Itoa(inc.Increment())))
	}
	fmt.Println()

	/* Composing interfaces together */
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello friends of civilization, this is a test"))
	wc.Close()
	fmt.Println()

	/* Type conversion
	 * WriterCloser --> BufferedWriterCloser
	 * As a WriterCloser, wc doesn't have access to the buffer
	 * it's an interface, so it's oblivious to the implementation.
	 * Converting to BufferedWriterCloser with bwc however,
	 * gives us access to the buffer in the implementation.
	 */
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)
	fmt.Println()

	/* Invalid conversions, can happen in any number of ways:
	 * - missing methods
	 * - typo
	 * - pointer vs value passed into the interface
	 *
	 * With the _, ok syntax, we can avoid our code panicking
	 * because panics are expensive.
	 */
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
	fmt.Println()

	/* Empty interfaces
	 * Almost always going to be an intermediate step.
	 * Need to figure out what it actually is, before
	 * doing anything useful with it.
	 */
	var myObj interface{} = NewBufferedWriterCloser()
	if wc2, ok := myObj.(WriterCloser); ok {
		wc2.Write([]byte("Oh hi friends of civilization, this is a test"))
		wc.Close()
	}

	/* Type switches with empty interfaces
	 * Usually empty interfaces are used this way
	 * s.t. we can determine the type, and then perform
	 * further ops with the item.
	 */
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}
}

/* Go doesn't have explicit interface "implements",
 * instead we do it implicitly by polymorphing.
 */
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/* Interfaces defining behaviors of custom types */
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

/* Interface composition */
type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
