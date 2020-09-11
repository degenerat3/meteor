package main

import (
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Core is running...\n")
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
