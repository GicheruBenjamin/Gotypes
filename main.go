package main

import (
	"fmt"
	"gotypes/src"
)


func main() {

	//Chose the type
	fmt.Println("Choose the type you want to demonstrate:")
    fmt.Println("1. Integer")
    fmt.Println("3. String")
    fmt.Print("Enter your choice: ")

    var choice int
    fmt.Scanln(&choice)
    switch choice {
    case 1:
        src.IntOperations()
    case 3:
		src.Str()
	default:
		fmt.Println("Invalid choice!")
		break
	}
}