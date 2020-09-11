package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	log.Fatal(http.ListenAndServe(":8888", nil))
	return
}
