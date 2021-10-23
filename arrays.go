package main

import "fmt"

func main(){
	var array1 [5]int 
	for i := 0; i < len(array1); i++ {
		array1[i] = i * 3
		
	}
	for i := 0; i < len(array1); i++ {
		fmt.Printf("Array at index %d is %d\n" , i ,array1[i])
		
	}

}