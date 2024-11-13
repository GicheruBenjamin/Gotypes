package main

import (
	"fmt"
	"gotypes/types"
)


func main() {

	//Welcome to the app
	fmt.Println("Welcome to the app")
	//Chose the type u wnat to create 
	var choice int
	fmt.Println("Chose the type U want")
	fmt.Println("1. Integer 2. Float 3. String")
	fmt.Scanln(&choice)

	// Execute the chosen type
	switch choice {
	case 1:
		types.Alias()
	case 2:
		types.Array()
	default:
		fmt.Println("Invalid choice")
	}

}