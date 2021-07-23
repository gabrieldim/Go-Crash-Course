package main

import "fmt"

func main() {
	// //Arrays
	// var fruitArr [2]string

	// //Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	//Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//We don't need to say immediately  how long array gonna be
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2]) //Range
}
