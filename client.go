package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8888"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//client to enter a text on console
		fmt.Printf("Enter some text:")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error while reading user input: %s\n", err.Error())
			continue
		}
		//connect
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("Error occurred while connecting: %s\n", err.Error())

			//write the data
			fmt.Fprintf(conn, data)

			status, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Printf("Error while reading back the response: %s\n", err.Error())

			}
			fmt.Printf("Received back the response: %s\n", status)
			conn.Close()
		}
	}
}
