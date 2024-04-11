package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey, What's your name?")
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Scan(&age)
	fmt.Println("Hello", name, "age", age)

	// reader := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n')
	// fmt.Println("Hello", name)
}
