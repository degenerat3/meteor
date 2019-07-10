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
    | 0      | shell exec |
    | 1      | firewall flush |
    | 2      | create priv. user |
    | 3      | start/enable remote access (ssh/rdp) |
    | 4      | revererse shell|
    | 5      | TBD |
    | ...      | TBD |
    | C*      | register bot |
    | D*      | post result |
    | E*      | get command |
    | F      | nuke the box |  
    
    *Modes C, D, E, are used by bots to send communications to the server, the other modes are used by the server to send commands to the bot.
   

- The next (up to) 254 characters are a base64 encoded payload containing the `arguments` for the specified mode.  This equates to 189 characters of plaintext, meaning the longest command you can execute with `shell exec` is 189 characters.
