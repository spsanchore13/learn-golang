package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) (result int) {
	result = a - b
	return
}

func main() {
	fmt.Println("Hello, World!")
	simpleFunction()
	ans := add(10, 20)
	fmt.Println("Addition of 2 numbers", ans)
	ans = sub(30, 20)
	fmt.Println("Subtraction of 2 numbers", ans)
}
