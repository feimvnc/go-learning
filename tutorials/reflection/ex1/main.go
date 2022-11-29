package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {

	// pass ...any which is ...interface{}, allows any types
	// func Println(a ...any) (n int, err error)
	//
	Println(true)
	Println(10)
	Println(3.14)
	Println(7 + 7i)
	Println(ID("12345"))
	Println([5]byte{})
	Println([]byte{})
	Println(map[string]int{})
	Println(Person{name: "alice"})
	Println(&Person{name: "bob"}) // print &{bob}
	Println(make(chan int))

	var foo interface{}
	Println(foo) // type is '<nil>', value: <nil>
	foo = 3.14
	Println(foo) // type is 'float64', value: 3.14
	foo = &Person{name: "charlie"}
	Println(foo)  // type is '*main.Person', value: &{charlie}
	Println(&foo) // type is '*interface {}', value: 0xc000014280

	Println2(7 + 7i)
	Println2(ID("12345"))
	Println2([5]byte{})

	Println3(true)
	Println3(10)
	Println3(3.14)
}

func Println(x interface{}) {
	fmt.Printf("type is '%T', value: %v\n", x, x)
}

func Println2(x interface{}) {
	if v, ok := x.(ID); ok { // check if x.value has type ID
		fmt.Printf("x has type ID, value is %v\n", v)
	} else {
		fmt.Printf("'%T' is not type ID\n", x)
	}
}

func Println3(x interface{}) { // switch based onn types
	switch x.(type) { // check type of being stored in interface{}
	case bool:
		fmt.Println("boolean type: ", x.(bool))
	case int:
		fmt.Println("int type: ", x.(int))
	default:
		fmt.Println("unknown type")
	}
}
