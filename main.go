package main

import (
	"fmt"
	"math"

	"github.com/RK-GCP/my-go-service/models"
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
	fmt.Println(getPi())

	//array
	t := [3]int{1, 4, 78}
	fmt.Println("Array ia:", t)

	//slice
	mySlice := t[:]
	mySlice[2] = 88
	fmt.Println("A Slice from Array: ", mySlice)

	mSlice2 := []string{"one", "two", "three"}
	mSlice2 = append(mSlice2, "four", "five")
	fmt.Println("Self Init Slice: ", mSlice2)

	s1 := mSlice2[0:3]
	s2 := mSlice2[:2]
	s3 := mSlice2[1:2]

	fmt.Println(s1, s2, s3)

	//map
	m1 := map[string]int{"foo": 40, "bar": 38}
	fmt.Println(m1)
	fmt.Println("Foo is: ", m1["foo"])
	delete(m1, "bar")
	fmt.Println("After deleting bar: ", m1)

	//custom type
	u := models.User{
		ID:        2,
		FirstName: "Ping",
		LastName:  "Pong",
	}

	fmt.Println(u)

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func getPi() float32 {
	return math.Pi
}
