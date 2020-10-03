package main

import (
	"fmt"
	"os"
	"ssh-tunneling/libs"
)

func main() {
	sshConn := "root:root@127.0.0.1:2222"
	remote := "192.168.10.3:6379"
	local := ":6379"

	t, err := libs.Tunnel(sshConn, remote, local)
	if err != nil {
		fmt.Println("ssh failed:", err)
		os.Exit(0)
	}
	defer t.Close()
	fmt.Println("Tunnel ready on:", t.LocalPort())

	err = libs.Telnet("127.0.0.1:" + t.LocalPort(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Telnet closed")
}
