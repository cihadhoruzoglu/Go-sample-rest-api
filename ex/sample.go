package main

import (
	"fmt"
	"reflect"
)

const (
	message = "Hello"
)

var (
	name            = "Cihad"
	course, surname string
	age             int
	orderNo         = 10.000
)

func main() {
	surname = "Horuzoğlu"
	age = 26
	// ageWithString := strconv.Itoa(age)
	fmt.Println(message, name, surname, age, "type of", reflect.TypeOf(age))

	sum := age + int(orderNo)
	ptr := &sum

	myArray := [...]int{1, 2, 3}

	fmt.Println(myArray, sum, "pointer value", ptr, "is", *ptr)
}

