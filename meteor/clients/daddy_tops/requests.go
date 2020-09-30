package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
)

func login() {
	fmt.Println("LOGIN")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	un, _ := reader.ReadString('\n')
	un = strings.TrimSuffix(un, "\n")
	un = strings.TrimSuffix(un, "\r")
	DTUSER = un
	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	pw := string(bytePassword)
	fmt.Println()

	authdat := &mcs.AuthDat{
		Username: un,
		Password: pw,
	}
	hostReg1 := &mcs.MCS{
		AuthDat: authdat,
	}
	hdata1, err := proto.Marshal(hostReg1)
	if err != nil {
		panic(err)
	}
	url := "http://" + DTSERVER + "/login"
	resp, err := http.Post(url, "", bytes.NewBuffer(hdata1))
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		panic(err)
	}
	tok := stat.GetDesc()
	if tok == "Invalid user or password" {
		fmt.Println("Error: Invalid username or password")
		os.Exit(0)
	}
	authToken = tok
}

func generateAuthData() *mcs.AuthDat {
	authdat := &mcs.AuthDat{
		Token: authToken,
	}
	return authdat
}

func registerHostReq(host string) {
	hostInfo := strings.Split(host, ":")
	fmt.Printf("Registering host: %s\n", hostInfo[0])
	hostReg := &mcs.MCS{
		Hostname:  hostInfo[0],
		Interface: hostInfo[1],
		AuthDat:   generateAuthData(),
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	url := "http://" + DTSERVER + "/register/host"
	resp, err := http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %d\n", stat.GetStatus())
	return
}

func registerGroupReq(group string) {
	groupInfo := strings.Split(group, "--")
	fmt.Printf("Registering group: %s\n", groupInfo[0])
	groupReg := &mcs.MCS{
		Groupname: groupInfo[0],
		Desc:      groupInfo[1],
		AuthDat:   generateAuthData(),
	}
	gdata, _ := proto.Marshal(groupReg)
	url := "http://" + DTSERVER + "/register/group"
	resp, err := http.Post(url, "", bytes.NewBuffer(gdata))
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %d\n", stat.GetStatus())
	return
}

func registerHostGroupReq(host string, group string) {
	gn := strings.Split(group, "--")[0]
	fmt.Printf("Adding '%s' to the '%s' group\n", host, gn)
	groupReg := &mcs.MCS{
		Groupname: gn,
		Hostname:  host,
		AuthDat:   generateAuthData(),
	}
	gdata, _ := proto.Marshal(groupReg)
	url := "http://" + DTSERVER + "/register/hostgroup"
	resp, err := http.Post(url, "", bytes.NewBuffer(gdata))
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %d\n", stat.GetStatus())
	return
}
