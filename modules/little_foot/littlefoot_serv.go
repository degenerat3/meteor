package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CORE is the address/port of the meteor core API
var CORE = "http://172.69.1.1:9999"

// PORT : port to listen on
var PORT = "8080"

//forward json from registration callbacks
func register(rw http.ResponseWriter, req *http.Request) {
	url := CORE + "/register/bot"
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	cli := http.Client{}
	jsonStr := []byte(reqBody)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(r)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	rw.Write(body)
	return
}

//forward json from getCommands callbacks
func getCommands(rw http.ResponseWriter, req *http.Request) {
	url := CORE + "/get/command"
	reqBody, err := ioutil.ReadAll(req.Body)
	cli := http.Client{}
	jsonStr := []byte(reqBody)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(r)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	rw.Write(body)
}

//forward json from addActionResult callbacks
func addActionResult(rw http.ResponseWriter, req *http.Request) {
	url := CORE + "/add/actionresult"
	reqBody, err := ioutil.ReadAll(req.Body)
	cli := http.Client{}
	jsonStr := []byte(reqBody)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(r)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	rw.Write(body)
}

func lfStatus(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("LittleFoot server is running.\n"))
}

func main() {
	fmt.Println("Listening for LittleFoot connections on port: " + PORT + "...")
	http.HandleFunc("/", lfStatus)
	http.HandleFunc("/register/bot", register)
	http.HandleFunc("/get/command", getCommands)
	http.HandleFunc("/add/actionresult", addActionResult)
	portStr := ":" + PORT
	http.ListenAndServe(portStr, nil)
}
