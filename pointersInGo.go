package main

import "fmt"

type car struct {
	make 	string
	model	string
	year	int
	color	string
}

func main() {
	// The pointer is an address in the computer's memory where the value is saved
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	// The type of the address is different from the type of the value
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// The operator * has two distict uses in Go
	// It can be used before a type to indicate a pointer type
	// It can also be used before a variable that store a pointer to reveal the value at the pointed address
	var b *int = &a
	fmt.Printf("the value of b: %v\n", b)
	fmt.Printf("the value at b: %v\n", *b)

	// The above can be demonstrated by eliminating b all together
	fmt.Println(*&a)

	// Once you have an address you can manipulate the value stored by using the dereferencing operator *
	*b = 1984
	fmt.Println(a)
	fmt.Printf("the value at b: %v\n", *b)

	// Here is a practical example why you might want to use pointers
	// The value of x is changed in foo and that change persists after foo is done executing
	x := 12
	foo(&x)
	fmt.Println(x)

	// When dereferencing a field of a struct that a passed pointer points to we use parenthesis around the struct identifier.
	// This is to make sure the compier doesn't confuse (*x).field with *(x.field).
	c := car {
		make: "Aston Martin",
		model: "DB5",
		year: 1963,
		color: "silver",
	}
	fmt.Println(c)
	changeCar(&c)	
	fmt.Println(c)
}

func foo (x *int) {
	fmt.Println(*x)
	*x = 13
	fmt.Println(*x)
}

func changeCar(c *car) {
	(*c).make = "Lotus"
	(*c).model = "Esprit"
	(*c).year = 1977
	(*c).color = "white"
}