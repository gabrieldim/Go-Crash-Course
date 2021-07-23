package main

import "fmt"

func main() {
	x := 15
	y := 10

	//You can add or ||, and && etc.
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if y < x {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else {
		fmt.Println("It's different color than red.")
	}

	//Switch
	switch color {
	case "red":
		fmt.Println("The color is red")
	case "blue":
		fmt.Println("The color is blue")
	default:
		fmt.Println("It's different color than red or blue.")
	}
}
