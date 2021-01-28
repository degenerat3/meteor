package main

import (
	"bufio"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

// DTSERVER is the Daddy_Tops listener address that comms will be sent to
var DTSERVER = ""

// DTUSER is the username associated with this session
var DTUSER string

var authToken string

func getDTServ() string {
	s := os.Getenv("DT_SERVER")
	if s == "" {
		if os.Args[1] == "--server" {
			return "placeholder"
		}
		fmt.Println("'DT_SERVER' env is undefined, please specify the upstream Daddy Tops server.")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter DT Server (ex 127.0.0.1:8888): ")
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		s = strings.TrimSuffix(s, "\r")
		setServer(s)
	}
	return s
}

func main() {
	DTSERVER = getDTServ()
	if len(os.Args) > 1 {
		if os.Args[1] == "--register-hosts" {
			if len(os.Args) < 3 {
				fmt.Println("Missing arg: config.yml")
				os.Exit(0)
			} else {
				registerHosts(os.Args[2])
				return
			}
		} else if os.Args[1] == "--server" {
			if len(os.Args) < 3 {
				fmt.Println("Missing arg: server")
				os.Exit(0)
			} else {
				setServer(os.Args[2])
			}
		} else if os.Args[1] == "--register-user" {
			registerUser()
		} else if os.Args[1] == "--change-password" {
			changePW()
			os.Exit(0)
		} else {
			fmt.Println("Unknown argument")
			os.Exit(0)
		}
		os.Exit(0)
	}
	fmt.Println(" ===============================")
	fmt.Println("| DADDY TOPS - METEOR COMMANDER |")
	fmt.Println(" ===============================")
	login()
	if authToken == "Invalid user or password" {
		fmt.Println(authToken)
		os.Exit(0)
	}
	prm()

}
