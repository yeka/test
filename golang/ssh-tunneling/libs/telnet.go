package libs

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func Telnet(addr string, dialFn func(network, addres string) (net.Conn, error)) (errr error) {
	if dialFn == nil {
		dialFn = net.Dial
	}
	c, err := dialFn("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		err := c.Close()
		if errr == nil {
			errr = err
		}
	}()
	fmt.Println("Telnet connected (type \"quit\" to exit)")

	exit := make(chan struct{})

	readTCPResponse := func() {
		r := bufio.NewReader(c)
		for {
			message, err := r.ReadString('\n')
			if errors.Is(err, io.EOF) {
				close(exit)
				return
			}
			fmt.Print(message)
		}
	}
	readStdin := func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_, err := fmt.Fprintf(c, text+"\n")
			if err != nil {
				fmt.Println("Telnet write error:", err)
			}
			switch strings.TrimSpace(text) {
			case "quit", "exit":
				close(exit)
			}
		}
	}

	go readTCPResponse()
	go readStdin()

	<-exit
	fmt.Println("Telnet exiting...")
	return
}
