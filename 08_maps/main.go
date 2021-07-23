package main

import "fmt"

func main() {
	//Define a map
	emails := make(map[string]string) // K, V

	//Assign key values
	emails["John"] = "john@gmail.com"
	emails["Brad"] = "brad@gmail.com"
	emails["Mark"] = "mark@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Brad"])
	fmt.Println(len(emails))

	//Delete from map
	fmt.Println("Deleting Brad from the map ...")
	delete(emails, "Brad")
	fmt.Println(emails)

	//Declare map and add key values
	emails2 := map[string]string{"Bob": "bob@gmail.com", "Mark": "mark@gmail.com"}
	fmt.Println(emails2)
}
