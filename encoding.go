package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "NCAtaWVjaG8gV2VsY29tZWV4aXQhIHwgYmFzZTY0IC1pQAo="

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("Error occurred %v", err)

	}
	fmt.Println("The decoded data is", decodedData)

}
