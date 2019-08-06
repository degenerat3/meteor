package main

import (
	//"fmt"
	"net"
	//"github.com/google/gopacket"
	//layers "github.com/google/gopacket/layers"
)

func main() {
	addr := net.UDPAddr{
		Port: 53,
		IP:   net.ParseIP("127.0.0.1"),
	}
	u, _ := net.ListenUDP("udp", &addr)
}
