package main

import (
	"fmt"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Core is running...\n")
}
