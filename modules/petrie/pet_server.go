// Shoutout to for a helpful blogpost/example @donutdan4114
package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
)

// HOST : server to listen on (pretty much always localhost)
var HOST = "localhost"

// PORT : port to listen on
var PORT = "5656"

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

func main() {

	l, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening for Petrie connections on " + HOST + ":" + PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go connHandle(conn)
	}
}

func connHandle(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString(magicStr)
	fmt.Printf("%s\n", message)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	resp := magicStr + "gottem" + magicStr
	conn.Write([]byte(resp))
	conn.Close()
}
