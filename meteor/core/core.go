package main

import (
	"context"
	"github.com/degenerat3/meteor/meteor/core/ent"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

// DBClient is the connection to the psql db
var (
	DBClient *ent.Client
	ctx      = context.Background()

	// LOGPATH is the output path (including fname) for the listener logs
	LOGPATH = os.Getenv("LOGPATH")

	// write info logs to this
	infoLog *log.Logger

	// write warning logs to this
	warnLog *log.Logger

	// write all errors to this
	errLog *log.Logger
)

func main() {
	infoLog, warnLog, errLog = InitLogger(LOGPATH)
	var err error
	DBClient, err = ent.Open("postgres", "host=172.16.77.3 port=5432 user=met dbname=meteor password=dbpassword sslmode=disable")
	if err != nil {
		errLog.Printf("Error connecting to DB: %v\n", err.Error())
	}
	if err := DBClient.Schema.Create(ctx); err != nil {
		errLog.Printf("failed creating schema resources: %s\n", err.Error())
	}

	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/register/bot", regBot)
	http.HandleFunc("/register/host", regHost)
	http.HandleFunc("/register/group", regGroup)
	http.HandleFunc("/register/hostgroup", regHG)
	http.HandleFunc("/bot/checkin", botCheckin)
	http.HandleFunc("/add/action/single", addActSingle)
	http.HandleFunc("/add/action/group", addActGroup)
	http.HandleFunc("/add/result", addResult)
	http.HandleFunc("/list/bots", listBots)
	http.HandleFunc("/list/hosts", listHosts)
	http.HandleFunc("/list/groups", listGroups)
	http.HandleFunc("/list/actions", listActions)
	http.HandleFunc("/list/result", listResult)
	http.HandleFunc("/list/host", listHost)
	http.HandleFunc("/list/group", listGroup)
	infoLog.Println(http.ListenAndServe(":9999", nil))
	return
}
