package main

import (
	"bytes"
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
)

var SERVER string

func main() {
	SERVER = os.Args[1]
	//testRegUser()
	testLogin()

}

func testRegUser() {

	authdat := &mcs.AuthDat{
		Username: "jim",
		Password: "letmein",
		Token:    "admin123!",
	}
	hostReg := &mcs.MCS{
		AuthDat: authdat,
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		panic(err)
	}
	url := "http://" + SERVER + "/register/user"
	_, err = http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		panic(err)
	}
}

func testLogin() {
	authdat := &mcs.AuthDat{
		Username: "jim",
		Password: "letmein",
	}
	hostReg1 := &mcs.MCS{
		AuthDat: authdat,
	}
	hdata1, err := proto.Marshal(hostReg1)
	if err != nil {
		panic(err)
	}
	url := "http://" + SERVER + "/login"
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
	fmt.Printf("Token: %s\n", tok)

	authdat2 := &mcs.AuthDat{
		Token: tok,
	}

	hostReg := &mcs.MCS{
		Hostname:  "192.168.206.197",
		Interface: "eth0",
		AuthDat:   authdat2,
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		panic(err)
	}
	url = "http://" + SERVER + "/register/host"
	_, err = http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		panic(err)
	}
	url = "http://" + SERVER + "/list/hosts"
	resp, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	stat = &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", stat.GetStatus())
	fmt.Println(string(stat.GetDesc()))

}

func testRegHost() {
	hostReg := &mcs.MCS{
		Hostname:  "192.168.206.197",
		Interface: "eth0",
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		panic(err)
	}
	url := "http://" + SERVER + "/register/host"
	_, err = http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		panic(err)
	}
	url = "http://" + SERVER + "/list/hosts"
	resp, err := http.Get(url)
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
	fmt.Printf("%d\n", stat.GetStatus())
	fmt.Println(string(stat.GetDesc()))

}
