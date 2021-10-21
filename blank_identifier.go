package main

import "fmt"

func main() {
	var value1 int
	var value2 float32

	value1, _, value2 = ThreeValues()
	fmt.Printf("First value: %d\n Second value: %f\n", value1, value2)
}

func ThreeValues() (int, int, float32) {
	return 34, 23, 23.0

}
