# Core API
Details on the to communicate with the Meteor Core API.  

All of these requests/responses will make use of the MCS protobuf as defined in `meteor/pbuf/mcs.proto`.  


### Registration Endpoints
---

### `/register/bot`
Desc: create new bot entry  
Method: `POST`   
Request Params:   
```
uuid:      the client-generated unique identifier for the new bot [string]
interval:  the interval (in seconds) between callbacks [int32]
delta:     the variation (in seconds) of each interval [int32]
hostname:  the host that the bot is running on [string]
```  
Response Params:   
```
status: HTTP response [int32]
```

---

### `/register/host`
Desc: create new host entry  
Method: `POST`   
Request Params:   
```
hostname:  the name of the host being registered [string]
interface: the primary interface used by the host [string]
```
Response Params: 
```
status: HTTP response [int32]
```

---

### `/register/group`
Desc: create new group entry  
Method: `POST`    
Request Params:   
```
groupname: the name of the group being registered [string]
desc:      the description of the group [string]
``` 
Response Params:   
```
status: HTTP response [int32]
```

---

### `/register/hostgroup`
Desc: assign a group to a host  
Method: `POST`    
Request Params:   
```
hostname:  the host that will be assigned [string]
groupname: the group that the host will be added to [string]
``` 
Response Params:   
```
status: HTTP response [int32]
```

---

## Action Endpoints

---

### `/add/action/single`	
Desc: queue a new action assigned to a specific host  
Method: `POST`  
Request Params:   
```
mode:      the action mode [string]
args:      required arg data for the mode type [string]
hostname:    the host to run the action against [string]
```
Response Params:   
```
status: HTTP response [int32]
```

---

### `/add/action/group`
Desc: queue a new action assigned to a group  
Method: `POST`   
Request Params:   
```
mode:      the action mode [string]
args:      required arg data for the mode type [string]
groupname:    the group to run the action against [string]     
```
Response Params:   
```
status: HTTP response [int32]
```

---

### `/add/result`
Desc: update an action with the included "result" (usually oputput of action)  
Method: `POST`   
Request Params:   
```
uuid:      the action id this result is associated with [string]
result:      the action result data to store (usually output of the action) [string]
```
Response Params:   
```
status: HTTP response [int32]
```

---

### `/bot/checkin`
Desc: the endpoint listeners will query when a bot "beacons." Checks if any action is pending, returns proto for any pending actions or "None"  
Method: `POST`  
Request Params:   
```
uuid:      the previously-registered unique identifier for the bot [string]
```
Response Params:   
```
status: HTTP response [int32]
actions: an array of actions to execute [Action]
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
Response Params:   
```
status: HTTP response [int32]
desc: newline-separated '<UUID> : <lastseen>' [string]
```

---

### `/list/hosts`
Desc: print all hosts and what group they're associated with  
Method: `GET`  
Request Params:   
```
None
```
Response Params:   
```
status: HTTP response [int32]
desc: newline-separated '<hostnme> : <group(s)> : <lastseen>' [string]
```

---

### `/list/groups`
Desc: print all group names and how many members they have  
Method: `GET`  
Request Params:   
```
None
```
Response Params:   
```
status: HTTP response [int32]
desc: newline-separated '<name> : <desc>' [string]
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
Response Params:   
```
status: HTTP response [int32]
```

---

### `/status` or `/`
Desc: get status of the Core server, either "Core is running." or no response  
Method: `Get`  
Request Params:   
```
None
```  
Response Params:   
```
status: HTTP response [int32]
desc: error, if any [string]
```
