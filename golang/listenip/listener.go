package main

import (
	"log"
	"net"
	"github.com/vishvananda/netlink"
)

func main() {
	lo, err := netlink.LinkByName("lo")
	log.Println("Err1:", err)
	addr, err := netlink.ParseAddr("169.254.169.254/32")
	log.Println("Err2:", err)
	netlink.AddrAdd(lo, addr)


	l, err := net.Listen("tcp", "169.254.169.254:1234")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		b := make([]byte, 1024)
		for {
			n, err := conn.Read(b)
			if err != nil {
				log.Println(err)
				close(done)
				return
			}
			log.Printf("Read %v bytes: %v\n", n, b[:n])
			if string(b[:n]) == "ping\n\n" {
				_, err = conn.Write([]byte("pong\n"))
			} else {
				_, err = conn.Write([]byte("only know ping\n"))
			}
			if err != nil {
				log.Println(err)
				close(done)
				return
			}
		}
	}()
	<-done
}
