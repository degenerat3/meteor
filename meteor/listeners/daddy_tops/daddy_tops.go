package main

import (
	"log"
	"net/http"
	"os"
)

// CORESERVER is the IP:Port of the Meteor core
var CORESERVER = os.Getenv("CORE_SERVER") // format: 9.9.9.9:9999

func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/register/bot", forwardReq)
	http.HandleFunc("/register/host", forwardReq)
	http.HandleFunc("/register/group", forwardReq)
	http.HandleFunc("/register/hostgroup", forwardReq)
	http.HandleFunc("/add/action/single", forwardReq)
	http.HandleFunc("/add/action/group", forwardReq)
	http.HandleFunc("/add/result", forwardReq)
	http.HandleFunc("/list/bots", listForward)
	http.HandleFunc("/list/hosts", listForward)
	http.HandleFunc("/list/groups", listForward)
	http.HandleFunc("/list/actions", listForward)
	log.Fatal(http.ListenAndServe(":8888", nil))
	return
}
