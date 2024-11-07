// src/str.go
package str

import (
	"fmt"
	"strings"
)

func Str() {
	fmt.Println("Strings are those enclosed in double quotes")
	fmt.Print("Provide a string: ")

	var input string
	fmt.Scanln(&input) // Scanf replaced by Scanln for correct input reading
	fmt.Println("The string is:", input)

	// String functions demonstration
	fmt.Println("The length of the string is:", len(input))
	fmt.Println("Does the string contain the substring 'in'?", strings.Contains(input, "in"))
	fmt.Println("Does the string start with 'The'?", strings.HasPrefix(input, "The"))
	fmt.Println("Does the string end with 'string'?", strings.HasSuffix(input, "string"))
	fmt.Println("The string in uppercase is:", strings.ToUpper(input))
	fmt.Println("The string in lowercase is:", strings.ToLower(input))
	fmt.Println("The string with 'i' replaced by 'y' is:", strings.ReplaceAll(input, "i", "y"))
	fmt.Println("The string split by space:", strings.Split(input, " "))
	fmt.Println("The string joined by space:", strings.Join(strings.Split(input, " "), " "))
	fmt.Println("The string with trimmed spaces:", strings.TrimSpace(input))
	fmt.Println("The string with leading spaces removed:", strings.TrimLeft(input, " "))
	fmt.Println("The string with trailing spaces removed:", strings.TrimRight(input, " "))
	fmt.Println("Index of first 'i':", strings.Index(input, "i"))
	fmt.Println("Index of last 'i':", strings.LastIndex(input, "i"))
	fmt.Println("Occurrences of 'i':", strings.Count(input, "i"))
	fmt.Println("Repeated string 3 times:", strings.Repeat(input, 3))
	fmt.Println("Title-cased string:", strings.Title(input))
	fmt.Println("ToTitle-cased string:", strings.ToTitle(input))
}
