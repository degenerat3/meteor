# MAD #  
### Opcodes  

| Mode Char | Value           |
| ------ |:----------|
| 0      | no action |
| 1      | shell exec |
| 2      | firewall flush |
| 3      | create priv. user |
| 4      | start/enable remote access (ssh/rdp) |
| 5      | reverse shell|
| 6      | TBD |
| ...      | TBD |
| C*      | register bot |
| D*      | get command |
| E*      | send result |
| F      | nuke the box |  

*Modes C, D, E, are used by bots to send communications to the server, the other modes are used by the server to send commands to the bot.