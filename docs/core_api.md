# Core API 

## Registration

### `/register/bot`
Desc: create new bot entry
Method: `POST`
Params: 
```
TBD
```

### `/register/host`
Desc: create new host entry
Method: `POST`
Params: 
```
TBD
```

### `/register/group`
Desc: create new group entry
Method: `POST`
Params: 
```
TBD
```

### `/register/hostgroup`
Desc: assign a group to a host
Method: `POST`
Params: 
```
TBD
```

## Actions

### `/add/action/single`	
Desc: queue a new action assigned to a specific host
Method: `POST`
Params: 
```
TBD
```

### `/add/action/group`
Desc: queue a new action assigned to a group
Method: `POST`
Params: 
```
TBD
```

### `/add/result`
	update an action with the included "result" (usually oputput of action)
Desc: queue a new action assigned to a group
Method: `POST`
Params: 
```
TBD
```

### `/bot/checkin
Desc: the endpoint listeners will query when a bot "beacons." Checks if any action is pending, returns proto for any pending actions or "None"
Method: `POST`
Params: 
```
TBD
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
Metho: `Get`
