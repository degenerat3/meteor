package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	log.Fatal(http.ListenAndServe(":9999", nil))
	return
}
