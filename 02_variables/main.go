package main

import "fmt"

//var name = "Gabriel" //Global Variable

func main() {

	// var name = "Gabriel"
	var myAge int32 = 21
	const isCool = true
	//isCool = false   // now allowed
	name, email := "Gabriel", "test@test.com"
	size := 1.3

	fmt.Println(name, myAge, isCool, size, email)
	fmt.Printf("%T\n", name) // take the type from name

}

//ALL TYPES:
// string
// bool
// int
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128
