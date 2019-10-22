// Shoutout to for a helpful blogpost/example @donutdan4114
package main

import (
	"encoding/json"
	"fmt"
	"github.com/degenerat3/metcli/server"
	"io/ioutil"
	"net/http"
)

// CORE is the address/port of the meteor core API
var CORE = "http://172.69.1.1:9999"

// PORT : port to listen on
var PORT = "80"

// MAGIC : the shared hex byte that will signify the start of each MAD payload
var MAGIC = []byte{0xAA}

// MAGICBYTE is the single byte (rather than a byte array)
var MAGICBYTE = MAGIC[0]

//MAGICSTR is the ascii representation of the magic byte
var MAGICSTR = "XXXXX" //string(MAGIC)

// MAGICTERM : the shared hex byte that will signify the end of each MAD payload
var MAGICTERM = []byte{0xAB}

// MAGICTERMBYTE is the single byte (rather than a byte array)
var MAGICTERMBYTE = MAGICTERM[0]

//MAGICTERMSTR is the ascii representation of the magic byte
var MAGICTERMSTR = "YYYYY" //string(MAGICTERM)

var m = server.GenMetserver(CORE, MAGIC, MAGICSTR, MAGICTERM, MAGICTERMSTR)

//take the MAD payload and do stuff with it
func connHandle(rw http.ResponseWriter, req *http.Request) {
	bytmessage, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	type Message struct {
		Comms string `json:"comms"`
	}
	var msg Message
	err := json.Unmarshal(bytmessage, &msg)
	result := server.HandlePayload(msg, m)
	rw.Write([]byte(result))
}

func lfStatus(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("LittleFoot server is running.\n"))
}

func main() {
	fmt.Println("Listening for LittleFoot connections on port: " + PORT + "...")
	http.HandleFunc("/", lfStatus)
	http.HandleFunc("/communicate", connHandle)
	portStr := ":" + PORT
	http.ListenAndServe(portStr, nil)
}
