package main

import (
	"fmt"
	"math"
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

	//function call
	fmt.Println(split(18))
	fmt.Println(math.Pi)

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
