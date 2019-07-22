package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
)

// HOST : server to call
var serv = "192.168.206.183:5656"

// MAGIC : the shared hex byte that will signify the start of each MAD payload
var MAGIC = []byte{0xAA}

// MAGICBYTE is the single byte (rather than a byte array)
var MAGICBYTE = MAGIC[0]

//MAGICSTR is the ascii representation of the magic byte
var MAGICSTR = string(MAGIC)

// MAGICTERM : the shared hex byte that will signify the end of each MAD payload
var MAGICTERM = []byte{0xAB}

// MAGICTERMBYTE is the single byte (rather than a byte array)
var MAGICTERMBYTE = MAGICTERM[0]

//MAGICTERMSTR is the ascii representation of the magic byte
var MAGICTERMSTR = string(MAGICTERM)

func genMagStr() string {
	dst := make([]byte, hex.DecodedLen(len(MAGIC)))
	n, _ := hex.Decode(dst, MAGIC)
	ret := string(dst[:n])
	return ret
}

func sendData(data string) string {
	payload := MAGICSTR + data + MAGICTERMSTR
	fmt.Printf("sending: %s\n", payload)
	conn, err := net.Dial("tcp4", serv)
	if err != nil {
		panic(err)
	}
	outText := []byte(payload)
	conn.Write(outText)
	//resp := make([]byte, 256)
	//conn.Read(resp)
	message, err := bufio.NewReader(conn).ReadString(MAGICTERMBYTE)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	respStr := string(message)
	conn.Close()
	return respStr
}

func main() {
	a := sendData("testing can you recv the payload")
	fmt.Printf("GOT: %s\n", a)
}
