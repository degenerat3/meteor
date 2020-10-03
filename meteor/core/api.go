package main

import (
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
	"io/ioutil"
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
		errLog.Println(err.Error())
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
		errLog.Println(err.Error())
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
		errLog.Println(err.Error())
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
		errLog.Println(err.Error())
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
		errLog.Println(err.Error())
	}
	newAct := &mcs.MCS{}
	proto.Unmarshal(data, newAct)
	stat, uid := addActSingleUtil(newAct)
	resp := &mcs.MCS{
		Status: stat,
		Uuid:   uid,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func addActGroup(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errLog.Println(err.Error())
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
		errLog.Println(err.Error())
	}
	newRes := &mcs.MCS{}
	proto.Unmarshal(data, newRes)
	stat := addResultUtil(newRes)
	resp := &mcs.MCS{
		Status: stat,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func botCheckin(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errLog.Println(err.Error())
	}
	newCheck := &mcs.MCS{}
	proto.Unmarshal(data, newCheck)
	stat, pendingActions := botCheckinUtil(newCheck)
	resp := &mcs.MCS{
		Status:  stat,
		Actions: pendingActions,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
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

func listActions(w http.ResponseWriter, r *http.Request) {
	actsList := listActionsUtil()
	resp := &mcs.MCS{
		Status: 200,
		Desc:   actsList,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func listResult(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errLog.Println(err.Error())
	}
	resProt := &mcs.MCS{}
	proto.Unmarshal(data, resProt)
	res := listResultUtil(resProt.GetUuid())
	resp := &mcs.MCS{
		Status: 200,
		Desc:   res,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func listHost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errLog.Println(err.Error())
	}
	resProt := &mcs.MCS{}
	proto.Unmarshal(data, resProt)
	res := listHostUtil(resProt)
	resp := &mcs.MCS{
		Status: 200,
		Desc:   res,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}

func listGroup(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errLog.Println(err.Error())
	}
	resProt := &mcs.MCS{}
	proto.Unmarshal(data, resProt)
	res := listGroupUtil(resProt)
	resp := &mcs.MCS{
		Status: 200,
		Desc:   res,
	}
	rdata, _ := proto.Marshal(resp)
	w.Write(rdata)
}
