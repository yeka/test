package main

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"

	"ssh-tunneling/libs"
)

func main() {
	sshUser := "root"
	sshPass := "root"
	sshConn := "127.0.0.1:2222"
	remote := "192.168.10.3:6379"

	cfg := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{ssh.Password(sshPass)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	sshClient, err := ssh.Dial("tcp", sshConn, cfg)
	if err != nil {
		fmt.Printf("server dial error: %v", err)
		return
	}
	defer sshClient.Close()

	if err := libs.Telnet(remote, sshClient.Dial); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Telnet closed")
}
