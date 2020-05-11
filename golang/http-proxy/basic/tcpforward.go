package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

var targetAddr = "127.0.0.1:8001"

func main() {
	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		fmt.Println("unable to open local port (:8002): %w", err)
		return
	}
	defer listener.Close()

	fmt.Println("Start listening on :8003")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accpting connection error: %w", err)
			return
		}

		go func() {
			handleConnection(conn)
		}()
	}

}

func handleConnection(conn net.Conn) {
	rconn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		fmt.Println("unable to connect to ", targetAddr)
		return
	}
	defer rconn.Close()

	wg := sync.WaitGroup{}
	copyConn := func(writer, reader net.Conn, name string) {
		defer wg.Done()
		if err := Copy(writer, reader, name); err != nil {
			err = fmt.Errorf("io.Copy error: %w", err)
		}
	}

	wg.Add(2)
	go copyConn(conn, rconn, "CLIENT")
	go copyConn(rconn, conn, "SERVER")
	wg.Wait()
}

func Copy(source io.Reader, dest io.Writer, name string) error {
	buf := make([]byte, 32*1024)
	for {
		n, err := source.Read(buf)
		if err != nil {
			return err
		}
		fmt.Println("-- " + name + " --\n" + string(buf[:n]))
		if n > 0 {
			w, err := dest.Write(buf[:n])
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			if w != n {
				return errors.New("read/write not same")
			}
		}
	}
}
