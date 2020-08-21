package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	tRate := 4.75
	fmt.Println("Current rate", tRate)

	//pointer examples
	var testString *string = new(string)

	*testString = "my current pointer"
	fmt.Println(*testString, testString)

	*testString = "hey new string"
	fmt.Println(*testString, testString)

	something := "I am not pointer"

	ptr := &something
	fmt.Println(*ptr, ptr)

	something = "I am a pointer now"
	fmt.Println(*ptr, ptr)

}
