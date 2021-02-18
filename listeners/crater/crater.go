package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var (
	// CORESERVER is the IP:Port of the Meteor core
	CORESERVER = os.Getenv("CORE_SERVER") // format: 9.9.9.9:9999
)

func main() {

	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/bot/checkin", forwardReq)
	http.HandleFunc("/register/bot", forwardReq)
	http.HandleFunc("/add/result", forwardReq)
	log.Fatal(http.ListenAndServe(":10999", nil))
	return
}
