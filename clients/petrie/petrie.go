package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

// HOST : server to call
var serv = "127.0.0.1:5656"

// MAGIC : the shared hex bytes that will signify the start/end of each MAD payload
var MAGIC uint16 = 0xAAAA
var magicStr = genMagStr()

func genMagStr() []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, MAGIC)
	mstr := make([]byte, hex.DecodedLen(len(b)))
	_, err := hex.Decode(mstr, b)
	if err != nil {
		log.Fatal(err)
	}
	return mstr
}

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
