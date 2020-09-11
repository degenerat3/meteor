package main

import (
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://172.69.1.1:9999/status")
	data, err := ioutil.ReadAll(r.Body)
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
