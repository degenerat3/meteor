package utils

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strings"
)

func checkRegStatus(regFile string) bool {
	if _, err := os.Stat(regFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckRegStatus is the exported version of registration status check
func CheckRegStatus(regfile string) bool {
	return checkRegStatus(regfile)
}

//get IP of default interface, which the DB will use as hostname
func getIP() string {
	conn, _ := net.Dial("udp4", "8.8.8.8:80")
	defer conn.Close()
	ad := conn.LocalAddr().(*net.UDPAddr)
	ipStr := ad.IP.String()
	return ipStr
}

func storeUUID(regFile string, uuid string, obfText string) {
	obf := obfuscateUUID(uuid, obfText)
	f, _ := os.Create(regFile)
	io.WriteString(f, obf)
	f.Close()
	return
}

func fetchUUID(regFile string, obfText string) string {
	obf, _ := ioutil.ReadFile(regFile)
	deobf := deobfuscateUUID(string(obf), obfText)
	return deobf
}

func obfuscateUUID(uuid string, obfText string) string {
	splituid := strings.Split(uuid, "-")
	seed := 4
	l1 := strings.Repeat(obfText, rand.Intn(seed))
	l2 := strings.Repeat(obfText, rand.Intn(seed))
	l3 := strings.Repeat(obfText, rand.Intn(seed))
	l4 := strings.Repeat(obfText, rand.Intn(seed))
	obf := splituid[0] + l1 + splituid[1] + splituid[2] + l2 + splituid[3] + l3 + splituid[4] + l4
	// its' really crappy obfuscation but it's a small deterrent
	return obf
}

func deobfuscateUUID(obfuscatedUUID string, obfText string) string {
	p := strings.Replace(obfuscatedUUID, obfText, "", -1)
	p = p[:8] + "-" + p[8:]
	p = p[:13] + "-" + p[13:]
	p = p[:18] + "-" + p[18:]
	p = p[:23] + "-" + p[23:]
	return p
}
