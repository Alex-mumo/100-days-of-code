package main

import "fmt"

func main() {
	var arr = [5]int{23, 89, 78, 34, 12}
	x := Sum(&arr) // pass a pointer to the array
	fmt.Printf("The sum is %d\n :", x)
}

func Sum(a *[5]int) (sum int) {
	for _, v := range a {
		sum += v

	}
	return

}
