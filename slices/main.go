package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 7, 8, 9, 10)
	fmt.Println("Numbers are : ", numbers)
	fmt.Printf("Length of array is : %d\n", len(numbers))
	fmt.Printf("Numbers has data type : %T\n", numbers)
}
