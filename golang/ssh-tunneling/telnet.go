package main

import (
	"fmt"
	"ssh-tunneling/libs"
)

func main() {
	err := libs.Telnet("127.0.0.1:6379")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Telnet closed")
}
