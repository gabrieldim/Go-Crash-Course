package main

import "fmt"

func main() {
	a := 5
	b := &a //the memory address of the a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) //we get *int (pointer to int)

	//User * to read value from memory address
	fmt.Println(*b)
	fmt.Println(*&b) // same thing as line 13

	//Change value with pointer
	*b = 10
	fmt.Println(a)

}
