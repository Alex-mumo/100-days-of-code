package main

import (
	"fmt"
)

type student struct {
	name       string
	course     string
	age        int
	department string
}

func newStudent(name string) *student {
	s := student{name: name}
	s.age = 23
	return &s
}

func main() {

	fmt.Println(student{"Alex", "IT", 24, "Information technology"})
	fmt.Println(student{name: "Mumo", course: "IT", age: 34, department: "Alex"})

}
