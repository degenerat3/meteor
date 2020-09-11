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
	log.Fatal(http.ListenAndServe(":9999", nil))
	return
}
