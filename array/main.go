package main

import "fmt"

func main() {
	fmt.Println("We are learning array in golang")

	var name [5]string
	name[0] = "Shantilal"
	name[2] = "Pintu"

	fmt.Println("Names of person is : ", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are : ", numbers)

	fmt.Println("Length of array is : ", len(numbers))

	fmt.Println("Value at 2nd index is :", name[2])

}
