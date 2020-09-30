package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// DTSERVER is the Daddy_Tops listener address that comms will be sent to
var DTSERVER = getDTServ()

// DTUSER is the username associated with this session
var DTUSER string

var authToken string

func getDTServ() string {
	s := os.Getenv("DT_SERVER")
	if s == "" {
		fmt.Println("'DT_SERVER' env is undefined, please specify the upstream Daddy Tops server.")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter DT Server (ex 127.0.0.1:8888): ")
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		s = strings.TrimSuffix(s, "\r")
	}
	return s
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--register-hosts" {
			if len(os.Args) < 3 {
				fmt.Println("Missing arg: config.yml")
				os.Exit(1)
			} else {
				registerHosts(os.Args[2])
				return
			}
		} else if os.Args[1] == "--server" {
			if len(os.Args) < 3 {
				fmt.Println("Missing arg: server")
				os.Exit(1)
			} else {
				setServer(os.Args[2])
			}
		} else if os.Args[1] == " --register-user" {
			registerUser()
		} else {
			fmt.Println("Unknow argument")
			os.Exit(1)
		}
	}
	fmt.Println(" ===============================")
	fmt.Println("| DADDY TOPS - METEOR COMMANDER |")
	fmt.Println(" ===============================")
	login()
	if authToken == "Invalid user or password" {
		fmt.Println(authToken)
		os.Exit(0)
	}
	prm()

}

func testRegUser() {

	authdat := &mcs.AuthDat{
		Username: "jim",
		Password: "letmein",
		Token:    "admin123!",
	}
	hostReg := &mcs.MCS{
		AuthDat: authdat,
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		panic(err)
	}
	url := "http://" + DTSERVER + "/register/user"
	_, err = http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		panic(err)
	}
}

func testLogin() {
	authdat := &mcs.AuthDat{
		Username: "jim",
		Password: "letmeinaaa",
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
	fmt.Printf("Token: %s\n", tok)

	authdat2 := &mcs.AuthDat{
		Token: tok,
	}

	hostReg := &mcs.MCS{
		Hostname:  "2.168.206.197",
		Interface: "eth0",
		AuthDat:   authdat2,
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		panic(err)
	}
	url = "http://" + DTSERVER + "/register/host"
	_, err = http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		panic(err)
	}
	url = "http://" + DTSERVER + "/list/hosts"
	resp, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	stat = &mcs.MCS{}
	proto.Unmarshal(data, stat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", stat.GetStatus())
	fmt.Println(string(stat.GetDesc()))

}

func testRegHost() {
	hostReg := &mcs.MCS{
		Hostname:  "192.168.206.197",
		Interface: "eth0",
	}
	hdata, err := proto.Marshal(hostReg)
	if err != nil {
		panic(err)
	}
	url := "http://" + DTSERVER + "/register/host"
	_, err = http.Post(url, "", bytes.NewBuffer(hdata))
	if err != nil {
		panic(err)
	}
	url = "http://" + DTSERVER + "/list/hosts"
	resp, err := http.Get(url)
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
	fmt.Printf("%d\n", stat.GetStatus())
	fmt.Println(string(stat.GetDesc()))

}
