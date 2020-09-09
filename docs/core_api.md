# Core API 

## Registration

### `/register/bot`
Desc: create new bot entry
Method: `POST`
Params: 
```
uuid:      the client-generated unique identifier for the new bot 
interval:  the interval (in seconds) between callbacks 
delta:     the variation (in seconds) of each interval 
hostname:  the host that the bot is running on (generally an IP address) 
```

### `/register/host`
Desc: create new host entry
Method: `POST`
Params: 
```
hostname:  the name of the host being registered 
interface: the primary interface used by the host 
```

### `/register/group`
Desc: create new group entry
Method: `POST`
Params: 
```
groupname: the name of the group being registered 
desc:      a description of the group 
```

### `/register/hostgroup`
Desc: assign a group to a host
Method: `POST`
Params: 
```
hostname:  the host that will be assigned 
groupname: the group that the host will be added to 
```

## Actions

### `/add/action/single`	
Desc: queue a new action assigned to a specific host
Method: `POST`
Params: 
```
mode:      the action mode 
args:      required arg data for the mode type 
target:    the host to run the action against 
```

### `/add/action/group`
Desc: queue a new action assigned to a group
Method: `POST`
Params: 
```
mode:      the action mode 
args:      required arg data for the mode type
target:    the group to run the action against     
```

### `/add/result`
Desc: update an action with the included "result" (usually oputput of action)
Method: `POST`
Params: 
```
aid:       the action id this result is associated with
data:      the action result data to store (usually output of the action)
```

### `/bot/checkin`
Desc: the endpoint listeners will query when a bot "beacons." Checks if any action is pending, returns proto for any pending actions or "None"
Method: `POST`
Params: 
```
uuid:      the previously-registered unique identifier for the bot 
```

## List  

### `/list/bots`	
Desc: print all bots and what host they're associated with
Method: `GET`
Params: 
```
TBD
```

### `/list/hosts`
Desc: print all hosts and what group they're associated with
Method: `GET`
Params: 
```
TBD
```

### `/list/groups`
Desc: print all group names and how many members they have
Method: `GET`
Params: 
```
TBD
```

## Misc

### `/cleardata`
Desc: delete all the current bots/hosts/groups/actions on the server
Method: `GET`

### `/`
Desc: get status of the Core server, either "Core is running." or no response
Method: `Get`
