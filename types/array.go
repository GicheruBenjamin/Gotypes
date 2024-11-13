package types

import (
	"fmt"
	"strconv"
)

func Array() {
	fmt.Println("Array")
	fmt.Println("---------------------------------")
	fmt.Println("Fixed size of same type list that is constrained by itself")

	// Step 1: Provide the type of the array
	fmt.Println("Choose the type of array: 1. int, 2. float64, 3. string")
	var choice int
	fmt.Scanln(&choice)

	// Step 2: Provide the size of the array
	fmt.Println("Provide the size of the array:")
	var size int
	fmt.Scanln(&size)

	// Step 3: Provide the values for the array
	fmt.Printf("Provide %d values for the array:\n", size)
	values := make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter value %d: ", i+1)
		fmt.Scanln(&values[i])
	}

	// Step 4: Print the array based on the chosen type
	fmt.Println("Here is your array:")
	fmt.Println("---------------------------------")
	switch choice {
	case 1: // int array
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			val, err := strconv.Atoi(values[i])
			if err != nil {
				fmt.Println("Invalid integer value:", values[i])
				return
			}
			arr[i] = val
		}
		fmt.Printf("Array of type int: %v\n", arr)

	case 2: // float64 array
		arr := make([]float64, size)
		for i := 0; i < size; i++ {
			val, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Println("Invalid float value:", values[i])
				return
			}
			arr[i] = val
		}
		fmt.Printf("Array of type float64: %v\n", arr)

	case 3: // string array
		arr := make([]string, size)
		for i := 0; i < size; i++ {
			arr[i] = values[i]
		}
		fmt.Printf("Array of type string: %v\n", arr)

	default:
		fmt.Println("Invalid choice")
	}

	fmt.Println("---------------------------------")
}
