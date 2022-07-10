package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	/* Can be used for example for validation frameworks.
	 * But then the framework would need to do the
	 * heavy-lifting of parsing the tags.
	 */
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
