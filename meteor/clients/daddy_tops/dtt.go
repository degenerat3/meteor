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
	testRegHost()
	testRegGroup()
	testRegHG()
	testRegBot()

}

func testRegHost() {
	hostReg := &mcs.MCS{
		Hostname:  "blackbox",
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

func testRegGroup() {
	groupReg := &mcs.MCS{
		Groupname: "web",
		Desc:      "Web servers",
	}
	gdata, _ := proto.Marshal(groupReg)
	url := "http://" + SERVER + "/register/group"
	_, err := http.Post(url, "", bytes.NewBuffer(gdata))
	if err != nil {
		panic(err)
	}
	url = "http://" + SERVER + "/list/groups"
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

func testRegHG() {
	groupReg := &mcs.MCS{
		Groupname: "web",
		Hostname:  "blackbox",
	}
	gdata, _ := proto.Marshal(groupReg)
	url := "http://" + SERVER + "/register/hostgroup"
	_, err := http.Post(url, "", bytes.NewBuffer(gdata))
	if err != nil {
		panic(err)
	}
	url = "http://" + SERVER + "/list/groups"
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

func testRegBot() {
	botReg := &mcs.MCS{
		Uuid:     "abcdefg",
		Interval: 60,
		Delta:    5,
		Hostname: "blackbox",
	}
	bdata, _ := proto.Marshal(botReg)
	url := "http://" + SERVER + "/register/bot"
	_, err := http.Post(url, "", bytes.NewBuffer(bdata))
	if err != nil {
		panic(err)
	}
	url = "http://" + SERVER + "/list/bots"
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
