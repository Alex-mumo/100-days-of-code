package main

import "fmt"

func main() {

	//calling the function
	fmt.Printf("The answer ins: %d\n", multiply(23, 34, 89))

}

func multiply(a int, b int, c int) int {
	return a * b * c

}
