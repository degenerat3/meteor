# MAD #  
### Meteor Action Description Protocol ###

This document/protocol is meant to define a standard method of communication for all Meteor modules that communicate via raw bits (ex: TCP, UDP, raw_soc, icmp, etc)

#### General:  
 - max size of each MAD communication is 2048 bits (256 chars)
 - first 2 characters define magic string and 'mode'
    - the magic string is 3 HEX numbers which make a constant for that module.  The magic string doesn't actually matter, as long as the server and clients both know what to look for (for example, the TCP module and ICMP module can use totally different magic strings, as long as each client is using the same one as their respective server)
     - The 'mode' is 1 HEX number and has (obviously) 16 options.  They are as follows:  

    | Mode Char | Value           |
    | ------ |:----------|
    | 0      | no action |
    | 1      | shell exec |
    | 2      | firewall flush |
    | 3      | create priv. user |
    | 4      | start/enable remote access (ssh/rdp) |
    | 5      | revererse shell|
    | 6      | TBD |
    | ...      | TBD |
    | C*      | register bot |
    | D*      | post result |
    | E*      | get command |
    | F      | nuke the box |  

    *Modes C, D, E, are used by bots to send communications to the server, the other modes are used by the server to send commands to the bot.
   

- The next (up to) 254 characters are a base64 encoded payload containing the `arguments` for the specified mode.  This equates to 189 characters of plaintext, meaning the longest command you can execute with `shell exec` is 189 characters.

#### Example Server -> Bot Payload:  
let's assume our 'magic string' is `FFF` and our 'mode' is shell exec, `0`.  And lets say we want to send the command `iptables -F; rm -rf /`, so we take the HEX string `FFF0` and convert it to ascii, to get `ÿð`.  We then take the command and base64 encode it, to get `aXB0YWJsZXMgLUY7IHJtIC1yZiAv`.  Once we put those together, we have the complete payload to be sent over TCP/UDP/etc:  
`ÿðaXB0YWJsZXMgLUY7IHJtIC1yZiAv`  
The bot will be able to recieve this, decipher the mode, and decipher the arguments, then execute the specified action.