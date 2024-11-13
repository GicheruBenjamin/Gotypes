package types

import (
	"fmt"
	"strconv"
)

func Alias() {
	fmt.Println("Aliases types")
	fmt.Println("---------------------------------")

	// Provide a name for creation of a type
	fmt.Println("Provide a name for creation of a type")

	// Choose the type you want to create an alias
	fmt.Println("Choose the type you want to create an alias")
	fmt.Println("---------------------------------")
	fmt.Println("1. int , 2. float64 , 3. string , 4. bool")
	fmt.Println("---------------------------------")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanln(&choice)

	// Define type alias and print the result based on user input
	var variable, val string
	fmt.Print("Provide the variable name: ")
	fmt.Scanln(&variable)
	fmt.Print("Provide the value: ")
	fmt.Scanln(&val)
	fmt.Println("---------------------------------")

	switch choice {
	case 1:
		type AliasInt int
		valInt, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid integer value")
			return
		}
		var aliasVar AliasInt = AliasInt(valInt)
		fmt.Printf("Variable: %s, Value: %d, Type: %T\n", variable, aliasVar, aliasVar)

	case 2:
		type AliasFloat float64
		valFloat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println("Invalid float value")
			return
		}
		var aliasVar AliasFloat = AliasFloat(valFloat)
		fmt.Printf("Variable: %s, Value: %f, Type: %T\n", variable, aliasVar, aliasVar)

	case 3:
		type AliasString string
		var aliasVar AliasString = AliasString(val)
		fmt.Printf("Variable: %s, Value: %s, Type: %T\n", variable, aliasVar, aliasVar)

	case 4:
		type AliasBool bool
		valBool, err := strconv.ParseBool(val)
		if err != nil {
			fmt.Println("Invalid boolean value (use true/false)")
			return
		}
		var aliasVar AliasBool = AliasBool(valBool)
		fmt.Printf("Variable: %s, Value: %t, Type: %T\n", variable, aliasVar, aliasVar)

	default:
		fmt.Println("Invalid choice")
	}
}

