package main

import (
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	addr, err := net.ResolveUDPAddr("udp4", "192.168.10.255:8829") // broadcast address of subnet 192.168.10.0/24
	if err != nil {
		panic(err)
	}

	_, err = pc.WriteTo([]byte("hello guys"), addr)
	if err != nil {
		panic(err)
	}
}
