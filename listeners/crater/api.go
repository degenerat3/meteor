package main

import (
	"bytes"
	"github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	resp := &mcs.MCS{
		Status: 200,
		Desc:   "Crater is running...\n",
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
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
	prot := &mcs.MCS{}
	proto.Unmarshal(data, prot)
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
