package main

import "fmt"

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

}