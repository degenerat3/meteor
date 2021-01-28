# Nest API
Details on how to communicate with the Meteor Nest API  

Unlike the Core API, which utilizes protobuf, the Nest API uses json for requests/responses (easier for a normal user to hit this)

### `/status` or `/`
Desc: get status of the Core server, either "Core is running." or no response  
Method: `Get`  
Request Params:   
```
None
```  
Response Params:   
```
a string (not encoded into JSON) stating the API is running
```

---

### `/buildreq`
Desc: pass in the parameters for a bot config, get a link to download a compiled bot with those attributes
Method: `POST`
Request Params:
```
ClientName: The name of the client type (ex: 'petrie') [string]
Server:     The server to communicate with, including any port number (Ex: 127.0.0.1:444) [string]
RegFile:    The location to store the bot's registration file [string]
ObfText:    The text to use for UUID obfuscation [string]
TargetOS:   Target OS to compile for, either 'windows' or 'linux' [string]
``` 
Response Params:
```
msg: A reference link to download the compiled bot, or an error message [string]
```