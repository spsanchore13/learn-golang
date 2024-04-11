package main

import "fmt"

// func divide(a, b float64) (float64, error) {

// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator cannot be ze+ro")
// 	}

// 	return a / b, nil
// }

func divide(a, b float64) (float64, string) {

	if b == 0 {
		return 0, "denominator cannot be zero"
	}

	return a / b, "nil"
}

func main() {
	fmt.Println("Error Handling In The Functions")
	ans, err := divide(10, 0)
	if err != "nil" {
		fmt.Println(err)
	}
	fmt.Println("Division of two numbers is", ans)
}
