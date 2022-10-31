package main

import "fmt"

// a variable includes a name, a type, a value
// a variable is stored somewhere in memory with an address
// e.g. var foo int = 10
// pointer p *int, can point to foo, or other variables

func main() {
	i, j := 2, 3
	fmt.Println(i, j)   // print int value
	fmt.Println(&i, &j) // print address of i, j by adding &, "address of"

	p := &i         // assign address of i to new pointer p
	fmt.Println(p)  // print address of i
	fmt.Println(*p) // * is operator, value at address , also is called deferencing
	// *int - * in front of type, whole thing is an pointer  type of int
	// *p - * in front of var, is operator return value p points to

	fmt.Printf("%T\n", p)
	*p = 22 // changing *p would change value of i
	fmt.Println(i)

	// use same pointer to point another var
	// you can share the var and update it in multiple places
	// more efficient than  copying the variable everytime you need
	p = &j      // assign p to address of j
	*p = *p / 3 // change value at address of j
	fmt.Println(j)
}

/*
memory allocation
when we try to execute code, a go routine is created
each go routine gets a stack of memory
whenever a go routine makes a function call, a part of memory
is allocated, and we called that a frame

go can only work inside the frame, can't get data from other frame or other stack
each frame is guaranteed immutability
it is safe variables not getting modified throughout the program
function by copy means frame value changes are only limited to the frame

pointers help to update specific address in main frame
which can access across stack boundary of frames

squareVal(v int) { }  // immutability, parameter is an int

squareAdd(p *int) { }  // efficiency, giving up safety fo immutability, parameter is an pointer

*/

/*
return value of m or &m, when heap allocation is needed

package main

import "fmt"

type person struct {
  name string
  age int
}

func initPerson() *person {
  m := person{name: "goname", age: 42}
  fmt.Printf("initPerson --> %p\n", &m)
  return &m
}

func main(){
  fmt.Println(initPerson())
  fmt.Printf("main --> %p\n", initPerson())
}

*/
