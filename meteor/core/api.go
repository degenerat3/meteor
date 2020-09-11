package main

import (
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	resp := &mcs.MCS{
		Status: 200,
		Desc:   "Core is running...\n",
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func regBot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	newBot := &mcs.MCS{}
	proto.Unmarshal(data, newBot)
	stat := regBotUtil(newBot)
	resp := &mcs.MCS{
		Status: stat,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func regHost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	newHost := &mcs.MCS{}
	proto.Unmarshal(data, newHost)
	stat := regHostUtil(newHost)
	resp := &mcs.MCS{
		Status: stat,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)

}

func regGroup(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	newGroup := &mcs.MCS{}
	proto.Unmarshal(data, newGroup)
	stat := regGroupUtil(newGroup)
	resp := &mcs.MCS{
		Status: stat,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func regHG(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	newHG := &mcs.MCS{}
	proto.Unmarshal(data, newHG)
	stat := regHGUtil(newHG)
	resp := &mcs.MCS{
		Status: stat,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func addActSingle(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	newAct := &mcs.MCS{}
	proto.Unmarshal(data, newAct)
	stat := addActSingleUtil(newAct)
	resp := &mcs.MCS{
		Status: stat,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func addActGroup(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	newAct := &mcs.MCS{}
	proto.Unmarshal(data, newAct)
	stat := addActGroupUtil(newAct)
	resp := &mcs.MCS{
		Status: stat,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func addResult(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	newRes := &mcs.MCS{}
	proto.Unmarshal(data, newRes)
}

func listBots(w http.ResponseWriter, r *http.Request) {
	botsList := listBotsUtil()
	resp := &mcs.MCS{
		Status: 200,
		Desc:   botsList,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)

}

func listHosts(w http.ResponseWriter, r *http.Request) {
	hostsList := listHostsUtil()
	resp := &mcs.MCS{
		Status: 200,
		Desc:   hostsList,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func listGroups(w http.ResponseWriter, r *http.Request) {
	groupsList := listGroupsUtil()
	resp := &mcs.MCS{
		Status: 200,
		Desc:   groupsList,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}
