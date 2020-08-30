package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"

	"github.com/RK-GCP/my-go-service/controllers"
	"github.com/RK-GCP/my-go-service/models"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}

//TryPanic - example
func TryPanic() {
	_, err := startwebserver(8200)
	if err != nil {
		_, errFile := os.Create("/tmp/file.txt")
		panic(errFile)
	}

}

//TryingLoops -Example
func TryingLoops() {
	//simple for loop
	s := []int{11, 21, 39}
	for i := 0; i < len(s); i++ {
		fmt.Println("Simple Loop: ", s[i])
	}

	//go way using indexer and value - works for slice, array, and map
	for i, v := range s {
		fmt.Println("Index Value Loop: ", i, v)
	}

	tp := map[string]int{"http": 80, "https": 443}
	//Value Only
	for _, val := range tp {
		fmt.Println("Value only: ", val)
	}

	//Keys only
	for k := range tp {
		fmt.Println("key only: ", k)
	}

	//both Key and Value
	for k, v := range tp {
		fmt.Println(k, ":", v)
	}

}

func manything2() {
	fmt.Println("Hello World")

	ec, err := startwebserver(3000)

	fmt.Println(ec, err)

	//if only need one piece of data from func data return use underscore for data not interested

	_, err2 := startwebserver(9000)
	fmt.Println("Single Value: ", err2)

	manything()
}

func startwebserver(port int) (int, error) {

	fmt.Println("Starting web server...")
	//code to start a web server.handler
	fmt.Println("web server started at port: ", port)
	er := errors.New("Not implemented")
	errorcode := 501
	return errorcode, er

}

func manything() {
	fmt.Println("******* MANY THINGS *********")
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
