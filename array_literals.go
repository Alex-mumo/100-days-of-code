package main

import "fmt"

func main() {

	//var age = [5]int{34, 34, 89, 78, 90}
	var ageKeyValue = []string{1: "alex", 2: "student", 3: "it", 4: "learning"}

	for i := 0; i < len(ageKeyValue); i++ {
		fmt.Printf("Name at %d\n is %s\n", i, ageKeyValue[i])

	}

}
