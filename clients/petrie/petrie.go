package main

import (
	"fmt"
	"net"
)

var serv = "127.0.0.1:5656"

func sendData(data string) string {
	fmt.Printf("sending: %s", data)
	conn, _ := net.Dial("tcp", serv)
	outText := []byte(data)
	conn.Write(outText)
	resp := make([]byte, 256)
	conn.Read(resp)
	respStr := string(resp)
	return respStr
}



func main() {
	a := sendData("")
	print("GOT: %s", a)
}
