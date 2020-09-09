# Core API 

### Registration Endpoints
---

### `/register/bot`
Desc: create new bot entry  
Method: `POST`  
Request Proto: `RegBot`  
Request Params:   
```
uuid:      the client-generated unique identifier for the new bot 
interval:  the interval (in seconds) between callbacks 
delta:     the variation (in seconds) of each interval 
hostname:  the host that the bot is running on (generally an IP address) 
```
Response Proto: `GenResp`  
Response Params:   
```
status: HTTP response
msg: error, if any
```

---

### `/register/host`
Desc: create new host entry  
Method: `POST`  
Request Proto: `RegHost`  
Request Params:   
```
hostname:  the name of the host being registered 
interface: the primary interface used by the host 
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```

---

### `/register/group`
Desc: create new group entry
Method: `POST`
Request Proto: `RegGroup`
Request Params: 
```
groupname: the name of the group being registered 
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```

---

### `/register/hostgroup`
Desc: assign a group to a host
Method: `POST`
Request Proto: `RegHG`
Request Params: 
```
hostname:  the host that will be assigned 
groupname: the group that the host will be added to 
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```

---

## Action Endpoints

---

### `/add/action/single`	
Desc: queue a new action assigned to a specific host
Method: `POST`
Request Proto: `AddAct`
Request Params: 
```
mode:      the action mode 
args:      required arg data for the mode type 
target:    the host to run the action against 
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```

---

### `/add/action/group`
Desc: queue a new action assigned to a group
Method: `POST`
Request Proto: `AddAct`
Request Params: 
```
mode:      the action mode 
args:      required arg data for the mode type
target:    the group to run the action against     
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```

---

### `/add/result`
Desc: update an action with the included "result" (usually oputput of action)
Method: `POST`
Request Proto: `AddRes`
Request Params: 
```
aid:       the action id this result is associated with
data:      the action result data to store (usually output of the action)
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```

---

### `/bot/checkin`
Desc: the endpoint listeners will query when a bot "beacons." Checks if any action is pending, returns proto for any pending actions or "None"
Method: `POST`
Request Proto: `CheckIn`
Request Params: 
```
uuid:      the previously-registered unique identifier for the bot 
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: proto for all pending actions for the specified host, separated by "|". ex: 'abcdef=|ABCDEF='
```

---

## List Endpoints

---

### `/list/bots`	
Desc: print all bots and what host they're associated with
Method: `GET`
Request Params: 
```
None
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: newline-separated '<UUID> : <lastseen>'
```

---

### `/list/hosts`
Desc: print all hosts and what group they're associated with
Method: `GET`
Request Params: 
```
None
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: newline-separated '<hostnme> : <group(s)> : <lastseen>'
```

---

### `/list/groups`
Desc: print all group names and how many members they have
Method: `GET`
Request Params: 
```
None
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: newline-separated '<name> : <desc>'
```

---

## Misc Endpoints

---

### `/cleardata`
Desc: delete all the current bots/hosts/groups/actions on the server
Method: `GET`
Request Params: 
```
None
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```

---

### `/status` or `/`
Desc: get status of the Core server, either "Core is running." or no response
Method: `Get`
Request Params: 
```
None
```
Response Proto: `GenResp`
Response Params: 
```
status: HTTP response
msg: error, if any
```
