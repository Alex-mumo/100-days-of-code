package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const address = "127.0.0.1:8888"

func capitalized(conn net.Conn) {
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print("error reading the file: %s\n", err.Error())
		return
	}
	fmt.Print("Received data: %s\n", data)
	conn.Write([]byte(strings.ToUpper(data)))
	//closing the connection
	conn.Close()
}

func main() {

	//specifying the communicatin protocol and
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	fmt.Printf("listening on: %s\n", address)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("error while accepting the connection : %s\n", err.Error())
			continue

		}
		go capitalized(conn)
	}

}
