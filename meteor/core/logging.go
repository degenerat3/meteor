package main

import (
	"io"
	"log"
	"os"
)

// InitLogger will take a filepath and generate logger type for info/warn/err that writes to stdout and the path
func InitLogger(logpath string) (*log.Logger, *log.Logger, *log.Logger) {
	file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	multiW := io.MultiWriter(os.Stdout, file)
	InfoLogger := log.New(multiW, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger := log.New(multiW, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(multiW, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return InfoLogger, WarningLogger, ErrorLogger
}
