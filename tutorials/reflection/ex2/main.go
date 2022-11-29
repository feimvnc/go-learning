package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = 3.14
	fmt.Printf("x: type = %T, value = %v\n", x, x)
	xx := x
	fmt.Printf("x: type = %T, value = %v\n", xx, xx)

	x = &struct{ name string }{}
	fmt.Printf("x: type = %T, value = %v\n", x, x)
	y := x
	fmt.Printf("x: type = %T, value = %v\n", y, y)

	var z interface{}
	z = 3.14
	t := reflect.TypeOf(z)
	v := reflect.ValueOf(z) // z.(<type>)
	fmt.Printf("z: type = %v, value = %v\n", t, v)

	zz := z
	fmt.Printf("zz: type = %v, value = %v\n", zz, zz)

	// kind of category
	fmt.Printf("t: type = %v, kind = %v\n", t, t.Kind())
	fmt.Printf("v: type = %v, kind = %v\n", v, v.Kind())

	printStructInfo(z)
	printStructInfo(Person{name: "alice"})

}

type Person struct {
	name    string
	address string
}

func printStructInfo(x interface{}) {
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Struct {
		fmt.Printf("not a struct type\n")
		return
	}

	n := t.NumField()
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		fmt.Printf("struct field %v: name: %v, type: %v\n", i, tt.Name, tt.Type)

	}

}
