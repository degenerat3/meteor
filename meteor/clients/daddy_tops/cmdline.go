package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"github.com/smallfish/simpleyaml"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func registerHosts(cfname string) {
	login()
	filename, _ := filepath.Abs(cfname)
	yamlFile, _ := ioutil.ReadFile(filename)
	y, err := simpleyaml.NewYaml(yamlFile)
	if err != nil {
		panic(err)
	}
	ymap, err := y.Map()
	if err != nil {
		panic(err)
	}
	keys, err := y.GetMapKeys()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
	hasHosts := checkHasHosts(keys)
	if hasHosts == false {
		fmt.Println("Missing required `hosts` block in yaml")
		os.Exit(0)
	}

	hostsDat := ymap["hosts"]
	hostsDatInter := hostsDat.([]interface{})
	for _, hostinfo := range hostsDatInter {
		registerHostReq(hostinfo.(string))
	}

	for _, key := range keys {
		if key == "hosts" {
			continue
		}
		registerGroupReq(key)
		keydat := ymap[key]
		keydatInter := keydat.([]interface{})
		for _, hostinfo := range keydatInter {
			registerHostGroupReq(hostinfo.(string), key)
		}
	}
	return
}

func setServer(sv string) {
	DTSERVER = sv
	return
}

func checkHasHosts(keys []string) bool {
	for _, val := range keys {
		if val == "hosts" {
			return true
		}
	}
	return false
}

func registerUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Current Admin Password: ")
	adminpwb, _ := terminal.ReadPassword(int(syscall.Stdin))
	adminpw := string(adminpwb)
	fmt.Print("New Username: ")
	un, _ := reader.ReadString('\n')
	un = strings.TrimSuffix(un, "\n")
	un = strings.TrimSuffix(un, "\r")
	DTUSER = un
	fmt.Print("New Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("Confirm Password: ")
	bytePassword2, _ := terminal.ReadPassword(int(syscall.Stdin))
	pw := string(bytePassword)
	pw2 := string(bytePassword2)
	if pw != pw2 {
		fmt.Println("Error: Passwords do not match")
		os.Exit(0)
	}
	authdat := &mcs.AuthDat{
		Username: un,
		Password: pw,
		Token:    adminpw,
	}
	hostReg := &mcs.MCS{
		AuthDat: authdat,
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		panic(err)
	}
	url := "http://" + DTSERVER + "/register/user"
	_, err = http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		panic(err)
	}
	return
}
