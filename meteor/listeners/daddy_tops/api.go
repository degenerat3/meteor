package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://172.69.1.1/status")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(body))
}
