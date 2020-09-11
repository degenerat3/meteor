package main

import (
	"context"
	"github.com/degenerat3/meteor/meteor/core/ent"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// DBClient is the connection to the psql db
var DBClient *ent.Client
var ctx = context.Background()

func main() {
	var err error
	DBClient, err = ent.Open("postgres", "host=172.69.1.2 port=5432 user=met dbname=meteor password=dbpassword")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/register/bot", regBot)
	http.HandleFunc("/register/host", regHost)
	http.HandleFunc("/register/group", regGroup)
	http.HandleFunc("/register/hostgroup", regHG)
	http.HandleFunc("/add/action/single", addActSingle)
	http.HandleFunc("/add/action/group", addActGroup)
	http.HandleFunc("/add/result", addResult)
	log.Fatal(http.ListenAndServe(":9999", nil))
	return
}
