# Module Spec
 Modules should be written in Golang so they can properly utilize the [metcli](https://github.com/degenerat3/metcli) go package.  

 The only functionality that must be actually created for a module to work is the communication function that can send/recieve a string to/from a client. All request building, core interaction, payload parsing, etc, is handled by the `metcli` package. Any new module must simply spin up some type of listener and have a "handler" function that looks something like the following (the example below is for TCP so it has a `net.Conn` object, this is obviously not required as not all transfer methods will utilize said objects).  

 ```golang
//take the MAD payload and do stuff with it
func connHandle(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString(MAGICTERMBYTE)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	result := server.HandlePayload(message, m)
	conn.Write([]byte(result))
	conn.Close()
}
 ```
The `HandlePayload` function is available via the `github.com/degenerat3/metcli/server` package. In the above context, the `message` is a string representation of the string that has been recieved, and the `m` is a `Metserver` object that is generated at the head of the file (not shown) using the `server.GenMetserver` function. Any communication channel can be used to create a module as long as it has the ability to send/recieve strings in a server-like fashion, and has an associated client that can send/recieve strings. For full source code of an example server, look at `meteor/modules/petrie/pet_server.go`.

