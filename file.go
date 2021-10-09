package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Create("file.dat")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hello world")
	file.Close()

	stream, err := ioutil.ReadFile("file.dat")
	if err != nil {
		log.Fatal(err)
	}
	streamString := string(stream)
	fmt.Printf(streamString)

}
