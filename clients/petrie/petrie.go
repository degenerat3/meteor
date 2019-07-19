package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

// HOST : server to call
var serv = "192.168.206.183:5656"

// MAGIC : the shared hex bytes that will signify the start/end of each MAD payload
var MAGIC = []byte("AAAA")
var magicStr = genMagStr()

func genMagStr() string {
	dst := make([]byte, hex.DecodedLen(len(MAGIC)))
	n, _ := hex.Decode(dst, MAGIC)
	ret := string(dst[:n])
	return ret
}

func sendData(data string) string {
	payload := magicStr + data
	fmt.Printf("sending: %s", payload)
	conn, err := net.Dial("tcp", serv)
	if err != nil {
		panic(err)
	}
	outText := []byte(payload)
	conn.Write(outText)
	resp := make([]byte, 256)
	conn.Read(resp)
	respStr := string(resp)
	return respStr
}

func main() {
	a := sendData("testing can you recv the payload")
	print("GOT: %s", a)
}
