package src

import (
	"fmt"
)

func IntOperations() {
	fmt.Println("Provide two integers:")

	var num1, num2 int
	fmt.Scan(&num1, &num2)

	fmt.Println("The first number is:", num1)
	fmt.Println("The second number is:", num2)

	// Basic operations
	fmt.Println("Sum:", num1+num2)
	fmt.Println("Difference:", num1-num2)
	fmt.Println("Product:", num1*num2)
	if num2 != 0 {
		fmt.Println("Quotient:", num1/num2)
		fmt.Println("Remainder:", num1%num2)
	} else {
		fmt.Println("Cannot divide by zero")
	}

	// Comparison
	fmt.Println("Is num1 greater than num2?", num1 > num2)
	fmt.Println("Is num1 less than num2?", num1 < num2)
	fmt.Println("Are num1 and num2 equal?", num1 == num2)
}

func BasicOperations(num1, num2 int) {
	fmt.Printf("Numbers: num1 = %d, num2 = %d\n", num1, num2)

	// Arithmetic
	fmt.Printf("Sum: %d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("Difference: %d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("Product: %d * %d = %d\n", num1, num2, num1*num2)

	if num2 != 0 { // Prevent division by zero
		fmt.Printf("Quotient: %d / %d = %d\n", num1, num2, num1/num2)
		fmt.Printf("Remainder: %d %% %d = %d\n", num1, num2, num1%num2)
	} else {
		fmt.Println("Division by zero is not allowed!")
	}

	// Comparisons
	fmt.Printf("num1 == num2: %t\n", num1 == num2)
	fmt.Printf("num1 != num2: %t\n", num1 != num2)
	fmt.Printf("num1 > num2: %t\n", num1 > num2)
	fmt.Printf("num1 < num2: %t\n", num1 < num2)
	fmt.Printf("num1 >= num2: %t\n", num1 >= num2)
	fmt.Printf("num1 <= num2: %t\n", num1 <= num2)
}