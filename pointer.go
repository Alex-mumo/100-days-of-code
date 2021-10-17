package main

import (
	"fmt"
)

func main() {
	var value = 34
	fmt.Println("An interger: %d,  location in memory: %p\n", value, &value)

	var IntegerValue *int
	IntegerValue = &value
	fmt.Println("Value at memory location is %p is:  %d\n", IntegerValue, &IntegerValue)

}
