package main

import (
	"fmt"
)

func main() {
	// Shorthand to declare a variable
	name := "Shantilal"
	age := 26
	height := 5.823558
	city := "Sanchore"
	state := "Rajasthan"

	fmt.Println("Hello, World!")
	fmt.Println("age: ", age, "name:", name, "height:", height)

	// f in Printf is used to format the output
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("name is %T\n", name)
	fmt.Printf("name is %s\n", name)
	fmt.Println(city)
	fmt.Println(state)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
